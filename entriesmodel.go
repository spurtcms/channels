package channels

import (
	"time"
)

type TblChannelEntries struct {
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


