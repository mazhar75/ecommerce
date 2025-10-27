package config

import (
	"fmt"
	"os"
	"strconv"
)

type ServerConfig struct {
	service  string
	version  string
	httpPORT int
}
type PostgresqlConfig struct {
	dbHost string
	dbPort string
	dbName string
	dbUser string
	dbPass string
}
type JWTConfig struct {
	jwtSecret string
}
type RedisConfig struct {
	redisIP   string
	redisPort int
	redisPass string
	redisDb   int
}

var serverConfig *ServerConfig
var postgresqlConfig *PostgresqlConfig
var jwtConfig *JWTConfig
var redisConfig *RedisConfig

func NewConfig() {

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
		fmt.Println("Give a valid server port")
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
	jwtSecrets := os.Getenv("JWT_SECRET")
	redisIP := os.Getenv("REDIS_IP")
	redis_port := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASSWORD")
	redis_db := os.Getenv("REDIS_DB")

	redisPort, err := strconv.Atoi(redis_port)
	if err != nil {
		fmt.Println("Give a valid redis port")
		os.Exit(1)
	}
	redisDb, err := strconv.Atoi(redis_db)
	if err != nil {
		fmt.Println("Give a valid redis db")
		os.Exit(1)
	}

	serverConfig = &ServerConfig{
		service:  service,
		version:  version,
		httpPORT: port,
	}
	postgresqlConfig = &PostgresqlConfig{
		dbHost: dbHost,
		dbPort: dbPort,
		dbName: dbName,
		dbUser: dbUser,
		dbPass: dbPass,
	}
	jwtConfig = &JWTConfig{
		jwtSecret: jwtSecrets,
	}
	redisConfig = &RedisConfig{
		redisIP:   redisIP,
		redisPort: redisPort,
		redisPass: redisPass,
		redisDb:   redisDb,
	}

}
func DbString() *string {
	if postgresqlConfig == nil {
		NewConfig()
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgresqlConfig.dbUser, postgresqlConfig.dbPass, postgresqlConfig.dbHost, postgresqlConfig.dbPort, postgresqlConfig.dbName)
	return &dsn

}
func Get_JWT_SECRET() *string {
	if jwtConfig == nil {
		NewConfig()
	}
	return &jwtConfig.jwtSecret
}
func Get_Redis_Config() (string, string, int) {
	if redisConfig == nil {
		NewConfig()
	}
	addr := redisConfig.redisIP + ":" + strconv.Itoa(redisConfig.redisPort)
	return addr, redisConfig.redisPass, redisConfig.redisDb
}
