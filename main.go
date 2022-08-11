package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manaty226/go-authorization-server/endpoint/authorization"
)

func main() {
	engine := gin.Default()
	ae := &authorization.AuthorizationEndpoint{}

	engine.GET("/", ae.Handle)
	_ = engine.Run(":3000")
}
