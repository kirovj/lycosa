package lycosa

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// init prepare to start
func init() {
	// load config from config
	loadConfig()

	// Confirm bash by os
	confirmBash()

	// load tasks from Task and run
	loadTask()
	for _, task := range Tasks {
		if task.Valid {
			runTask(task)
		}
	}
}

func Start() {
	r := gin.Default()

	// api for all task
	r.GET("/all", func(c *gin.Context) {
		c.JSON(http.StatusOK, Tasks)
	})

	var getTaskParams = func(c *gin.Context) (string, string, string) {
		return c.PostForm("name"), c.PostForm("scheduling"), c.PostForm("command")
	}

	// api for add task
	r.POST("/add", func(c *gin.Context) {
		addTask(getTaskParams(c))
		c.String(http.StatusOK, "add task success.")
	})

	// api for change task valid, valid -> invalid, invalid -> valid
	r.POST("/valid", func(c *gin.Context) {
		err := changeTaskValid(c.PostForm("name"))
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.String(http.StatusOK, "change task valid success.")
		}
	})

	// api for change task
	r.POST("/change", func(c *gin.Context) {
		err := changeTask(getTaskParams(c))
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.String(http.StatusOK, "change task success.")
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "aaa")
	})

	_ = r.Run(":10003")
}
