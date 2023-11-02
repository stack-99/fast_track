package main

import (
	"fast_track/client/client"
	"fast_track/client/cmd"

	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	client := client.GetClient()

	if err := client.Initialize("127.0.0.1:4000"); err != nil {
		log.Fatal().Msgf(err.Error())
	}

	cmd.Execute()
}
