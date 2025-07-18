package channels

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spurtcms/categories"
	"github.com/spurtcms/member"
	access "github.com/spurtcms/member-access"
	"github.com/spurtcms/team"
	"gorm.io/datatypes"
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
	MemebrshipLevelIds   string                       `gorm:"column:memebrship_level_ids"`
	RelatedArticles      string                       `gorm:"column:related_articles"`
	CreatedDate          string                       `gorm:"-:migration;<-:false"`
	ModifiedDate         string                       `gorm:"-:migration;<-:false"`
	Username             string                       `gorm:"<-:false"`
	CategoryGroup        string                       `gorm:"-:migration;<-:false"`
	ChannelName          string                       `gorm:"-:migration;<-:false"`
	Cno                  string                       `gorm:"<-:false"`
	ProfileImagePath     string                       `gorm:"<-:false"`
	EntryStatus          string                       `gorm:"-"`
	TblChannelEntryField []TblChannelEntryField       `gorm:"<-:false; foreignKey:ChannelEntryId"`
	Categories           [][]categories.TblCategories `gorm:"-"`
	AdditionalData       string                       `gorm:"-"`
	AuthorDetail         team.TblUser                 `gorm:"-"`
	Sections             []tblfield                   `gorm:"-"`
	Fields               []tblfield                   `gorm:"-"`
	MemberProfiles       member.TblMemberProfile      `gorm:"-"`
	Feature              int                          `gorm:"column:feature;DEFAULT:0"`
	ViewCount            int                          `gorm:"column:view_count;DEFAULT:0"`
	Author               string                       `gorm:"column:author"`
	SortOrder            int                          `gorm:"column:sort_order"`
	CreateTime           time.Time                    `gorm:"column:create_time;default:null"`
	PublishedTime        time.Time                    `gorm:"column:published_time;default:null"`
	ReadingTime          int                          `gorm:"column:reading_time;DEFAULT:0"`
	Tags                 string                       `gorm:"column:tags"`
	Excerpt              string                       `gorm:"column:excerpt"`
	ImageAltTag          string                       `gorm:"column:image_alt_tag"`
	TenantId             string                       `gorm:"column:tenant_id"`
	Uuid                 string                       `gorm:"column:uuid"`
	ParentId             int                          `gorm:"column:parent_id"`
	ChildrenList         []Tblchannelentries          `gorm:"-"`
	OrderIndex           int                          `gorm:"column:order_index"`
	MembergroupId        string                       `gorm:"type:membergroup_id"`
	FirstName            string                       `gorm:"<-:false"`
	LastName             string                       `gorm:"<-:false"`
	NameString           string                       `gorm:"<-:false"`
	CtaId                int                          `gorm:"column:cta_id"`
	SavedFlag            bool                         `gorm:"<-:false"`
	LanguageId           int                          `gorm:"column:language_id"`
	EntryReferenceId     int                          `gorm:"column:entry_reference_id"`
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
	Publishedonly            bool   //if you want published entries only set true
	ActiveChannelEntriesonly bool   //if you want active channel entries only set true
	MemberProfile            bool   //if you want member profile set true
	AdditionalFields         bool   //if you want additionalfields set true
	AuthorDetails            bool   //if you want authordetails set true
	ContentHide              bool   //if you want hide content only for memberaccesscontrol set true otherwise it doesn't fetch the entry
	MemberAccessControl      bool
	MemberId                 int
	ImageUrlPath             string
	FieldTypeId              int
	MemberFieldTypeId        int
	Sorting                  string
	EntriesTitle             string
	LanguageId               int
}

type IndivEntriesReq struct {
	ChannelName       string
	EntryId           int
	MemberProfile     bool //if you want member profile set true
	AdditionalFields  bool //if you want additionalfields set true
	AuthorDetails     bool //if you want authordetails set true
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
	Title              string
	Content            string
	CoverImage         string
	SEODetails         SEODetails
	CategoryIds        string
	MembershipLevelids string
	ChannelName        string
	Status             int
	ChannelId          int
	Author             string
	CreateTime         time.Time
	PublishTime        time.Time
	ReadingTime        int
	SortOrder          int
	Tag                string
	Excerpt            string
	CreatedBy          int
	ModifiedBy         int
	IsActive           int
	OrderIndex         int
	CtaId              int
	LanguageId         int
}

type RecentActivities struct {
	Contenttype string
	Title       string
	User        string
	Imagepath   string
	Createdon   time.Time
	Active      string
	Channelname string
	NameString  string
}

type EntriesModel struct {
	Userid     int
	Dataaccess int
}

type EntriesInputs struct {
	Id                     int
	Slug                   string
	Limit                  int
	Offset                 int
	SortBy                 string
	Order                  int
	Keyword                string
	Title                  string
	Status                 string
	ChannelId              int
	CategoryId             int
	CategorySlug           string
	TenantId               string
	SelectedCategoryFilter bool
	ActiveEntriesonly      bool
	GetMemberProfile       bool
	GetAdditionalFields    bool
	GetAuthorDetails       bool
	GetLinkedCategories    bool
	ContentHide            bool
	MemberAccessControl    bool
	MemberId               int
	SectionFieldTypeId     int
	MemberFieldTypeId      int
	TotalCount             bool
	Location               string
	Experience             string
	Posteddate             string
	ChannelName            string
	GetSavedEntryList      bool
	Uuid                   string
}

