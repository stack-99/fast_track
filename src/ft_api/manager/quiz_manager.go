package manager

import (
	"fast_track/common/models"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type QuizManager struct {
	questions       []*models.QuizQuestion
	questionsPath   string
	questionAnswers map[string]string
}

func (qz *QuizManager) initialize() {
	loadedQuestions, err := loadQuestions(qz.questionsPath)

	if err != nil {
		return
	}

	qz.questionAnswers = map[string]string{}

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

func (qz *QuizManager) ValidateAnswers(answerResponses []*models.QuizAnswer) (int32, error) {
	var correctCount int32 = 0
	for _, answerResponse := range answerResponses {
		if correctAnswerId, ok := qz.questionAnswers[answerResponse.QuestionId]; ok {
			if correctAnswerId == answerResponse.ChoiceId {
				correctCount++
			}
		} else {
			return 0, fmt.Errorf("unable to identify question")
		}
	}

	return correctCount, nil
}

var quizManager *QuizManager
var clientOnce sync.Once

func GetQuizManager() *QuizManager {
	clientOnce.Do(func() {
		quizManager = &QuizManager{questionsPath: "questions.json"}
		quizManager.initialize()
	})

	return quizManager
}
