package lycosa

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "aaa")
	})
}
