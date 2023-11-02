/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"fast_track/client/client"
	"fast_track/common/models"
)

func startQuiz() {
	log.Info().Msg("Hello Player!")
	log.Info().Msg("Thank you for joining us for tonights quiz!")
}

// startQuizCmd represents the startQuiz command
var rootCmd = &cobra.Command{
	Use:   "startQuiz",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var questionChoices []*models.QuizAnswer
		client := client.GetClient()

		questionsResp, err := client.GetQuestions()

		if err != nil {
			log.Error().Msg("Failed to get questions")
			return
		}

		startTime := time.Now()
		for questIdx, question := range questionsResp.Questions {
			log.Info().Msgf("Question %d: %s", questIdx+1, question.Question.Value)

			for i, choice := range question.Choices {
				log.Info().Msgf("%d) %s", i+1, choice.Value)
			}

			for {
				var chosenChoice uint
				_, err := fmt.Scanln(&chosenChoice)

				if err == nil && chosenChoice-1 < uint(len(question.Choices)) {
					questionChoices = append(questionChoices, &models.QuizAnswer{QuestionId: question.Question.Id,
						ChoiceId: question.Choices[chosenChoice-1].Id})
					break
				} else {
					log.Warn().Msgf("Invalid input!")
				}
			}
		}
		endTime := time.Now()

		answerResp, err := client.GrpcClient.AnswerQuiz(context.Background(), &models.AnswerRequest{Answers: questionChoices,
			SecondsTaken: int32(endTime.Sub(startTime).Seconds())})

		if err != nil {
			log.Error().Msg("Failed to post answers")
			return
		}

		log.Info().Msgf("You got %d correct out of %d", answerResp.CorrectAnswerCount, len(questionsResp.Questions))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//rootCmd.AddCommand(startQuizCmd)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startQuizCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startQuizCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
