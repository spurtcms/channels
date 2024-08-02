package channels

import (
	"sort"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// return all parentid to child ids aray
func Categories(categoryid int, DB *gorm.DB, tenantid int) string {

	var id string

	categoreis, _ := EntryModel.GetChildCategories(categoryid, DB,tenantid)

	for _, val := range categoreis {

		id += strconv.Itoa(val.Id) + ","
	}

	return id
}

// using name
func CategoriesByUsingName(categoryname string, DB *gorm.DB,tenantid int) string {

	var id string

	categoreid, _ := EntryModel.GetCategoryIdByName(categoryname, DB,tenantid)

	categoreis, _ := EntryModel.GetChildCategories(categoreid.Id, DB,tenantid)

	for _, val := range categoreis {

		id += strconv.Itoa(val.Id) + ","
	}

	return id
}

// DashboardEntry count function
func (channel *Channel) DashboardEntriesCount(tenantid int) (totalcount int, lasttendayscount int, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return 0, 0, autherr
	}

	allentrycount, err := EntryModel.AllentryCount(channel.DB,tenantid)

	if err != nil {

		return 0, 0, err
	}

	entrycount, err := EntryModel.NewentryCount(channel.DB,tenantid)

	if err != nil {

		return 0, 0, err
	}

	return int(allentrycount), int(entrycount), nil
}

func (channel *Channel) DashboardChannellist(tenantid int) (channelList []Tblchannel, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []Tblchannel{}, autherr
	}

	Newchannels, err := EntryModel.Newchannels(channel.DB,tenantid)

	if err != nil {

		return []Tblchannel{}, err

	}

	return Newchannels, nil

}

/*DashboardEntries */
func (channel *Channel) DashboardEntrieslist(tenantid int) (entries []Tblchannelentries, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []Tblchannelentries{}, autherr
	}

	Newentries, err := EntryModel.Newentries(channel.DB,tenantid)

	if err != nil {

		return []Tblchannelentries{}, err

	}

	return Newentries, nil

}

/*Recent activites for dashboard*/
func (channel *Channel) DashboardRecentActivites(tenantid int) (entries []RecentActivities, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []RecentActivities{}, autherr
	}

	Newentries, _ := EntryModel.Newentries(channel.DB,tenantid)

	var Newrecords []RecentActivities

	for _, val := range Newentries {

		newrecord := RecentActivities{Contenttype: "entry", Title: val.Title, User: val.Username, Imagepath: val.ProfileImagePath, Createdon: val.CreatedOn, Channelname: val.ChannelName}

		Newrecords = append(Newrecords, newrecord)
	}

	Newchannel, _ := EntryModel.Newchannels(channel.DB,tenantid)

	for _, val := range Newchannel {

		newrecord := RecentActivities{Contenttype: "channel", Title: val.ChannelName, User: val.AuthorDetails.Username, Imagepath: val.ProfileImagePath, Createdon: val.CreatedOn, Channelname: val.ChannelName}

		Newrecords = append(Newrecords, newrecord)
	}
	sort.Slice(Newrecords, func(i, j int) bool {

		return Newrecords[i].Createdon.After(Newrecords[j].Createdon)

	})
	maxRec := 5

	if len(Newrecords) < maxRec {

		maxRec = len(Newrecords)

	}
	recentActive := Newrecords[:maxRec]

	var newactive RecentActivities

	var NewActive []RecentActivities

	for _, val := range recentActive {

		difference := time.Now().Sub(val.Createdon)

		hour := int(difference.Hours())

		min := int(difference.Minutes())

		if hour >= 1 {

			newactive.Contenttype = val.Contenttype

			newactive.Title = val.Title

			newactive.User = val.User

			newactive.Imagepath = val.Imagepath

			newactive.Createdon = val.Createdon

			newactive.Channelname = val.Channelname

			newactive.Active = strconv.Itoa(hour) + " " + "hrs"
		} else {
			newactive.Contenttype = val.Contenttype

			newactive.Title = val.Title

			newactive.User = val.User

			newactive.Imagepath = val.Imagepath

			newactive.Createdon = val.Createdon

			newactive.Channelname = val.Channelname

			newactive.Active = strconv.Itoa(min) + " " + "mins"

		}

		NewActive = append(NewActive, newactive)

	}

	return NewActive, nil
}

/*Remove entries cover image if media image delete*/
func (channel *Channel) RemoveEntriesCoverImage(ImagePath string,tenantid int) error {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return autherr
	}

	err := EntryModel.UpdateImagePath(ImagePath, channel.DB,tenantid)

	if err != nil {

		return err
	}

	return nil

}

func (channel *Channel) GetPermissionChannel(tenantid int) (channels []Tblchannel, errr error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []Tblchannel{}, autherr
	}

	channelss, err := CH.GetPermissionChannel(channel, channel.DB,tenantid)

	var chnallist []Tblchannel

	for _, val := range channelss {

		val.SlugName = val.ChannelDescription
		val.ChannelDescription = TruncateDescription(val.ChannelDescription, 130)

		entriescount := true

		if entriescount {
			_, entrcount, _ := EntryModel.ChannelEntryList(Entries{ChannelId: val.Id}, channel, Empty, true, channel.DB,tenantid)
			val.EntriesCount = int(entrcount)
		}

		chnallist = append(chnallist, val)

	}

	if err != nil {

		return []Tblchannel{}, err
	}

	return chnallist, nil

}
