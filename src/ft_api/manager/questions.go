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

func loadQuestions(filePath string) ([]Question, error) {
	bytes, err := os.ReadFile(filePath)

	var questions QuestionJson

	if err != nil {
		return []Question{}, err
	}

	if err := json.Unmarshal(bytes, &questions); err != nil {
		return []Question{}, err
	}

	return questions.Questions, nil
}
