package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	QueueParams struct {
		Address  string `mapstructure:"queue_addr"`
		Password string `mapstructure:"queue_pass"`
		Number   int    `mapstructure:"queue_num"`
		Protocol int    `mapstructure:"queue_protocol"`
	} `mapstructure:"queue_params"`
	ServerParams struct {
		ListenPort string `mapstructure:"listen_port"`
	} `mapstructure:"server_params"`
	DBParams struct {
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"db_params"`
}

func LoadConfig() *Config {
	viper.AddConfigPath("../internal/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	AppConfig := new(Config)
	viper.Unmarshal(AppConfig)
	return AppConfig
}
