package server

import (
	"log"
	"os"

	"ws_notifications_email/application"
	"ws_notifications_email/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

type ServerImpl struct {
	router *gin.Engine
}

func InitServer() Server {
	serverImpl := &ServerImpl{}
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	application.InitController(router)
	serverImpl.router = router
	return serverImpl
}

func (api *ServerImpl) RunServer() {
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = os.Getenv("LOCAL_PORT") //localhost
	}
	log.Fatal(api.router.Run(":" + appPort))
}
