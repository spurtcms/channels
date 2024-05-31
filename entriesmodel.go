package channels

import (
	"strconv"
	"time"

	"github.com/spurtcms/categories"
	"github.com/spurtcms/member"
	access "github.com/spurtcms/member-access"
	"github.com/spurtcms/team"
	"gorm.io/gorm"
)

type Tblchannelentries struct {
	Id                   int                          `gorm:"column:id"`
	Title                string                       `gorm:"column:title"`
	Slug                 string                       `gorm:"column:slug"`
	Description          string                       `gorm:"column:description"`
	UserId               int                          `gorm:"column:user_id"`
	ChannelId            int                          `gorm:"column:channel_id"`
	Status               int                          `gorm:"column:status"` //0-draft 1-publish 2-unpublish
	IsActive             int                          `gorm:"column:is_active"`
	IsDeleted            int                          `gorm:"column:is_deleted;DEFAULT:0"`
	DeletedBy            int                          `gorm:"column:deleted_by;DEFAULT:NULL"`
	DeletedOn            time.Time                    `gorm:"column:deleted_on;DEFAULT:NULL"`
	CreatedOn            time.Time                    `gorm:"column:created_on"`
	CreatedBy            int                          `gorm:"column:created_by"`
	ModifiedBy           int                          `gorm:"column:modified_by;DEFAULT:NULL"`
	ModifiedOn           time.Time                    `gorm:"column:modified_on;DEFAULT:NULL"`
	CoverImage           string                       `gorm:"column:cover_image"`
	ThumbnailImage       string                       `gorm:"column:thumbnail_image"`
	MetaTitle            string                       `gorm:"column:meta_title"`
	MetaDescription      string                       `gorm:"column:meta_description"`
	Keyword              string                       `gorm:"column:keyword"`
	CategoriesId         string                       `gorm:"column:categories_id"`
	RelatedArticles      string                       `gorm:"column:related_articles"`
	CreatedDate          string                       `gorm:"-;<-false"`
	ModifiedDate         string                       `gorm:"-;<-false"`
	Username             string                       `gorm:"<-:false"`
	CategoryGroup        string                       `gorm:"-:migration;<-:false"`
	ChannelName          string                       `gorm:"-:migration;<-:false"`
	Cno                  string                       `gorm:"<-:false"`
	ProfileImagePath     string                       `gorm:"<-:false"`
	EntryStatus          string                       `gorm:"-"`
	TblChannelEntryField []TblChannelEntryField       `gorm:"<-:false; foreignKey:ChannelEntryId"`
	Categories           [][]categories.TblCategories `gorm:"-"`
	AdditionalData       string                       `gorm:"-"`
	AuthorDetail         Author                       `gorm:"-"`
	Sections             []tblfield                   `gorm:"-"`
	Fields               []tblfield                   `gorm:"-"`
	MemberProfiles       member.TblMemberProfile      `gorm:"-"`
	Feature              int                          `gorm:"column:feature;DEFAULT:0"`
	ViewCount            int                          `gorm:"column:view_count;DEFAULT:0"`
	Author               string                       `gorm:"column:author"`
	SortOrder            int                          `gorm:"column:sort_order"`
	CreateTime           time.Time                    `gorm:"column:created_date"`
	PublishedTime        time.Time                    `gorm:"column:published_time;default:null"`
	ReadingTime          int                          `gorm:"column:reading_time;DEFAULT:0"`
	Tags                 string                       `gorm:"column:tags"`
	Excerpt              string                       `gorm:"column:excerpt"`
	ImageAltTag          string                       `gorm:"column:image_alt_tag"`
}

type Author struct {
	AuthorID         int       `json:"AuthorId" gorm:"column:id"`
	FirstName        string    `json:"FirstName"`
	LastName         string    `json:"LastName"`
	Email            string    `json:"Email"`
	MobileNo         *string   `json:"MobileNo,omitempty"`
	IsActive         *int      `json:"IsActive,omitempty"`
	ProfileImage     *string   `json:"ProfileImage,omitempty"`
	ProfileImagePath *string   `json:"ProfileImagePath,omitempty"`
	CreatedOn        time.Time `json:"CreatedOn"`
	CreatedBy        int       `json:"CreatedBy"`
}

