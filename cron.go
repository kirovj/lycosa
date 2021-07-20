package lycosa

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var cronManager = cron.New(cron.WithSeconds())

func runCron() {
	tasks, _ := getTasks()
	for _, t := range *tasks {
		if _, err := cronManager.AddFunc(t.Cron, t.runTask); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("task: %s start run with %s\n", t.Name, t.Cron)
	}
	cronManager.Start()
}
