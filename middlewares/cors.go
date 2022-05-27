package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {

	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"*"},
			AllowHeaders:    []string{"Origin"},
			ExposeHeaders:   []string{"Content-Length", "Authorization"},
			// AllowCredentials: true,
			MaxAge: 12 * time.Hour,
		})
	}
}
