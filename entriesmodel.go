package channels

import (
	"time"

	"gorm.io/gorm"
)

type tblchannelentries struct {
	Id               int
	Title            string `form:"title" binding:"required"`
	Slug             string `form:"slug" binding:"required"`
	Description      string
	UserId           int
	ChannelId        int
	Status           int //0-draft 1-publish 2-unpublish
	IsActive         int
	IsDeleted        int       `gorm:"DEFAULT:0"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"DEFAULT:NULL"`
	CreatedOn        time.Time
	CreatedBy        int
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	ModifiedOn       time.Time `gorm:"DEFAULT:NULL"`
	CoverImage       string
	ThumbnailImage   string
	MetaTitle        string `form:"metatitle" binding:"required"`
	MetaDescription  string `form:"metadesc" binding:"required"`
	Keyword          string `form:"keywords" binding:"required"`
	CategoriesId     string
	RelatedArticles  string
	CreatedDate      string `gorm:"-"`
	ModifiedDate     string `gorm:"-"`
	Username         string `gorm:"<-:false"`
	CategoryGroup    string `gorm:"-:migration;<-:false"`
	ChannelName      string `gorm:"-:migration;<-:false"`
	Cno              string `gorm:"<-:false"`
	ProfileImagePath string `gorm:"<-:false"`
	EntryStatus      string `gorm:"-"`
	// Categories       [][]TblCategory    `gorm:"-"`
	AdditionalData string     `gorm:"-"`
	AuthorDetail   Author     `gorm:"-"`
	Sections       []TblField `gorm:"-"`
	Fields         []TblField `gorm:"-"`
	// MemberProfiles   []TblMemberProfile `gorm:"-"`
	Feature       int `gorm:"DEFAULT:0"`
	ViewCount     int `gorm:"DEFAULT:0"`
	Author        string
	SortOrder     int
	CreateDate    string
	PublishedTime string
	ReadingTime   int `gorm:"DEFAULT:0"`
	Tags          string
	Excerpt       string
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

type EntriesFilter struct {
	Keyword     string
	Title       string
	ChannelName string
	Status      string
	UserName    string
	CategoryId  string
}

type EntriesModel struct{}

var EntryModel EntriesModel

/*List Channel Entry*/
func (Ch EntriesModel) ChannelEntryList(limit, offset, chid int, filter EntriesFilter, publishedflg bool, RoleId int, activechannel bool, DB *gorm.DB) (chentry []tblchannelentries, chentcount int64, err error) {

	query := DB.Table("tbl_channel_entries").Select("tbl_channel_entries.*,tbl_users.username,tbl_channels.channel_name").Joins("inner join tbl_users on tbl_users.id = tbl_channel_entries.created_by").Joins("inner join tbl_channels on tbl_channels.id = tbl_channel_entries.channel_id").Where("tbl_channel_entries.is_deleted=0").Order("id desc")

	if RoleId != 1 {

		query = query.Where("channel_id in (select id from tbl_channels where channel_name in (select display_name from tbl_module_permissions inner join tbl_modules on tbl_modules.id = tbl_module_permissions.module_id inner join tbl_role_permissions on tbl_role_permissions.permission_id = tbl_module_permissions.id where role_id =(?) and tbl_modules.module_name='Entries' )) ", RoleId)

	}

	if activechannel {

		query = query.Where("tbl_channels.is_active =1")
	}

	if publishedflg {

		query = query.Where("tbl_channel_entries.status=1")

	}

	if chid != 0 {

		query = query.Where("tbl_channel_entries.channel_id=?", chid)
	}

	if filter.UserName != "" {

		query = query.Debug().Where("LOWER(TRIM(tbl_users.username)) ILIKE LOWER(TRIM(?))", "%"+filter.UserName+"%")

	}

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(title)) ILIKE LOWER(TRIM(?)) OR LOWER(TRIM(channel_name)) ILIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

	}

	if filter.Status != "" {

		query = query.Where("tbl_channel_entries.status=?", filter.Status)

	}
	if filter.Title != "" {

		query = query.Where("LOWER(TRIM(title)) ILIKE LOWER(TRIM(?))", "%"+filter.Title+"%")

	}

	if filter.ChannelName != "" {

		query = query.Where("LOWER(TRIM(channel_name)) ILIKE LOWER(TRIM(?))", "%"+filter.ChannelName+"%")

	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Order("id asc").Find(&chentry)

	} else {

		query.Find(&chentry).Count(&chentcount)

		return chentry, chentcount, nil
	}

	return chentry, 0, nil
}
