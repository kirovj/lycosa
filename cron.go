package lycosa

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	_, err := c.AddFunc("*/5 * * * * *", print5)
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
	defer c.Stop()
	select { /* select: to keep this running */
	}
}

//执行函数
func print5() {
	fmt.Println("每5s执行一次cron")
}
