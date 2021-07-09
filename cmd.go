package lycosa

import (
	"fmt"
	"os/exec"
	"runtime"
)

const (
	BASH_WIN   = "C:\\Program Files\\Git\\git-bash.exe"
	BASH_LINUX = "/usr/bin/bash"
)

func RunBash(command string) {

	var (
		cmd  *exec.Cmd
		out  []byte
		err  error
		bash string
	)

	switch runtime.GOOS {
	case "windows":
		bash = BASH_WIN
	case "linux":
		bash = BASH_LINUX
	}

	cmd = exec.Command(bash, "-c", command)

	if out, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
