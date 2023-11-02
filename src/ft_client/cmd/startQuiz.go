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

var username string

func answerQuestions(questionsResp *models.QuestionResponse) []*models.QuizAnswer {
	var questionChoices []*models.QuizAnswer

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

	return questionChoices
}

func startQuiz() {
	log.Info().Msgf("Hello Player %s!", username)
	log.Info().Msg("Thank you for joining us for tonight's quiz!")

	client := client.GetClient()

	questionsResp, err := client.GetQuestions()

	if err != nil {
		log.Error().Msgf("Failed to get questions, %v", err)
		return
	}

	startTime := time.Now()
	questionChoices := answerQuestions(questionsResp)
	endTime := time.Now()

	answerResp, err := client.GrpcClient.AnswerQuiz(context.Background(), &models.AnswerRequest{Answers: questionChoices,
		SecondsTaken: int32(endTime.Sub(startTime).Seconds()), Username: username})

	if err != nil {
		log.Error().Msg("Failed to post answers")
		return
	}

	log.Info().Msgf("You got %d correct out of %d with a score compared to others of %.2f", answerResp.CorrectAnswerCount, len(questionsResp.Questions), answerResp.UserComparedScore)
}

// startQuizCmd represents the startQuiz command
var startQuizCmd = &cobra.Command{
	Use:   "startQuiz",
	Short: "Starts a quiz",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		startQuiz()
	},
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	startQuizCmd.Flags().StringVarP(&username, "username", "u", "default", "Username ")
	rootCmd.AddCommand(startQuizCmd)
}
