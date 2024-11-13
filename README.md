# Channels Package

The Channels work in harmony to enhance content management in websites. Channels provide structured containers for organizing diverse data, each tailored to a specific content type. Our Channels package offers a robust solution for organizing and managing data within your Golang projects, providing a structured framework for seamless information storage and retrieval.


## Features

- Enables retrieving channels, fetching channels based on user permissions, and obtaining details about specific channels by their IDs.  
- Administrators can create, edit, and delete channels, as well as modify their status.
- Facilitates the management of channel entries by allowing administrators to retrieve all entries across channels, fetch published entries, create new entries, and delete existing ones.
- Retrieving additional field data for specific channels, fetching details about individual entries by their IDs, updating entry details, modifying entry statuses.
- Retrieval of master field types for channels and provides a list of channels for display on the admin dashboard.



# Installation

``` bash
go get github.com/spurtcms/channels
```


# Usage Example
``` bash
import(
	"github.com/spurtcms/auth"
	"github.com/spurtcms/channel"
)

func main() {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Channels", auth.CRUD)

	channel := ChannelSetup(Config{
		DB:               &gorm.DB{},
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	//channels
	if permisison {

		//list channel
		channellist, count, err := channel.ListChannel(10, 0, Filter{Keyword: ""}, true, true)
		fmt.Println(channellist, count, err)

		//create channel
		cchannel, err := channel.CreateChannel(ChannelCreate{ChannelName: "demo", ChannelDescription: "demo", CategoryIds: []string{"56,77"}, CreatedBy: 1},1,1)
		fmt.Println(cchannel, err)

		//update channel
		uerr := channel.EditChannel("demo2", "demo2", 2, 1, []string{"55,44"},1)
		fmt.Println(uerr)

		//delete channel
		derr := channel.DeleteChannel(1, 1,"",1)
		fmt.Println(derr)

		//create additionfield
		field1value := Fiedlvalue{FieldName: "text", MasterFieldId: 2, OrderIndex: 1, IconPath: "/public/img/text.svg"}
		field2value := Fiedlvalue{FieldName: "date&time", MasterFieldId: 4, OrderIndex: 2, IconPath: "/public/img/date-time.svg",DateFormat: "DD/MM/YYYY",TimeFormat: "12"}

		err := channel.CreateAdditionalFields(ChannelAddtionalField{FieldValues: []Fiedlvalue{field1value, field2value}, CreatedBy: 1}, 10, 1)

		if err != nil {

			panic(err)
		}

	}

	cpermisison, _ := Auth.IsGranted("Entries", auth.CRUD)

	if cpermisison {

		//channelentries list
		entries, filtercount, overallcount, err := channel.ChannelEntriesList(Entries{
			ChannelId:                1,
			Limit:                    10,
			Offset:                   0,
			Keyword:                  "",
			ChannelName:              "",
			SelectedCategoryFilter:   true,
			Publishedonly:            true,
			ActiveChannelEntriesonly: true,
			MemberProfile:            true,
		},1)

		fmt.Println(entries, filtercount, overallcount, err)

		//create entry
		entries, flg, err := channel.CreateEntry(EntriesRequired{Title: "golang", Content: "about go", ChannelId: 1,CreatedBy: 1}, 1)

		if err != nil {

			panic(err)
		}

		fmt.Println(entry, flg, err)

		//update entry
		channel.UpdateEntry(EntriesRequired{
			Title:      "demo2",
			Content:    "demo2",
			CoverImage: "",
		}, "demo", 1,1)

		//delete entry
		flg, derr := channel.DeleteEntry("demo", 1, 1,1)

		if derr != nil {

			fmt.Println(derr)
		}

	}
}

```




# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/channels/issues]. 
