package main

import (
	"github.com/govies/framework/config"
	"github.com/govies/framework/logger"
	"github.com/govies/framework/server"
	"github.com/govies/onboard/domain"
	"github.com/govies/onboard/routes"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.DefaultAppConf()
	l := logger.New(cfg)

	l.Info().Msg("connecting to database")
	err := mgm.SetDefaultConfig(nil, "onboard", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		l.Fatal().Err(err).Msg("failed to connect database")
	}

	user := domain.User{
		Email:        "testing@email.com",
		PasswordHash: "thisIsAPasswordHash",
	}
	if err := mgm.Coll(&user).Create(&user); err != nil {
		l.Info().Err(err).Msg("failed to create sample user.")
	}

	s := server.New(cfg, l)
	s.InitDefaultMiddlewares()

	l.Info().Msg("initializing routes")
	routes.Initialize(s.Engine, l)

	s.Run()
}