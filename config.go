package main

import "github.com/spf13/viper"

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.AddConfigPath(".")
	Config.SetConfigName("config")
	Config.SetConfigType("json")
	if err := Config.ReadInConfig(); err != nil {
		print("读取配置失败")
		panic(err)
	}
}
