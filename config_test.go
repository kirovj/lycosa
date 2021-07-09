package lycosa

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	init()
	fmt.Println(Conf.User, Conf.Pass)
}
