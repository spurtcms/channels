package channels

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

var SecretKey = "Secret123"

func TestChannelList(t *testing.T) {

	channel := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
	})

	chanlist, count, err := channel.ListChannel(10, 0, Filter{}, true, true)

	if err != nil {

		panic(err)
	}

	fmt.Println(chanlist, count)
}

func TestGetchannelByName(t *testing.B) {

}

func TestGetChannelById(t *testing.T) {

}
