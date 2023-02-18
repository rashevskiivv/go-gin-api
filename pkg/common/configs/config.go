package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBHost string `mapstructure:"MONGODB_HOST"`
	DBUser string `mapstructure:"MONGODB_USER"`
	DBPass string `mapstructure:"MONGODB_PASS"`
	DBName string `mapstructure:"MONGODB_NAME"`
	DBPort string `mapstructure:"MONGODB_PORT"`
}

var config Config

func LoadConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}
	log.Println("Config loaded")
	return nil
}

func GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
}

func GetConfig() (c Config) {
	return config
}
