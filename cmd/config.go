package cmd

import (
	"github.com/spf13/viper"
)

type ConfigType struct {
	DBConn string `mapstructure:"db_conn"`
}

func LoadConfig() (config ConfigType) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	// TODO: logger
	if err != nil {
		panic(err)
	}

	return
}
