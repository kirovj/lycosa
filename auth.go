package lycosa

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func onlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			err   error
			user  *User
		)

		if token, err = c.Cookie("user"); err != nil || token == "" {
			c.AbortWithStatus(http.StatusForbidden)
		}

		if user = getUserByToken(token); user == nil {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
