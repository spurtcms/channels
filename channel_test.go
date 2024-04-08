package channels

import (
	"fmt"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/gorm"
)

var SecretKey = "Secret123"

func TestChannelList(t *testing.T) {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
	})

	token, _ := Auth.CreateToken()

	channel := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
	})

	chanlist, count, err := channel.ListChannel(10, 0, Filter{}, true)

	if err != nil {

		panic(err)
	}

	fmt.Println(chanlist, count)
}