type JoinEntries struct {
	Id                int `gorm:"column:entry_id"`
	Title             string
	Slug              string
	Description       string
	UserID            int
	ChannelID         int
	Status            int
	IsActive          int
	CreatedOn         time.Time `gorm:"column:entry_created_on"`
	CreatedBy         int       `gorm:"column:entry_created_by"`
	ModifiedBy        int       `gorm:"column:entry_modified_by"`
	ModifiedOn        time.Time `gorm:"column:entry_modified_on"`
	IsDeleted         int
	DeletedOn         time.Time
	DeletedBy         int
	CoverImage        string
	ThumbnailImage    string
	MetaTitle         string
	MetaDescription   string
	Keyword           string
	CategoriesID      string
	RelatedArticles   string
	TenantId          string
	Feature           int
	Author            string
	SortOrder         int
	CreateTime        time.Time
	PublishedTime     time.Time
	ReadingTime       int
	Tags              string
	Excerpt           string
	ViewCount         int
	ImageAltTag       string
	ProfileId         int `gorm:"column:prof_id"`
	MemberID          int
	ProfileName       string
	ProfileSlug       string
	ProfilePage       string
	MemberDetails     datatypes.JSONMap
	CompanyName       string
	CompanyLocation   string
	CompanyLogo       string
	About             string
	SeoTitle          string
	SeoDescription    string
	SeoKeyword        string
	Linkedin          string
	Twitter           string
	Website           string
	ProfCreatedBy     int       `gorm:"column:prof_created_by"`
	ProfCreatedOn     time.Time `gorm:"column:prof_created_on"`
	ProfModifiedOn    time.Time `gorm:"column:prof_modified_on"`
	ProfModifiedBy    int       `gorm:"column:prof_modified_by"`
	ClaimStatus       int
	ProfIsDeleted     int       `gorm:"column:prof_is_deleted"`
	ProfDeletedOn     time.Time `gorm:"column:prof_deleted_on"`
	ProfDeletedBy     int       `gorm:"column:prof_deleted_by"`
	ClaimDate         time.Time
	ProfTenantId      string `gorm:"column:prof_tenant_id"`
	ProfStorageType   string `gorm:"column:prof_storage_type"`
	AuthorId          int    `gorm:"column:author_id"`
	FirstName         string
	LastName          string
	Email             string
	MobileNo          string
	AuthorActive      int `gorm:"column:author_active"`
	ProfileImagePath  string
	ProfileImage      string
	AuthorStorageType string    `gorm:"column:author_storage_type"`
	AuthorCreatedOn   time.Time `gorm:"column:author_created_on"`
	AuthorCreatedBy   int       `gorm:"column:author_created_by"`
	AuthorModifiedOn  time.Time `gorm:"column:author_modified_on"`
	AuthorModifiedBy  int       `gorm:"column:author_modified_by"`
	RoleId            int
	Username          string
	DataAccess        int
	LastLogin         time.Time
	AuthorIsDeleted   int       `gorm:"column:author_is_deleted"`
	AuthorDeletedOn   time.Time `gorm:"column:author_deleted_on"`
	AuthorDeletedBy   int       `gorm:"column:author_deleted_by"`
	DefaultLanguageId int
	UserTenantId      string `gorm:"column:user_tenant_id"`
	RoleName          string
	CtaId             int
	SavedFlag         bool
	Entry_Uuid        string
}
type EntrySave struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	EntryId   int       `gorm:"type:integer"`
	UserId    int       `gorm:"type:integer"`
	TenantId  string    `gorm:"type:character varying"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone"`
	IsDeleted int       `gorm:"type:integer"`
}

var EntryModel EntriesModel

/*List Channel Entry*/
func (Ch EntriesModel) ChannelEntryList(filter Entries, channel *Channel, categoryid string, createonly bool, DB *gorm.DB, tenantid string) (chentry []Tblchannelentries, chentcount int64, err error) {

	query := DB.Model(TblChannelEntries{}).Select("tbl_channel_entries.*,tbl_users.username,tbl_users.first_name,tbl_users.last_name,tbl_users.profile_image_path,tbl_channels.channel_name").Joins("inner join tbl_users on tbl_users.id = tbl_channel_entries.created_by").Joins("left join tbl_channels on tbl_channels.id = tbl_channel_entries.channel_id").Where("tbl_channel_entries.is_deleted=0 and tbl_channel_entries.tenant_id=?", tenantid)

	if filter.Sorting == "lastUpdated" {

		query = query.Order("modified_on desc")

	} else if filter.Sorting == "createdDate" {

		query = query.Order("created_on asc")

	} else if filter.Sorting == "asc" {

		query = query.Order("title asc")

	} else if filter.Sorting == "desc" {

		query = query.Order("title desc")

	} else {

		query = query.Order("order_index asc")

	}

	if channel.PermissionEnable && (channel.Auth.RoleId != 1 && channel.Auth.RoleId != 2) {

		query = query.Where("channel_id in (select id from tbl_channels where slug_name in (select display_name from tbl_module_permissions inner join tbl_modules on tbl_modules.id = tbl_module_permissions.module_id inner join tbl_role_permissions on tbl_role_permissions.permission_id = tbl_module_permissions.id where role_id =(?) and tbl_modules.module_name='Entries' )) ", channel.Auth.RoleId)

	}

	if Ch.Dataaccess == 1 {

		query = query.Where("tbl_channel_entries.created_by=?", Ch.Userid)

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

	if filter.LanguageId != 0 {

		if filter.LanguageId == 1 {
			query = query.Where("tbl_channel_entries.language_id=? or tbl_channel_entries.language_id is null", filter.LanguageId)
		} else {
			query = query.Where("tbl_channel_entries.language_id=?", filter.LanguageId)
		}
	}

	if filter.UserName != "" {

		query = query.Where("LOWER(TRIM(tbl_users.username)) LIKE LOWER(TRIM(?))", "%"+filter.UserName+"%")

	}

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?)) OR LOWER(TRIM(channel_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

	}

	if filter.EntriesTitle != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?)) OR LOWER(TRIM(channel_name)) LIKE LOWER(TRIM(?))", "%"+filter.EntriesTitle+"%", "%"+filter.EntriesTitle+"%")

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

		_, entryid := EntryModel.MemberAccessCheck(filter.MemberId, DB, tenantid)

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

