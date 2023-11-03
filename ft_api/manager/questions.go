package manager

import (
	"encoding/json"
	"os"
)

type Question struct {
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
	Answer   string   `json:"answer"`
}

type QuestionJson struct {
	Questions []Question
}

type QuestionJsonStorage struct {
	filePath string
}

func (qjs *QuestionJsonStorage) Initialize(filePath string) {
	qjs.filePath = filePath
}

func (qjs *QuestionJsonStorage) LoadQuestions() ([]Question, error) {
	bytes, err := os.ReadFile(qjs.filePath)

	var questions QuestionJson

	if err != nil {
		return []Question{}, err
	}

	if err := json.Unmarshal(bytes, &questions); err != nil {
		return []Question{}, err
	}

	return questions.Questions, nil
}
