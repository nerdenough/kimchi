package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configure sets all the config for the app.
func Configure() {
	viper.SetConfigFile("bot.config.json")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %s", err))
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)

	level, err := log.ParseLevel(viper.GetString("logLevel"))
	if err != nil {
		log.Error("error parsing log level from config")
	} else {
		log.SetLevel(level)
	}
}
