package channels

import (
	"github.com/spurtcms/auth"
	"gorm.io/gorm"
)

type Config struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Auth             *auth.Auth
}

type Channel struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Auth             *auth.Auth
}
