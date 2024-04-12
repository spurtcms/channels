package channels

import (
	"github.com/spurtcms/auth"
	role "github.com/spurtcms/team-roles"
	"gorm.io/gorm"
)

type Config struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Auth             *auth.Auth
	Permissions      *role.PermissionConfig
}

type Channel struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Auth             *auth.Auth
	Permissions      *role.PermissionConfig
}