type Entries struct {
	ChannelId                int //if pass the channelid it will return that particular channel entries only otherwise return all
	Limit                    int
	Offset                   int
	Keyword                  string //filter
	Title                    string //filter
	ChannelName              string //filter
	Status                   string //filter
	UserName                 string //filter
	CategoryId               int    //filter
	CategoryName             string //filter
	SelectedCategoryFilter   bool   //selected category filter or selected category child filter also
	Publishedonly            bool   //if this will be enable published entries only show
	ActiveChannelEntriesonly bool   //if this will be enable active channel entries only show
	MemberProfile            bool   //if you want member profile pls enable memberprofile true
	AdditionalFields         bool   //if you want additionalfields pls enable additionalfields true
	AuthorDetails            bool   //if you want authordetails pls enable authordetails true
	ContentHide              bool   //if you want hide content only for memberaccesscontrol enable true otherwise it doesn't fetch the entry
	MemberAccessControl      bool
	MemberId                 int
	ImageUrlPath             string
	FieldTypeId              int
	MemberFieldTypeId        int
}

type IndivEntriesReq struct {
	ChannelName       string
	EntryId           int
	MemberProfile     bool //if you want member profile pls enable memberprofile true
	AdditionalFields  bool //if you want additionalfields pls enable additionalfields true
	AuthorDetails     bool //if you want authordetails pls enable authordetails true
	ContentHide       bool //if you want show entries name without content enable true, ensure want hide content only must be restricted in memberaccess
	CategoriesEnable  bool //
	MemberId          int
	ImageUrlPath      string
	FieldTypeId       int
	MemberFieldTypeId int
}

type SEODetails struct {
	MetaTitle       string
	MetaDescription string
	MetaKeywords    string
	MetaSlug        string
	ImageAltTag     string
}

type AdditionalFields struct {
	Id            int
	FieldName     string
	FieldValue    string
	FieldId       int
	MultipleValue []string
	ModifiedBy    int
}

type EntriesRequired struct {
	Title       string
	Content     string
	CoverImage  string
	SEODetails  SEODetails
	CategoryIds string
	ChannelName string
	Status      int
	ChannelId   int
	Author      string
	CreateTime  time.Time
	PublishTime time.Time
	ReadingTime int
	SortOrder   int
	Tag         string
	Excerpt     string
	CreatedBy   int
	ModifiedBy  int
}

type RecentActivities struct {
	Contenttype string
	Title       string
	User        string
	Imagepath   string
	Createdon   time.Time
	Active      string
	Channelname string
}

type EntriesModel struct{}

var EntryModel EntriesModel

