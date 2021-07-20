package lycosa

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func onlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := tokenInCookie(c); token != "" {
			if user := getUserByToken(token); user != nil {
				c.Next()
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func tokenInCookie(c *gin.Context) string {
	s, _ := c.Cookie("user")
	return s
}
