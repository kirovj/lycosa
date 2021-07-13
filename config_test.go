package lycosa

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	loadConfig()
	fmt.Println(Conf.User, Conf.Pass)
}
