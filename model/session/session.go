package session

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Session interface {
	CreateID()
	String() (string, error)
}

type AuthzSession struct {
	SessionID    string
	UserID       string
	ResponseType string
	ClientID     string
	RedirectURI  string
	Scope        string
	State        string
	createdAt    time.Time
	modifiedAt   time.Time
}

func NewSession(responseType, clientID, redirectURI, scope, state string) *AuthzSession {
	return &AuthzSession{
		SessionID:    uuid.New().String(),
		ResponseType: responseType,
		ClientID:     clientID,
		RedirectURI:  redirectURI,
		Scope:        scope,
		State:        state,
		createdAt:    time.Now(),
		modifiedAt:   time.Now(),
	}
}

func (as *AuthzSession) SetUserID(userID string) {
	as.UserID = userID
}

func (as *AuthzSession) CreateID() {
	as.SessionID = uuid.New().String()
}

func (as AuthzSession) String() (string, error) {
	str, err := json.Marshal(as)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

type LoginSession struct {
	SessionID string
	UserID    string
}

func (ls *LoginSession) CreateID() {
	ls.SessionID = uuid.New().String()
}

func (ls LoginSession) String() (string, error) {
	str, err := json.Marshal(ls)
	if err != nil {
		return "", err
	}
	return string(str), nil
}
