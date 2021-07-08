package lycosa

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	Load()
	for _, task := range Tasks {
		fmt.Println(task)
	}
}
