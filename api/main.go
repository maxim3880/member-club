package main

import (
	"api/config"
	"api/handler"
	"api/persistence/repo"
	"api/server"
	"api/service"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

const (
	envFilePath = "./config/.env"
)

func main() {
	if err := godotenv.Load(envFilePath); err != nil {
		log.Panic().Err(err).Msgf("failed to read data from .env file")
	}
	cfg, err := config.FromEnv()
	if err != nil {
		log.Panic().Err(err).Msgf("failed to parse .env file")
	}

	rp := repo.NewDatabase()
	svc := service.NewService(rp)
	h := handler.NewServerHandler(cfg, svc)

	serv := server.NewServer(h, cfg)
	serv.Run(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port))
}
