package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Service  string
	Version  string
	HttpPORT int
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("env not loaded")
		os.Exit(1)
	}
	service := os.Getenv("SERVICE_NAME")
	if service == "" {
		fmt.Println("env not loaded")
		os.Exit(1)
	}
	httpPort := os.Getenv("PORT")
	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("Give a valid port")
		os.Exit(1)
	}
	return Config{
		Service:  service,
		Version:  version,
		HttpPORT: port,
	}

}
