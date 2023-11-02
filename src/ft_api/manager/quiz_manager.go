package manager

import (
	"fast_track/common/models"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type QuestionStorage interface {
	Initialize(path string)
	LoadQuestions() ([]Question, error)
}

type QuizManager struct {
	questionsPath string

	questions       []*models.QuizQuestion
	questionAnswers map[string]string
	userScores      map[string]int32

	questionStorage QuestionStorage
}

func (qz *QuizManager) initialize() {
	qz.questionStorage.Initialize(qz.questionsPath)

	loadedQuestions, err := qz.questionStorage.LoadQuestions()

	if err != nil {
		return
	}

	qz.questionAnswers = map[string]string{}
	qz.userScores = map[string]int32{}

	for _, question := range loadedQuestions {
		var quizQuestion models.QuizQuestion
		quizQuestion.Question = &models.KeyValue{Id: uuid.New().String(), Value: question.Question}

		for _, choice := range question.Choices {
			quizChoice := models.KeyValue{Id: uuid.New().String(), Value: choice}
			quizQuestion.Choices = append(quizQuestion.Choices, &quizChoice)

			if question.Answer == choice {
				qz.questionAnswers[quizQuestion.Question.Id] = quizChoice.Id
			}
		}

		qz.questions = append(qz.questions, &quizQuestion)
	}
}

func (qz *QuizManager) GetQuestions() []*models.QuizQuestion {
	return qz.questions
}

func (qz *QuizManager) ValidateAnswers(username string, answerResponses []*models.QuizAnswer) (int32, error) {
	var correctCount int32 = 0
	for _, answerResponse := range answerResponses {
		correctAnswerId, ok := qz.questionAnswers[answerResponse.QuestionId]

		if !ok {
			return 0, fmt.Errorf("unable to identify question")
		}

		if correctAnswerId == answerResponse.ChoiceId {
			correctCount++
		}
	}

	qz.userScores[username] = correctCount

	return correctCount, nil
}

func (qz *QuizManager) CalculateUserScore(reqUsername string) float32 {
	reqScore, ok := qz.userScores[reqUsername]

	if !ok {
		return 0
	}

	var betterCount int32 = 0
	for username, score := range qz.userScores {
		if username == reqUsername {
			continue
		}

		if reqScore > score {
			betterCount++
		}
	}

	userScoreLen := int32(len(qz.userScores))

	return (float32(betterCount+1) / float32(userScoreLen)) * 100
}

var quizManager *QuizManager
var clientOnce sync.Once

func GetQuizManager() *QuizManager {
	clientOnce.Do(func() {
		quizManager = &QuizManager{questionsPath: "questions.json", questionStorage: &QuestionJsonStorage{}}
		quizManager.initialize()
	})

	return quizManager
}
