package channels

import (
	"log"
	"strings"
	"time"
)

/*all channel Entries List*/
//if channelid 0 get all channel entries
// if channelid not eq 0 to get particular entries of the channel
func (channel *Channel) GetAllChannelEntriesList(channelid int, limit, offset int, filter EntriesFilter) (entries []tblchannelentries, filterentriescount int, totalentriescount int, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []tblchannelentries{}, 0, 0, autherr
	}

	if filter.Status == "Draft" {

		filter.Status = "0"

	} else if filter.Status == "Published" {

		filter.Status = "1"

	} else if filter.Status == "Unpublished" {

		filter.Status = "2"
	}

	chnentry, _, _ := EntryModel.ChannelEntryList(limit, offset, channelid, filter, channel.Permissions.RoleId, true, channel.DB)

	_, filtercount, _ := EntryModel.ChannelEntryList(0, 0, channelid, filter, channel.Permissions.RoleId, true, channel.DB)

	_, entrcount, _ := EntryModel.ChannelEntryList(0, 0, channelid, EntriesFilter{}, channel.Permissions.RoleId, true, channel.DB)

	return chnentry, int(filtercount), int(entrcount), nil

}

// get entry details
func (channel *Channel) GetEntryDetailsById(ChannelName string, EntryId int) (tblchannelentries, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return tblchannelentries{}, autherr
	}

	Entry, err := EntryModel.GetChannelEntryById(EntryId, channel.DB)

	if err != nil {

		return tblchannelentries{}, nil
	}

	return Entry, nil
}

// create entry
func (channel *Channel) CreateEntry(entriesrequired EntriesRequired) (entry tblchannelentries, flg bool, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return tblchannelentries{}, false, autherr
	}

	var Entries tblchannelentries

	Entries.Title = entriesrequired.Title

	Entries.Description = entriesrequired.Content

	Entries.CoverImage = entriesrequired.CoverImage

	Entries.MetaTitle = entriesrequired.SEODetails.MetaTitle

	Entries.MetaDescription = entriesrequired.SEODetails.MetaDescription

	Entries.Keyword = entriesrequired.SEODetails.MetaKeywords

	if entriesrequired.SEODetails.MetaSlug == "" {

		Entries.Slug = strings.ReplaceAll(strings.ToLower(entriesrequired.Title), " ", "_")

	} else {

		Entries.Slug = entriesrequired.SEODetails.MetaSlug

	}

	Entries.Status = entriesrequired.Status

	Entries.ChannelId = entriesrequired.ChannelId

	Entries.CategoriesId = entriesrequired.CategoryIds

	Entries.CreatedBy = entriesrequired.CreatedBy

	Entries.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Entries.Author = entriesrequired.Author

	Entries.CreateTime = entriesrequired.CreateTime

	Entries.PublishedTime = entriesrequired.PublishTime

	Entries.ReadingTime = entriesrequired.ReadingTime

	Entries.SortOrder = entriesrequired.SortOrder

	Entries.Tags = entriesrequired.Tag

	Entries.Excerpt = entriesrequired.Excerpt

	Entries.ImageAltTag = entriesrequired.SEODetails.ImageAltTag

	Entriess, err := EntryModel.CreateChannelEntry(Entries, channel.DB)

	if err != nil {

		log.Println(err)
	}

	return Entriess, true, nil
}

func (channel *Channel) CreateChannelEntryFields(entryid int, createdby int, AdditionalFields []AdditionalFields) error {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return autherr
	}

	var EntriesField []TblChannelEntryField

	for _, val := range AdditionalFields {

		var Entrfield TblChannelEntryField

		Entrfield.ChannelEntryId = entryid

		Entrfield.FieldName = val.FieldName

		Entrfield.FieldValue = val.FieldValue

		Entrfield.FieldId = val.FieldId

		Entrfield.CreatedBy = createdby

		Entrfield.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		EntriesField = append(EntriesField, Entrfield)

	}

	ferr := EntryModel.CreateEntrychannelFields(&EntriesField, channel.DB)

	if ferr != nil {

		return ferr
	}

	return nil

}

/**/
func (channel *Channel) DeleteEntry(ChannelName string, modifiedby int, Entryid int) (bool, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return false, autherr
	}

	var entries tblchannelentries

	entries.Id = Entryid

	entries.IsDeleted = 1

	entries.DeletedBy = modifiedby

	entries.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := EntryModel.DeleteChannelEntryId(&entries, Entryid, channel.DB)

	var field TblChannelEntryField

	field.DeletedBy = modifiedby

	field.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := EntryModel.DeleteChannelEntryFieldId(&field, Entryid, channel.DB)

	if err != nil {

		log.Println(err)
	}

	if err1 != nil {

		log.Println(err)
	}

	return true, nil

}
