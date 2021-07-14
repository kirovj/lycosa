package lycosa

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
)

// init prepare to start
func init() {
	// load config from config
	loadConfig()

	// confirm bash by os
	confirmBash()

	// load tasks from Task and run
	loadTask()

	// cron run
	go runCron()
}

func setCookie(c *gin.Context, k, v string) {
	c.SetCookie(k, v, 86400*7, "", "", true, true)
}

func rmCookie(c *gin.Context, k string) {
	c.SetCookie(k, "", 0, "", "", true, true)
}

func getTaskParams(c *gin.Context) (string, string, string) {
	return c.PostForm("name"), c.PostForm("scheduling"), c.PostForm("command")
}

func Start() {
	r := gin.Default()

	// api for all task
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, Tasks)
	})

	// api for login
	r.POST("/login", func(c *gin.Context) {
		user := c.PostForm("user")
		pass := c.PostForm("pass")
		if user == Conf.User && pass == Conf.Pass {
			sessionId = uuid.New()
			setCookie(c, "user", sessionId)
			c.String(http.StatusOK, "login success")
		} else {
			c.String(http.StatusOK, "failed, check user or password!")
		}
	})

	admin := r.Group("/admin")
	admin.Use(onlyAdmin())

	admin.GET("/logout", func(c *gin.Context) {
		sessionId = ""
		rmCookie(c, "user")
		c.String(http.StatusOK, "log out.")
	})

	admin.GET("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, Conf)
	})

	// api for add task: /admin/add
	admin.POST("/add", func(c *gin.Context) {
		addTask(getTaskParams(c))
		c.String(http.StatusOK, "add task success.")
	})

	// api for change task valid, valid -> invalid, invalid -> valid
	admin.POST("/valid", func(c *gin.Context) {
		err := changeTaskValid(c.PostForm("name"))
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.String(http.StatusOK, "change task valid success.")
		}
	})

	// api for change task
	admin.POST("/change", func(c *gin.Context) {
		err := changeTask(getTaskParams(c))
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.String(http.StatusOK, "change task success.")
		}
	})

	_ = r.Run(":10003")
}
