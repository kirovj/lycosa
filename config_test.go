package lycosa

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	LoadConfig()
	fmt.Println(Conf.User, Conf.Pass)
}
