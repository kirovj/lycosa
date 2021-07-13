package lycosa

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	loadTask()
	for _, task := range Tasks {
		fmt.Println(task)
	}
}
