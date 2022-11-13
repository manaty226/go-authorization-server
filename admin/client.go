package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manaty226/go-authorization-server/handler/request"
	"github.com/manaty226/go-authorization-server/model/client"
)

type AdminClientHandler struct {
	ClientStore client.ClientStore
}

// GET /clients
func (ap *AdminClientHandler) HandleGet(c *gin.Context) {
	clients := ap.ClientStore.GetAll()
	c.JSON(http.StatusOK, gin.H{"clients": clients})
}

// POST /clients
type AdminClientRequest struct {
	ClientName  string `json:"client_name"`
	ClientType  string `json:"client_type"`
	RedirectURI string `json:"redirect_uri"`
}

func (ap *AdminClientHandler) HandlePost(c *gin.Context) {
	req, err := request.ParseJSONRequest(c, AdminClientRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad_request"})
		return
	}
	client, err := client.New(req.ClientName, req.RedirectURI)
	if err != nil {
		fmt.Printf("error is %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if err := ap.ClientStore.Add(client); err != nil {
		fmt.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error when adding client"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"client_id": client.ClientID})
}

// DELETE /clients/{id}
func (ap *AdminClientHandler) HandleDelete(c *gin.Context) {
	id := c.Param("id")
	if err := ap.ClientStore.Delete(id); err != nil {
		fmt.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error when adding client"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
