package session

type Repository interface {
	Get(key interface{}) interface{}
	Set(key interface{}, val interface{})
	Save() error
}

type SessionStore interface {
	SetStore(store Repository)
	GetSessionByID(id string, model Session) error
	Add(id string, session Session) error
	Delete(id string) error
}