// Fetching the channel entries data
func (Ch EntriesModel) GetFlexibleEntriesData(input EntriesInputs, channel *Channel, db *gorm.DB, joinData *[]JoinEntries, commoncount, totalCount *int64) error {

	selectData := "en.*, en.id as entry_id,en.uuid as entry_uuid, en.created_on as entry_created_on,en.created_by as entry_created_by,en.modified_by as entry_modified_by,en.modified_on as entry_modified_on"

	query := db.Distinct("en.id").Table("tbl_channel_entries as en").Joins("inner join tbl_channels as tc on tc.id = en.channel_id").Where("en.is_deleted = 0 and tc.is_deleted = 0")

	if input.TotalCount {

		if err := query.Count(totalCount).Error; err != nil {

			return err
		}
	}

	if channel.PermissionEnable && (channel.Auth.RoleId != 1 && channel.Auth.RoleId != 2) {

		query = query.Where("channel_id in (select id from tbl_channels where channel_name in (select display_name from tbl_module_permissions inner join tbl_modules on tbl_modules.id = tbl_module_permissions.module_id inner join tbl_role_permissions on tbl_role_permissions.permission_id = tbl_module_permissions.id where role_id =(?) and tbl_modules.module_name='Entries' )) ", channel.Auth.RoleId)

	}

	if Ch.Dataaccess == 1 {

		query = query.Where("en.created_by=?", Ch.Userid)

	}

	if input.TenantId != "" {

		query = query.Where("en.tenant_id=?", input.TenantId)
	}

	if input.ChannelId != 0 {

		query = query.Where("en.channel_id = ?", input.ChannelId)
	}

	if input.ChannelName != "" {
		query = query.Where("en.channel_id IN (SELECT id FROM tbl_channels WHERE channel_name = ? AND is_deleted = 0)", input.ChannelName)
	}

	if input.Status != "" {

		status, _ := strconv.Atoi(input.Status)

		query = query.Where("en.status = ?", status)
	}

	if input.ActiveEntriesonly {

		query = query.Where("en.is_active = ?", 1)
	}

	if input.Keyword != "" {

		query = query.Where("TRIM(LOWER(en.title)) LIKE TRIM(LOWER(?))", "%"+input.Keyword+"%")
	}

	if input.Location != "" {
		query = query.Joins("INNER JOIN tbl_channel_entry_fields AS ef ON ef.channel_entry_id = en.id").
			Where("TRIM(LOWER(ef.field_value)) LIKE TRIM(LOWER(?))", "%"+strings.ToLower(input.Location)+"%")
	}

	// Filter by Experience (Fix: Use 'ef2' alias to prevent duplicate alias issue)
	if input.Experience != "" {
		query = query.Joins("INNER JOIN tbl_channel_entry_fields AS ef2 ON ef2.channel_entry_id = en.id").
			Where("TRIM(LOWER(ef2.field_value)) = TRIM(LOWER(?))", input.Experience) // Exact match
	}

	if input.Posteddate != "" {
		now := time.Now()
		startOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, -int(now.Weekday())+1)
		startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
		startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)

		switch strings.ToLower(input.Posteddate) {
		case "today":
			query = query.Joins("INNER JOIN tbl_channel_entry_fields AS ef3 ON ef3.channel_entry_id = en.id").
				Where("(ef3.field_value) = ?", now.Format("2006-01-02"))

		case "this week":
			fmt.Println(input.Posteddate, "postedate")
			query = query.Joins("INNER JOIN tbl_channel_entry_fields AS ef3 ON ef3.channel_entry_id = en.id").
				Where(" ef3.field_value>=?  AND ef3.field_value<=? ", startOfWeek.Format("2006-01-02"), now.Format("2006-01-02"))

		case "this month":
			query = query.Joins("INNER JOIN tbl_channel_entry_fields AS ef3 ON ef3.channel_entry_id = en.id").
				Where("ef3.field_value>=?  AND ef3.field_value<=?", startOfMonth.Format("2006-01-02"), now.Format("2006-01-02"))

		case "this year":
			query = query.Joins("INNER JOIN tbl_channel_entry_fields AS ef3 ON ef3.channel_entry_id = en.id").
				Where("ef3.field_value>=?  AND ef3.field_value<=?", startOfYear.Format("2006-01-02"), now.Format("2006-01-02"))

		default:
			fmt.Println("Invalid posted date filter:", input.Posteddate) // Optional logging
		}
	}

	if input.Title != "" {

		query = query.Where("en.title = ?", input.Title)
	}

	var joinCondition, profileCondition string

	if input.CategoryId != 0 || input.CategorySlug != "" {

		if db.Config.Dialector.Name() == "mysql" {

			joinCondition = `find_in_set(cat.id,en.categories_id) > 0`

		} else if db.Config.Dialector.Name() == "postgres" {

			joinCondition = `cat.id = any(string_to_array(en.categories_id,',')::Integer[])`

		}
	}

	if input.GetMemberProfile {

		if db.Config.Dialector.Name() == "mysql" {

			profileCondition = `find_in_set(tmp.member_id,cef.field_value) > 0`

		} else if db.Config.Dialector.Name() == "postgres" {

			profileCondition = `tmp.member_id = any(string_to_array(cef.field_value,',')::Integer[])`

		}
	}

	if input.CategoryId != 0 {

		switch {

		case input.SelectedCategoryFilter:

			query = query.Joins("inner join tbl_categories as cat on "+joinCondition+" and cat.id = ?", input.CategoryId)

		default:

			subQuery := db.Table("tbl_categories as cat").Select("cat.id").Where("cat.is_deleted = 0 and tenant_id = ? and cat.id = (?) or cat.parent_id in (?)", input.TenantId, input.CategoryId, input.CategoryId)

			query = query.Joins("inner join tbl_categories as cat on "+joinCondition+" and cat.id in (?)", subQuery)
		}

	} else if input.CategorySlug != "" {

		switch {

		case input.SelectedCategoryFilter:

			query = query.Joins("inner join tbl_categories as cat on "+joinCondition+" and cat.category_slug = ?", input.CategorySlug)

		default:

			innerSubQuery := db.Table("tbl_categories as cat").Select("cat.id").Where("cat.is_deleted = 0 and tenant_id = ? and cat.category_slug = ?", input.TenantId, input.CategorySlug)

			subQuery := db.Table("tbl_categories as cat").Select("cat.id").Where("cat.is_deleted = 0 and cat.id = (?) or cat.parent_id in (?)", innerSubQuery, innerSubQuery)

			query = query.Joins("inner join tbl_categories as cat on "+joinCondition+" and cat.id in (?)", subQuery)
		}

	}

	if input.GetMemberProfile {

		selectData += ",mj.*,mj.created_by as prof_created_by,mj.created_on as prof_created_on,mj.modified_on as prof_modified_on,mj.modified_by as prof_modified_by,mj.id as prof_id,mj.is_deleted as prof_is_deleted,mj.deleted_on as prof_deleted_on,mj.deleted_by as prof_deleted_by,mj.storage_type as prof_storage_type,mj.tenant_id as prof_tenant_id"

		joinSubQuery := db.Select("tmp.*,cef.channel_entry_id,cef.field_value").Table("tbl_channel_entry_fields as cef").Joins("inner join tbl_fields tf on tf.id = cef.field_id").Joins("inner join tbl_member_profiles tmp on " + profileCondition).Where("tmp.is_deleted=0 and tf.is_deleted=0")

		query = query.Joins("left join (?) mj on mj.channel_entry_id = en.id", joinSubQuery)
	}

	if input.GetAuthorDetails {

		selectData += ", tu.*, " +
			"tu.id as author_id, " +
			"tu.is_active as author_active, " +
			"tu.created_on as author_created_on, " +
			"tu.created_by as author_created_by, " +
			"tu.modified_on as author_modified_on, " +
			"tu.modified_by as author_modified_by, " +
			"tu.deleted_on as author_deleted_on, " +
			"tu.deleted_by as author_deleted_by, " +
			"tu.is_deleted as author_is_deleted, " +
			"tu.storage_type as author_storage_type, " +
			"tu.tenant_id as user_tenant_id, " +
			"tu.profile_image_path as author_profile_image_path, " +
			"tr.name as role_name"

		query = query.Select(selectData).
			Joins("LEFT JOIN tbl_users AS tu ON tu.id = en.created_by").
			Joins("LEFT JOIN tbl_roles AS tr ON tr.id = tu.role_id").
			Where("tu.is_deleted = 0")
	}

	if input.GetSavedEntryList && input.MemberId != 0 {
		selectData += ", se.entry_id as saved_entry_id, true AS saved_flag"
		query = query.Joins("INNER JOIN tbl_saved_entries AS se ON se.entry_id = en.id").
			Where("se.user_id = ? AND se.tenant_id = ? AND se.is_deleted = 0", input.MemberId, input.TenantId)
	} else if input.MemberId != 0 {
		// Fetch all entries with savedFlag
		selectData += ", CASE WHEN se.entry_id IS NOT NULL THEN true ELSE false END AS saved_flag"
		query = query.Joins("LEFT JOIN tbl_saved_entries AS se ON se.entry_id = en.id AND se.user_id = ? AND se.tenant_id = ? AND se.is_deleted = 0", input.MemberId, input.TenantId)
	}
	if input.MemberAccessControl && input.MemberId != 0 && input.ContentHide {

		restrictQuery := db.Select("acp.entry_id").Table("tbl_access_control_pages as acp").Joins("inner join tbl_access_control_user_groups as acu on acu.id = acp.access_control_user_group_id").Joins("inner join tbl_members as tm on tm.member_group_id = acu.member_group_id").Where("tm.is_deleted = 0 and tm.id = ? and acu.is_deleted= 0 and acp.is_deleted = 0", input.MemberId)

		query = query.Where("en.id not in (?)", restrictQuery)
	}

	if err := query.Count(commoncount).Error; err != nil {

		return err
	}

	if input.SortBy != "" {

		if input.Order > 0 {

			query = query.Order(input.SortBy + " desc")

		} else {

			query = query.Order(input.SortBy)
		}

	} else {

		query = query.Order("en.id desc")
	}

	if input.Limit > 0 {

		query = query.Limit(input.Limit)
	}

	if input.Offset > -1 {

		query = query.Offset(input.Offset)
	}

	if err := query.Select(selectData).Find(&joinData).Error; err != nil {

		return err

	}

	return nil

}

