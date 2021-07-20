package lycosa

import "fmt"

type User struct {
	Id    int    `json:"id"`
	Ctime int64  `json:"ctime"`
	Mtime int64  `json:"mtime"`
	Valid bool   `json:"valid"`
	Name  string `json:"name"`
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
		var user *User
		if err = rows.Scan(user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func getUserByName(name string) (*User, error) {
	row := db.QueryRow(fmt.Sprintf(`select * from user where name=%s`, name))
	var user *User
	if err := row.Scan(user); err != nil {
		return nil, err
	}
	return user, nil
}

func getUserByToken(token string) (*User, error) {
	row := db.QueryRow(fmt.Sprintf(`select * from user where token=%s`, token))
	var user *User
	if err := row.Scan(user); err != nil {
		return nil, err
	}
	return user, nil
}

func updateToken(user *User, token string) {
	user.Token = token
	_, _ = db.Exec(fmt.Sprintf(`update user set token=%s where id=%d`, token, user.Id))
}
