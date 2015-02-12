package auth

type User struct {
	Username string
	Password string
}

func (u User) Auth(passwd string) (pass bool) {
	if string == u.Password {
		pass = true
	}
	return
}

