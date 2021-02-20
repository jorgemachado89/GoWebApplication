package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	id int
	name string
	password string	
}

func (u User) GetUsername() string {
	return u.name
}

func Login(user, password string) (*User, error) {
	result := &User{};
	row := db.QueryRow (`
		SELECT id, name
		FROM public.user
		WHERE name = $1 AND password = $2
	`, user, password)
	err := row.Scan(&result.id, &result.name)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}
	return result, nil
}