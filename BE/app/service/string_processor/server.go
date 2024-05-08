package string_processor

import (
	"context"
	"github.com/Nanhtu187/VNG/BE/app/pkg/logger"
	BE "github.com/Nanhtu187/VNG/BE/proto/rpc/BE/v1"
	"go.uber.org/zap"
)

type Server struct {
	BE.UnimplementedBeServiceServer
	service IService
}

func NewServer(service IService) *Server {
	return &Server{
		service: service,
	}
}

func InitServer() *Server {
	service := NewService()
	return NewServer(service)
}

func (s *Server) StringProcess(ctx context.Context, req *BE.StringProcessRequest) (*BE.StringProcessResponse, error) {
	if err := req.Validate(); err != nil {
		logger.Extract(ctx).Error("Invalid request", zap.Error(err))
		return nil, err
	}

	resp, err := s.service.ProcessString(ctx, StringProcessRequest{Value: req.Value})
	if err != nil {
		logger.Extract(ctx).Error("Error:", zap.Error(err))
		return nil, err
	}
	return &BE.StringProcessResponse{
		Code:    200,
		Message: "Success",
		Data: &BE.StringData{
			Value: resp.Value,
		},
	}, nil
}
