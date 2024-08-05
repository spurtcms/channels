package channels

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spurtcms/categories"
	"github.com/spurtcms/channels/migration"
	permission "github.com/spurtcms/team-roles"
	"gorm.io/gorm"
)

// Channelsetup used to initialie channel configuration
func ChannelSetup(config Config) *Channel {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &Channel{
		DB:               config.DB,
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
	}

}

// get all channel list
func (channel *Channel) ListChannel(inputs Channels) (channelList []Tblchannel, channelcount int, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {
		return []Tblchannel{}, 0, autherr
	}

	CH.Userid = channel.Userid
	CH.Dataaccess = channel.DataAccess

	var (
		channellist  []Tblchannel
		count        int64
	)

	err = CH.Channellist(channel.DB,inputs,&channellist,&count)

	if err != nil {
		return []Tblchannel{}, 0, err
	}

	return channellist, int(count), nil
}

func (channel *Channel) ChannelDetail(inputs Channels)(channelDetails Tblchannel,err error){

	autherr := AuthandPermission(channel)

	if autherr != nil {
		return  Tblchannel{}, autherr
	}

	CH.Userid = channel.Userid
	CH.Dataaccess = channel.DataAccess

	if err = CH.ChannelDetail(channel.DB,inputs,&channelDetails);err!=nil{

		return Tblchannel{},err
	}

	return channelDetails, nil
}


/*create channel*/
func (channel *Channel) CreateChannel(channelcreate ChannelCreate, tenantid int) (TblChannel, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return TblChannel{}, autherr
	}

	/*create channel*/
	var cchannel TblChannel
	cchannel.ChannelName = channelcreate.ChannelName
	cchannel.ChannelDescription = channelcreate.ChannelDescription
	cchannel.SlugName = strings.ToLower(strings.ReplaceAll(channelcreate.ChannelName, " ", "-"))
	cchannel.IsActive = 1
	cchannel.CreatedBy = channelcreate.CreatedBy
	cchannel.TenantId = tenantid
	cchannel.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	ch, chanerr := CH.CreateChannel(&cchannel, channel.DB)

	if chanerr != nil {

		fmt.Println(chanerr)
	}

	/*This is for module permission creation*/
	var modperms permission.TblModulePermission
	modperms.DisplayName = ch.ChannelName
	modperms.RouteName = "/channel/entrylist/" + strconv.Itoa(ch.Id)
	modperms.SlugName = strings.ReplaceAll(strings.ToLower(ch.ChannelName), " ", "_")
	modperms.CreatedBy = channelcreate.CreatedBy
	modperms.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	modperms.ModuleId = 8
	modperms.AssignPermission = 1
	modperms.OrderIndex = 2
	modperms.FullAccessPermission = 1
	modperms.TenantId = tenantid

	permission.AS.CreateModulePermission(&modperms, channel.DB)

	for _, categoriesid := range channelcreate.CategoryIds {
		var channelcategory TblChannelCategorie
		channelcategory.ChannelId = ch.Id
		channelcategory.CategoryId = categoriesid
		channelcategory.CreatedAt = channelcreate.CreatedBy
		channelcategory.TenantId = tenantid
		channelcategory.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		err := CH.CreateChannelCategory(&channelcategory, channel.DB)

		if err != nil {

			fmt.Println(err)

		}

	}
	return ch, nil
}

