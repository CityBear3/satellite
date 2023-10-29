package driver

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	DBConfig       DBConfig
	ServerConfig   ServerConfig
	AuthConfig     AuthConfig
	RabbitMQConfig RabbitMQConfig
}

type DBConfig struct {
	Driver   string `env:"DB_DRIVER" envDefault:"mysql"`
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Port     int    `env:"DATABASE_PORT" envDefault:"3306"`
	DbName   string `env:"DATABASE_NAME" envDefault:"satellite"`
	User     string `env:"DATABASE_USER" envDefault:"satellite"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"satellite"`
}

type MinioConfig struct {
	Host string `env:"MINIO_HOST" envDefault:"localhost"`
	Port int    `env:"MINIO_PORT" envDefault:"9000"`
}

type ServerConfig struct {
	Host      string `env:"SERVER_HOST" envDefault:"localhost"`
	Port      int    `env:"SERVER_PORT" envDefault:"8080"`
	IsDevelop bool   `env:"SERVER_IS_DEVELOP" envDefault:"false"`
}

type AuthConfig struct {
	HMACSecret string `env:"HMAC_SECRET"`
}

type RabbitMQConfig struct {
	Host     string `env:"RABBITMQ_HOST" envDefault:"localhost"`
	Port     int    `env:"RABBITMQ_PORT" envDefault:"5672"`
	User     string `env:"RABBITMQ_USER" envDefault:"satellite"`
	Password string `env:"RABBITMQ_PASSWORD" envDefault:"satellite"`
}

func LoadConfig(config *Config) error {
	if err := env.Parse(config); err != nil {
		return err
	}

	return nil
}
