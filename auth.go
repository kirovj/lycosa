package lycosa

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var sessionId string

func onlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if id := sessionInCookie(c); id != "" {
			sessionId = id
			c.Next()
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func sessionInCookie(c *gin.Context) string {
	s, _ := c.Cookie("user")
	return s
}
