package authenticate

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/manaty226/go-authorization-server/handler/oauth"
	"github.com/manaty226/go-authorization-server/model/session"
	"github.com/manaty226/go-authorization-server/model/user"
	"github.com/manaty226/go-authorization-server/store"
)

type SignInEndpoint struct {
	UserStore    user.UserStore
	SessionStore session.SessionStore
}

type SignInRequest struct {
	UserName string `form:"user_name" valudate:"required"`
	Password string `form:"password" validate:"required"`
}

func (s *SignInEndpoint) Handle(c *gin.Context) {
	// parse username and password from request body
	req, err := parseRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "badrequest"})
	}
	// validate request
	if ok := validateRequest(req); !ok {
		fmt.Printf("error in validate: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "badrequest"})
		return
	}

	// check user password
	// TODO: pasword should be hashed.
	u, err := s.UserStore.GetUserByID(req.UserName)
	if err != nil || u.Password != req.Password {
		fmt.Printf("invalid username or password")
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid username or password"})
		return
	}

	// get session id from cookie
	sessionID, err := c.Cookie("az-session-id")
	if err != nil {
		fmt.Printf("error in validate: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "server_error"})
		return
	}
	// get session information
	var authzSession session.AuthzSession
	if err := s.SessionStore.GetSessionByID(sessionID, &authzSession); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not found session"})
		return
	}

	switch authzSession.ResponseType {
	case "code":
		authorization_code_flow := oauth.AuthorizationCodeFlow{
			SessionStore: &store.SessionStoreInMemory{Store: s.SessionStore},
		}
		authorization_code_flow.FlowHandler(c)
	case "token":
		implicit_flow := oauth.ImplicitFlow{
			SessionStore: &store.SessionStoreInMemory{Store: s.SessionStore},
		}
		implicit_flow.FlowHandler(c)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "badrequest"})
	}
}

func parseRequest(c *gin.Context) (*SignInRequest, error) {
	var req SignInRequest
	if err := c.Bind(&req); err != nil {
		fmt.Printf("error in request bind: %v", err)
		return nil, err
	}
	return &req, nil
}

func validateRequest(req *SignInRequest) bool {
	vd := validator.New()
	if err := vd.Struct(req); err != nil {
		return false
	}
	return true
}
