package channels

import (
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spurtcms/categories"
)

// get channel Entries List
func (channel *Channel) ChannelEntriesList(entry Entries) (entries []tblchannelentries, filterentriescount int, totalentriescount int, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []tblchannelentries{}, 0, 0, autherr
	}

	if entry.Status == "Draft" {

		entry.Status = "0"

	} else if entry.Status == "Published" {

		entry.Status = "1"

	} else if entry.Status == "Unpublished" {

		entry.Status = "2"
	}

	var categoryid string

	if entry.CategoryId != 0 && entry.CategoryId > 0 {

		categoryid = Categories(entry.CategoryId, channel.DB)

	}

	if entry.CategoryName != "" {

		categoryid = CategoriesByUsingName(entry.CategoryName, channel.DB)
	}

	chnentry, _, _ := EntryModel.ChannelEntryList(entry, channel, categoryid, channel.DB)

	_, filtercount,_  := EntryModel.ChannelEntryList(Entries{Limit: 0,Offset: 0,Keyword: entry.Keyword,ChannelName: entry.ChannelName,Status: entry.Status,Title: entry.Title,UserName: entry.UserName,Publishedonly: entry.Publishedonly,ActiveChannelEntriesonly: entry.ActiveChannelEntriesonly,CategoryId: entry.ChannelId,MemberAccessControl: entry.MemberAccessControl,ChannelId: entry.ChannelId}, channel, categoryid, channel.DB)

	_, entrcount, _ := EntryModel.ChannelEntryList(Entries{Limit: 0,Offset: 0,Keyword: entry.Keyword,ChannelName: entry.ChannelName,Status: entry.Status,Title: entry.Title,UserName: entry.UserName,Publishedonly: entry.Publishedonly,ActiveChannelEntriesonly: entry.ActiveChannelEntriesonly,CategoryId: entry.ChannelId,MemberAccessControl: entry.MemberAccessControl,ChannelId: entry.ChannelId}, channel, categoryid, channel.DB)

	var final_entries_list []tblchannelentries

	for _, entries := range chnentry {

		if entry.AuthorDetails {

			authorDetails, _ := EntryModel.GetAuthorDetails(channel.DB, entries.CreatedBy)

			if authorDetails.AuthorID != 0 {

				var modified_profileImage string

				if authorDetails.ProfileImagePath != nil {

					modified_profileImage = entry.ImageUrlPath + *authorDetails.ProfileImagePath
				}

				authorDetails.ProfileImagePath = &modified_profileImage

				entries.AuthorDetail = authorDetails
			}

		}

		var memberId string

		var final_fieldsList []tblfield

		if entry.AdditionalFields {

			sections, _ := EntryModel.GetSectionsUnderEntries(channel.DB, entries.ChannelId, entry.FieldTypeId)

			entries.Sections = sections

			fields, _ := EntryModel.GetFieldsInEntries(channel.DB, entries.ChannelId, entry.FieldTypeId)

			for _, field := range fields {

				var modified_field_path string

				if field.ImagePath != "" {

					modified_field_path = entry.ImageUrlPath + strings.TrimPrefix(field.ImagePath, "/")
				}

				field.ImagePath = modified_field_path

				fieldValue, _ := EntryModel.GetFieldValue(channel.DB, field.Id, entries.Id)

				if fieldValue.Id != 0 {

					field.FieldValue = fieldValue

					if field.FieldTypeId == entry.MemberFieldTypeId {

						memberId = fieldValue.FieldValue
					}
				}

				fieldOptions, _ := EntryModel.GetFieldOptions(channel.DB, field.Id)

				if len(fieldOptions) > 0 {

					field.FieldOptions = fieldOptions

				}

				final_fieldsList = append(final_fieldsList, field)
			}
		}

		if entry.MemberProfile {

			entries.Fields = final_fieldsList

			conv_memid, _ := strconv.Atoi(memberId)

			memberProfile, _ := EntryModel.GetMemberProfile(channel.DB, conv_memid)

			var modified_profile_path string

			if memberProfile.CompanyLogo != "" {

				modified_profile_path = entry.ImageUrlPath + strings.TrimPrefix(memberProfile.CompanyLogo, "/")
			}

			memberProfile.CompanyLogo = modified_profile_path

			entries.MemberProfiles = memberProfile

		}

		splittedArr := strings.Split(entries.CategoriesId, ",")

		var parentCatId int

		var indivCategories [][]categories.TblCategories

		for _, catId := range splittedArr {

			conv_id, _ := strconv.Atoi(catId)

			var indivCategory []categories.TblCategories

			category, _ := EntryModel.GetGraphqlEntriesCategoryByParentId(channel.DB, conv_id)

			var modified_category_path string

			if category.ImagePath != "" {

				modified_category_path = entry.ImageUrlPath + strings.TrimPrefix(category.ImagePath, "/")
			}

			category.ImagePath = modified_category_path

			if category.Id != 0 {

				indivCategory = append(indivCategory, category)
			}

			parentCatId = category.ParentId

			if parentCatId != 0 {

				var count int

			LOOP:

				for {

					count = count + 1 //count increment used to check how many times the loop gets executed

					parentCategory, _ := EntryModel.GetGraphqlEntriesCategoryByParentId(channel.DB, parentCatId)

					var modified_category_path string

					if parentCategory.ImagePath != "" {

						modified_category_path = entry.ImageUrlPath + strings.TrimPrefix(parentCategory.ImagePath, "/")
					}

					parentCategory.ImagePath = modified_category_path

					if parentCategory.Id != 0 {

						indivCategory = append(indivCategory, parentCategory)
					}

					parentCatId = parentCategory.ParentId

					if parentCatId != 0 { //mannuall condition to break the loop in overlooping situations

						goto LOOP

					} else if count > 49 {

						break //use to break the loop if infinite loop doesn't break ,So forcing the loop to break at overlooping conditions

					} else {

						break
					}

				}

			}

			if len(indivCategory) > 0 {

				sort.SliceStable(indivCategory, func(i, j int) bool {

					return indivCategory[i].Id < indivCategory[j].Id

				})

				indivCategories = append(indivCategories, indivCategory)
			}

		}

		entries.Categories = indivCategories

		final_entries_list = append(final_entries_list, entries)
	}

	return final_entries_list, int(filtercount), int(entrcount), nil

}

