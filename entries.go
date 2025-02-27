package channels

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/spurtcms/categories"
	"github.com/spurtcms/member"
	"github.com/spurtcms/team"
)

// get channel Entries List
func (channel *Channel) ChannelEntriesList(entry Entries, tenantid int) (entries []Tblchannelentries, filterentriescount int, totalentriescount int, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []Tblchannelentries{}, 0, 0, autherr
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

		categoryid = Categories(entry.CategoryId, channel.DB, tenantid)
	}

	if entry.CategoryName != "" {
		categoryid = CategoriesByUsingName(entry.CategoryName, channel.DB, tenantid)
	}

	EntryModel.Dataaccess = channel.DataAccess
	EntryModel.Userid = channel.Userid

	chnentry, _, _ := EntryModel.ChannelEntryList(entry, channel, categoryid, true, channel.DB, tenantid)

	_, filtercount, _ := EntryModel.ChannelEntryList(Entries{Limit: 0, Offset: 0, Keyword: entry.Keyword, ChannelName: entry.ChannelName, Status: entry.Status, Title: entry.Title, UserName: entry.UserName, Publishedonly: entry.Publishedonly, ActiveChannelEntriesonly: entry.ActiveChannelEntriesonly, CategoryId: entry.CategoryId, MemberAccessControl: entry.MemberAccessControl, ChannelId: entry.ChannelId}, channel, categoryid, true, channel.DB, tenantid)

	// _, entrcount, _ := EntryModel.ChannelEntryList(Entries{Limit: 0, Offset: 0, Keyword: entry.Keyword, ChannelName: entry.ChannelName, Status: entry.Status, Title: entry.Title, UserName: entry.UserName, Publishedonly: entry.Publishedonly, ActiveChannelEntriesonly: entry.ActiveChannelEntriesonly, CategoryId: entry.CategoryId, MemberAccessControl: entry.MemberAccessControl, ChannelId: entry.ChannelId}, channel, categoryid, true, channel.DB)

	var final_entries_list []Tblchannelentries

	for _, entries := range chnentry {

		var first = entries.FirstName
		var last = entries.LastName
		var firstn = strings.ToUpper(first[:1])
		var lastn string
		if entries.LastName != "" {
			lastn = strings.ToUpper(last[:1])
		}

		var Name = firstn + lastn
		entries.NameString = Name

		if entry.AuthorDetails {

			authorDetails, _ := EntryModel.GetAuthorDetails(channel.DB, entries.CreatedBy, tenantid)
			if authorDetails.Id != 0 {

				var modified_profileImage string
				if authorDetails.ProfileImagePath != "" {
					modified_profileImage = entry.ImageUrlPath + authorDetails.ProfileImagePath
				}

				authorDetails.ProfileImagePath = modified_profileImage
				entries.AuthorDetail = authorDetails
			}

		}

		var (
			memberId         string
			final_fieldsList []tblfield
		)

		if entry.AdditionalFields {

			sections, _ := EntryModel.GetSectionsUnderEntries(channel.DB, entries.ChannelId, entry.FieldTypeId, tenantid)
			entries.Sections = sections
			fields, _ := EntryModel.GetFieldsInEntries(channel.DB, entries.ChannelId, entry.FieldTypeId, tenantid)

			for _, field := range fields {

				var modified_field_path string
				if field.ImagePath != "" {
					modified_field_path = entry.ImageUrlPath + strings.TrimPrefix(field.ImagePath, "/")
				}

				field.ImagePath = modified_field_path
				fieldValue, _ := EntryModel.GetFieldValue(channel.DB, field.Id, entries.Id, tenantid)

				if fieldValue.Id != 0 {
					field.FieldValue = fieldValue
					if field.FieldTypeId == entry.MemberFieldTypeId {
						memberId = fieldValue.FieldValue
					}
				}

				fieldOptions, _ := EntryModel.GetFieldOptions(channel.DB, field.Id, tenantid)
				if len(fieldOptions) > 0 {
					field.FieldOptions = fieldOptions

				}
				final_fieldsList = append(final_fieldsList, field)
			}
		}

		if entry.MemberProfile {

			entries.Fields = final_fieldsList
			conv_memid, _ := strconv.Atoi(memberId)
			memberProfile, _ := EntryModel.GetMemberProfile(channel.DB, conv_memid, tenantid)
			var modified_profile_path string
			if memberProfile.CompanyLogo != "" {
				modified_profile_path = entry.ImageUrlPath + strings.TrimPrefix(memberProfile.CompanyLogo, "/")
			}

			memberProfile.CompanyLogo = modified_profile_path
			entries.MemberProfiles = memberProfile
		}

		splittedArr := strings.Split(entries.CategoriesId, ",")

		var (
			parentCatId     int
			indivCategories [][]categories.TblCategories
		)

		for _, catId := range splittedArr {

			conv_id, _ := strconv.Atoi(catId)

			var (
				indivCategory          []categories.TblCategories
				modified_category_path string
			)

			category, _ := EntryModel.GetGraphqlEntriesCategoryByParentId(channel.DB, conv_id, tenantid)

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
					parentCategory, _ := EntryModel.GetGraphqlEntriesCategoryByParentId(channel.DB, parentCatId, tenantid)
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

	return final_entries_list, int(filtercount), int(filtercount), nil

}

// Channel entries list retrieval function with a feature to get entries related datas
func (channel *Channel) FlexibleChannelEntriesList(input EntriesInputs) (ChannelEntries []Tblchannelentries, FilterCount int, TotalCount int, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []Tblchannelentries{}, 0, 0, autherr
	}

	if input.Status == "Draft" {

		input.Status = "0"

	} else if input.Status == "Publish" {

		input.Status = "1"

	} else if input.Status == "Unpublish" {

		input.Status = "2"

	} else {

		input.Status = ""
	}

	var (
		EntriesData             []JoinEntries
		commonCount, totalCount int64
	)

	err = EntryModel.GetFlexibleEntriesData(input, channel, channel.DB, &EntriesData, &commonCount, &totalCount)

	if err != nil {

		return []Tblchannelentries{}, 0, 0, autherr
	}

	var channelEntries []Tblchannelentries

	for _, data := range EntriesData {

		var memberProfile member.TblMemberProfile

		if input.GetMemberProfile {

			memberProfile = member.TblMemberProfile{
				Id:              data.ProfileId,
				MemberId:        data.MemberID,
				ProfilePage:     data.ProfilePage,
				ProfileName:     data.ProfileName,
				ProfileSlug:     data.ProfileSlug,
				CompanyLogo:     data.CompanyLogo,
				StorageType:     data.ProfStorageType,
				CompanyName:     data.CompanyName,
				CompanyLocation: data.CompanyLocation,
				About:           data.About,
				Linkedin:        data.Linkedin,
				Website:         data.Website,
				Twitter:         data.Twitter,
				SeoTitle:        data.SeoTitle,
				SeoDescription:  data.SeoDescription,
				SeoKeyword:      data.SeoKeyword,
				MemberDetails:   data.MemberDetails,
				ClaimStatus:     data.ClaimStatus,
				CreatedBy:       data.ProfCreatedBy,
				CreatedOn:       data.ProfCreatedOn,
				ModifiedBy:      data.ProfModifiedBy,
				ModifiedOn:      data.ProfModifiedOn,
				IsDeleted:       data.ProfIsDeleted,
				DeletedOn:       data.ProfDeletedOn,
				DeletedBy:       data.ProfDeletedBy,
				ClaimDate:       data.ClaimDate,
				TenantId:        data.ProfTenantId,
			}
		}

		var authorDetails team.TblUser

		if input.GetAuthorDetails {

			authorDetails = team.TblUser{
				Id:                data.AuthorId,
				FirstName:         data.FirstName,
				LastName:          data.LastName,
				RoleId:            data.RoleId,
				Email:             data.Email,
				Username:          data.Username,
				MobileNo:          data.MobileNo,
				IsActive:          data.AuthorActive,
				ProfileImage:      data.ProfileImage,
				ProfileImagePath:  data.ProfileImagePath,
				StorageType:       data.AuthorStorageType,
				DataAccess:        data.DataAccess,
				CreatedOn:         data.AuthorCreatedOn,
				CreatedBy:         data.AuthorCreatedBy,
				ModifiedOn:        data.AuthorModifiedOn,
				ModifiedBy:        data.AuthorModifiedBy,
				LastLogin:         data.LastLogin,
				IsDeleted:         data.AuthorIsDeleted,
				DeletedOn:         data.AuthorDeletedOn,
				DeletedBy:         data.AuthorDeletedBy,
				DefaultLanguageId: data.DefaultLanguageId,
				TenantId:          data.UserTenantId,
				RoleName:          data.RoleName,
			}

		}

		var categoryHierarchy [][]categories.TblCategories

		if input.GetLinkedCategories && data.CategoriesID != "" {

			var categoriez []categories.TblCategories

			splitArr := strings.Split(data.CategoriesID, ",")

			err := categories.Categorymodel.GetHierarchicalCategoriesMappedInEntries(splitArr, &categoriez, channel.DB, channel.DB.Dialector.Name())

			if err != nil {

				fmt.Println("category retrieval error: ", err)
			}

			for _, mapId := range splitArr {

				IntId, _ := strconv.Atoi(mapId)

				var categoryStream []categories.TblCategories

				for _, category := range categoriez {

					if category.Id == IntId {

						parentId := category.ParentId

						categoryStream = append(categoryStream, category)

					LOOP:

						for _, parent := range categoriez {

							if parentId == parent.Id {

								parentId = parent.ParentId

								categoryStream = append(categoryStream, parent)

								if parent.ParentId != 0 {

									goto LOOP

								} else {

									break
								}
							}
						}
					}
				}

				categoryHierarchy = append(categoryHierarchy, categoryStream)

			}

			fmt.Println("catdata", categoryHierarchy)

		}

		var sections, fields []tblfield

		if input.GetAdditionalFields {

			additionalFields, _ := EntryModel.GetChannelAdditionalFields(channel.DB, data.ChannelID)

			for _, field := range additionalFields {

				if field.FieldTypeId != input.SectionFieldTypeId {

					if field.OptionExist == 1 {

						field.FieldOptions, _ = EntryModel.GetFieldOptions(channel.DB, field.Id, input.TenantId)
					}

					field.FieldValue, _ = EntryModel.GetFieldValue(channel.DB, field.Id, data.Id, input.TenantId)

					fields = append(fields, field)

				} else {

					sections = append(sections, field)
				}
			}

		}

		channnel_entry := Tblchannelentries{
			Id:              data.Id,
			Title:           data.Title,
			Slug:            data.Slug,
			Description:     data.Description,
			UserId:          data.UserID,
			ChannelId:       data.ChannelID,
			Status:          data.Status,
			IsActive:        data.IsActive,
			CreatedOn:       data.CreatedOn,
			CreatedBy:       data.CreatedBy,
			ModifiedBy:      data.ModifiedBy,
			ModifiedOn:      data.ModifiedOn,
			CoverImage:      data.CoverImage,
			ThumbnailImage:  data.ThumbnailImage,
			PublishedTime:   data.PublishedTime,
			MetaDescription: data.MetaDescription,
			MetaTitle:       data.MetaTitle,
			Keyword:         data.Keyword,
			ImageAltTag:     data.ImageAltTag,
			CategoriesId:    data.CategoriesID,
			RelatedArticles: data.RelatedArticles,
			Feature:         data.Feature,
			ViewCount:       data.ViewCount,
			Author:          data.Author,
			SortOrder:       data.SortOrder,
			CreateTime:      data.CreateTime,
			ReadingTime:     data.ReadingTime,
			Tags:            data.Tags,
			Excerpt:         data.Excerpt,
			IsDeleted:       data.IsDeleted,
			DeletedOn:       data.DeletedOn,
			DeletedBy:       data.DeletedBy,
			AuthorDetail:    authorDetails,
			MemberProfiles:  memberProfile,
			Categories:      categoryHierarchy,
			Sections:        sections,
			Fields:          fields,
			TenantId:        data.TenantId,
		}

		channelEntries = append(channelEntries, channnel_entry)
	}

	return channelEntries, int(commonCount), int(totalCount), nil

}

