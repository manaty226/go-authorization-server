package oauth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/manaty226/go-authorization-server/model/session"
	"github.com/manaty226/go-authorization-server/model/token"
)

type ImplicitFlow struct {
	SessionStore session.SessionStore
}

func (ipf *ImplicitFlow) FlowHandler(c *gin.Context) {
	ipf.SessionStore.SetStore(sessions.Default(c))

	sessionID, err := c.Cookie("az-session-id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "session timeout"})
	}

	var authzSession session.AuthzSession
	if err := ipf.SessionStore.GetSessionByID(sessionID, &authzSession); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not found session"})
		return
	}

	redirectUri := authzSession.RedirectURI
	accessToken, _ := token.AccessToken{
		AccessTokenID: uuid.New().String(),
		UserID:        uuid.New().String(),
		ClientID:      authzSession.ClientID,
		IssuedAt:      time.Now(),
		ExpiredAt:     time.Now().Add(30 * time.Minute),
		Scopes:        authzSession.Scope,
	}.GenerateJWT()
	tokenType := "bearer"
	state := authzSession.State
	location := fmt.Sprintf("%s#access_token=%s&token_type=%s&state=%s", redirectUri, accessToken, tokenType, state)

	c.Redirect(http.StatusMovedPermanently, location)
	c.Abort()
}
