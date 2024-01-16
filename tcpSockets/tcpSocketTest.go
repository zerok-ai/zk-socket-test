package tcpSockets

import (
	"awesome-go/config"
	"fmt"
	"github.com/zerok-ai/zk-utils-go/socket"
	"time"
)

type TestTCPSocketClientServer struct {
	client *socket.TCPClient
	server *socket.TCPServer

	configs config.Configs
}

func StartTestSocketClientServer(configs config.Configs) *TestTCPSocketClientServer {
	testSocketClientServer := TestTCPSocketClientServer{configs: configs}
	return &testSocketClientServer
}

func (test TestTCPSocketClientServer) StartTest() {
	if test.configs.AppConfig.StartServerOnBoot {
		go test.startServer()
	}

	if test.configs.AppConfig.StartClientOnBoot {
		go test.startClient()
	}
}

func (test TestTCPSocketClientServer) Close() error {
	if test.client != nil {
		test.client.Close()
	}

	if test.server != nil {
		test.server.Close()
	}
	return nil
}

func (test TestTCPSocketClientServer) startServer() {
	tcpServer := socket.CreateTCPServer(test.configs.TCPServerConfig, func(data []byte) string {
		return "Server received: " + string(data[:])
	})
	test.server = tcpServer
	test.server.Start()
}

var messageId = 0

func (test TestTCPSocketClientServer) sendData() {
	messageId += 1
	test.client.SendData([]byte(fmt.Sprintf("Message from client %d", messageId)))
}

func (test TestTCPSocketClientServer) startClient() {
	client := socket.CreateTCPClient(test.configs.TCPClientConfig)
	test.client = client

	interval := 1 // seconds
	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	// Start a goroutine to call the function at regular intervals
	go func() {
		var index = 0
		for {
			select {
			case <-ticker.C:
				if index == 0 {
					if !test.client.Connect() {
						continue
					}
				} else {
					test.sendData()
				}
			}
			index += 1
		}
	}()
}
