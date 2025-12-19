package channels

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SecretKey = "Secret123"

// Db connection
func DBSetup() (*gorm.DB, error) {

	dbConfig := map[string]string{
		"username": "postgres",
		"password": "*****",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "spurtcms",
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=" + dbConfig["username"] + " password=" + dbConfig["password"] +
			" dbname=" + dbConfig["dbname"] + " host=" + dbConfig["host"] +
			" port=" + dbConfig["port"] + " sslmode=disable TimeZone=Asia/Kolkata",
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err != nil {
		return nil, err
	}

	return db, nil
}

// test list channel
func TestChannelList(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		chanlist, count, err := channel.ListChannel(Channels{Limit: 10, Offset: 0, TenantId: TenantId})

		if err != nil {

			panic(err)
		}

		fmt.Println(chanlist, count)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test create channel
func TestCreateChennal(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		chanlist, err := channel.CreateChannel(ChannelCreate{ChannelName: "life style", ChannelDescription: "collections", CategoryIds: []string{"1", "2"}}, 9, "1")

		if err != nil {

			panic(err)
		}

		fmt.Println(chanlist, err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test create additionalfields
func TestCreateAdditionalFields(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		field1 := Section{SectionId: 1, SectionName: "date", MasterFieldId: 2, OrderIndex: 1}
		field2 := Section{SectionId: 2, SectionName: "date&time", OrderIndex: 2}

		optval1 := OptionValues{Value: "1"}
		optval2 := OptionValues{Value: "2"}

		field1value := Fiedlvalue{Url: "/defaultchannel", OrderIndex: 1, OptionValue: []OptionValues{optval1, optval2}}
		field2value := Fiedlvalue{Url: "/defaultchannel", OrderIndex: 2}

		err := channel.CreateAdditionalFields(ChannelAddtionalField{Sections: []Section{field1, field2}, FieldValues: []Fiedlvalue{field1value, field2value}}, 5, "1")

		if err != nil {

			panic(err)
		}

		fmt.Println(err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test getchannelbyname
func TestGetchannelByName(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		channel, err := channel.GetchannelByName("life style", TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(channel, err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test getchannelbyid
func TestGetChannelsById(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		channel, chn_categories, err := channel.GetChannelsById(3, TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(channel, chn_categories, err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test getchannelfieldsbyid
func TestGetChannelsFieldsById(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		section, fields, err := channel.GetChannelsFieldsById(3, TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(section, fields, err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test deletechannel
func TestDeleteChannel(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		err := channel.DeleteChannel(3, 1, "", TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test deletechannelpermissions function
func TestDeleteChannelPermissions(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		err := channel.DeleteChannelPermissions(3, TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test chanegechannelstatus function
func TestChangeChannelStatus(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		flg, err := channel.ChangeChannelStatus(3, 0, 1, TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(flg, err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test getallmasterfieldtype function
func TestGetAllMasterFieldType(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		fields, err := channel.GetAllMasterFieldType(TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(fields, err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test editchannel function
func TestEditChannel(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		err := channel.EditChannel("go", "", "", "about golang", "", "", "", 1, 3, []string{"1", "2"}, TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

// test updatechannelfileds function
func TestUpdateChannelField(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD, TenantId)

	channel := ChannelSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		field1 := Section{SectionId: 1, SectionName: "text", MasterFieldId: 2, OrderIndex: 1}
		field2 := Section{SectionId: 2, SectionName: "date&time", OrderIndex: 2}

		optval1 := OptionValues{Value: "1", FieldId: 7}
		optval2 := OptionValues{Value: "2", FieldId: 7}

		field1value := Fiedlvalue{Url: "/defaultchannel", OrderIndex: 1, OptionValue: []OptionValues{optval1, optval2}, FieldId: 7}
		field2value := Fiedlvalue{Url: "/defaultchannel", OrderIndex: 2, FieldId: 8}

		err := channel.UpdateChannelField(ChannelUpdate{Sections: []Section{field1}, FieldValues: []Fiedlvalue{field1value, field2value}, Deletesections: []Section{field2}, DeleteFields: []Fiedlvalue{field2value}, DeleteOptionsvalue: []OptionValues{optval2}, CategoryIds: []string{"1", "2"}, ModifiedBy: 1}, 3, "1")

		if err != nil {

			panic(err)
		}

		fmt.Println(err)
	} else {

		log.Println("permissions enabled not initialised")

	}
}
