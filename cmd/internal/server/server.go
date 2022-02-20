package server

import (
	"github.com/Hitsugaya/rest-api-project/cmd/internal/config"
	"github.com/Hitsugaya/rest-api-project/cmd/internal/controllers"
	"github.com/Hitsugaya/rest-api-project/cmd/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Run() {
	RunServer()
}

func RunServer() {

	serverConfig := config.GetServerConfig()
	circumferenceConfig := config.GetDatabaseConfig()

	router := gin.New()
	allServices := services.NewAllServices(circumferenceConfig)
	controller := controllers.NewController(allServices)

	//register routers
	controller.Register(router)

	server := &http.Server{
		Addr:           ":" + serverConfig.Server.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
