package lycosa

import (
	"runtime"
)

var Bash string

const (
	BashWin   = "C:\\Program Files\\Git\\git-bash.exe"
	BashLinux = "/usr/bin/bash"
	NotFound  = "Not Found: "
	ProjName  = "lycosa"
)

const (
	CreateTables = `
CREATE TABLE task (
	id     INTEGER PRIMARY KEY,
	ctime  DATE,
	mtime  DATE,
	valid  BOOLEAN,
	name   VARCHAR(255)    NOT NULL,
	cron   CHAR(20) NOT NULL,
	cmd    TEXT
	);
CREATE TABLE user (
	id      INTEGER PRIMARY KEY,
	ctime   DATE,
	mtime   DATE,
	valid   BOOLEAN,
	name    CHAR(20) NOT NULL,
	pass    CHAR(20) NOT NULL,
	token   CHAR(36)
);
INSERT INTO user
(ctime, mtime, valid, name, pass)
VALUES(datetime('now', 'localtime'), datetime('now', 'localtime'), 1, 'admin', 'admin');`
	InsertTaskSql = `insert into task(ctime, mtime, valid, name, cron, cmd) 
values(datetime('now', 'localtime'), datetime('now', 'localtime'), 1, '%s', '%s', '%s');`
	UpdateTaskSql      = `update task set name='%s', cron='%s', cmd='%s', mtime=datetime('now', 'localtime') where id = %d;`
	UpdateTaskValidSql = `update task set valid=1, mtime=datetime('now', 'localtime') where id = %d;`
)

func init() {
	switch runtime.GOOS {
	case "windows":
		Bash = BashWin
	case "linux":
		Bash = BashLinux
	}
}
