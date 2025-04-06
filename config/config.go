package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

var c config

type config struct {
	System   SystemConfig
	Database DatabaseConfig
}

func Init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.Unmarshal(&c)
}

func GetSystemConfig() SystemConfig {
	return c.System
}

func GetDatabaseConfig() DatabaseConfig {
	return c.Database
}