/*create additional fields*/
func (channel *Channel) CreateAdditionalFields(channelcreate ChannelAddtionalField, channelid int, tenantid int) error {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return autherr
	}

	/*Temp store section id*/
	type tempsection struct {
		Id           int
		SectionId    int
		NewSectionId int
	}

	var TempSections []tempsection
	/*create Section*/
	for _, sectionvalue := range channelcreate.Sections {
		var cfld TblField
		cfld.FieldName = strings.TrimSpace(sectionvalue.SectionName)
		cfld.FieldTypeId = sectionvalue.MasterFieldId
		cfld.CreatedBy = channelcreate.CreatedBy
		cfld.TenantId = tenantid
		cfld.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		cfid, fiderr := CH.CreateFields(&cfld, channel.DB)
		if fiderr != nil {

			fmt.Println(fiderr)
		}

		/*create group field*/
		var grpfield TblGroupField
		grpfield.ChannelId = channelid
		grpfield.FieldId = cfid.Id
		grpfield.TenantId = tenantid
		grpfielderr := CH.CreateGroupField(&grpfield, channel.DB)
		if grpfielderr != nil {

			fmt.Println(grpfielderr)

		}

		var TempSection tempsection
		TempSection.Id = cfid.Id
		TempSection.SectionId = sectionvalue.SectionId
		TempSection.NewSectionId = sectionvalue.SectionNewId
		TempSections = append(TempSections, TempSection)

	}

	/*create field*/
	for _, val := range channelcreate.FieldValues {

		var cfld TblField
		cfld.FieldName = strings.TrimSpace(val.FieldName)
		cfld.FieldTypeId = val.MasterFieldId
		cfld.CreatedBy = channelcreate.CreatedBy
		cfld.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		cfld.OrderIndex = val.OrderIndex
		cfld.ImagePath = val.IconPath
		cfld.CharacterAllowed = val.CharacterAllowed
		cfld.Url = val.Url

		if val.MasterFieldId == 4 {

			cfld.DatetimeFormat = val.DateFormat

			cfld.TimeFormat = val.TimeFormat

		}
		if val.MasterFieldId == 6 {

			cfld.DatetimeFormat = val.DateFormat
		}

		if len(val.OptionValue) > 0 {

			cfld.OptionExist = 1
		}

		for _, sectionid := range TempSections {

			if sectionid.SectionId == val.SectionId && sectionid.NewSectionId == val.SectionNewId {

				cfld.SectionParentId = sectionid.Id

			}

		}

		cfid, fiderr := CH.CreateFields(&cfld, channel.DB)

		if fiderr != nil {

			fmt.Println(fiderr)

		}

		/*option value create*/
		for _, opt := range val.OptionValue {

			var fldopt TblFieldOption

			fldopt.OptionName = opt.Value

			fldopt.OptionValue = opt.Value

			fldopt.FieldId = cfid.Id

			fldopt.CreatedBy = channelcreate.CreatedBy
			fldopt.TenantId = tenantid

			fldopt.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

			fopterr := CH.CreateFieldOption(&fldopt, channel.DB)

			if fopterr != nil {

				fmt.Println(fopterr)

			}

		}

		/*create group field*/
		var grpfield TblGroupField

		grpfield.ChannelId = channelid

		grpfield.FieldId = cfid.Id
		grpfield.TenantId = tenantid

		grpfielderr := CH.CreateGroupField(&grpfield, channel.DB)

		if grpfielderr != nil {

			fmt.Println(grpfielderr)

		}

	}

	return nil
}

/*Get channel by name*/
func (channel *Channel) GetchannelByName(channelname string, tenantid int) (channels Tblchannel, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return Tblchannel{}, autherr
	}

	channellist, err1 := CH.GetChannelByChannelName(channelname, channel.DB, tenantid)

	if err1 != nil {

		return Tblchannel{}, err1
	}

	return channellist, nil

}

/*Get Channels By Id*/
func (channel *Channel) GetChannelsById(channelid int, tenantid int) (channelList Tblchannel, SelectedCategories []categories.Arrangecategories, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return Tblchannel{}, []categories.Arrangecategories{}, autherr
	}

	channellist, err := CH.GetChannelById(channelid, channel.DB, tenantid)

	if err != nil {

		return Tblchannel{}, []categories.Arrangecategories{}, err

	}

	GetSelectedChannelCateogry, err1 := CH.GetSelectedCategoryChannelById(channelid, channel.DB, tenantid)

	if err1 != nil {

		fmt.Println(err)
	}

	var FinalSelectedCategories []categories.Arrangecategories

	for _, val := range GetSelectedChannelCateogry {

		var id []int

		ids := strings.Split(val.CategoryId, ",")

		for _, cid := range ids {

			convid, _ := strconv.Atoi(cid)

			id = append(id, convid)
		}

		GetSelectedCategory, _ := CH.GetCategoriseById(id, channel.DB, tenantid)

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
func (channel *Channel) GetChannelsFieldsById(channelid int, tenantid int) (section []Section, fields []Fiedlvalue, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []Section{}, []Fiedlvalue{}, autherr
	}

	groupfield, _ := CH.GetFieldIdByGroupId(channelid, channel.DB, tenantid)

	var ids []int

	for _, val := range groupfield {

		ids = append(ids, val.FieldId)
	}

	fieldValue, _ := CH.GetFieldAndOptionValue(ids, channel.DB, tenantid)

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
func (channel *Channel) DeleteChannel(channelid, modifiedby int, tenantid int) error {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return autherr
	}

	if channelid <= 0 {

		return ErrorChannelId
	}

	CH.DeleteEntryByChannelId(channelid, channel.DB, tenantid)

	CH.DeleteChannelById(channelid, channel.DB, tenantid)

	chdel, _ := CH.GetChannelById(channelid, channel.DB, tenantid)

	var delfidgrp TblFieldGroup

	delfidgrp.IsDeleted = 1

	delfidgrp.DeletedBy = modifiedby

	delfidgrp.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	CH.DeleteFieldGroupById(&delfidgrp, chdel.FieldGroupId, channel.DB, tenantid)

	return nil

}

