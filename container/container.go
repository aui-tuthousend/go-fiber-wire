package container

import (
	"go-fiber-wire/internal/features/auth"
	"go-fiber-wire/internal/features/role"
	"go-fiber-wire/internal/features/user"
)

type AppContainer struct {
	AuthHandler *auth.AuthHandler
	RoleHandler *role.RoleHandler
	UserHandler *user.UserHandler
}