// get entry details
func (channel *Channel) EntryDetailsById(Ent IndivEntriesReq, tenantid int) (Tblchannelentries, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return Tblchannelentries{}, autherr
	}

	Entry, err := EntryModel.GetChannelEntryById(Ent, channel.DB, tenantid)
	if err != nil {

		return Tblchannelentries{}, nil
	}

	if Ent.AuthorDetails {

		authorDetails, _ := EntryModel.GetAuthorDetails(channel.DB, Entry.CreatedBy, tenantid)

		if authorDetails.Id != 0 {

			var modified_profileImage string
			if authorDetails.ProfileImagePath != "" {
				modified_profileImage = Ent.ImageUrlPath + authorDetails.ProfileImagePath
			}
			authorDetails.ProfileImagePath = modified_profileImage

			Entry.AuthorDetail = authorDetails
		}

	}

	var memberId string

	if Ent.AdditionalFields {

		sections, _ := EntryModel.GetSectionsUnderEntries(channel.DB, Entry.ChannelId, Ent.FieldTypeId, tenantid)

		Entry.Sections = sections

		var final_fieldsList []tblfield

		fields, _ := EntryModel.GetFieldsInEntries(channel.DB, Entry.ChannelId, Ent.FieldTypeId, tenantid)

		for _, field := range fields {

			var modified_field_path string

			if field.ImagePath != "" {

				modified_field_path = Ent.ImageUrlPath + strings.TrimPrefix(field.ImagePath, "/")
			}

			field.ImagePath = modified_field_path

			fieldValue, _ := EntryModel.GetFieldValue(channel.DB, field.Id, Entry.Id, tenantid)

			if fieldValue.Id != 0 {

				field.FieldValue = fieldValue

				if field.FieldTypeId == Ent.MemberFieldTypeId {

					memberId = fieldValue.FieldValue
				}
			}

			fieldOptions, _ := EntryModel.GetFieldOptions(channel.DB, field.Id, tenantid)

			if len(fieldOptions) > 0 {

				field.FieldOptions = fieldOptions

			}

			final_fieldsList = append(final_fieldsList, field)
		}
	}

	if Ent.MemberProfile {

		conv_memid, _ := strconv.Atoi(memberId)

		memberProfile, _ := EntryModel.GetMemberProfile(channel.DB, conv_memid, tenantid)

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

			category, _ := EntryModel.GetGraphqlEntriesCategoryByParentId(channel.DB, conv_id, tenantid)

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

					parentCategory, _ := EntryModel.GetGraphqlEntriesCategoryByParentId(channel.DB, parentCatId, tenantid)

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
func (channel *Channel) CreateEntry(entriesrequired EntriesRequired, tenantid int) (entry Tblchannelentries, flg bool, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return Tblchannelentries{}, false, autherr
	}

	var Entries Tblchannelentries

	uuid := (uuid.New()).String()

	arr := strings.Split(uuid, "-")

	Entries.Uuid = arr[len(arr)-1]

	Entries.Title = entriesrequired.Title

	Entries.Description = entriesrequired.Content

	Entries.CoverImage = entriesrequired.CoverImage

	Entries.MetaTitle = entriesrequired.SEODetails.MetaTitle

	Entries.IsActive = entriesrequired.IsActive

	Entries.MetaDescription = entriesrequired.SEODetails.MetaDescription

	Entries.Keyword = entriesrequired.SEODetails.MetaKeywords

	// Entries.Slug = strings.ReplaceAll(strings.ToLower(entriesrequired.Title), " ", "-")

	// re := regexp.MustCompile(`[^a-z0-9\-]`)

	// Entries.Slug = re.ReplaceAllString(Entries.Slug, "-")

	// Entries.Slug = regexp.MustCompile(`-+`).ReplaceAllString(Entries.Slug, "-")

	// Entries.Slug = strings.Trim(Entries.Slug, "-")
	if entriesrequired.SEODetails.MetaSlug == "" {

		Entries.Slug = strings.ReplaceAll(strings.ToLower(entriesrequired.Title), " ", "-")

		re := regexp.MustCompile(`[^a-z0-9\-]`)

		Entries.Slug = re.ReplaceAllString(Entries.Slug, "-")

		Entries.Slug = regexp.MustCompile(`-+`).ReplaceAllString(Entries.Slug, "-")

		Entries.Slug = strings.Trim(Entries.Slug, "-")

	} else if entriesrequired.SEODetails.MetaSlug != "" {

		Entries.Slug = entriesrequired.SEODetails.MetaSlug

	}

	Entries.Status = entriesrequired.Status

	Entries.ChannelId = entriesrequired.ChannelId

	Entries.CategoriesId = entriesrequired.CategoryIds
	
	Entries.MemebrshipLevelIds = entriesrequired.MembershipLevelids

	Entries.CreatedBy = entriesrequired.CreatedBy

	Entries.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Entries.ModifiedOn = Entries.CreatedOn

	Entries.Author = entriesrequired.Author

	Entries.CreateTime = entriesrequired.CreateTime

	Entries.PublishedTime = entriesrequired.PublishTime

	Entries.ReadingTime = entriesrequired.ReadingTime

	Entries.SortOrder = entriesrequired.SortOrder

	Entries.Tags = entriesrequired.Tag

	Entries.Excerpt = entriesrequired.Excerpt

	Entries.ImageAltTag = entriesrequired.SEODetails.ImageAltTag

	Entries.OrderIndex = 1

	Entries.TenantId = tenantid

	Entries.CtaId = entriesrequired.CtaId

	Entriess, err := EntryModel.CreateChannelEntry(Entries, channel.DB)

	if err != nil {

		fmt.Println(err)
	}

	return Entriess, true, nil
}

