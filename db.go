package lycosa

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func loadDB() {

	var (
		err    error
		fn     = ProjName + ".db"
		create bool
	)

	// comfirm is need to create tables
	_, err = os.Lstat(fn)
	create = os.IsNotExist(err)

	if db, err = sql.Open("sqlite", fn); err != nil {
		fmt.Println(err)
		panic("create or load db err!")
	}

	// if db not exists, create tables when open db.
	if create {
		db.Exec(`
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
	session CHAR(36)
);
INSERT INTO user
(ctime, mtime, valid, name, pass)
VALUES(datetime('now', 'localtime'), datetime('now', 'localtime'), 1, 'admin', 'admin');`)
	}
}