func (channel *Channel) DeleteChannelPermissions(channelid int) error {

	checkid, _ := permission.AS.GetIdByRouteName(strconv.Itoa(channelid), channel.DB,TenantId)

	permission.AS.Deleterolepermission(checkid.Id, channel.DB,TenantId)

	// permission.AS.DeleteModulePermissioninEntries(channelid, channel.DB)

	return nil
}

/*Change Channel status*/
// status 0 = inactive
// status 1 = active
func (channel *Channel) ChangeChannelStatus(channelid int, status, modifiedby int, tenantid int) (bool, error) {

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

	CH.ChannelIsActive(&channelstatus, channelid, status, channel.DB, tenantid)

	return true, nil

}

/*Get All Master Field type */
func (channel *Channel) GetAllMasterFieldType(tenantid int) (field []TblFieldType, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []TblFieldType{}, autherr
	}

	fid, err := CH.GetAllField(channel.DB, tenantid)

	if err != nil {

		return []TblFieldType{}, err
	}

	return fid, nil

}

/*Edit channel*/
func (channel *Channel) EditChannel(ChannelName string, ChannelDescription string, modifiedby int, channelid int, CategoryIds []string, tenantid int) error {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return autherr
	}

	var chn TblChannel

	chn.ChannelName = ChannelName

	chn.ChannelDescription = ChannelDescription

	chn.ModifiedBy = modifiedby

	chn.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	CH.UpdateChannelDetails(&chn, channelid, channel.DB, tenantid)

	var modpermissionupdate permission.TblModulePermission

	modpermissionupdate.SlugName = strings.ReplaceAll(strings.ToLower(ChannelName), " ", "-")

	modpermissionupdate.RouteName = "/channel/entrylist/" + strconv.Itoa(channelid)

	modpermissionupdate.DisplayName = ChannelName

	CH.UpdateChannelNameInEntries(&modpermissionupdate, channel.DB, tenantid)

	/*channel category create if not exist*/
	for _, val := range CategoryIds {

		err := CH.CheckChannelCategoryAlreadyExitst(channelid, val, channel.DB, tenantid)

		if errors.Is(err, gorm.ErrRecordNotFound) {

			var createCateogry TblChannelCategorie

			createCateogry.ChannelId = channelid

			createCateogry.CategoryId = val

			createCateogry.CreatedAt = modifiedby
			createCateogry.TenantId = tenantid

			createCateogry.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

			CH.CreateChannelCategory(&createCateogry, channel.DB)
		}

	}

	/*delete categoryid if not exist in array*/
	var notexistcategory []tblchannelcategory

	CH.GetChannelCategoryNotExist(&notexistcategory, channelid, CategoryIds, channel.DB, tenantid)

	for _, val := range notexistcategory {

		var deletechannelcategory tblchannelcategory

		CH.DeleteChannelCategoryByValue(&deletechannelcategory, val.Id, channel.DB, tenantid)

	}

	return nil
}

