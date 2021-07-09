package lycosa

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	Tasks []*Task
	lock  = new(sync.RWMutex)
)

type Task struct {
	Valid      bool   `json:"valid"`      // is task valid
	Name       string `json:"name"`       // task name
	Scheduling string `json:"scheduling"` // scheduling: crontab-like
	Command    string `json:"command"`    // task shell command
}

func NewTask(valid bool, name, scheduling, command string) *Task {
	return &Task{
		Valid:      valid,
		Name:       name,
		Scheduling: scheduling,
		Command:    command,
	}
}

func defaultTask(name, scheduling, command string) *Task {
	return NewTask(true, name, scheduling, command)
}

func newTaskFromBytes(bs []byte) *Task {
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
	if file, err = os.Open(Conf.TaskFile); err != nil {
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
				Tasks = append(Tasks, newTaskFromBytes(bytes))
			}
			break
		}

		if err != nil {
			fmt.Println(err)
		}
		Tasks = append(Tasks, newTaskFromBytes(bytes))
	}
}

func AddTask(name, scheduling, command string) {
	lock.Lock()
	Tasks = append(Tasks, defaultTask(name, scheduling, command))
	lock.Unlock()
}

func changeTask(name, scheduling, command string) error {
	for _, task := range Tasks {
		if task.Name == name {
			task.Scheduling = scheduling
			task.Command = command
			return nil
		}
	}
	return NotFound(name)
}

func changeTaskValid(name string) error {
	for _, task := range Tasks {
		if task.Name == name {
			task.Valid = !task.Valid
			return nil
		}
	}
	return NotFound(name)
}
