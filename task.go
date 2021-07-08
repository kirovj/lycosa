package lycosa

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	Tasks    []*Task
	TaskFile = "Task"
	lock     = new(sync.RWMutex)
)

type Task struct {
	Valid      bool   // is task valid
	Name       string // task name
	Scheduling string // scheduling: crontab-like
	Command    string // task shell command
}

func NewTask(valid bool, name, scheduling, command string) *Task {
	return &Task{
		Valid:      valid,
		Name:       name,
		Scheduling: scheduling,
		Command:    command,
	}
}

func NewDefaultTask(name, scheduling, command string) *Task {
	return NewTask(true, name, scheduling, command)
}

func NewTaskFromBytes(bs []byte) *Task {
	list := strings.Split(strings.Trim(string(bs), "\r"), "\t")
	return NewTask(list[0] == "1", list[1], list[2], list[3])
}

func (t *Task) String() string {
	return fmt.Sprintf("Task{valid: %t, name: %s, scheduling: %s, command: %s}", t.Valid, t.Name, t.Scheduling, t.Command)
}

func LoadTask() {
	var (
		file   *os.File
		err    error
		reader *bufio.Reader
	)

	lock.RLock()
	if file, err = os.Open(TaskFile); err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		file.Close()
		lock.RUnlock()
	}()

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

func AddTask(task *Task) {
	lock.Lock()
	Tasks = append(Tasks, task)
	lock.Unlock()
}

func StopTask(name string) error {
	for _, task := range Tasks {
		if task.Name == name {
			task.Valid = false
			return nil
		}
	}
	return errors.New("task not found: " + name)
}
