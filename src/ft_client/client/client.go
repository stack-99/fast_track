package client

import (
	"context"
	"fast_track/common/models"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type QuizClient struct {
	GrpcClient models.QuizClient
}

func (qz *QuizClient) Initialize(host string) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		error_ret := fmt.Errorf("fail to dial: %v", err)

		return error_ret
	}

	qz.GrpcClient = models.NewQuizClient(conn)

	return nil
}

func (qz *QuizClient) GetQuestions() (*models.QuestionResponse, error) {
	questions_req := &models.QuestionRequest{}

	return qz.GrpcClient.GetQuizQuestions(context.Background(), questions_req)
}

var quizClient *QuizClient
var clientOnce sync.Once

func GetClient() *QuizClient {
	clientOnce.Do(func() {
		quizClient = &QuizClient{}
	})

	return quizClient
}