// get entry details
func (channel *Channel) EntryDetailsById(Ent IndivEntriesReq) (tblchannelentries, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return tblchannelentries{}, autherr
	}
	
	Entry, err := EntryModel.GetChannelEntryById(Ent, channel.DB)

	if err != nil {

		return tblchannelentries{}, nil
	}

	if Ent.AuthorDetails {

		authorDetails, _ := EntryModel.GetAuthorDetails(channel.DB, Entry.CreatedBy)

		if authorDetails.AuthorID != 0 {

			var modified_profileImage string

			if authorDetails.ProfileImagePath != nil {

				modified_profileImage = Ent.ImageUrlPath + *authorDetails.ProfileImagePath
			}

			authorDetails.ProfileImagePath = &modified_profileImage

			Entry.AuthorDetail = authorDetails
		}

	}

	var memberId string

	if Ent.AdditionalFields {

		sections, _ := EntryModel.GetSectionsUnderEntries(channel.DB, Entry.ChannelId, Ent.FieldTypeId)

		Entry.Sections = sections

		var final_fieldsList []tblfield

		fields, _ := EntryModel.GetFieldsInEntries(channel.DB, Entry.ChannelId, Ent.FieldTypeId)

		for _, field := range fields {

			var modified_field_path string

			if field.ImagePath != "" {

				modified_field_path = Ent.ImageUrlPath + strings.TrimPrefix(field.ImagePath, "/")
			}

			field.ImagePath = modified_field_path

			fieldValue, _ := EntryModel.GetFieldValue(channel.DB, field.Id, Entry.Id)

			if fieldValue.Id != 0 {

				field.FieldValue = fieldValue

				if field.FieldTypeId == Ent.MemberFieldTypeId {

					memberId = fieldValue.FieldValue
				}
			}

			fieldOptions, _ := EntryModel.GetFieldOptions(channel.DB, field.Id)

			if len(fieldOptions) > 0 {

				field.FieldOptions = fieldOptions

			}

			final_fieldsList = append(final_fieldsList, field)
		}
	}

	if Ent.MemberProfile {

		conv_memid, _ := strconv.Atoi(memberId)

		memberProfile, _ := EntryModel.GetMemberProfile(channel.DB, conv_memid)

		var modified_profile_path string

		if memberProfile.CompanyLogo != "" {

			modified_profile_path = Ent.ImageUrlPath + strings.TrimPrefix(memberProfile.CompanyLogo, "/")
		}

		memberProfile.CompanyLogo = modified_profile_path

		Entry.MemberProfiles = memberProfile

	}

	if Ent.CategoriesEnable {

		splittedArr := strings.Split(Entry.CategoriesId, ",")

		var parentCatId int

		var indivCategories [][]categories.TblCategories

		for _, catId := range splittedArr {

			var indivCategory []categories.TblCategories

			conv_id, _ := strconv.Atoi(catId)

			category, _ := EntryModel.GetGraphqlEntriesCategoryByParentId(channel.DB, conv_id)

			var modified_category_path string

			if category.ImagePath != "" {

				modified_category_path = Ent.ImageUrlPath + strings.TrimPrefix(category.ImagePath, "/")
			}

			category.ImagePath = modified_category_path

			if category.Id != 0 {

				indivCategory = append(indivCategory, category)
			}

			parentCatId = category.ParentId

			if parentCatId != 0 {

				var count int

			LOOP:

				for {

					count = count + 1 //count increment used to check how many times the loop gets executed

					parentCategory, _ := EntryModel.GetGraphqlEntriesCategoryByParentId(channel.DB, parentCatId)

					var modified_category_path string

					if parentCategory.ImagePath != "" {

						modified_category_path = Ent.ImageUrlPath + strings.TrimPrefix(parentCategory.ImagePath, "/")
					}

					parentCategory.ImagePath = modified_category_path

					if parentCategory.Id != 0 {

						indivCategory = append(indivCategory, parentCategory)
					}

					parentCatId = parentCategory.ParentId

					if parentCatId != 0 { //mannuall condition to break the loop in overlooping situations

						goto LOOP

					} else if count > 49 {

						break //use to break the loop if infinite loop doesn't break ,So forcing the loop to break at overlooping conditions

					} else {

						break
					}

				}

			}

			if len(indivCategory) > 0 {

				sort.SliceStable(indivCategory, func(i, j int) bool {

					return indivCategory[i].Id < indivCategory[j].Id

				})

				indivCategories = append(indivCategories, indivCategory)
			}

		}

		Entry.Categories = indivCategories

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

// update entry
func (channel *Channel) UpdateEntry(entriesrequired EntriesRequired, ChannelName string, EntryId int) (bool, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return false, autherr
	}

	var Entries TblChannelEntries

	Entries.Title = entriesrequired.Title

	Entries.Description = entriesrequired.Content

	Entries.CoverImage = entriesrequired.CoverImage

	Entries.MetaTitle = entriesrequired.SEODetails.MetaTitle

	Entries.MetaDescription = entriesrequired.SEODetails.MetaDescription

	Entries.Keyword = entriesrequired.SEODetails.MetaKeywords

	Entries.ImageAltTag = entriesrequired.SEODetails.ImageAltTag

	if entriesrequired.SEODetails.MetaSlug == "" {

		Entries.Slug = strings.ReplaceAll(strings.ToLower(entriesrequired.Title), " ", "_")

	} else {

		Entries.Slug = entriesrequired.SEODetails.MetaSlug

	}

	Entries.Status = entriesrequired.Status

	Entries.ChannelId = entriesrequired.ChannelId

	Entries.CategoriesId = entriesrequired.CategoryIds

	Entries.ModifiedBy = entriesrequired.ModifiedBy

	Entries.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Entries.Author = entriesrequired.Author

	Entries.CreateTime = entriesrequired.CreateTime

	Entries.PublishedTime = entriesrequired.PublishTime

	Entries.ReadingTime = entriesrequired.ReadingTime

	Entries.SortOrder = entriesrequired.SortOrder

	Entries.Tags = entriesrequired.Tag

	Entries.Excerpt = entriesrequired.Excerpt

	err := EntryModel.UpdateChannelEntryDetails(&Entries, EntryId, channel.DB)

	if err != nil {

		return false, err
	}

	return true, nil
}

