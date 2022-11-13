package store

import (
	"testing"

	"github.com/manaty226/go-authorization-server/model/user"
	"github.com/stretchr/testify/require"
)

func TestUserAddToMemory(t *testing.T) {
	userStore := UserStoreInMemory{UserList: map[string]*user.User{}}
	u, _ := user.New("test", "test")
	err := userStore.Add(u)
	require.NoError(t, err)

	got, _ := userStore.GetUserByID("test")
	require.Equal(t, u.UserID, got.UserID)
}