func (channel *Channel) CreateChannelEntryFields(entryid int, createdby int, AdditionalFields []AdditionalFields, tenantid int) error {

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
		Entrfield.TenantId = tenantid

		EntriesField = append(EntriesField, Entrfield)

	}

	ferr := EntryModel.CreateEntrychannelFields(&EntriesField, channel.DB)

	if ferr != nil {

		return ferr
	}

	return nil

}

// update entry
func (channel *Channel) UpdateEntry(entriesrequired EntriesRequired, ChannelName string, EntryId int, tenantid int) (bool, error) {

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

	Entries.MemebrshipLevelIds=entriesrequired.MembershipLevelids

	// Entries.Slug = strings.ReplaceAll(strings.ToLower(entriesrequired.Title), " ", "-")

	// re := regexp.MustCompile(`[^a-z0-9\-]`)

	// Entries.Slug = re.ReplaceAllString(Entries.Slug, "-")

	// Entries.Slug = regexp.MustCompile(`-+`).ReplaceAllString(Entries.Slug, "-")

	// Entries.Slug = strings.Trim(Entries.Slug, "-")

	Entries.Slug = entriesrequired.SEODetails.MetaSlug

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

	Entries.OrderIndex = entriesrequired.OrderIndex

	Entries.CtaId = entriesrequired.CtaId

	err := EntryModel.UpdateChannelEntryDetails(&Entries, EntryId, channel.DB, tenantid)

	if err != nil {

		return false, err
	}

	return true, nil
}

