package lycosa

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func loadDB() {
	db, err := sql.Open("sqlite", ProjName+".db")
	if err != nil {
		fmt.Println(err)
		panic("create or load db err!")
	}

	if _, err = db.Exec(`
drop table if exists t;
create table t(i);
insert into t values(42), (314);
`); err != nil {
		return
	}
}
