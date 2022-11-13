package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manaty226/go-authorization-server/handler/request"
	"github.com/manaty226/go-authorization-server/model/user"
)

type AdminUserHandler struct {
	UserStore user.UserStore
}

func (ah AdminUserHandler) HandleGet(c *gin.Context) {
	clients := ah.UserStore.GetAll()
	c.JSON(http.StatusOK, gin.H{"clients": clients})
}

type AdminUserRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// POST /clients
func (ah AdminUserHandler) HandlePost(c *gin.Context) {
	req, err := request.ParseJSONRequest(c, AdminUserRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad_request"})
		return
	}
	fmt.Printf("request is %v \n", req)
	user, err := user.New(req.UserName, req.Password)
	if err != nil {
		fmt.Printf("error is %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if err := ah.UserStore.Add(user); err != nil {
		fmt.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error when adding user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_id": user.ID})
}

// DELETE /users/{id}
func (ah AdminUserHandler) HandleDelete(c *gin.Context) {
	id := c.Param("id")
	if err := ah.UserStore.Delete(id); err != nil {
		fmt.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error when delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
