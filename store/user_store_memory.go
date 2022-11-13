package store

import (
	"fmt"

	"github.com/manaty226/go-authorization-server/model/user"
)

type UserStoreInMemory struct {
	UserList map[string]*user.User
}

func CreateUserStore() *UserStoreInMemory {
	return &UserStoreInMemory{
		map[string]*user.User{},
	}
}

func (ur UserStoreInMemory) GetAll() []user.User {
	users := []user.User{}
	for _, u := range ur.UserList {
		users = append(users, *u)
	}
	return users
}

func (ur *UserStoreInMemory) GetUserByID(id string) (user.User, error) {
	u, exist := ur.UserList[id]
	if !exist {
		return user.User{}, fmt.Errorf("user not found")
	}
	return *u, nil
}
func (ur *UserStoreInMemory) Add(user user.User) error {
	_, exist := ur.UserList[user.UserID]
	if exist {
		return fmt.Errorf("same user id is already exists.")
	}
	ur.UserList[user.UserID] = &user
	return nil
}

func (ur *UserStoreInMemory) UpdatePassword(id, password string) error {
	u, exist := ur.UserList[id]
	if !exist {
		return fmt.Errorf("user not found")
	}
	u.Password = password

	return nil
}

func (ur *UserStoreInMemory) Delete(id string) error {
	delete(ur.UserList, id)
	return nil
}