// update entry additional fields
func (channel *Channel) UpdateAdditionalField(AdditionalFields []AdditionalFields, EntryId int) (bool, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return false, autherr
	}

	for _, val := range AdditionalFields {

		if val.Id == 0 {

			var Entrfield TblChannelEntryField

			Entrfield.ChannelEntryId = EntryId

			Entrfield.FieldName = val.FieldName

			Entrfield.FieldValue = val.FieldValue

			Entrfield.FieldId = val.FieldId

			Entrfield.CreatedBy = val.ModifiedBy

			Entrfield.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

			EntryModel.CreateSingleEntrychannelFields(&Entrfield, channel.DB)

		} else {

			var Entrfield TblChannelEntryField

			Entrfield.Id = val.Id

			Entrfield.ChannelEntryId = EntryId

			Entrfield.FieldName = val.FieldName

			Entrfield.FieldValue = val.FieldValue

			Entrfield.FieldId = val.FieldId

			Entrfield.ModifiedBy = val.ModifiedBy

			Entrfield.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

			EntryModel.UpdateChannelEntryAdditionalDetails(Entrfield, channel.DB)

		}

	}

	return true, nil

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

// Makefeature helps to highlights entry, only one entry should be featured of each channel and that is also optional
func (channel *Channel) MakeFeatureEntry(channelid, entryid, status int) (flag bool, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return false, autherr
	}

	merr := EntryModel.MakeFeature(channelid, entryid, status, channel.DB)

	if merr != nil {

		return false, merr
	}

	return true, nil

}

// change entries status
func (channel *Channel) EntryStatus(ChannelName string, EntryId int, status int, modifiedby int) (bool, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return false, autherr
	}

	var Entries TblChannelEntries

	Entries.Status = status

	Entries.ModifiedBy = modifiedby

	Entries.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	EntryModel.PublishQuery(&Entries, EntryId, channel.DB)

	return true, nil

}
