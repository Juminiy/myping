package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func AppConfig() {
	viper.SetConfigName("myping")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("🌐 Config has already been mastered globally!")
		} else {
			fmt.Println("Error is : ", err)
		}
	}
}
