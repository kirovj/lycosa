package lycosa

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	loadDB()
}

func setCookie(c *gin.Context, k, v string) {
	c.SetCookie(k, v, 86400*7, "", "", true, true)
}

func rmCookie(c *gin.Context, k string) {
	c.SetCookie(k, "", 0, "", "", true, true)
}

func getTaskParams(c *gin.Context) (string, string, string) {
	return c.PostForm("name"), c.PostForm("cron"), c.PostForm("cmd")
}

func Start() {
	r := gin.Default()

	// api for all task
	r.GET("/", func(c *gin.Context) {
		tasks, _ := getTasks()
		c.JSON(http.StatusOK, tasks)
	})

	// api for login
	// r.POST("/login", func(c *gin.Context) {
	// 	name := c.PostForm("user")
	// 	pass := c.PostForm("pass")
	// 	if user := getUser(name); user != nil && user.Pass == pass {
	// 		sessionId = uuid.New()
	// 		setCookie(c, "user", sessionId)
	// 		c.String(http.StatusOK, "login success")
	// 	} else {
	// 		c.String(http.StatusOK, "failed, check user or password!")
	// 	}
	// })

	// admin := r.Group("/admin")
	// admin.Use(onlyAdmin())

	// admin.GET("/logout", func(c *gin.Context) {
	// 	sessionId = ""
	// 	rmCookie(c, "user")
	// 	c.String(http.StatusOK, "log out.")
	// })

	// api for add task: /admin/add
	r.POST("/add", func(c *gin.Context) {
		if err := addTask(getTaskParams(c)); err != nil {
			c.String(http.StatusOK, err.Error())
		}
		c.String(http.StatusOK, "add task success.")
	})

	// api for change task valid, valid -> invalid, invalid -> valid
	r.POST("/valid", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		if err := updateTaskValid(id); err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.String(http.StatusOK, "change task valid success.")
		}
	})

	// api for change task
	r.POST("/change", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		name, cron, cmd := getTaskParams(c)
		if err := updateTask(id, name, cron, cmd); err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.String(http.StatusOK, "change task success.")
		}
	})

	_ = r.Run(":10003")
}
