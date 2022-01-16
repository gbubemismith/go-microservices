package domain

import (
	"fmt"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Gbubemi", LastName: "Smith", Email: "ggg@gmail.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		fmt.Println("Here 1")
		return user, nil
	}

	return nil, fmt.Errorf("user %v was not found", userId)
}
