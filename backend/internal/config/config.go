package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

const (
	EnvLocal = "local"
	EnvProd  = "production"
)

type Config struct {
	Env          string
	ServerDomain string           `mapstructure:"server_domain"`
	Postgres     PostgresDatabase `yaml:"postgres"`
	HTTPServer   HTTPServer       `mapstructure:"http_server"`
	Gpt          YandexGptApi
}

type YandexGptApi struct {
	OAuth2Token string
	Dir         string
}
type HTTPServer struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type PostgresDatabase struct {
	RetryInterval time.Duration `mapstructure:"retry_interval"`
	MaxRetries    int           `mapstructure:"max_retries"`
	Port          int           `yaml:"port"`
	Host          string        `yaml:"host"`
	Database      string        `yaml:"database"`
	Username      string        `yaml:"username"`
	Password      string        `yaml:"password"`
}

func MustLoad() *Config {

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config"
	}

	viper.SetConfigName(os.Getenv("ENV"))
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("config file doesn't exists: %s", configPath)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode config into struct: %v", err)
	}

	config.Postgres.Password = os.Getenv("PG_PWD")
	config.Postgres.Username = os.Getenv("PG_USR")
	config.Postgres.Host = os.Getenv("PG_HOST")
	config.Postgres.Database = os.Getenv("PG_DB")

	portStr := os.Getenv("PG_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid postgres port: %s", portStr)
	}
	config.Postgres.Port = port
	config.Env = os.Getenv("ENV")

	return &config
}
