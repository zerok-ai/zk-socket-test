package main

import (
	"awesome-go/config"
	"awesome-go/tcpSockets"
	"github.com/zerok-ai/zk-utils-go/common"
	zkConfig "github.com/zerok-ai/zk-utils-go/config"
)

func main() {

	// read configuration from the file and environment variables
	var cfg config.Configs
	if err := zkConfig.ProcessArgs[config.Configs](&cfg); err != nil {
		panic(err)
	}

	// tcp socket test
	test := tcpSockets.StartTestSocketClientServer(cfg)
	test.StartTest()

	common.BlockUntilChannelClosed(func() error {
		return test.Close()
	})
}
