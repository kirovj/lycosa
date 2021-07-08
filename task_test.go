package lycosa

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	LoadTask()
	for _, task := range Tasks {
		fmt.Println(task)
	}
}