// update entry additional fields
func (channel *Channel) UpdateAdditionalField(AdditionalFields []AdditionalFields, EntryId int, tenantid int) (bool, error) {

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

			EntryModel.UpdateChannelEntryAdditionalDetails(Entrfield, channel.DB, tenantid)

		}

	}

	return true, nil

}

/**/
func (channel *Channel) DeleteEntry(ChannelName string, modifiedby int, Entryid int, tenantid int) (bool, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return false, autherr
	}

	var entries Tblchannelentries

	entries.Id = Entryid

	entries.IsDeleted = 1

	entries.DeletedBy = modifiedby

	entries.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := EntryModel.DeleteChannelEntryId(&entries, Entryid, channel.DB, tenantid)

	var field TblChannelEntryField

	field.DeletedBy = modifiedby

	field.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := EntryModel.DeleteChannelEntryFieldId(&field, Entryid, channel.DB, tenantid)

	if err != nil {

		fmt.Println(err)
	}

	if err1 != nil {

		fmt.Println(err)
	}

	return true, nil

}

// Makefeature helps to highlights entry, only one entry should be featured of each channel and that is also optional
func (channel *Channel) MakeFeatureEntry(channelid, entryid, status int, tenantid int) (flag bool, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return false, autherr
	}

	merr := EntryModel.MakeFeature(channelid, entryid, status, channel.DB, tenantid)

	if merr != nil {

		return false, merr
	}

	return true, nil

}

