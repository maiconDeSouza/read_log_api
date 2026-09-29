package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maiconDeSouza/read_log_api/internal/config"
)

func main() {
	env, mux, _ := config.InitConfig()

	fmt.Println("🚀 Servidor iniciado com sucesso na porta " + env.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.ServerPort), mux))
}
