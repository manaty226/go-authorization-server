package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       string
	UserID   string
	UserName UserName
	Password string
}

func New(userName, password string) (User, error) {
	return User{
		ID:       uuid.New().String(),
		UserID:   userName,
		Password: password,
	}, nil
}

func (u User) FullName() string {
	return u.UserName.fullName()
}

type UserName struct {
	FirstName string
	LastName  string
}

func (un UserName) fullName() string {
	return un.FirstName + un.LastName
}
