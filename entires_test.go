package channels

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
)

// test channelentrieslist functon
func TestChannelEntriesList(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		entries, fitlercount, totolcount, err := channel.ChannelEntriesList(Entries{Limit: 10, Offset: 0, Publishedonly: false})

		if err != nil {

			panic(err)
		}

		fmt.Println(entries, fitlercount, totolcount, err)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// test entriesdetailsbyid function
func TestEntryDetailsById(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		entries, err := channel.EntryDetailsById(IndivEntriesReq{EntryId: 1, MemberProfile: true})

		if err != nil {

			panic(err)
		}

		fmt.Println(entries, err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

func TestCreateEntry(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		entries, flg, err := channel.CreateEntry(EntriesRequired{Title: "java", Content: "about java", ChannelId: 3})

		if err != nil {

			panic(err)
		}

		fmt.Println(entries, flg, err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test createchannelentryfields function
func TestCreateChannelEntryFields(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		err := channel.CreateChannelEntryFields(1, 2, []AdditionalFields{
			{Id: 1, FieldName: "title", FieldValue: "demouser"},
		})

		if err != nil {

			panic(err)
		}

		fmt.Println(err)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// test updateentry function
func TestUpdateEntry(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		flg, err := channel.UpdateEntry(EntriesRequired{Title: "default2", Content: "default2"}, "default", 1)

		if err != nil {

			panic(err)
		}

		fmt.Println(flg, err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

func TestUpdateAdditionalField(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		flg, err := channel.UpdateAdditionalField([]AdditionalFields{{Id: 1, FieldName: "title", FieldValue: "sam"}}, 1)

		if err != nil {

			panic(err)
		}

		fmt.Println(flg, err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test deleteentry function
func TestDeleteEntry(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		flg,err := channel.DeleteEntry("Default_Channel",1, 1)

		if err != nil {

			panic(err)
		}

		fmt.Println(flg)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test makefeatureentry function
func TestMakeFeatureEntry(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		flg, err := channel.MakeFeatureEntry(1, 1, 1)

		if err != nil {

			panic(err)
		}

		fmt.Println(flg)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// test entrystatus function
func TestEntryStatus(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		flg, err := channel.EntryStatus("default",1, 1, 1)

		if err != nil {

			panic(err)
		}

		fmt.Println(flg)
	} else {

		log.Println("permissions enabled not initialised")

	}
}