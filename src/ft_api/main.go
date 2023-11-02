package main

import (
	"fmt"
	"net"
	"os"

	"fast_track/api/service"
	"fast_track/common/models"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return
	}
}

func main() {
	port := viper.GetInt("Port")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatal().Msgf("Failed to listen: %v", err)

		return
	}

	log.Info().Msgf("Started listening on port %d", port)

	var opts []grpc.ServerOption

	if viper.GetBool("TLS") {
		var certFile string = viper.GetString("ServerCertPath")
		var keyFile string = viper.GetString("ServerKeyPath")

		if certFile == "" {
			certFile = "certs/server-cert.pem"
		}
		if keyFile == "" {
			keyFile = "certs/server-key.pem"
		}

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatal().Msgf("Failed to initialize tls from file: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}

		log.Info().Msg("Configured TLS")
	}

	grpcServer := grpc.NewServer(opts...)
	quizService := service.NewQuizService()
	models.RegisterQuizServer(grpcServer, quizService)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal().Msg("Failed to serve grpc server")
	}
}
