package channels

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func TestChannelEntriesList(t *testing.T) {

	chann := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
	})

	entries, fitlercount, totolcount, err := chann.ChannelEntriesList(Entries{Limit: 10, Offset: 0, Publishedonly: true})

	fmt.Println(entries, fitlercount, totolcount, err)

}

func TestEntryDetailsById(t *testing.T) {

	chann := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
	})

	entry, err := chann.EntryDetailsById(IndivEntriesReq{EntryId: 1, MemberProfile: true})

	fmt.Println(entry, err)
}

func TestCreateEntry(t *testing.T) {

	chann := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
	})

	entry, flg, err := chann.CreateEntry(EntriesRequired{Title: "default", Content: "default"})

	fmt.Println(entry, flg, err)
}

func TestCreateChannelEntryFields(t *testing.T) {

	chann := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
	})

	err := chann.CreateChannelEntryFields(1, 2, []AdditionalFields{
		{Id: 1, FieldName: "title", FieldValue: "demouser"},
	})

	fmt.Println(err)

}

func TestUpdateEntry(t *testing.T) {

	chann := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
	})

	flg, err := chann.UpdateEntry(EntriesRequired{Title: "default2", Content: "default2"}, "default", 1)

	fmt.Println(flg, err)
}

func TestUpdateAdditionalField(t *testing.T) {

	chann := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
	})

	flg, err := chann.UpdateAdditionalField([]AdditionalFields{{Id: 1, FieldName: "title", FieldValue: "sam"}}, 1)

	fmt.Println(flg, err)
}

func TestDeleteEntry(t *testing.T) {

	chann := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
	})

	err := chann.DeleteChannel(1, 2)

	fmt.Println(err)
}

func TestMakeFeatureEntry(t *testing.T) {

	chann := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
	})

	flg, err := chann.MakeFeatureEntry(1, 1, 1)

	fmt.Println(flg, err)

}

func TestEntryStatus(t *testing.T) {

	chann := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
	})

	flg, err := chann.EntryStatus("default", 1, 1, 2)

	fmt.Println(flg, err)
}