// change entries status
func (channel *Channel) EntryStatus(ChannelName string, EntryId int, status int, modifiedby int, tenantid int) (bool, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return false, autherr
	}

	var Entries TblChannelEntries

	Entries.Status = status

	Entries.ModifiedBy = modifiedby

	Entries.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	EntryModel.PublishQuery(&Entries, EntryId, channel.DB, tenantid)

	return true, nil

}

// make feature function
func (channel *Channel) MakeFeature(channelid, entryid, status int, tenantid int) (flag bool, err error) {

	merr := CH.MakeFeature(channelid, entryid, status, channel.DB, tenantid)

	if merr != nil {

		return false, merr
	}

	return true, nil

}

// MULTI SELECT ENTRY DELETE FUNCTION//
func (channel *Channel) DeleteSelectedEntry(Entryid []int, modifiedby int, tenantid int) (bool, error) {

	var entries TblChannelEntries

	entries.IsDeleted = 1

	entries.DeletedBy = modifiedby

	entries.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := EntryModel.DeleteSelectedChannelEntryId(&entries, Entryid, channel.DB, tenantid)

	var field TblChannelEntryField

	field.DeletedBy = modifiedby

	field.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := EntryModel.DeleteSelectedChannelEntryFieldId(&field, Entryid, channel.DB, tenantid)

	if err != nil {

		fmt.Println(err)

		return false, err

	}

	if err1 != nil {

		fmt.Println(err)

		return false, err1
	}

	return true, nil

}

// MULTI SELECTE ENTRY UNPUBLISHED FUNCTION//
func (channel *Channel) UnpublishSelectedEntry(entryid []int, status int, modifiedby int, tenantid int) (bool, error) {

	autherr := AuthandPermission(channel)
	if autherr != nil {
		return false, autherr
	}

	var entries TblChannelEntries
	entries.Status = status
	entries.ModifiedBy = modifiedby
	entries.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	err := EntryModel.UnpublishSelectedChannelEntryId(&entries, entryid, channel.DB, tenantid)

	if err != nil {
		return false, err
	}

	return true, nil

}

// Entry Preview
func (channel *Channel) EntryPreview(uuid string) (Entry TblChannelEntries, flg bool, err error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return TblChannelEntries{}, false, autherr
	}

	var entry TblChannelEntries

	EntryModel.GetPreview(&entry, channel.DB, uuid)

	return entry, true, nil

}

// Entry is_active
func (channel *Channel) EntryIsActive(entryid int, status int, modifiedby int, tenantid int) (bool, error) {

	if AuthErr := AuthandPermission(channel); AuthErr != nil {
		return false, AuthErr
	}

	var entryisactive Tblchannelentries
	entryisactive.ModifiedBy = modifiedby
	entryisactive.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := EntryModel.EntryIsActive(entryisactive, entryid, status, channel.DB, tenantid)
	if err != nil {
		return false, err
	}

	return true, nil

}

// update channel entry view count
func (channel *Channel) UpdateChannelEntryViewCount(id int, slug string, tenantId int) (int, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return 0, autherr
	}

	var viewCount int

	err := EntryModel.UpdateEntryViewCount(channel.DB, id, slug, tenantId, &viewCount)

	if err != nil {

		return int(viewCount), err
	}

	return int(viewCount), nil

}

