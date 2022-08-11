package authorization

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthorizationEndpoint struct {
	Validator *validator.Validate
}

type AuthzRequest struct {
	ResponseType string `form:"response_type" validate:"required"`
	ClientID     string `form:"client_id" validate:"required"`
	RedirectURI  string `form:"redirect_uri"`
	Scope        string `form:"scope"`
	State        string `form:"state"`
}

func (ae *AuthorizationEndpoint) Handle(c *gin.Context) {
	vd := validator.New()
	var authzRequest AuthzRequest
	if err := c.Bind(&authzRequest); err != nil {
		fmt.Printf("error in request bind: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "badrequest"})
		return
	}

	err := vd.Struct(authzRequest)
	if err != nil {
		fmt.Printf("error in validate: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "badrequest"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
