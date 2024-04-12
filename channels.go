package channels

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spurtcms/categories"
	permission "github.com/spurtcms/team-roles"
)

// Channelsetup used to initialie channel configuration
func ChannelSetup(config Config) *Channel {

	MigrateTables(config.DB)

	return &Channel{
		DB:               config.DB,
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
		Permissions:      config.Permissions,
	}

}

// get all channel list
func (channel *Channel) ListChannel(limit, offset int, filter Filter, activestatus bool, entriescount bool) (channelList []tblchannel, channelcount int, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []tblchannel{}, 0, autherr
	}

	channellist, _, _ := CH.Channellist(limit, offset, filter, activestatus, channel.DB)

	var chnallist []tblchannel

	for _, val := range channellist {

		val.SlugName = val.ChannelDescription

		val.ChannelDescription = TruncateDescription(val.ChannelDescription, 130)

		if entriescount {

			_, entrcount, _ := EntryModel.ChannelEntryList(0, 0, val.Id, EntriesFilter{}, false, channel.Permissions.RoleId, false, channel.DB)

			val.EntriesCount = int(entrcount)
		}

		chnallist = append(chnallist, val)

	}

	_, chcount, _ := CH.Channellist(0, 0, filter, activestatus, channel.DB)

	return chnallist, int(chcount), nil
}

/*Get channel by name*/
func (channel *Channel) GetchannelByName(channelname string) (channels tblchannel, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return tblchannel{}, autherr
	}

	channellist, err1 := CH.GetChannelByChannelName(channelname, channel.DB)

	if err1 != nil {

		return tblchannel{}, err1
	}

	return channellist, nil

}

/*Get Channels By Id*/
func (channel *Channel) GetChannelsById(channelid int) (channelList tblchannel, SelectedCategories []categories.Arrangecategories, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return tblchannel{}, []categories.Arrangecategories{}, autherr
	}

	channellist, err := CH.GetChannelById(channelid, channel.DB)

	if err != nil {

		return tblchannel{}, []categories.Arrangecategories{}, err

	}

	GetSelectedChannelCateogry, err1 := CH.GetSelectedCategoryChannelById(channelid, channel.DB)

	if err1 != nil {

		log.Println(err)
	}

	var FinalSelectedCategories []categories.Arrangecategories

	for _, val := range GetSelectedChannelCateogry {

		var id []int

		ids := strings.Split(val.CategoryId, ",")

		for _, cid := range ids {

			convid, _ := strconv.Atoi(cid)

			id = append(id, convid)
		}

		GetSelectedCategory, _ := CH.GetCategoriseById(id, channel.DB)

		var addcat categories.Arrangecategories

		var individualid []categories.CatgoriesOrd

		for _, CategoriesArrange := range GetSelectedCategory {

			var individual categories.CatgoriesOrd

			individual.Id = CategoriesArrange.Id

			individual.Category = CategoriesArrange.CategoryName

			individualid = append(individualid, individual)

		}

		addcat.Categories = individualid

		FinalSelectedCategories = append(FinalSelectedCategories, addcat)

	}

	return channellist, FinalSelectedCategories, nil
}

/*get channel fields by channel id*/
func (channel *Channel) GetChannelsFieldsById(channelid int) (section []Section, fields []Fiedlvalue, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []Section{}, []Fiedlvalue{}, autherr
	}

	groupfield, _ := CH.GetFieldIdByGroupId(channelid, channel.DB)

	var ids []int

	for _, val := range groupfield {

		ids = append(ids, val.FieldId)
	}

	fieldValue, _ := CH.GetFieldAndOptionValue(ids, channel.DB)

	var sections []Section

	var Fieldvalue []Fiedlvalue

	for _, val := range fieldValue {

		var section Section

		var fieldvalue Fiedlvalue

		if val.FieldTypeId == 12 {

			section.SectionId = val.Id

			section.SectionNewId = 0

			section.MasterFieldId = val.FieldTypeId

			section.SectionName = val.FieldName

			sections = append(sections, section)

		} else {

			var optionva []OptionValues

			for _, optionval := range val.TblFieldOption {

				var optiovalue OptionValues

				optiovalue.Id = optionval.Id

				optiovalue.FieldId = optionval.FieldId

				optiovalue.Value = optionval.OptionValue

				optionva = append(optionva, optiovalue)
			}

			fieldvalue.FieldId = val.Id

			fieldvalue.FieldName = val.FieldName

			fieldvalue.CharacterAllowed = val.CharacterAllowed

			fieldvalue.DateFormat = val.DatetimeFormat

			fieldvalue.TimeFormat = val.TimeFormat

			fieldvalue.IconPath = val.ImagePath

			fieldvalue.MasterFieldId = val.FieldTypeId

			fieldvalue.Mandatory = val.MandatoryField

			fieldvalue.OrderIndex = val.OrderIndex

			fieldvalue.SectionId = val.SectionParentId

			fieldvalue.OptionValue = optionva

			Fieldvalue = append(Fieldvalue, fieldvalue)

		}

	}

	return sections, Fieldvalue, nil
}

/*Delete Channel*/
func (channel *Channel) DeleteChannel(channelid, modifiedby int) error {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return autherr
	}

	if channelid <= 0 {

		return ErrorChannelId
	}

	CH.DeleteEntryByChannelId(channelid, channel.DB)

	CH.DeleteChannelById(channelid, channel.DB)

	chdel, _ := CH.GetChannelById(channelid, channel.DB)

	var delfidgrp TblFieldGroup

	delfidgrp.IsDeleted = 1

	delfidgrp.DeletedBy = modifiedby

	delfidgrp.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	CH.DeleteFieldGroupById(&delfidgrp, chdel.FieldGroupId, channel.DB)

	return nil

}

func (channel *Channel) DeleteChannelPermissions(channelid int) error {

	checkid, _ := permission.AS.GetIdByRouteName(strconv.Itoa(channelid), channel.DB)

	permission.AS.Deleterolepermission(checkid.Id, channel.DB)

	// permission.AS.DeleteModulePermissioninEntries(channelid, channel.DB)

	return nil
}

/*Change Channel status*/
// status 0 = inactive
// status 1 = active
func (channel *Channel) ChangeChannelStatus(channelid int, status, modifiedby int) (bool, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return false, autherr
	}

	if channelid <= 0 {

		return false, ErrorChannelId
	}

	var channelstatus TblChannel

	channelstatus.ModifiedBy = modifiedby

	channelstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	CH.ChannelIsActive(&channelstatus, channelid, status, channel.DB)

	return true, nil

}
