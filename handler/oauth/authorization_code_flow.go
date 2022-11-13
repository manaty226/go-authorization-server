package oauth

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/manaty226/go-authorization-server/model/session"
)

type AuthorizationCodeFlow struct {
	SessionStore session.SessionStore
}

func (acf *AuthorizationCodeFlow) FlowHandler(c *gin.Context) {
	acf.SessionStore.SetStore(sessions.Default(c))
	sessionID, err := c.Cookie("az-session-id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "session timeout"})
	}
	var authzSession session.AuthzSession
	if err := acf.SessionStore.GetSessionByID(sessionID, &authzSession); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not found session"})
		return
	}

	redirectUri := authzSession.RedirectURI
	state := authzSession.State
	location := fmt.Sprintf("%s?code=%s&state=%s", redirectUri, uuid.New().String(), state)

	c.Redirect(http.StatusMovedPermanently, location)
	c.Abort()
}
