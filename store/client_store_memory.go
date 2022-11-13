package store

import (
	"fmt"

	"github.com/manaty226/go-authorization-server/model/client"
)

type ClientStoreInMemory struct {
	ClientList map[string]client.Client
}

func CreateClientStore() *ClientStoreInMemory {
	return &ClientStoreInMemory{
		ClientList: map[string]client.Client{},
	}
}

func (cs ClientStoreInMemory) GetAll() []client.Client {
	clients := []client.Client{}
	for _, c := range cs.ClientList {
		clients = append(clients, c)
	}
	return clients
}

func (cs *ClientStoreInMemory) GetClientByID(id string) (client.Client, error) {
	c, exist := cs.ClientList[id]
	if !exist {
		return client.Client{}, fmt.Errorf("client not found")
	}
	return c, nil
}
func (cs *ClientStoreInMemory) Add(client client.Client) error {
	_, exist := cs.ClientList[client.ClientID]
	if exist {
		return fmt.Errorf("same client id is already exists.")
	}
	cs.ClientList[client.ClientID] = client
	return nil
}

func (cs *ClientStoreInMemory) Delete(id string) error {
	delete(cs.ClientList, id)
	return nil
}
