package lycosa

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// init prepare to start
func init() {
	// load config from config
	LoadConfig()

	// load tasks from Task
	LoadTask()
}

func Start() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "aaa")
	})
}
