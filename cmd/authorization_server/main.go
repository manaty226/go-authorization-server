package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/manaty226/go-authorization-server/admin"
	"github.com/manaty226/go-authorization-server/handler/authenticate"
	"github.com/manaty226/go-authorization-server/handler/oauth"
	"github.com/manaty226/go-authorization-server/model/client"
	"github.com/manaty226/go-authorization-server/store"
)

func main() {
	engine := gin.Default()
	engine.Static("/assets", "./web/theme")
	engine.LoadHTMLGlob("web/theme/*.html")

	userStore := store.CreateUserStore()
	clientStore := store.CreateClientStore()
	testClient, _ := client.New("test", "https://oauthdebugger.com/debug")
	if err := (*clientStore).Add(testClient); err != nil {
		panic(err)
	}

	sessionStore := &store.SessionStoreInMemory{}

	ae := &oauth.AuthorizationEndpoint{
		ClientStore:  clientStore,
		SessionStore: sessionStore,
	}
	se := &authenticate.SignInEndpoint{
		UserStore:    userStore,
		SessionStore: sessionStore,
	}
	tk := &oauth.TokenEndPoint{}

	ac := &admin.AdminClientHandler{
		ClientStore: clientStore,
	}

	au := &admin.AdminUserHandler{
		UserStore: userStore,
	}

	store := memstore.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("authz", store))

	engine.GET("/authorize", ae.Handle)
	engine.POST("/signin", se.Handle)
	engine.POST("/token", tk.Handle)
	engine.POST("/api/admin/clients", ac.HandlePost)
	engine.GET("/api/admin/clients", ac.HandleGet)
	engine.DELETE("/api/admin/clients/:id", ac.HandleDelete)
	engine.POST("/api/admin/users", au.HandlePost)
	engine.GET("/api/admin/users", au.HandleGet)
	engine.DELETE("/api/admin/users/:id", au.HandleDelete)
	_ = engine.Run(":3000")
}
