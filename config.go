package lycosa

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
)

const (
	BashWin   = "C:\\Program Files\\Git\\git-bash.exe"
	BashLinux = "/usr/bin/bash"
	NotFound  = "Not Found: "
	ProjName  = "lycosa"
)

type Config struct {
	User     string
	Pass     string
	TaskFile string
}

var (
	Conf = &Config{}
	Bash string
)

func setField(v reflect.Value, kvStr string) {
	kv := strings.Split(kvStr, "=")
	field := v.FieldByNameFunc(func(s string) bool { return strings.ToLower(s) == kv[0] })
	field.SetString(kv[1])
}

// LoadConfig load config from file config
func loadConfig() {
	var (
		file   *os.File
		line   []byte
		err    error
		reader *bufio.Reader
		v      = reflect.ValueOf(Conf).Elem() // 反射获取Conf的指针
	)

	if file, err = os.Open("config"); err != nil {
		fmt.Println(err)
		return
	}

	reader = bufio.NewReader(file)
	defer file.Close()
	for {
		line, _, err = reader.ReadLine()
		if err == io.EOF {
			if len(line) > 0 {
				setField(v, string(line))
			}
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		setField(v, string(line))
	}
}

func confirmBash() {
	switch runtime.GOOS {
	case "windows":
		Bash = BashWin
	case "linux":
		Bash = BashLinux
	}
}
