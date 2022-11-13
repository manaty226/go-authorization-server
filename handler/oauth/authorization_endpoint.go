package oauth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/manaty226/go-authorization-server/handler/response"
	"github.com/manaty226/go-authorization-server/model/client"
	"github.com/manaty226/go-authorization-server/model/session"
)

type AuthorizationEndpoint struct {
	Validator    *validator.Validate
	ClientStore  client.ClientStore
	SessionStore session.SessionStore
}

type AuthzRequest struct {
	ResponseType string `form:"response_type"`
	ClientID     string `form:"client_id"`
	RedirectURI  string `form:"redirect_uri"`
	Scope        string `form:"scope"`
	State        string `form:"state"`
}

func (ae *AuthorizationEndpoint) Handle(c *gin.Context) {
	ae.SessionStore.SetStore(sessions.Default(c))

	// Check if required parameters exist.
	req, err := parseRequest(c)
	if err != nil {
		response.ErrorResponse(c, response.SERVER_ERROR, "", "")
		return
	}

	// Check if valid client id
	client, err := ae.ClientStore.GetClientByID(req.ClientID)
	if err != nil {
		response.ErrorResponse(c, response.INVALID_REQUEST, req.State, "")
		return
	}

	// Check if valid redirect uri
	redirectUri, err := validateRedirectUri(req.RedirectURI, client)
	if err != nil {
		response.ErrorResponse(c, response.INVALID_REQUEST, req.State, "")
	}

	// Check if valid response type
	if ok := validateResponseType(req.ResponseType); !ok {
		response.ErrorResponse(c, response.INVALID_REQUEST, req.State, redirectUri)
		return
	}

	// filtering supported scope
	scope := filterScope(req.Scope, client)

	// All check passed and continue flow
	sessionID := uuid.New().String()
	c.SetCookie("az-session-id", sessionID, 5*60, "/", "", true, true)
	authzSession := session.NewSession(
		req.ResponseType,
		client.ClientID,
		redirectUri,
		scope,
		req.State,
	)
	if err := ae.SessionStore.Add(sessionID, authzSession); err != nil {
		response.ErrorResponse(c, response.SERVER_ERROR, req.State, redirectUri)
	}
	c.HTML(http.StatusOK, "signin.html", nil)
}

func parseRequest(c *gin.Context) (AuthzRequest, error) {
	var authzRequest AuthzRequest
	if err := c.Bind(&authzRequest); err != nil {
		return AuthzRequest{}, err
	}
	return authzRequest, nil
}

func validateResponseType(responseType string) bool {
	switch responseType {
	case "code":
		return true
	case "token":
		return true
	}
	return false
}

func validateRedirectUri(redirectUri string, client client.Client) (string, error) {
	// If redirect uri is not registered on client, use uri of request params.
	if len(client.RedirectURI) == 0 {
		return redirectUri, nil
	}

	for _, uri := range client.RedirectURI {
		if uri.String() == redirectUri {
			return redirectUri, nil
		}
	}
	return "", fmt.Errorf("not validate uri")
}

func filterScope(scope string, client client.Client) string {
	scopeArray := strings.Split(scope, " ")
	filteredScope := ""
	for _, s := range scopeArray {
		if strings.Contains(client.Scopes, s) {
			filteredScope += s + " "
		}
	}
	return filteredScope[:len(filteredScope)-1]
}
