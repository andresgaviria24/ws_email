package application

import (
	"ws_notifications_email/domain/dto"
	"ws_notifications_email/domain/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	controller service.Service
}

func InitController(router *gin.Engine) {
	controller := Controller{
		controller: service.InitServiceImpl(),
	}

	email := router.Group("/email")

	email.POST("", // middleware.AuthMiddleware(),
		controller.SendEmail)

}

func (cdc *Controller) SendEmail(c *gin.Context) {

	email := dto.Email{}

	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.Response{})
		return
	}

	response := cdc.controller.SendEmail(email)

	if response == nil {
		c.JSON(response.Status, response.Message)
		return
	}

	c.JSON(http.StatusCreated, response)
}
