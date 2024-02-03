package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/odanaraujo/user-api/config/env"
	"github.com/odanaraujo/user-api/internal/database"
	"github.com/odanaraujo/user-api/internal/database/sqlc"
	"github.com/odanaraujo/user-api/internal/handler/routes"
	"github.com/odanaraujo/user-api/internal/handler/userhandler"
	"github.com/odanaraujo/user-api/internal/repository/userepository"
	"github.com/odanaraujo/user-api/internal/service/userservice"
	"log/slog"
	"net/http"
)

func main() {
	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load environment variables", err, slog.String("package", "main"))
		return
	}

	dbConnection, err := database.NewDBConnection()

	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
		return
	}

	router := chi.NewRouter()
	queries := sqlc.New(dbConnection)

	//user
	userRepo := userepository.NewUserRepository(dbConnection, queries)
	service := userservice.NewUserService(userRepo)
	userHandler := userhandler.NewUserHandler(service)

	routes.InitUserRoutes(router, userHandler)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	if err = http.ListenAndServe(port, router); err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}

}
