package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Log    LogConfig    `json:"log" mapstructure:"log"`
	Server ServerConfig `json:"server" mapstructure:"server"`
}

// ServerConfig for configure HTTP & gRPC host & port
type ServerConfig struct {
	HTTP ServerListen `json:"http" mapstructure:"http"`
	GRPC ServerListen `json:"grpc" mapstructure:"grpc"`
}

// ServerListen for specifying host & port
type ServerListen struct {
	Host string `json:"host" mapstructure:"host"`
	Port uint16 `json:"port" mapstructure:"port"`
}

func (s ServerListen) String() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// ListenString for listen to 0.0.0.0
func (s ServerListen) ListenString() string {
	return fmt.Sprintf(":%d", s.Port)
}

func ServerDefaultConfig() ServerConfig {
	return ServerConfig{
		HTTP: ServerListen{
			Host: "localhost",
			Port: 10080,
		},
		GRPC: ServerListen{
			Host: "localhost",
			Port: 10443,
		},
	}
}

func Load() (*Config, error) {
	c := &Config{
		Server: ServerDefaultConfig(),
		Log:    LogDefaultConfig(),
	}
	// --- hacking to load reflect structure config into env ----//
	viper.SetConfigType("json")
	configBuffer, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	viper.ReadConfig(bytes.NewBuffer(configBuffer))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	// -- end of hacking --//
	viper.AutomaticEnv()
	err = viper.Unmarshal(c)
	return c, err
}
