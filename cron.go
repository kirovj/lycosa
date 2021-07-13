package lycosa

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func runCron() {
	var (
		c   *cron.Cron
		err error
	)
	c = cron.New(cron.WithSeconds())
	defer c.Stop()

	for _, t := range Tasks {
		if _, err = c.AddFunc(t.Scheduling, t.runTask); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("task: %s start run with %s\n", t.Name, t.Scheduling)
	}

	c.Start()
	select { /* select: to keep this running */
	}
}
