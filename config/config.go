package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"log"
)

const fileType = "yaml"

func Init(output io.Writer, file string) error {
	viper.SetConfigType(fileType)
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		_, _ = fmt.Fprintf(output, "Config file changed")
	})
	return nil
}

func MustInit(output io.Writer, cfg string) {
	if err := Init(output, cfg); err != nil {
		log.Fatal("Init config error: ", err)
	}
}
