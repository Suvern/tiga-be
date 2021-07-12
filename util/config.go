package util

import "github.com/spf13/viper"

var Config *viper.Viper

func init() {
	config := viper.New()
	config.AddConfigPath(".")
	config.SetConfigName("config")
	config.SetConfigType("json")
	if err := config.ReadInConfig(); err != nil {
		print("读取配置失败")
		panic(err)
	}
	Config = config
}
