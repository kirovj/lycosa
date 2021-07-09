package lycosa

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	Init()
	fmt.Println(Conf.User, Conf.Pass)
}
