package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maiconDeSouza/read_log_api/internal/config"
	"github.com/maiconDeSouza/read_log_api/internal/handlers"
	"github.com/maiconDeSouza/read_log_api/internal/repositories"
	"github.com/maiconDeSouza/read_log_api/internal/router"
	"github.com/maiconDeSouza/read_log_api/internal/services"
)

func main() {
	env, mux, db := config.InitConfig()
	userRepo := repositories.NewUserRepo(db)
	userServices := services.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandlers(userServices)
	router := router.NewRoutes(mux, userHandlers, env.VersionAPI)
	router.InitRoutes()

	fmt.Println("🚀 Servidor iniciado com sucesso na porta " + env.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.ServerPort), mux))
}
