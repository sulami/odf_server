package auth

import "fmt"

type User struct {
	Username string
	Password string
}

func (u User) String() string {
	return fmt.Sprintf("User: %s", u.Username)
}

