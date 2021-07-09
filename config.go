package lycosa

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	User     string
	Pass     string
	TaskFile string
}

var Conf = &Config{}

func setField(v reflect.Value, kvStr string) {
	kv := strings.Split(kvStr, "=")
	field := v.FieldByNameFunc(func(s string) bool { return strings.ToLower(s) == kv[0] })
	field.SetString(kv[1])
}

// LoadConfig load config from file config
func LoadConfig() {
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
