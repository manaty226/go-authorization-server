package oauth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/manaty226/go-authorization-server/model/session"
	"github.com/manaty226/go-authorization-server/model/token"
)

type TokenEndPoint struct {
	Validator    *validator.Validate
	SessionStore session.SessionStore
}

type AccessTokenRequest struct {
	GrantType   string `form:"grant_type" json:"grant_type" validate:"required"`
	Code        string `form:"code" validate:"required"`
	RedirectURI string `form:"redirect_uri"`
	ClientID    string `form:"client_id"`
}

func (te *TokenEndPoint) Handle(c *gin.Context) {
	sessionID, err := c.Cookie("az-session-id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "session timeout"})
	}

	var authzSession session.AuthzSession
	if err := te.SessionStore.GetSessionByID(sessionID, &authzSession); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not found session"})
		return
	}

	accessToken, _ := token.AccessToken{
		AccessTokenID: uuid.New().String(),
		UserID:        authzSession.UserID,
		ClientID:      authzSession.ClientID,
		IssuedAt:      time.Now(),
		ExpiredAt:     time.Now().Add(30 * time.Minute),
		Scopes:        authzSession.Scope,
	}.GenerateJWT()

	c.JSON(http.StatusOK,
		gin.H{
			"access_token":  accessToken,
			"token_type":    "bearer",
			"refresh_token": uuid.New().String(),
		},
	)
}
