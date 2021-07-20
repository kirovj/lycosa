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
		db.Exec(CreateTables)
	}
}
