package initializer

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"GOTODO_POSTGRES_HOST"`
	DBUser     string `mapstructure:"GOTODO_POSTGRES_USER"`
	DBPassword string `mapstructure:"GOTODO_POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"GOTODO_POSTGRES_DB"`
	DBPort     string `mapstructure:"GOTODO_POSTGRES_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	return
}
