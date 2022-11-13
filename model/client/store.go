package client

type ClientStore interface {
	GetAll() []Client
	GetClientByID(id string) (Client, error)
	Add(client Client) error
	Delete(id string) error
}
