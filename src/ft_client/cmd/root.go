package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "cli",
	Short:        "Welcome to our game!",
	Long:         `Run startQuiz to start the game, supplying the username if you want to keep track of your score!`,
	SilenceUsage: true,
	Hidden:       true,
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("Run startQuiz to start the game, supplying the username if you want to keep track of your score!")
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