// fetch channel netry details
func (channel *Channel) FetchChannelEntryDetail(inputs EntriesInputs, multiFetchIds []int) (Tblchannelentries, []Tblchannelentries, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return Tblchannelentries{}, []Tblchannelentries{}, autherr
	}

	var (
		data           JoinEntries
		multiData      []JoinEntries
		channelEntry   Tblchannelentries
		channelEntries []Tblchannelentries
	)

	err := EntryModel.FlexibleChannelEntryDetail(channel.DB, inputs, multiFetchIds, &data, &multiData)

	if err != nil {

		return Tblchannelentries{}, []Tblchannelentries{}, err
	}

	switch {

	case len(multiData) > 0:

		for _, data := range multiData {

			var memberProfile member.TblMemberProfile

			if inputs.GetMemberProfile {

				memberProfile = member.TblMemberProfile{
					Id:              data.ProfileId,
					MemberId:        data.MemberID,
					ProfilePage:     data.ProfilePage,
					ProfileName:     data.ProfileName,
					ProfileSlug:     data.ProfileSlug,
					CompanyLogo:     data.CompanyLogo,
					StorageType:     data.ProfStorageType,
					CompanyName:     data.CompanyName,
					CompanyLocation: data.CompanyLocation,
					About:           data.About,
					Linkedin:        data.Linkedin,
					Website:         data.Website,
					Twitter:         data.Twitter,
					SeoTitle:        data.SeoTitle,
					SeoDescription:  data.SeoDescription,
					SeoKeyword:      data.SeoKeyword,
					MemberDetails:   data.MemberDetails,
					ClaimStatus:     data.ClaimStatus,
					CreatedBy:       data.ProfCreatedBy,
					CreatedOn:       data.ProfCreatedOn,
					ModifiedBy:      data.ProfModifiedBy,
					ModifiedOn:      data.ProfModifiedOn,
					IsDeleted:       data.ProfIsDeleted,
					DeletedOn:       data.ProfDeletedOn,
					DeletedBy:       data.ProfDeletedBy,
					ClaimDate:       data.ClaimDate,
					TenantId:        data.ProfTenantId,
				}
			}

			var authorDetails team.TblUser

			if inputs.GetAuthorDetails {

				authorDetails = team.TblUser{
					Id:                data.AuthorId,
					FirstName:         data.FirstName,
					LastName:          data.LastName,
					RoleId:            data.RoleId,
					Email:             data.Email,
					Username:          data.Username,
					MobileNo:          data.MobileNo,
					IsActive:          data.AuthorActive,
					ProfileImage:      data.ProfileImage,
					ProfileImagePath:  data.ProfileImagePath,
					StorageType:       data.AuthorStorageType,
					DataAccess:        data.DataAccess,
					CreatedOn:         data.AuthorCreatedOn,
					CreatedBy:         data.AuthorCreatedBy,
					ModifiedOn:        data.AuthorModifiedOn,
					ModifiedBy:        data.AuthorModifiedBy,
					LastLogin:         data.LastLogin,
					IsDeleted:         data.AuthorIsDeleted,
					DeletedOn:         data.AuthorDeletedOn,
					DeletedBy:         data.AuthorDeletedBy,
					DefaultLanguageId: data.DefaultLanguageId,
					TenantId:          data.UserTenantId,
					RoleName:          data.RoleName,
				}

			}

			var categoryHierarchy [][]categories.TblCategories

			if inputs.GetLinkedCategories && data.CategoriesID != "" {

				var categoriez []categories.TblCategories

				splitArr := strings.Split(data.CategoriesID, ",")

				categories.Categorymodel.GetHierarchicalCategoriesMappedInEntries(splitArr, &categoriez, channel.DB, channel.DB.Dialector.Name())

				for _, mapId := range splitArr {

					IntId, _ := strconv.Atoi(mapId)

					var categoryStream []categories.TblCategories

					for _, category := range categoriez {

						if category.Id == IntId {

							parentId := category.ParentId

							categoryStream = append(categoryStream, category)

						LOOPM:

							for _, parent := range categoriez {

								if parentId == parent.Id {

									parentId = parent.ParentId

									categoryStream = append(categoryStream, parent)

									if parent.ParentId != 0 {

										goto LOOPM

									} else {

										break
									}
								}
							}
						}
					}

					categoryHierarchy = append(categoryHierarchy, categoryStream)

				}

			}

			var sections, fields []tblfield

			if inputs.GetAdditionalFields {

				additionalFields, _ := EntryModel.GetChannelAdditionalFields(channel.DB, data.ChannelID)

				for _, field := range additionalFields {

					if field.FieldTypeId != inputs.SectionFieldTypeId {

						if field.OptionExist == 1 {

							field.FieldOptions, _ = EntryModel.GetFieldOptions(channel.DB, field.Id, inputs.TenantId)
						}

						field.FieldValue, _ = EntryModel.GetFieldValue(channel.DB, field.Id, data.Id, inputs.TenantId)

						fields = append(fields, field)

					} else {

						sections = append(sections, field)
					}
				}

			}

			channnel_entry := Tblchannelentries{
				Id:              data.Id,
				Title:           data.Title,
				Slug:            data.Slug,
				Description:     data.Description,
				UserId:          data.UserID,
				ChannelId:       data.ChannelID,
				Status:          data.Status,
				IsActive:        data.IsActive,
				CreatedOn:       data.CreatedOn,
				CreatedBy:       data.CreatedBy,
				ModifiedBy:      data.ModifiedBy,
				ModifiedOn:      data.ModifiedOn,
				CoverImage:      data.CoverImage,
				ThumbnailImage:  data.ThumbnailImage,
				PublishedTime:   data.PublishedTime,
				MetaDescription: data.MetaDescription,
				MetaTitle:       data.MetaTitle,
				Keyword:         data.Keyword,
				ImageAltTag:     data.ImageAltTag,
				CategoriesId:    data.CategoriesID,
				RelatedArticles: data.RelatedArticles,
				Feature:         data.Feature,
				ViewCount:       data.ViewCount,
				Author:          data.Author,
				SortOrder:       data.SortOrder,
				CreateTime:      data.CreateTime,
				ReadingTime:     data.ReadingTime,
				Tags:            data.Tags,
				Excerpt:         data.Excerpt,
				IsDeleted:       data.IsDeleted,
				DeletedOn:       data.DeletedOn,
				DeletedBy:       data.DeletedBy,
				AuthorDetail:    authorDetails,
				TenantId:        data.TenantId,
				MemberProfiles:  memberProfile,
				Categories:      categoryHierarchy,
				Sections:        sections,
				Fields:          fields,
			}

			channelEntries = append(channelEntries, channnel_entry)
		}

	case data.Id != 0:

		var memberProfile member.TblMemberProfile

		if inputs.GetMemberProfile {

			memberProfile = member.TblMemberProfile{
				Id:              data.ProfileId,
				MemberId:        data.MemberID,
				ProfilePage:     data.ProfilePage,
				ProfileName:     data.ProfileName,
				ProfileSlug:     data.ProfileSlug,
				CompanyLogo:     data.CompanyLogo,
				StorageType:     data.ProfStorageType,
				CompanyName:     data.CompanyName,
				CompanyLocation: data.CompanyLocation,
				About:           data.About,
				Linkedin:        data.Linkedin,
				Website:         data.Website,
				Twitter:         data.Twitter,
				SeoTitle:        data.SeoTitle,
				SeoDescription:  data.SeoDescription,
				SeoKeyword:      data.SeoKeyword,
				MemberDetails:   data.MemberDetails,
				ClaimStatus:     data.ClaimStatus,
				CreatedBy:       data.ProfCreatedBy,
				CreatedOn:       data.ProfCreatedOn,
				ModifiedBy:      data.ProfModifiedBy,
				ModifiedOn:      data.ProfModifiedOn,
				IsDeleted:       data.ProfIsDeleted,
				DeletedOn:       data.ProfDeletedOn,
				DeletedBy:       data.ProfDeletedBy,
				ClaimDate:       data.ClaimDate,
				TenantId:        data.ProfTenantId,
			}
		}

		var authorDetails team.TblUser

		if inputs.GetAuthorDetails {

			authorDetails = team.TblUser{
				Id:                data.AuthorId,
				FirstName:         data.FirstName,
				LastName:          data.LastName,
				RoleId:            data.RoleId,
				Email:             data.Email,
				Username:          data.Username,
				MobileNo:          data.MobileNo,
				IsActive:          data.AuthorActive,
				ProfileImage:      data.ProfileImage,
				ProfileImagePath:  data.ProfileImagePath,
				StorageType:       data.AuthorStorageType,
				DataAccess:        data.DataAccess,
				CreatedOn:         data.AuthorCreatedOn,
				CreatedBy:         data.AuthorCreatedBy,
				ModifiedOn:        data.AuthorModifiedOn,
				ModifiedBy:        data.AuthorModifiedBy,
				LastLogin:         data.LastLogin,
				IsDeleted:         data.AuthorIsDeleted,
				DeletedOn:         data.AuthorDeletedOn,
				DeletedBy:         data.AuthorDeletedBy,
				DefaultLanguageId: data.DefaultLanguageId,
				TenantId:          data.UserTenantId,
				RoleName:          data.RoleName,
			}

		}

		var categoryHierarchy [][]categories.TblCategories

		if inputs.GetLinkedCategories && data.CategoriesID != "" {

			var categoriez []categories.TblCategories

			splitArr := strings.Split(data.CategoriesID, ",")

			categories.Categorymodel.GetHierarchicalCategoriesMappedInEntries(splitArr, &categoriez, channel.DB, channel.DB.Dialector.Name())

			for _, mapId := range splitArr {

				IntId, _ := strconv.Atoi(mapId)

				var categoryStream []categories.TblCategories

				for _, category := range categoriez {

					if category.Id == IntId {

						parentId := category.ParentId

						categoryStream = append(categoryStream, category)

					LOOP:

						for _, parent := range categoriez {

							if parentId == parent.Id {

								parentId = parent.ParentId

								categoryStream = append(categoryStream, parent)

								if parent.ParentId != 0 {

									goto LOOP

								} else {

									break
								}
							}
						}
					}
				}

				categoryHierarchy = append(categoryHierarchy, categoryStream)

			}

		}

		var sections, fields []tblfield

		if inputs.GetAdditionalFields {

			additionalFields, _ := EntryModel.GetChannelAdditionalFields(channel.DB, data.ChannelID)

			for _, field := range additionalFields {

				if field.FieldTypeId != inputs.SectionFieldTypeId {

					if field.OptionExist == 1 {

						field.FieldOptions, _ = EntryModel.GetFieldOptions(channel.DB, field.Id, inputs.TenantId)
					}

					field.FieldValue, _ = EntryModel.GetFieldValue(channel.DB, field.Id, data.Id, inputs.TenantId)

					fields = append(fields, field)

				} else {

					sections = append(sections, field)
				}
			}

		}

		channelEntry = Tblchannelentries{
			Id:              data.Id,
			Title:           data.Title,
			Slug:            data.Slug,
			Description:     data.Description,
			UserId:          data.UserID,
			ChannelId:       data.ChannelID,
			Status:          data.Status,
			IsActive:        data.IsActive,
			CreatedOn:       data.CreatedOn,
			CreatedBy:       data.CreatedBy,
			ModifiedBy:      data.ModifiedBy,
			ModifiedOn:      data.ModifiedOn,
			CoverImage:      data.CoverImage,
			ThumbnailImage:  data.ThumbnailImage,
			PublishedTime:   data.PublishedTime,
			MetaDescription: data.MetaDescription,
			MetaTitle:       data.MetaTitle,
			Keyword:         data.Keyword,
			ImageAltTag:     data.ImageAltTag,
			CategoriesId:    data.CategoriesID,
			RelatedArticles: data.RelatedArticles,
			Feature:         data.Feature,
			ViewCount:       data.ViewCount,
			Author:          data.Author,
			SortOrder:       data.SortOrder,
			CreateTime:      data.CreateTime,
			ReadingTime:     data.ReadingTime,
			Tags:            data.Tags,
			Excerpt:         data.Excerpt,
			IsDeleted:       data.IsDeleted,
			DeletedOn:       data.DeletedOn,
			DeletedBy:       data.DeletedBy,
			AuthorDetail:    authorDetails,
			MemberProfiles:  memberProfile,
			Categories:      categoryHierarchy,
			Sections:        sections,
			Fields:          fields,
			TenantId:        data.TenantId,
		}

	}

	return channelEntry, channelEntries, nil
}

