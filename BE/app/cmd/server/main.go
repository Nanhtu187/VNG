package main

import (
	"context"
	"fmt"
	"github.com/Nanhtu187/VNG/BE/app/config"
	"github.com/Nanhtu187/VNG/BE/app/pkg/errors"
	log "github.com/Nanhtu187/VNG/BE/app/pkg/logger"
	"github.com/Nanhtu187/VNG/BE/app/service/string_processor"
	BE "github.com/Nanhtu187/VNG/BE/proto/rpc/BE/v1"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	_ "github.com/rs/cors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

func main() {
	errors.FinishNewErrors()

	rootCmd := cobra.Command{
		Use: "server",
	}
	rootCmd.AddCommand(
		startServerCommand(),
	)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}

func startServer() {
	conf, err := config.Load()
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	logger := config.NewLogger(conf.Log)
	grpcServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryHandlerFunc)),
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
		),
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryHandlerFunc)),
			grpc_prometheus.UnaryServerInterceptor,
			log.SetTraceInfoInterceptor(logger),
			errors.UnaryServerInterceptor,
		),
	)

	StringProcessServer := string_processor.InitServer()
	BE.RegisterBeServiceServer(grpcServer, StringProcessServer)
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcServer)

	startHTTPAndGRPCServers(*conf, grpcServer)
}

func startHTTPAndGRPCServers(conf config.Config, grpcServer *grpc.Server) {
	fmt.Println("GRPC:", conf.Server.GRPC.ListenString())
	fmt.Println("HTTP:", conf.Server.HTTP.ListenString())

	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(errors.CustomerHTTPError),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}),
	)

	ctx := context.Background()
	grpcHost := conf.Server.GRPC.String()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	registerGRPCGateway(ctx, mux, grpcHost, opts)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust this to your needs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:            true, // Enable debug to see what's happening
	})

	httpMux := http.NewServeMux()
	httpMux.Handle("/metrics", promhttp.Handler())
	httpMux.Handle("/", c.Handler(mux))

	httpServer := &http.Server{
		Addr:    conf.Server.HTTP.ListenString(),
		Handler: httpMux,
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		fmt.Println("Shutdown HTTP server successfully")
	}()

	go func() {
		defer wg.Done()

		listener, err := net.Listen("tcp", conf.Server.GRPC.ListenString())
		if err != nil {
			panic(err)
		}

		err = grpcServer.Serve(listener)
		if err != nil {
			panic(err)
		}
		fmt.Println("Shutdown gRPC server successfully")
	}()

	//--------------------------------
	// Graceful Shutdown
	//--------------------------------
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	grpcServer.GracefulStop()
	err := httpServer.Shutdown(ctx)
	if err != nil {
		panic(err)
	}

	wg.Wait()
}

func registerGRPCGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	_ = BE.RegisterBeServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func recoveryHandlerFunc(p interface{}) error {
	fmt.Println("stacktrace from panic:\n" + string(debug.Stack()))
	return status.Errorf(codes.Internal, "panic: %s", p)
}

func startServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start the server",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
}
