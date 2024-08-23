package config

import "github.com/spf13/viper"

func InitConfig() {

	viper.AddConfigPath("config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
