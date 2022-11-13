package store

import (
	"encoding/json"

	"github.com/koron/go-dproxy"
	az_session "github.com/manaty226/go-authorization-server/model/session"
)

type SessionStoreInMemory struct {
	Store az_session.Repository
}

func (s *SessionStoreInMemory) SetStore(store az_session.Repository) {
	s.Store = store
}

func (s SessionStoreInMemory) GetSessionByID(id string, model az_session.Session) error {
	session, err := dproxy.New(s.Store.Get(id)).String()
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(session), &model); err != nil {
		return err
	}
	return nil
}

func (s SessionStoreInMemory) Add(id string, session az_session.Session) error {
	str, err := session.String()
	if err != nil {
		return err
	}
	s.Store.Set(id, str)
	if err := s.Store.Save(); err != nil {
		return err
	}
	return nil
}

func (s SessionStoreInMemory) Delete(id string) error {
	return nil
}
