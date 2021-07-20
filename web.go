package lycosa

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
)

func init() {
	loadDB()

	go runCron()
}

func setCookie(c *gin.Context, k, v string) {
	c.SetCookie(k, v, 86400, "", "", true, true)
}

func rmCookie(c *gin.Context, k string) {
	c.SetCookie(k, "", 0, "", "", true, true)
}

func getTaskParams(c *gin.Context) (string, string, string) {
	return c.PostForm("name"), c.PostForm("cron"), c.PostForm("cmd")
}

func Start() {
	engine := gin.Default()

	// api for all task
	engine.GET("/", func(c *gin.Context) {
		tasks, _ := getTasks()
		c.JSON(http.StatusOK, tasks)
	})

	// api for login
	engine.POST("/login", func(c *gin.Context) {
		name := c.PostForm("user")
		pass := c.PostForm("pass")
		if user := getUserByName(name); user != nil && user.Pass == pass {
			token := uuid.New()
			setCookie(c, "user", token)
			updateToken(user, token)
			c.String(http.StatusOK, "login success")
		} else {
			c.String(http.StatusOK, "failed, check user or password!")
		}
	})

	admin := engine.Group("/admin")
	admin.Use(onlyAdmin())

	admin.GET("/logout", func(c *gin.Context) {
		rmCookie(c, "user")
		c.String(http.StatusOK, "log out.")
	})

	engine.GET("/users", func(c *gin.Context) {
		users, _ := getUsers()
		c.JSON(http.StatusOK, users)
	})

	// api for add task: /admin/add
	admin.POST("/add", func(c *gin.Context) {
		if err := addTask(getTaskParams(c)); err != nil {
			c.String(http.StatusOK, err.Error())
		}
		c.String(http.StatusOK, "add task success.")
	})

	// api for change task valid, valid -> invalid, invalid -> valid
	admin.POST("/valid", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		if err := updateTaskValid(id); err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.String(http.StatusOK, "change task valid success.")
		}
	})

	// api for change task
	admin.POST("/change", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		name, cron, cmd := getTaskParams(c)
		if err := updateTask(id, name, cron, cmd); err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.String(http.StatusOK, "change task success.")
		}
	})

	_ = engine.Run(":10003")
}
