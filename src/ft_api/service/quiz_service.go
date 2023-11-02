package service

import (
	"context"
	"fast_track/api/manager"
	"fast_track/common/models"
)

type QuizService struct {
	models.UnimplementedQuizServer
}

func NewQuizService() *QuizService {
	return &QuizService{}
}

func (qs *QuizService) GetQuizQuestions(ctx context.Context, req *models.QuestionRequest) (*models.QuestionResponse, error) {
	ques_resp := &models.QuestionResponse{Questions: manager.GetQuizManager().GetQuestions()}

	return ques_resp, nil
}

func (qs *QuizService) AnswerQuiz(ctx context.Context, req *models.AnswerRequest) (*models.AnswerResponse, error) {
	correctCount, err := manager.GetQuizManager().ValidateAnswers(req.Answers)

	return &models.AnswerResponse{CorrectAnswerCount: correctCount}, err
}