func (channel *Channel) EntryParentIdUpdate(parentid, entriid, tenantid, userid int) error {
	autherr := AuthandPermission(channel)
	if autherr != nil {
		return autherr
	}

	var entries Tblchannelentries
	entries.ParentId = parentid
	entries.ModifiedBy = userid
	entries.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	entries.Id = entriid
	entries.TenantId = tenantid
	err := EntryModel.EntryParentIdUpdate(&entries, channel.DB)

	if err != nil {

		fmt.Println(err)
	}

	return nil

}

func (channel *Channel) EntrylistByParentId(entryid int, tenantid int) ([]Tblchannelentries, error) {

	autherr := AuthandPermission(channel)
	if autherr != nil {
		return []Tblchannelentries{}, autherr
	}

	var entries []Tblchannelentries
	err := EntryModel.EntrylistByParentId(&entries, entryid, channel.DB, tenantid)

	if err != nil {

		fmt.Println(err)
	}

	return entries, nil
}

// Update entry Reorder
func (channel *Channel) UpdateEntryOrder(Entryids []int, tenantid, userid, offset int) error {
	autherr := AuthandPermission(channel)
	if autherr != nil {
		return autherr
	}

	var entries Tblchannelentries

	entries.ModifiedBy = userid
	entries.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	entries.TenantId = tenantid

	for index, id := range Entryids {

		if id != 0 {

			entries.Id = id

			entries.OrderIndex = index + offset

			err := EntryModel.UpdateEntryOrder(&entries, channel.DB)

			if err != nil {

				return err
			}
		}

	}

	return nil
}

