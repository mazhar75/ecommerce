package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Service  string
	Version  string
	HttpPORT int
	DbHost   string
	DbPort   string
	DbName   string
	DbUser   string
	DbPass   string
}

func NewConfig() Config {
	loadENV()
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
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	if dbHost == "" || dbPort == "" || dbName == "" || dbUser == "" || dbPass == "" {
		fmt.Println("Check your db configurations in env file")
		os.Exit(1)
	}
	return Config{
		Service:  service,
		Version:  version,
		HttpPORT: port,
		DbHost:   dbHost,
		DbPort:   dbPort,
		DbName:   dbName,
		DbUser:   dbUser,
		DbPass:   dbPass,
	}

}
func DbString() *string {
	cnf := NewConfig()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cnf.DbUser, cnf.DbPass, cnf.DbHost, cnf.DbPort, cnf.DbName)
	return &dsn

}
