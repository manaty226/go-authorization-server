package client

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type ClientConfig struct {
	SessionMax          time.Time
	AccessTokenLifespan time.Time
}

type Client struct {
	ClientID     string
	ClientName   string
	ClientSecret string
	Scopes       string
	RedirectURI  []*url.URL
	Config       ClientConfig
}

func New(name, redirectUri string) (Client, error) {
	id := uuid.New().String()
	secret := uuid.New().String()
	uri, err := url.ParseRequestURI(redirectUri)
	if err != nil {
		return Client{}, err
	}
	return Client{
		ClientID:     id,
		ClientName:   name,
		ClientSecret: secret,
		Scopes:       "profile",
		RedirectURI:  []*url.URL{uri},
	}, nil
}

func (c *Client) AddRedirectUri(redirectUri string) error {
	uri, err := url.ParseRequestURI(redirectUri)
	if err != nil {
		return err
	}
	c.RedirectURI = append(c.RedirectURI, uri)
	return nil
}
