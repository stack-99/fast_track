/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fast_track/client/client"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// getScoreCmd represents the getScore command
var getScoreCmd = &cobra.Command{
	Use:   "getScore",
	Short: "Gets the score compared to all other quizzers ",
	Run: func(cmd *cobra.Command, args []string) {
		client := client.GetClient()
		score, _ := client.GetUserScore(username)

		log.Info().Msgf("Hello user %s, your score compared to others is %.2f%%", username, score)
	},
}

func init() {
	getScoreCmd.Flags().StringVarP(&username, "username", "u", "default", "Username ")
	rootCmd.AddCommand(getScoreCmd)
}
