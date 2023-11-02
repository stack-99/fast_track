package main

import (
	"fast_track/client/client"
	"fast_track/client/cmd"
	"fmt"

	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return
	}
}

func main() {
	client := client.GetClient()

	if err := client.Initialize(fmt.Sprintf("%s:%d", viper.GetString("Host"), viper.GetInt("Port"))); err != nil {
		log.Fatal().Msgf(err.Error())
	}

	cmd.Execute()
}
