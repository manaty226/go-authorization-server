package token

import (
	"io/ioutil"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type AccessToken struct {
	AccessTokenID string    `json:"jti"`
	UserID        string    `json:"user_id"`
	ClientID      string    `json:"client_id"`
	Scopes        string    `json:"scopes"`
	IssuedAt      time.Time `json:"iat"`
	ExpiredAt     time.Time `json:"exp"`
}

func (at AccessToken) GenerateJWT() ([]byte, error) {
	token, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/manaty226/go-authorization-server`).
		Subject(at.ClientID).
		IssuedAt(at.IssuedAt).
		Audience([]string{at.ClientID}).
		Expiration(at.ExpiredAt).
		Claim("user_name", "test").
		Claim("scopes", at.Scopes).
		Build()
	if err != nil {
		return nil, err
	}

	bytes, _ := ioutil.ReadFile(".cert/secret.pem")
	key, _ := jwk.ParseKey(bytes, jwk.WithPEM(true))

	signed, _ := jwt.Sign(token, jwt.WithKey(jwa.RS256, key))

	return signed, nil
}
