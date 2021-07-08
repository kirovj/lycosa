package lycosa

import (
	"fmt"
	"os/exec"
)

const BASH_PATH = "C:\\Program Files\\Git\\git-bash.exe"

func RunBash() {
	var (
		cmd *exec.Cmd
		out []byte
		err error
	)
	cmd = exec.Command(BASH_PATH, "-c", "sleep 5;ls -l")

	if out, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
