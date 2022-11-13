package user

type UserStore interface {
	GetAll() []User
	GetUserByID(id string) (User, error)
	Add(user User) error
	UpdatePassword(id, password string) error
	Delete(id string) error
}
