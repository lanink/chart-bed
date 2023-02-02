package configs

import (
	"github.com/spf13/viper"
	"log"
)

var config *Config

type Config struct {
	host    string
	port    string
	dir     string
	ginMode string
}

func (config *Config) Host() string {
	return config.host
}
func (config *Config) Port() string {
	return config.port
}
func (config *Config) Dir() string {
	return config.dir
}
func (config *Config) Server() string {
	return config.host + ":" + config.port
}
func (config *Config) GinMode() string {
	return config.ginMode
}

func ReadConfig() *Config {
	if config != nil {
		return config
	}

	setConfigInfo()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read configs failed: %v", err)
	}

	config = &Config{
		host: viper.GetString("server.host"),
		port: viper.GetString("server.port"),
		dir:  viper.GetString("image.dir"),
	}

	return config
}

func setConfigInfo() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("server.host", "")
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("image.dir", "./images/")

}
