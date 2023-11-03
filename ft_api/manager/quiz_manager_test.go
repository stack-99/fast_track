package manager

import (
	"fast_track/common/models"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type questionMemoryStorage struct {
}

func (qjs *questionMemoryStorage) Initialize(filePath string) {
}

func (qjs *questionMemoryStorage) LoadQuestions() ([]Question, error) {
	var questions []Question

	questions = append(questions, Question{Question: "Q1", Choices: []string{"1", "2", "3"}, Answer: "1"})
	questions = append(questions, Question{Question: "Q2", Choices: []string{"1", "2", "3"}, Answer: "2"})
	questions = append(questions, Question{Question: "Q3", Choices: []string{"1", "2", "3"}, Answer: "3"})

	return questions, nil
}

func getTestQuizManager() *QuizManager {
	quizManager = &QuizManager{questionStorage: &questionMemoryStorage{}}
	quizManager.initialize()

	return quizManager
}

func TestGetQuestions(t *testing.T) {
	quizManager := getTestQuizManager()

	assert.True(t, len(quizManager.questions) != 0)

	for _, question := range quizManager.questions {
		answerId, ok := quizManager.questionAnswers[question.Question.Id]
		assert.True(t, ok)

		foundChoice := false
		for _, choice := range question.Choices {
			if choice.Id == answerId {
				foundChoice = true
				break
			}
		}

		assert.True(t, foundChoice)
	}
}

func TestAllCorrectValidateAnswer(t *testing.T) {
	quizManager := getTestQuizManager()

	username := "Test001"

	var answers []*models.QuizAnswer

	for questId, answId := range quizManager.questionAnswers {
		answers = append(answers, &models.QuizAnswer{QuestionId: questId, ChoiceId: answId})
	}

	correctCount, err := quizManager.ValidateAnswers(username, answers)

	assert.Nil(t, err)

	assert.Equal(t, correctCount, int32(len(quizManager.questions)))
}

func TestBadValidateAnswer(t *testing.T) {
	quizManager := getTestQuizManager()

	username := "Test001"

	var answers []*models.QuizAnswer

	for questId, _ := range quizManager.questionAnswers {
		answers = append(answers, &models.QuizAnswer{QuestionId: questId, ChoiceId: ""})
	}

	correctCount, err := quizManager.ValidateAnswers(username, answers)

	assert.Nil(t, err)

	assert.True(t, correctCount == 0)
	assert.NotEqual(t, correctCount, int32(len(quizManager.questions)))
}

func TestCalculateAboveAverageUserScores(t *testing.T) {
	quizManager := getTestQuizManager()

	username := "Test999"

	var answers []*models.QuizAnswer

	for questId, answId := range quizManager.questionAnswers {
		answers = append(answers, &models.QuizAnswer{QuestionId: questId, ChoiceId: answId})
	}

	var badAnswers []*models.QuizAnswer

	for questId := range quizManager.questionAnswers {
		badAnswers = append(badAnswers, &models.QuizAnswer{QuestionId: questId, ChoiceId: ""})
	}

	correctCount, _ := quizManager.ValidateAnswers(username, answers)
	assert.Equal(t, correctCount, int32(len(quizManager.questions)))

	userScore := quizManager.CalculateUserScore(username)

	assert.True(t, userScore == 100)

	badScoreAmount := 0
	for i := 0; i < 9; i++ {
		otherUsername := "Test00" + strconv.Itoa(i)

		if i%4 == 0 {
			correctCount, _ := quizManager.ValidateAnswers(otherUsername, answers)

			assert.Equal(t, correctCount, int32(len(quizManager.questions)))
		} else {
			badScoreAmount++
			correctCount, _ := quizManager.ValidateAnswers(otherUsername, badAnswers)

			assert.Equal(t, int32(0), correctCount)
		}

		assert.Equal(t, int32(len(quizManager.questions)), correctCount)
	}

	userScore = quizManager.CalculateUserScore(username)

	assert.Equal(t, 6, badScoreAmount)

	// 10 Players, 6 of them which have bad scores
	assert.Equal(t, 10, len(quizManager.userScores))

	// Our user is better than 6 people
	t.Logf("User amount: %d, Players with Bad scores = %d, User Scores: %.2f better than others",
		len(quizManager.userScores), badScoreAmount, userScore)

	assert.True(t, userScore == 70)
}
