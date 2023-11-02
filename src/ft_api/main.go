package main

import (
	"fmt"
	"net"
	"os"

	"fast_track/api/service"
	"fast_track/common/models"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4000))

	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)

		return
	}

	log.Info().Msgf("Started listening on port %d", 4000)

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	quizService := service.NewQuizService()
	models.RegisterQuiz(grpcServer, quizService)
	grpcServer.Serve(lis)
}
