package middleware

import (
	"net/http"
	"os"
	"ws_notifications_email/domain/dto"
	"ws_notifications_email/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	var language string
	var responseDto = dto.Response{}

	return func(c *gin.Context) {
		language = c.Request.Header[os.Getenv("LENGUAGE_HEADER")][0]
		descripcion, code := utils.VerifyToken(c.Request, language, os.Getenv("PUBLIC_RSA"))
		if code != http.StatusOK {
			responseDto.Status = code
			responseDto.Description = http.StatusText(code)
			responseDto.Message = descripcion
			c.JSON(code, responseDto)
			c.AbortWithStatus(code)
			return
		}
		c.Next()
	}
}
