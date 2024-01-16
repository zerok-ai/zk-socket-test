package config

import (
	zkLogsConfig "github.com/zerok-ai/zk-utils-go/logs/config"
	"github.com/zerok-ai/zk-utils-go/socket"
)

type AppConfig struct {
	StartServerOnBoot bool `yaml:"startServerOnBoot" env:"START_SOCKET_SERVER" env-default:"true"`
	StartClientOnBoot bool `yaml:"runTestClientOnBoot" env:"RUN_TEST_CLIENT" env-default:"false"`
}

// Configs is an application configuration structure
type Configs struct {
	TCPServerConfig socket.TCPServerConfig  `yaml:"server"`
	TCPClientConfig socket.TCPClientConfig  `yaml:"client"`
	LogsConfig      zkLogsConfig.LogsConfig `yaml:"logs"`
	AppConfig       AppConfig               `yaml:"app"`
}
