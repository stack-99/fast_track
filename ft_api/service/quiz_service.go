package service

import (
	"context"

	"github.com/stack-99/gRPC-example/api/manager"
	"github.com/stack-99/gRPC-example/common/models"
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
	correctCount, err := manager.GetQuizManager().ValidateAnswers(req.Username, req.Answers)

	return &models.AnswerResponse{CorrectAnswerCount: correctCount,
		UserComparedScore: manager.GetQuizManager().CalculateUserScore(req.Username)}, err
}

func (qs *QuizService) GetUserScore(ctx context.Context, req *models.UserScoreRequest) (*models.UserScoreResponse, error) {
	return &models.UserScoreResponse{
		UserComparedScore: manager.GetQuizManager().CalculateUserScore(req.Username)}, nil
}
