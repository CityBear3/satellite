package driver

import (
	"github.com/CityBear3/satellite/pkg/env"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"strings"
)

type DBConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type MinioConfig struct {
	Host string `yaml:"host"`
	Port int
}

type ServerConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	IsDevelop bool   `yaml:"is_develop"`
}

type Config struct {
	ServerConfig ServerConfig `yaml:"server"`
	DBConfig     DBConfig     `yaml:"database"`
}

func LoadConfig(path string) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			key := strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")
			if strings.Contains(key, ":-") {
				values := strings.Split(key, ":-")
				v := env.GetStrEnv(values[0], values[1])
				viper.Set(k, v)
				continue
			}
			v, err := env.GetRequiredStrEnv(key)
			if err != nil {
				return Config{}, err
			}
			viper.Set(k, v)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "yaml"
	}); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
