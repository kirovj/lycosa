package lycosa

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	Tasks    []*Task
	TaskFile = "Task"
	// pattern  = regexp.MustCompile("\t")
)

type Task struct {
	Name       string // task name
	Scheduling string // scheduling: crontab-like
	Command    string // task shell command
}

func NewTask(name, scheduling, command string) *Task {
	return &Task{
		Name:       name,
		Scheduling: scheduling,
		Command:    command,
	}
}

func NewTaskFromBytes(bs []byte) *Task {
	list := strings.Split(strings.Trim(string(bs), "\r"), "\t")
	return NewTask(list[0], list[1], list[2])
}

func (t *Task) String() string {
	return fmt.Sprintf("Task  name: %s, scheduling: %s, command: %s", t.Name, t.Scheduling, t.Command)
}

func Load() {
	var (
		file   *os.File
		err    error
		reader *bufio.Reader
	)

	if file, err = os.Open(TaskFile); err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader = bufio.NewReader(file)
	for {
		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			if len(bytes) > 0 {
				Tasks = append(Tasks, NewTaskFromBytes(bytes))
			}
			break
		}

		if err != nil {
			fmt.Println(err)
		}
		Tasks = append(Tasks, NewTaskFromBytes(bytes))
	}
}
