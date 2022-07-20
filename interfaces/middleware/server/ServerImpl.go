package server

import (
	"log"
	"os"

	"ws_notifications_email/application"
	"ws_notifications_email/interfaces/middleware"

	"github.com/gin-gonic/gin"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type ServerImpl struct {
	router *gin.Engine
}

func InitServer() Server {
	serverImpl := &ServerImpl{}

	app, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("ws_email"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
	)

	router := gin.Default()
	router.Use(nrgin.Middleware(app))
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
