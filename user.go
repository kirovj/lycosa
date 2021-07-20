package lycosa

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	Id    int       `json:"id"`
	Ctime time.Time `json:"ctime"`
	Mtime time.Time `json:"mtime"`
	Valid bool      `json:"valid"`
	Name  string    `json:"name"`
	Pass  string
	Token string
}

func getUsers() (*[]*User, error) {
	rows, err := db.Query(`select * from user;`)
	if err != nil {
		return nil, err
	}

	var users []*User

	for rows.Next() {
		var user = User{}
		if err = rows.Scan(&user.Id, &user.Ctime, &user.Mtime, &user.Valid, &user.Name, &user.Pass, &user.Token); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return &users, nil
}

func getUserByName(name string) *User {
	row := db.QueryRow(fmt.Sprintf(SelectUserByName, name))
	if user := scanUser(row); user != nil {
		return user
	}
	return nil
}

func getUserByToken(token string) *User {
	row := db.QueryRow(fmt.Sprintf(SelectUserByToken, token))
	if user := scanUser(row); user != nil {
		return user
	}
	return nil
}

func updateToken(user *User, token string) {
	user.Token = token
	_, _ = db.Exec(fmt.Sprintf(UpdateUserToken, token, user.Id))
}

func scanUser(row *sql.Row) *User {
	var user = User{}
	if err := row.Scan(&user.Id, &user.Ctime, &user.Mtime, &user.Valid, &user.Name, &user.Pass, &user.Token); err == nil {
		return &user
	}
	return nil
}
