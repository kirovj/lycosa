package lycosa

import "fmt"

type User struct {
	Id      int    `json:"id"`
	Ctime   int64  `json:"ctime"`
	Mtime   int64  `json:"mtime"`
	Valid   bool   `json:"valid"`
	Name    string `json:"name"`
	Pass    string
	Session string
}

func getUsers() ([]*User, error) {
	rows, err := db.Query(`select * from user;`)
	if err != nil {
		return nil, err
	}

	var users []*User

	for rows.Next() {
		var user *User
		if err = rows.Scan(user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func getUser(name string) *User {
	db.QueryRow(fmt.Sprintf(`select * from user where name=%s`, name))
	return nil
}
