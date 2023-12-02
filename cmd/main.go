package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/dmytrodemianchuk/go-crud-csv/internal/config"
	"github.com/dmytrodemianchuk/go-crud-csv/internal/repository"
	"github.com/dmytrodemianchuk/go-crud-csv/internal/server"
	"github.com/dmytrodemianchuk/go-crud-csv/internal/services"
	"github.com/dmytrodemianchuk/go-crud-csv/internal/transport"
	"github.com/dmytrodemianchuk/go-crud-csv/pkg/db/postgres"
	"github.com/dmytrodemianchuk/go-crud-csv/pkg/signaler"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// @title Store API
// @version 1.0

// @host localhost:8080
// @BasePath /

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	cfg := config.Init()

	db, err := postgres.Open(cfg.DB)
	if err != nil {
		logrus.Fatalln(err)
	}

	repo := repository.New(db)
	service := services.New(repo)
	handler := transport.NewHandler(service)

	srv := server.New(cfg.HTTP.Port, handler.Init(cfg))

	go func() {
		if err = srv.Run(); err != nil {
			logrus.Errorf("error start server: %s", err.Error())
		}
	}()

	logrus.Println("rest-api started")

	signaler.Wait()

	if err = srv.Stop(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
