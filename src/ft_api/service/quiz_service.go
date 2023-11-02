package service

import (
	"fast_track/common/models"
)

type QuizService struct {
	models.UnimplementedQuiz
}

func NewQuizService() *QuizService {
	return &QuizService{}
}