/*List Channel Entry*/
func (Ch EntriesModel) ChannelEntryList(filter Entries, channel *Channel, categoryid string, DB *gorm.DB) (chentry []Tblchannelentries, chentcount int64, err error) {

	query := DB.Model(TblChannelEntries{}).Select("tbl_channel_entries.*,tbl_users.username,tbl_channels.channel_name").Joins("inner join tbl_users on tbl_users.id = tbl_channel_entries.created_by").Joins("inner join tbl_channels on tbl_channels.id = tbl_channel_entries.channel_id").Where("tbl_channel_entries.is_deleted=0").Order("id desc")

	if channel.PermissionEnable && channel.Auth.RoleId != 1 {

		query = query.Where("channel_id in (select id from tbl_channels where channel_name in (select display_name from tbl_module_permissions inner join tbl_modules on tbl_modules.id = tbl_module_permissions.module_id inner join tbl_role_permissions on tbl_role_permissions.permission_id = tbl_module_permissions.id where role_id =(?) and tbl_modules.module_name='Entries' )) ", channel.Auth.RoleId)

	}

	if filter.ActiveChannelEntriesonly {

		query = query.Where("tbl_channels.is_active =1")
	}

	if filter.Publishedonly {

		query = query.Where("tbl_channel_entries.status=1")

	}

	if filter.ChannelId != 0 {

		query = query.Where("tbl_channel_entries.channel_id=?", filter.ChannelId)
	}

	if filter.UserName != "" {

		query = query.Where("LOWER(TRIM(tbl_users.username)) LIKE LOWER(TRIM(?))", "%"+filter.UserName+"%")

	}

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?)) OR LOWER(TRIM(channel_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

	}

	if filter.Status != "" {

		query = query.Where("tbl_channel_entries.status=?", filter.Status)

	}
	if filter.Title != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?))", "%"+filter.Title+"%")

	}

	if filter.ChannelName != "" {

		query = query.Where("LOWER(TRIM(channel_name)) LIKE LOWER(TRIM(?))", "%"+filter.ChannelName+"%")

	}

	if filter.CategoryId != 0 && filter.CategoryId > 0 {

		query = query.Where("STRING_TO_ARRAY(categories_id, ',')::integer[] && ARRAY[" + categoryid + "]")

	}

	if filter.MemberAccessControl {

		_, entryid := EntryModel.MemberAccessCheck(filter.MemberId, DB)

		if !filter.ContentHide {

			query = query.Where("id not in (?)", entryid)

		}

	}

	if filter.Limit != 0 {

		query.Limit(filter.Limit).Offset(filter.Offset).Order("id asc").Find(&chentry)

	} else {

		query.Find(&chentry).Count(&chentcount)

		return chentry, chentcount, nil
	}

	return chentry, 0, nil
}

func (ch EntriesModel) MemberAccessCheck(memberid int, DB *gorm.DB) ([]int, []int) {

	var channelid, entryid []int

	var mem member.TblMember

	//get membergroup id
	DB.Table("tbl_members").Select("member_group_id").Where("is_deleted=0 and id=?", memberid).First(&mem)

	SUB := `select id from tbl_access_control_user_groups where is_deleted=0 and member_group_id=` + strconv.Itoa(mem.Id)

	var accessgroup []access.TblAccessControlPages

	DB.Table("tbl_access_control_pages").Where("access_control_user_group_id in (?)", SUB).Find(&accessgroup)

	for _, val := range accessgroup {

		channelid = append(channelid, val.ChannelId)

		entryid = append(entryid, val.EntryId)

	}

	return channelid, entryid
}

/*Create channel entry*/
func (Ch EntriesModel) CreateChannelEntry(entry Tblchannelentries, DB *gorm.DB) (Tblchannelentries, error) {

	if err := DB.Table("tbl_channel_entries").Create(&entry).Error; err != nil {

		return Tblchannelentries{}, err

	}

	return entry, nil

}

/*create channel entry field*/
func (Ch EntriesModel) CreateEntrychannelFields(entryfield *[]TblChannelEntryField, DB *gorm.DB) error {

	if err := DB.Table("tbl_channel_entry_fields").Create(&entryfield).Error; err != nil {

		return err
	}

	return nil

}

