package lycosa

import "testing"

func TestCmd(t *testing.T) {
	RunBash("echo a;sleep 3")
}
