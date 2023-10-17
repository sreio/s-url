package core

import (
	"github.com/spf13/viper"
	"github.com/sreio/s-url/define"
	"log"
)

var ViperLib *viper.Viper

func InitConfig() {
	ViperLib = viper.New()
	ViperLib.SetConfigName("config.ini")
	ViperLib.SetConfigType("ini")
	ViperLib.AddConfigPath(".")

	err := ViperLib.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = ViperLib.Unmarshal(define.Config)
	if err != nil {
		log.Fatal(err)
	}
	ViperLib.WatchConfig()
}
