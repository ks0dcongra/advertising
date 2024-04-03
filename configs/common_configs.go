package configs

import (
	"fmt"
	"os"
)

type commonConfig struct {
	ServerPort        string
}

func getCommonConfig() (*commonConfig, error) {
	var serverPort string
	
	serverPort = os.Getenv("SERVERPORT")
	if len(serverPort) == 0 {
		serverPort = "9520"
	}

	fmt.Println("Success to load the Service configuration")

	return &commonConfig{
		ServerPort:        serverPort,
	}, nil
}