func (ch EntriesModel) MemberAccessCheck(memberid int, DB *gorm.DB, tenantid string) ([]int, []int) {

	var channelid, entryid []int

	var mem member.TblMember

	//get membergroup id
	DB.Table("tbl_members").Select("member_group_id").Where("is_deleted=0 and id=? and tenant_id=?", memberid, tenantid).First(&mem)

	SUB := `select id from tbl_access_control_user_groups where is_deleted=0 and member_group_id=` + strconv.Itoa(mem.Id)

	var accessgroup []access.TblAccessControlPages

	DB.Table("tbl_access_control_pages").Where("access_control_user_group_id in (?) and tenant_id=?", SUB, tenantid).Find(&accessgroup)

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
func (Ch EntriesModel) DeleteChannelEntryId(chentry *Tblchannelentries, id int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entries").Where("id=? and tenant_id=?", chentry.Id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": chentry.IsDeleted, "deleted_by": chentry.DeletedBy, "deleted_on": chentry.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*Delete Channel Entry Field*/
func (Ch EntriesModel) DeleteChannelEntryFieldId(chentry *TblChannelEntryField, id int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entry_fields").Where("channel_entry_id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"deleted_by": chentry.DeletedBy, "deleted_on": chentry.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*Edit Channel Entry*/
func (Ch EntriesModel) GetChannelEntryById(ent IndivEntriesReq, DB *gorm.DB, tenantid string) (tblchanentry Tblchannelentries, err error) {

	query := DB.Table("tbl_channel_entries").Select("tbl_channel_entries.*,tbl_users.username,tbl_users.profile_image_path,tbl_channels.channel_name").Joins("inner join tbl_users on tbl_users.id = tbl_channel_entries.created_by").Joins("left join tbl_channels on tbl_channels.id = tbl_channel_entries.channel_id").Where("tbl_channel_entries.is_deleted=0 and tbl_channel_entries.id=? and tbl_channel_entries.tenant_id=?", ent.EntryId, tenantid)

	if ent.ContentHide {

		_, entryid := EntryModel.MemberAccessCheck(ent.MemberId, DB, tenantid)

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

func (Ch EntriesModel) GetAuthorDetails(DB *gorm.DB, authorId int, tenantid string) (authorDetail team.TblUser, err error) {

	if err := DB.Model(team.TblUser{}).Where("tbl_users.is_deleted = 0 and tbl_users.id = ? and tbl_users.tenant_id=?", authorId, tenantid).First(&authorDetail).Error; err != nil {

		return team.TblUser{}, err
	}

	return authorDetail, nil
}

func (Ch EntriesModel) GetSectionsUnderEntries(DB *gorm.DB, channelId, sectionTypeId int, tenantid string) (sections []tblfield, err error) {

	if err = DB.Table("tbl_group_fields").Select("tbl_fields.*,tbl_field_types.type_name").Joins("inner join tbl_fields on tbl_fields.id = tbl_group_fields.field_id").Joins("inner join tbl_field_types on tbl_field_types.id = tbl_fields.field_type_id").
		Where("tbl_fields.is_deleted = 0 and tbl_field_types.is_deleted = 0 and tbl_fields.field_type_id = ? and tbl_group_fields.channel_id = ? and tbl_group_fields.tenant_id=?", sectionTypeId, channelId, tenantid).Find(&sections).Error; err != nil {

		return []tblfield{}, err
	}

	return sections, nil
}

func (Ch EntriesModel) GetFieldsInEntries(DB *gorm.DB, channelId, sectionTypeId int, tenantid string) (fields []tblfield, err error) {

	if err = DB.Table("tbl_group_fields").Select("tbl_fields.*,tbl_field_types.type_name").Joins("inner join tbl_fields on tbl_fields.id = tbl_group_fields.field_id").Joins("inner join tbl_field_types on tbl_field_types.id = tbl_fields.field_type_id").
		Where("tbl_fields.is_deleted = 0 and tbl_field_types.is_deleted = 0 and tbl_fields.field_type_id != ? and tbl_group_fields.channel_id = ? and tbl_group_fields.tenant_id=?", sectionTypeId, channelId, tenantid).Find(&fields).Error; err != nil {

		return []tblfield{}, err
	}

	return fields, nil
}

func (Ch EntriesModel) GetFieldValue(DB *gorm.DB, fieldId, entryId int, tenantid string) (fieldvalue TblChannelEntryField, err error) {

	if err = DB.Model(TblChannelEntryField{}).Where("tbl_channel_entry_fields.field_id = ? and tbl_channel_entry_fields.channel_entry_id = ? and tbl_channel_entry_fields.tenant_id=?", fieldId, entryId, tenantid).First(&fieldvalue).Error; err != nil {

		return TblChannelEntryField{}, err
	}

	return fieldvalue, nil
}

func (ch EntriesModel) GetFieldOptions(DB *gorm.DB, fieldId int, tenantid string) (fieldOptions []TblFieldOption, err error) {

	if err = DB.Model(TblFieldOption{}).Where("tbl_field_options.is_deleted = 0 and tbl_field_options.field_id = ? and tbl_field_options.tenant_id=?", fieldId, tenantid).Find(&fieldOptions).Error; err != nil {

		return []TblFieldOption{}, err
	}

	return fieldOptions, nil
}

func (ch EntriesModel) GetMemberProfile(DB *gorm.DB, memberid int, tenantid string) (memberProfile member.TblMemberProfile, err error) {

	if err = DB.Model(member.TblMemberProfile{}).Select("tbl_member_profiles.*").Joins("inner join tbl_members on tbl_members.id = tbl_member_profiles.member_id").Where("tbl_members.is_deleted = 0 and tbl_members.id = ? and tbl_members.tenant_id=?", memberid, tenantid).First(&memberProfile).Error; err != nil {

		return member.TblMemberProfile{}, err
	}

	return memberProfile, nil

}

func (ch EntriesModel) GetGraphqlEntriesCategoryByParentId(DB *gorm.DB, categoryId int, tenantid string) (category categories.TblCategories, err error) {

	if err = DB.Model(categories.TblCategories{}).Where("is_deleted = 0 and id = ? and tenant_id=?", categoryId, tenantid).First(&category).Error; err != nil {

		return categories.TblCategories{}, err
	}

	return category, nil
}

func (ch EntriesModel) GetCategoryIdByName(name string, DB *gorm.DB, tenantid string) (category categories.TblCategories, er error) {

	if err := DB.Model(categories.TblCategories{}).Where("category_name= ? and tenant_id=?", name, tenantid).First(&category).Error; err != nil {

		return category, err
	}

	return category, nil
}

func (ch EntriesModel) GetChildCategories(categoryid int, DB *gorm.DB, tenantid string) (categories []categories.TblCategories, er error) {

	if err := DB.Raw(`WITH RECURSIVE cat_tree AS (SELECT id, category_name, category_slug,image_path, parent_id,created_on,modified_on,is_deleted
		FROM tbl_categories
		WHERE id = (?)
		UNION
		SELECT cat.id, cat.category_name, cat.category_slug, cat.image_path ,cat.parent_id,cat.created_on,cat.modified_on,
		cat.is_deleted
		FROM tbl_categories AS cat
		JOIN cat_tree ON cat.parent_id = cat_tree.id )
		SELECT * FROM cat_tree where is_deleted = 0 and tenant_id=?`, tenantid, categoryid).Find(&categories).Error; err != nil {

		return categories, err
	}

	return categories, nil
}

func (ch EntriesModel) MakeFeature(channelid, entryid, status int, DB *gorm.DB, tenantid string) (err error) {

	DB.Model(TblChannelEntries{}).Where("channel_id=?", channelid).UpdateColumns(map[string]interface{}{"feature": 0})

	if err := DB.Model(TblChannelEntries{}).Where("id=? and channel_id=? and tenant_id=?", entryid, channelid, tenantid).UpdateColumns(map[string]interface{}{"feature": status}).Error; err != nil {

		return err
	}

	return nil
}

func (ch EntriesModel) PublishQuery(chl *TblChannelEntries, id int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entries").Where("id =? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"status": chl.Status, "modified_on": chl.ModifiedOn, "modified_by": chl.ModifiedBy}).Error; err != nil {

		return err

	}

	return nil
}

/*Update Channel Entry Details*/
func (Ch EntriesModel) UpdateChannelEntryDetails(entry *TblChannelEntries, entryid int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entries").Where("id=? and tenant_id=?", entryid, tenantid).UpdateColumns(map[string]interface{}{"title": entry.Title, "description": entry.Description, "slug": entry.Slug, "cover_image": entry.CoverImage, "thumbnail_image": entry.ThumbnailImage, "meta_title": entry.MetaTitle, "meta_description": entry.MetaDescription, "keyword": entry.Keyword, "categories_id": entry.CategoriesId, "related_articles": entry.RelatedArticles, "status": entry.Status, "modified_on": entry.ModifiedOn, "modified_by": entry.ModifiedBy, "user_id": entry.UserId, "channel_id": entry.ChannelId, "author": entry.Author, "create_time": entry.CreateTime, "published_time": entry.PublishedTime, "reading_time": entry.ReadingTime, "sort_order": entry.SortOrder, "tags": entry.Tags, "excerpt": entry.Excerpt, "image_alt_tag": entry.ImageAltTag, "order_index": entry.OrderIndex, "cta_id": entry.CtaId, "memebrship_level_ids": entry.MemebrshipLevelIds, "language_id": entry.LanguageId}).Error; err != nil {

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
func (Ch EntriesModel) UpdateChannelEntryAdditionalDetails(entry TblChannelEntryField, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entry_fields").Where("id=? and tenant_id=?", entry.Id, tenantid).UpdateColumns(map[string]interface{}{"field_name": entry.FieldName, "field_value": entry.FieldValue, "modified_by": entry.ModifiedBy, "modified_on": entry.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

func (Ch EntriesModel) AllentryCount(DB *gorm.DB, tenantid string) (count int64, err error) {

	if err := DB.Table("tbl_channel_entries").Where("is_deleted = 0 and tenant_id=?", tenantid).Count(&count).Error; err != nil {

		return 0, err
	}

	return count, nil
}

func (Ch EntriesModel) NewentryCount(DB *gorm.DB, tenantid string) (count int64, err error) {

	if err := DB.Table("tbl_channel_entries").Where("is_deleted = 0 AND created_on >=? and tenant_id=?", time.Now().AddDate(0, 0, -10), tenantid).Count(&count).Error; err != nil {

		return 0, err
	}

	return count, nil
}

func (Ch EntriesModel) Newchannels(DB *gorm.DB, tenantid string) (chn []Tblchannel, err error) {

	if err := DB.Table("tbl_channels").Select("tbl_channels.*,tbl_users.username,tbl_users.profile_image_path,tbl_users.first_name,tbl_users.last_name").
		Joins("inner join tbl_users on tbl_users.id = tbl_channels.created_by").
		Where("tbl_channels.is_deleted=0 and tbl_channels.is_active=1 and tbl_channels.created_on >= ? and tbl_channels.tenant_id=?", time.Now().Add(-24*time.Hour).Format("2006-01-02 15:04:05"), tenantid).
		Order("created_on desc").Limit(6).Find(&chn).Error; err != nil {

		return []Tblchannel{}, err
	}

	return chn, nil

}

func (Ch EntriesModel) Newentries(DB *gorm.DB, tenantid string) (entries []Tblchannelentries, err error) {

	if err := DB.Table("tbl_channel_entries").Select("tbl_channel_entries.*,tbl_users.username,tbl_users.profile_image_path,tbl_users.first_name,tbl_users.last_name").
		Joins("inner join tbl_users on tbl_users.id = tbl_channel_entries.created_by").Where("tbl_channel_entries.is_deleted=0 and tbl_channel_entries.created_on >=? and tbl_channel_entries.tenant_id=?", time.Now().Add(-24*time.Hour).Format("2006-01-02 15:04:05"), tenantid).
		Order("created_on desc").Limit(6).Find(&entries).Error; err != nil {

		return []Tblchannelentries{}, err
	}

	return entries, nil

}

// update imagepath
func (Ch EntriesModel) UpdateImagePath(Imagepath string, DB *gorm.DB, tenantid string) error {

	if err := DB.Model(TblChannelEntries{}).Where("cover_image=? and tenant_id=?", Imagepath, tenantid).UpdateColumns(map[string]interface{}{
		"cover_image": ""}).Error; err != nil {

		return err
	}

	return nil

}

// make feature function
func (ch ChannelModel) MakeFeature(channelid, entryid, status int, DB *gorm.DB, tenantid string) (err error) {

	DB.Model(TblChannelEntries{}).Where("channel_id=? and tenant_id=?", channelid, tenantid).UpdateColumns(map[string]interface{}{"feature": 0})

	if err := DB.Model(TblChannelEntries{}).Where("id=? and channel_id=? and tenant_id=?", entryid, channelid, tenantid).UpdateColumns(map[string]interface{}{"feature": status}).Error; err != nil {

		return err
	}

	return nil
}

/*Delete MULTI Channel Entry Field*/
func (Ch EntriesModel) DeleteSelectedChannelEntryId(chentry *TblChannelEntries, id []int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entries").Where("id in (?) and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": chentry.IsDeleted, "deleted_by": chentry.DeletedBy, "deleted_on": chentry.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*Delete MULTI Channel Entry Field*/
func (Ch EntriesModel) DeleteSelectedChannelEntryFieldId(chentry *TblChannelEntryField, id []int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entry_fields").Where("channel_entry_id IN (?) and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"deleted_by": chentry.DeletedBy, "deleted_on": chentry.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

func (ch ChannelModel) GetChannelEntriesByChannelId(channel_entries *[]TblChannelEntries, channel_id int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entries").Where("tbl_channel_entries.is_deleted = 0 and tbl_channel_entries.status = 1 and tbl_channel_entries.channel_id = ? and tbl_channel_entries.tenant_id=?", channel_id, tenantid).Find(&channel_entries).Error; err != nil {

		return err
	}

	return nil
}

// UNPUBLISH MULTI SELECT ENTRY//
func (Ch EntriesModel) UnpublishSelectedChannelEntryId(chentry *TblChannelEntries, id []int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entries").Where("id IN (?) and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"status": chentry.Status, "modified_by": chentry.ModifiedBy, "modified_on": chentry.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil

}

func (Ch EntriesModel) GetChannelAdditionalFields(DB *gorm.DB, channelId int) (fields []tblfield, err error) {

	if err = DB.Table("tbl_group_fields as tgf").Select("tf.*").Joins("inner join tbl_fields as tf on tf.id = tgf.field_id").Where("tf.is_deleted = 0 and tgf.channel_id = ?", channelId).Find(&fields).Error; err != nil {

		return []tblfield{}, err
	}

	return fields, nil
}

// Entry Preview
func (Ch EntriesModel) GetPreview(chentry *TblChannelEntries, DB *gorm.DB, uuid string) (err error) {

	if err = DB.Table("tbl_channel_entries").Where("uuid = ? and is_deleted=0", uuid).Find(&chentry).Error; err != nil {

		return err
	}

	return nil
}

// Entry  IsActive Function
func (Ch EntriesModel) EntryIsActive(entryisactive Tblchannelentries, entryid int, status int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entries").Where("id=? and tenant_id=?", entryid, tenantid).UpdateColumns(map[string]interface{}{"is_active": status, "modified_by": entryisactive.ModifiedBy, "modified_on": entryisactive.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

// Update channel Entry View Count
func (En *EntriesModel) UpdateEntryViewCount(db *gorm.DB, id int, slug string, tenantId string, count *int) error {

	pipeline := db.Transaction(func(tx *gorm.DB) error {

		query := tx.Table("tbl_channel_entries").Where("is_deleted=0")

		switch {

		case id != 0:

			query = query.Where("id=?", id)

		case slug != "":

			query = query.Where("slug=?", slug)

		}

		if tenantId != "" {

			switch {

			case tenantId == "":

				query = query.Where("tenant_id=?", tenantId)

			default:

				query = query.Where("tenant_id=?", tenantId)
			}

		}

		if rowsAffected := query.Update("view_count", gorm.Expr("view_count + ?", 1)).RowsAffected; rowsAffected == 0 {

			return errors.New("no rows affected")
		}

		if err := query.Select("view_count").Scan(count).Error; err != nil {

			return err
		}

		return nil

	})

	if err := pipeline; err != nil {

		return err
	}

	return nil
}

func (En *EntriesModel) FlexibleChannelEntryDetail(db *gorm.DB, inputs EntriesInputs, multiFetchIds []int, channelEntryDetails *JoinEntries, multiEntryDetails *[]JoinEntries) error {

	selectData := "en.*, en.id as entry_id,en.uuid as entry_uuid,en.created_on as entry_created_on,en.created_by as entry_created_by,en.modified_by as entry_modified_by,en.modified_on as entry_modified_on"

	query := db.Table("tbl_channel_entries as en").Where("en.is_deleted=0")

	switch {

	case len(multiFetchIds) > 0:

		query = query.Where("en.id in (?)", multiFetchIds)

	case inputs.Id != 0:

		query = query.Where("en.id=?", inputs.Id)

	}

	if inputs.Slug != "" {

		if inputs.Slug != "" && inputs.Uuid != "" {

			query = query.Where("en.slug=? and en.uuid=?", inputs.Slug, inputs.Uuid)

		} else {

			query = query.Where("en.slug=?", inputs.Slug)

		}

	}

	if inputs.TenantId != "" {

		query = query.Where("en.tenant_id = ?", inputs.TenantId)
	}

	if inputs.ChannelId > 0 {

		query = query.Where("en.channel_id = ?", inputs.ChannelId)
	}

	if inputs.MemberId > 0 {

		selectData += ", CASE WHEN se.entry_id IS NOT NULL THEN true ELSE false END AS saved_flag"
		query = query.Joins("LEFT JOIN tbl_saved_entries AS se ON se.entry_id = en.id AND se.user_id = ? AND se.tenant_id = ? AND se.is_deleted = 0", inputs.MemberId, inputs.TenantId)
	}

	var profileCondition string

	if inputs.GetMemberProfile {

		if db.Config.Dialector.Name() == "mysql" {

			profileCondition = `find_in_set(tmp.member_id,cef.field_value) > 0`

		} else if db.Config.Dialector.Name() == "postgres" {

			profileCondition = `tmp.member_id = any(string_to_array(cef.field_value,',')::Integer[])`

		}
	}

	if inputs.GetMemberProfile {

		selectData += ",mj.*,mj.created_by as prof_created_by,mj.created_on as prof_created_on,mj.modified_on as prof_modified_on,mj.modified_by as prof_modified_by,mj.id as prof_id,mj.is_deleted as prof_is_deleted,mj.deleted_on as prof_deleted_on,mj.deleted_by as prof_deleted_by,mj.storage_type as prof_storage_type,mj.tenant_id as prof_tenant_id"

		joinSubQuery := db.Select("tmp.*,cef.channel_entry_id,cef.field_value").Table("tbl_channel_entry_fields as cef").Joins("inner join tbl_fields tf on tf.id = cef.field_id").Joins("inner join tbl_member_profiles tmp on " + profileCondition).Where("tmp.is_deleted=0 and tf.is_deleted=0")

		query = query.Joins("left join (?) mj on mj.channel_entry_id = en.id", joinSubQuery)
	}

	if inputs.GetAuthorDetails {

		selectData += ", tu.*, " +
			"tu.id as author_id, " +
			"tu.is_active as author_active, " +
			"tu.created_on as author_created_on, " +
			"tu.created_by as author_created_by, " +
			"tu.modified_on as author_modified_on, " +
			"tu.modified_by as author_modified_by, " +
			"tu.deleted_on as author_deleted_on, " +
			"tu.deleted_by as author_deleted_by, " +
			"tu.is_deleted as author_is_deleted, " +
			"tu.storage_type as author_storage_type, " +
			"tu.tenant_id as user_tenant_id, " +
			"tu.profile_image_path as author_profile_image_path, " +
			"tr.name as role_name"

		query = query.Select(selectData).
			Joins("LEFT JOIN tbl_users AS tu ON tu.id = en.created_by").
			Joins("LEFT JOIN tbl_roles AS tr ON tr.id = tu.role_id").
			Where("tu.is_deleted = 0")
	}
	switch {

	case len(multiFetchIds) > 0:

		query = query.Select(selectData).Find(&multiEntryDetails)

	default:

		query = query.Select(selectData).Scan(&channelEntryDetails)

	}

	if err := query.Error; err != nil {

		return err
	}

	return nil
}

func (Ch EntriesModel) EntryParentIdUpdate(Entries *Tblchannelentries, db *gorm.DB) error {

	if err := db.Table("tbl_channel_entries").Where("id=? and tenant_id=?", Entries.Id, Entries.TenantId).UpdateColumns(map[string]interface{}{"parent_id": Entries.ParentId, "modified_by": Entries.ModifiedBy, "modified_on": Entries.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil

}

func (Ch EntriesModel) EntrylistByParentId(channel_entries *[]Tblchannelentries, parentid int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entries").Select("tbl_channel_entries.*,tbl_users.username,tbl_users.profile_image_path,tbl_channels.channel_name").Joins("inner join tbl_users on tbl_users.id = tbl_channel_entries.created_by").Joins("left join tbl_channels on tbl_channels.id = tbl_channel_entries.channel_id").Where("tbl_channel_entries.is_deleted = 0  and tbl_channel_entries.parent_id = ? and tbl_channel_entries.tenant_id=?", parentid, tenantid).Find(&channel_entries).Error; err != nil {

		return err
	}

	return nil
}
func (Ch EntriesModel) UpdateEntryOrder(Entries *Tblchannelentries, DB *gorm.DB) error {

	if err := DB.Table("tbl_channel_entries").Where("id=? and tenant_id=?", Entries.Id, Entries.TenantId).UpdateColumns(map[string]interface{}{"order_index": Entries.OrderIndex, "modified_by": Entries.ModifiedBy, "modified_on": Entries.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}
func (Ch EntriesModel) UpdateEntryMemberGroupIds(Entries *Tblchannelentries, DB *gorm.DB) error {

	if err := DB.Table("tbl_channel_entries").Where("id=? and tenant_id=?", Entries.Id, Entries.TenantId).UpdateColumns(map[string]interface{}{"membergroup_id": Entries.MembergroupId, "modified_by": Entries.ModifiedBy, "modified_on": Entries.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

func (Ch EntriesModel) UpdateEntryOrderIndex(entry *TblChannelEntries, entryid int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_channel_entries").Where("id=? and tenant_id=?", entryid, tenantid).UpdateColumns(map[string]interface{}{"order_index": entry.OrderIndex}).Error; err != nil {

		return err
	}

	return nil
}
func (Ch EntriesModel) GetEntriesById(id int, DB *gorm.DB, tenantid string) (*TblChannelEntries, error) {

	var entries TblChannelEntries

	if err := DB.Table("tbl_channel_entries").Where("id=? and tenant_id=?", id, tenantid).Find(&entries).Error; err != nil {

		return nil, err
	}

	return &entries, nil

}
func (Ch EntriesModel) defaultchannelid(slug string, DB *gorm.DB, tenantid string) (int, error) {

	var Channels TblChannel

	if err := DB.Table("tbl_channels").Where("slug_name=? and tenant_id=?", slug, tenantid).Find(&Channels).Error; err != nil {

		return 0, err
	}

	return Channels.Id, nil
}

func (Ch EntriesModel) EntryAuthors(tenantid string, DB *gorm.DB) ([]Author, error) {

	var authors []Author
	err := DB.Table("tbl_channel_entries").
		Select("tbl_users.*, COUNT(tbl_channel_entries.id) as entry_count").
		Joins("JOIN tbl_users ON tbl_users.id = tbl_channel_entries.created_by").
		Where("tbl_channel_entries.tenant_id = ? and tbl_users.tenant_id is not null", tenantid).
		Group("tbl_users.id").
		Order("entry_count DESC").
		Limit(10).
		Find(&authors).Error

	if err != nil {
		return nil, err
	}
	return authors, nil

}

func (Ch EntriesModel) EntrySave(entrydata *EntrySave, DB *gorm.DB) error {

	if err := DB.Table("tbl_saved_entries").Create(&entrydata).Error; err != nil {

		return err

	}

	return nil
}
func (Ch EntriesModel) EntryUnsave(entryId int, userId int, tenantId string, DB *gorm.DB) error {

	if err := DB.Table("tbl_saved_entries").
		Where("entry_id = ? AND user_id = ? AND tenant_id = ?", entryId, userId, tenantId).
		Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}
