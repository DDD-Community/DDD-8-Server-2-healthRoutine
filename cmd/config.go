package cmd

import (
	"github.com/spf13/viper"
	"healthRoutine/pkgs/log"
)

const (
	named = "CONFIG"
)

type ConfigType struct {
	DBConn        string `mapstructure:"db_conn"`
	SigningSecret string `mapstructure:"signing_secret"`
}

func LoadConfig() (config ConfigType) {
	logger := log.Get()
	defer logger.Sync()

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Named(named).Error("failed to load config")
		logger.Named(named).Error(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		logger.Named(named).Error("failed to unmarshal config")
		logger.Named(named).Error(err)
	}

	return
}
