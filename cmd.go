package lycosa

import (
	"fmt"
	"os/exec"
	"runtime"
)

// RunBash run bash cmd
func RunBash(command string) {
	var (
		cmd  *exec.Cmd
		out  []byte
		err  error
		bash string
	)

	switch runtime.GOOS {
	case "windows":
		bash = BashWin
	case "linux":
		bash = BashLinux
	}

	cmd = exec.Command(bash, "-c", command)
	if out, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
