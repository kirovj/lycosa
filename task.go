package lycosa

import (
	"fmt"
	"os/exec"
	"sync"
)

var lock = new(sync.RWMutex)

type Task struct {
	Id    int    `json:"id"`
	Ctime int64  `json:"ctime"`
	Mtime int64  `json:"mtime"`
	Valid bool   `json:"valid"` // is task valid
	Name  string `json:"name"`  // task name
	Cron  string `json:"cron"`  // cron: crontab-like
	Cmd   string `json:"cmd"`   // task shell cmd
}

func NewTask(valid bool, name, cron, cmd string) *Task {
	return &Task{
		Valid: valid,
		Name:  name,
		Cron:  cron,
		Cmd:   cmd,
	}
}

func (t *Task) String() string {
	return fmt.Sprintf("Task{valid: %t, name: %s, cron: %s, cmd: %s}", t.Valid, t.Name, t.Cron, t.Cmd)
}

// getTasks get tasks from db
// it only run once when service start
func getTasks() ([]*Task, error) {
	rows, err := db.Query(`select * from task;`)
	if err != nil {
		return nil, err
	}

	var tasks []*Task

	for rows.Next() {
		var task *Task
		if err = rows.Scan(task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
		fmt.Println(task)
	}
	return tasks, nil
}

// AddTask add task from web api and then save to db
func addTask(name, cron, cmd string) error {
	lock.Lock()
	defer lock.Unlock()
	if _, err := db.Exec(fmt.Sprintf(InsertTaskSql, name, cron, cmd)); err != nil {
		return err
	}
	return nil
}

func updateTask(id int, name, cron, cmd string) error {
	lock.Lock()
	defer lock.Unlock()
	if _, err := db.Exec(fmt.Sprintf(UpdateTaskSql, name, cron, cmd, id)); err != nil {
		return err
	}
	return nil
}

func updateTaskValid(id int) error {
	lock.Lock()
	defer lock.Unlock()
	if _, err := db.Exec(fmt.Sprintf(UpdateTaskValidSql, id)); err != nil {
		return err
	}
	return nil
}

// RunTask run bash cmd
func (t *Task) runTask() {
	var (
		cmd *exec.Cmd
		// out []byte
		err error
	)

	cmd = exec.Command(Bash, "-c", t.Cmd)
	if _, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}
}
