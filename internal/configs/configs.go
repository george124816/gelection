package configs

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

type HttpConfig struct {
	Port uint16
}

func (h *HttpConfig) GetStringPort() string {
	return fmt.Sprintf(":%d", h.Port)

}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	SslMode  string
}

func (c DatabaseConfig) String() string {
	pgUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.Database, c.SslMode)
	u, err := url.Parse(pgUrl)
	if err != nil {
		log.Fatal(err)
	}
	return u.String()
}

func GetPostgressConfig() DatabaseConfig {
	user := GetEnvOrDefault("DATABASE_USER", "postgres")
	password := url.QueryEscape(GetEnvOrDefault("DATABASE_PASSWORD", "postgres"))
	host := GetEnvOrDefault("DATABASE_HOST", "localhost")
	port := GetEnvOrDefault("DATABASE_PORT", "5555")
	database := GetEnvOrDefault("DATABASE_NAME", "postgres")
	sslMode := GetEnvOrDefault("DATABASE_SSL_MODE", "disable")

	return DatabaseConfig{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
		SslMode:  sslMode,
	}

}

func GetEnvOrDefault(envName string, defaultValue string) string {
	resultValue := os.Getenv(envName)
	if resultValue != "" {
		return resultValue
	} else {
		return defaultValue
	}
}