/*Delete Channel Entry Field*/
func (Ch EntriesModel) DeleteChannelEntryId(chentry *Tblchannelentries, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_channel_entries").Where("id=?", chentry.Id).UpdateColumns(map[string]interface{}{"is_deleted": chentry.IsDeleted, "deleted_by": chentry.DeletedBy, "deleted_on": chentry.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*Delete Channel Entry Field*/
func (Ch EntriesModel) DeleteChannelEntryFieldId(chentry *TblChannelEntryField, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_channel_entry_fields").Where("channel_entry_id=?", id).UpdateColumns(map[string]interface{}{"deleted_by": chentry.DeletedBy, "deleted_on": chentry.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*Edit Channel Entry*/
func (Ch EntriesModel) GetChannelEntryById(ent IndivEntriesReq, DB *gorm.DB) (tblchanentry Tblchannelentries, err error) {

	query := DB.Table("tbl_channel_entries").Where("is_deleted=0 and id=?", ent.EntryId)

	if ent.ContentHide {

		_, entryid := EntryModel.MemberAccessCheck(ent.MemberId, DB)

		for _, val := range entryid {

			if val == ent.EntryId {

				query = query.Omit("tbl_channel_entries.description")

				break
			}
		}

	}

	query = query.Preload("TblChannelEntryField", func(db *gorm.DB) *gorm.DB {
		return db.Select("tbl_channel_entry_fields.*,tbl_fields.field_type_id").Joins("inner join tbl_fields on tbl_fields.Id = tbl_channel_entry_fields.field_id")
	})

	query.Find(&tblchanentry)

	if err := query.Error; err != nil {

		return Tblchannelentries{}, err

	}

	return tblchanentry, nil
}

func (Ch EntriesModel) GetAuthorDetails(DB *gorm.DB, authorId int) (authorDetail Author, err error) {

	if err := DB.Model(team.TblUser{}).Where("tbl_users.is_deleted = 0 and tbl_users.id = ?", authorId).First(&authorDetail).Error; err != nil {

		return Author{}, err
	}

	return authorDetail, nil
}

func (Ch EntriesModel) GetSectionsUnderEntries(DB *gorm.DB, channelId, sectionTypeId int) (sections []tblfield, err error) {

	if err = DB.Table("tbl_group_fields").Select("tbl_fields.*,tbl_field_types.type_name").Joins("inner join tbl_fields on tbl_fields.id = tbl_group_fields.field_id").Joins("inner join tbl_field_types on tbl_field_types.id = tbl_fields.field_type_id").
		Where("tbl_fields.is_deleted = 0 and tbl_field_types.is_deleted = 0 and tbl_fields.field_type_id = ? and tbl_group_fields.channel_id = ?", sectionTypeId, channelId).Find(&sections).Error; err != nil {

		return []tblfield{}, err
	}

	return sections, nil
}

func (Ch EntriesModel) GetFieldsInEntries(DB *gorm.DB, channelId, sectionTypeId int) (fields []tblfield, err error) {

	if err = DB.Table("tbl_group_fields").Select("tbl_fields.*,tbl_field_types.type_name").Joins("inner join tbl_fields on tbl_fields.id = tbl_group_fields.field_id").Joins("inner join tbl_field_types on tbl_field_types.id = tbl_fields.field_type_id").
		Where("tbl_fields.is_deleted = 0 and tbl_field_types.is_deleted = 0 and tbl_fields.field_type_id != ? and tbl_group_fields.channel_id = ?", sectionTypeId, channelId).Find(&fields).Error; err != nil {

		return []tblfield{}, err
	}

	return fields, nil
}

func (Ch EntriesModel) GetFieldValue(DB *gorm.DB, fieldId, entryId int) (fieldvalue TblChannelEntryField, err error) {

	if err = DB.Model(TblChannelEntryField{}).Where("tbl_channel_entry_fields.field_id = ? and tbl_channel_entry_fields.channel_entry_id = ?", fieldId, entryId).First(&fieldvalue).Error; err != nil {

		return TblChannelEntryField{}, err
	}

	return fieldvalue, nil
}

func (ch EntriesModel) GetFieldOptions(DB *gorm.DB, fieldId int) (fieldOptions []tblfieldoption, err error) {

	if err = DB.Model(TblFieldOption{}).Where("tbl_field_options.is_deleted = 0 and tbl_field_options.field_id = ?", fieldId).Find(&fieldOptions).Error; err != nil {

		return []tblfieldoption{}, err
	}

	return fieldOptions, nil
}

func (ch EntriesModel) GetMemberProfile(DB *gorm.DB, memberid int) (memberProfile member.TblMemberProfile, err error) {

	if err = DB.Model(member.TblMemberProfile{}).Select("tbl_member_profiles.*").Joins("inner join tbl_members on tbl_members.id = tbl_member_profiles.member_id").Where("tbl_members.is_deleted = 0 and tbl_members.id = ?", memberid).First(&memberProfile).Error; err != nil {

		return member.TblMemberProfile{}, err
	}

	return memberProfile, nil

}

func (ch EntriesModel) GetGraphqlEntriesCategoryByParentId(DB *gorm.DB, categoryId int) (category categories.TblCategories, err error) {

	if err = DB.Model(categories.TblCategories{}).Where("is_deleted = 0 and id = ?", categoryId).First(&category).Error; err != nil {

		return categories.TblCategories{}, err
	}

	return category, nil
}

func (ch EntriesModel) GetCategoryIdByName(name string, DB *gorm.DB) (category categories.TblCategories, er error) {

	if err := DB.Model(categories.TblCategories{}).Where("category_name= ?", name).First(&category).Error; err != nil {

		return category, err
	}

	return category, nil
}

func (ch EntriesModel) GetChildCategories(categoryid int, DB *gorm.DB) (categories []categories.TblCategories, er error) {

	if err := DB.Raw(`WITH RECURSIVE cat_tree AS (SELECT id, category_name, category_slug,image_path, parent_id,created_on,modified_on,is_deleted
		FROM tbl_categories
		WHERE id = (?)
		UNION
		SELECT cat.id, cat.category_name, cat.category_slug, cat.image_path ,cat.parent_id,cat.created_on,cat.modified_on,
		cat.is_deleted
		FROM tbl_categories AS cat
		JOIN cat_tree ON cat.parent_id = cat_tree.id )
		SELECT * FROM cat_tree where is_deleted = 0`, categoryid).Find(&categories).Error; err != nil {

		return categories, err
	}

	return categories, nil
}

func (ch EntriesModel) MakeFeature(channelid, entryid, status int, DB *gorm.DB) (err error) {

	DB.Model(TblChannelEntries{}).Where("channel_id=?", channelid).UpdateColumns(map[string]interface{}{"feature": 0})

	if err := DB.Model(TblChannelEntries{}).Where("id=? and channel_id=?", entryid, channelid).UpdateColumns(map[string]interface{}{"feature": status}).Error; err != nil {

		return err
	}

	return nil
}

func (ch EntriesModel) PublishQuery(chl *TblChannelEntries, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_channel_entries").Where("id =?", id).UpdateColumns(map[string]interface{}{"status": chl.Status, "modified_on": chl.ModifiedOn, "modified_by": chl.ModifiedBy}).Error; err != nil {

		return err

	}

	return nil
}

/*Update Channel Entry Details*/
func (Ch EntriesModel) UpdateChannelEntryDetails(entry *TblChannelEntries, entryid int, DB *gorm.DB) error {

	if err := DB.Debug().Table("tbl_channel_entries").Where("id=?", entryid).UpdateColumns(map[string]interface{}{"title": entry.Title, "description": entry.Description, "slug": entry.Slug, "cover_image": entry.CoverImage, "thumbnail_image": entry.ThumbnailImage, "meta_title": entry.MetaTitle, "meta_description": entry.MetaDescription, "keyword": entry.Keyword, "categories_id": entry.CategoriesId, "related_articles": entry.RelatedArticles, "status": entry.Status, "modified_on": entry.ModifiedOn, "modified_by": entry.ModifiedBy, "user_id": entry.UserId, "channel_id": entry.ChannelId, "author": entry.Author, "create_time": entry.CreateTime, "published_time": entry.PublishedTime, "reading_time": entry.ReadingTime, "sort_order": entry.SortOrder, "tags": entry.Tags, "excerpt": entry.Excerpt, "image_alt_tag": entry.ImageAltTag}).Error; err != nil {

		return err
	}

	return nil

}

/*create channel entry field*/
func (Ch EntriesModel) CreateSingleEntrychannelFields(entryfield *TblChannelEntryField, DB *gorm.DB) error {

	if err := DB.Table("tbl_channel_entry_fields").Create(&entryfield).Error; err != nil {

		return err
	}

	return nil

}

/*Update Channel Entry Details*/
func (Ch EntriesModel) UpdateChannelEntryAdditionalDetails(entry TblChannelEntryField, DB *gorm.DB) error {

	if err := DB.Table("tbl_channel_entry_fields").Where("id=?", entry.Id).UpdateColumns(map[string]interface{}{"field_name": entry.FieldName, "field_value": entry.FieldValue, "modified_by": entry.ModifiedBy, "modified_on": entry.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

func (Ch EntriesModel) AllentryCount(DB *gorm.DB) (count int64, err error) {

	if err := DB.Table("tbl_channel_entries").Where("is_deleted = 0 ").Count(&count).Error; err != nil {

		return 0, err
	}

	return count, nil
}

func (Ch EntriesModel) NewentryCount(DB *gorm.DB) (count int64, err error) {

	if err := DB.Table("tbl_channel_entries").Where("is_deleted = 0 AND created_on >=?", time.Now().AddDate(0, 0, -10)).Count(&count).Error; err != nil {

		return 0, err
	}

	return count, nil
}

func (Ch EntriesModel) Newchannels(DB *gorm.DB) (chn []Tblchannel, err error) {

	if err := DB.Table("tbl_channels").Select("tbl_channels.*,tbl_users.username,tbl_users.profile_image_path").
		Joins("inner join tbl_users on tbl_users.id = tbl_channels.created_by").
		Where("tbl_channels.is_deleted=0 and tbl_channels.is_active=1 and tbl_channels.created_on >= ?", time.Now().Add(-24*time.Hour).Format("2006-01-02 15:04:05")).
		Order("created_on desc").Limit(6).Find(&chn).Error; err != nil {

		return []Tblchannel{}, err
	}

	return chn, nil

}

func (Ch EntriesModel) Newentries(DB *gorm.DB) (entries []Tblchannelentries, err error) {

	if err := DB.Table("tbl_channel_entries").Select("tbl_channel_entries.*,tbl_users.username,tbl_users.profile_image_path").
		Joins("inner join tbl_users on tbl_users.id = tbl_channel_entries.created_by").Where("tbl_channel_entries.is_deleted=0 and tbl_channel_entries.created_on >=?", time.Now().Add(-24*time.Hour).Format("2006-01-02 15:04:05")).
		Order("created_on desc").Limit(6).Find(&entries).Error; err != nil {

		return []Tblchannelentries{}, err
	}

	return entries, nil

}

// update imagepath
func (Ch EntriesModel) UpdateImagePath(Imagepath string, DB *gorm.DB) error {

	if err := DB.Model(TblChannelEntries{}).Where("cover_image=?", Imagepath).UpdateColumns(map[string]interface{}{
		"cover_image": ""}).Error; err != nil {

		return err
	}

	return nil

}

// make feature function
func (ch ChannelModel) MakeFeature(channelid, entryid, status int, DB *gorm.DB) (err error) {

	DB.Model(TblChannelEntries{}).Where("channel_id=?", channelid).UpdateColumns(map[string]interface{}{"feature": 0})

	if err := DB.Model(TblChannelEntries{}).Where("id=? and channel_id=?", entryid, channelid).UpdateColumns(map[string]interface{}{"feature": status}).Error; err != nil {

		return err
	}

	return nil
}

/*Delete MULTI Channel Entry Field*/
func (Ch EntriesModel) DeleteSelectedChannelEntryId(chentry *TblChannelEntries, id []int, DB *gorm.DB) error {

	if err := DB.Debug().Table("tbl_channel_entries").Where("id in (?)", id).UpdateColumns(map[string]interface{}{"is_deleted": chentry.IsDeleted, "deleted_by": chentry.DeletedBy, "deleted_on": chentry.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*Delete MULTI Channel Entry Field*/
func (Ch EntriesModel) DeleteSelectedChannelEntryFieldId(chentry *TblChannelEntryField, id []int, DB *gorm.DB) error {

	if err := DB.Debug().Table("tbl_channel_entry_fields").Where("channel_entry_id IN (?)", id).UpdateColumns(map[string]interface{}{"deleted_by": chentry.DeletedBy, "deleted_on": chentry.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}
