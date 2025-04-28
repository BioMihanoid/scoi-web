package main

import (
	"log"
	"scoi-web/internal/api"
	"scoi-web/internal/config"
	"scoi-web/internal/repository"
	"scoi-web/internal/repository/postgres"
	"scoi-web/internal/service"
)

func main() {
	cfg := config.MustLoad()
	db, err := postgres.NewPostgresDb(cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := api.NewHandler(services)

	r := handler.InitRoutes()
	err = r.Run(cfg.Server.Port)
	if err != nil {
		panic(err)
	}
}