func (channel *Channel) UpdateChannelField(channelupt ChannelUpdate, channelid int, tenantid int) error {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return autherr
	}

	//delete sections & fields
	var delid []int //temp array for delid
	var optiondelid []int

	for _, val := range channelupt.Deletesections {

		delid = append(delid, val.SectionId)
	}

	for _, val := range channelupt.DeleteFields {

		delid = append(delid, val.FieldId)
	}

	for _, val := range channelupt.DeleteOptionsvalue {

		optiondelid = append(optiondelid, val.Id)

	}

	if len(delid) > 0 || len(optiondelid) > 0 {

		var delsection TblField

		delsection.DeletedBy = channelupt.ModifiedBy

		delsection.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		delsection.IsDeleted = 1

		CH.DeleteFieldById(&delsection, delid, channel.DB, tenantid)

		var deloption TblFieldOption

		deloption.DeletedBy = channelupt.ModifiedBy

		deloption.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		deloption.IsDeleted = 1

		CH.DeleteOptionById(&deloption, optiondelid, delid, channel.DB, tenantid)

	}

	/*Temp store section id*/
	type tempsection struct {
		Id           int
		SectionId    int
		NewSectionId int
	}

	var TempSections []tempsection

	for _, val := range channelupt.Sections {

		var cfld TblField

		cfld.FieldName = strings.TrimSpace(val.SectionName)

		cfld.FieldTypeId = val.MasterFieldId

		cfld.CreatedBy = channelupt.ModifiedBy
		cfld.TenantId = tenantid

		cfld.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		if val.SectionId != 0 {

			CH.UpdateFieldDetails(&cfld, val.SectionId, channel.DB, tenantid)

			var TempSection tempsection

			TempSection.Id = val.SectionId

			TempSection.SectionId = val.SectionId

			TempSection.NewSectionId = val.SectionNewId

			TempSections = append(TempSections, TempSection)

		} else {

			cfid, fiderr := CH.CreateFields(&cfld, channel.DB)

			if fiderr != nil {

				fmt.Println(fiderr)
			}

			/*create group field*/
			var grpfield TblGroupField

			grpfield.ChannelId = channelid

			grpfield.FieldId = cfid.Id
			grpfield.TenantId = tenantid

			grpfielderr := CH.CreateGroupField(&grpfield, channel.DB)

			if grpfielderr != nil {

				fmt.Println(grpfielderr)

			}

			var TempSection tempsection

			TempSection.Id = cfid.Id

			TempSection.SectionId = val.SectionId

			TempSection.NewSectionId = val.SectionNewId

			TempSections = append(TempSections, TempSection)

		}

	}

	for _, val := range channelupt.FieldValues {

		var cfld TblField

		cfld.FieldName = strings.TrimSpace(val.FieldName)

		cfld.FieldTypeId = val.MasterFieldId

		cfld.CreatedBy = channelupt.ModifiedBy

		cfld.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		cfld.OrderIndex = val.OrderIndex

		cfld.ImagePath = val.IconPath

		cfld.MandatoryField = val.Mandatory

		cfld.Url = val.Url

		cfld.CharacterAllowed = val.CharacterAllowed

		if val.MasterFieldId == 4 {

			cfld.DatetimeFormat = val.DateFormat

			cfld.TimeFormat = val.TimeFormat

		}
		if val.MasterFieldId == 6 {

			cfld.DatetimeFormat = val.DateFormat
		}

		if len(val.OptionValue) > 0 {

			cfld.OptionExist = 1
		}

		for _, sectionid := range TempSections {

			if sectionid.SectionId == val.SectionId && sectionid.NewSectionId == val.SectionNewId {

				cfld.SectionParentId = sectionid.Id

			}

		}

		var createdchannelid int

		if val.FieldId != 0 {

			cfld.ModifiedBy = channelupt.ModifiedBy

			cfld.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

			CH.UpdateFieldDetails(&cfld, val.FieldId, channel.DB, tenantid)

			createdchannelid = val.FieldId

		} else {
			cfld.TenantId = tenantid
			cfid, fiderr := CH.CreateFields(&cfld, channel.DB)

			if fiderr != nil {

				fmt.Println(fiderr)

			}

			/*create group field*/
			var grpfield TblGroupField

			grpfield.ChannelId = channelid

			grpfield.FieldId = cfid.Id
			grpfield.TenantId = tenantid

			grpfielderr := CH.CreateGroupField(&grpfield, channel.DB)

			if grpfielderr != nil {

				fmt.Println(grpfielderr)

			}

			createdchannelid = cfld.Id

		}
		for _, optv := range val.OptionValue {

			var fldopt TblFieldOption

			fldopt.OptionName = optv.Value

			fldopt.OptionValue = optv.Value

			fldopt.CreatedBy = channelupt.ModifiedBy

			fldopt.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

			if optv.Id != 0 {

				fldopt.FieldId = optv.FieldId

				fldopt.ModifiedBy = channelupt.ModifiedBy

				fldopt.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

				CH.UpdateFieldOption(&fldopt, optv.Id, channel.DB, tenantid)

			} else {

				fldopt.FieldId = createdchannelid
				fldopt.TenantId = tenantid

				fopterr := CH.CreateFieldOption(&fldopt, channel.DB)

				if fopterr != nil {

					fmt.Println(fopterr)

				}

			}

		}

	}
	return nil
}

// Get channel count
func (channel *Channel) GetChannelCount(tenantid int) (count int, err error) {

	var chcount int64

	err = CH.GetChannelCount(&chcount, channel.DB, tenantid)

	if err != nil {

		return 0, err
	}

	return int(chcount), nil

}

func (channel *Channel) GetChannelsWithEntries(tenantid int) ([]Tblchannel, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []Tblchannel{}, autherr
	}

	var channel_contents []Tblchannel

	err := CH.GetChannels(&channel_contents, channel.DB, tenantid)

	if err != nil {

		fmt.Println(err)

		return []Tblchannel{}, err
	}

	var FinalChannellist []Tblchannel

	for _, chn := range channel_contents {

		var channel_entries []TblChannelEntries

		CH.GetChannelEntriesByChannelId(&channel_entries, chn.Id, channel.DB, tenantid)

		if len(channel_entries) > 0 {

			chn.ChannelEntries = channel_entries

			FinalChannellist = append(FinalChannellist, chn)

		}
	}

	return FinalChannellist, nil

}
