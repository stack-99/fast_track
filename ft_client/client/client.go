package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/stack-99/fast_track/common/models"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type QuizClient struct {
	GrpcClient models.QuizClient
}

func (qz *QuizClient) Initialize(host string) error {
	var opts []grpc.DialOption

	if viper.GetBool("TLS") {
		var certFile string = viper.GetString("ServerRootCertPath")

		if certFile == "" {
			certFile = "certs/ca-cert.pem"
		}

		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			return fmt.Errorf("failed to initialize tls from file: %v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		error_ret := fmt.Errorf("failed to connect to host %s, error = %v", host, err)

		return error_ret
	}

	qz.GrpcClient = models.NewQuizClient(conn)

	return nil
}

func (qz *QuizClient) GetQuestions() (*models.QuestionResponse, error) {
	questionsReq := &models.QuestionRequest{}

	return qz.GrpcClient.GetQuizQuestions(context.Background(), questionsReq)
}

func (qz *QuizClient) GetUserScore(username string) (float32, error) {
	userScoreReq := &models.UserScoreRequest{Username: username}

	userScoreRes, err := qz.GrpcClient.GetUserScore(context.Background(), userScoreReq)

	if err != nil {
		return 0, err
	}

	return userScoreRes.UserComparedScore, nil
}

var quizClient *QuizClient
var clientOnce sync.Once

func GetClient() *QuizClient {
	clientOnce.Do(func() {
		quizClient = &QuizClient{}
	})

	return quizClient
}