// Update  Access Permission MemberGroupIds in Entry
func (channel *Channel) UpdateMemberGroupIds(membergrbids string, entryid int, tenantid int, userid int) error {

	autherr := AuthandPermission(channel)
	if autherr != nil {
		return autherr
	}

	var entries Tblchannelentries

	entries.ModifiedBy = userid
	entries.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	entries.TenantId = tenantid
	entries.MembergroupId = membergrbids
	entries.Id = entryid

	err := EntryModel.UpdateEntryMemberGroupIds(&entries, channel.DB)

	if err != nil {

		return err
	}

	return nil
}

// update entries orderindex
func (channel *Channel) UpdateEntryOrderIndex(Ordenindex int, EntryId int, userid int, tenantid int) (bool, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return false, autherr
	}

	var Entries TblChannelEntries

	Entries.OrderIndex = Ordenindex

	Entries.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Entries.ModifiedBy = userid

	err := EntryModel.UpdateEntryOrderIndex(&Entries, EntryId, channel.DB, tenantid)

	if err != nil {

		return false, err
	}

	return true, nil
}
func (channel *Channel) GetChannelEntriesById(id int, tenantid int) (*TblChannelEntries, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return nil, autherr
	}

	entries, err := EntryModel.GetEntriesById(id, channel.DB, tenantid)
	if err != nil {
		fmt.Println(err)
	}

	return entries, nil

}
func (channel *Channel) DefaultChannel(slug string, tenantid int) (int, error) {
	autherr := AuthandPermission(channel)

	if autherr != nil {

		return 0, autherr
	}

	id, err := EntryModel.defaultchannelid(slug, channel.DB, tenantid)
	if err != nil {
		fmt.Println(err)
	}

	return id, nil
}

func (channel *Channel) EntryAuthors(tenantid int) ([]Author, error) {

	autherr := AuthandPermission(channel)

	if autherr != nil {

		return []Author{}, autherr
	}

	Authors, err := EntryModel.EntryAuthors(tenantid, channel.DB)
	if err != nil {
		fmt.Println(err)
	}

	return Authors, nil

}
