package channels

import (
	"time"

	"gorm.io/gorm"
)

type Filter struct {
	Keyword string
}

type tblchannel struct {
	Id                 int
	ChannelName        string
	ChannelDescription string
	SlugName           string
	FieldGroupId       int
	IsActive           int
	IsDeleted          int
	CreatedOn          time.Time
	CreatedBy          int
	ModifiedOn         time.Time           `gorm:"DEFAULT:NULL"`
	ModifiedBy         int                 `gorm:"DEFAULT:NULL"`
	DateString         string              `gorm:"-"`
	EntriesCount       int                 `gorm:"-"`
	ChannelEntries     []TblChannelEntries `gorm:"-"`
	ProfileImagePath   string              `gorm:"<-:false"`
	Username           string              `gorm:"<-:false"`
}

type tblchannelcategory struct {
	Id         int
	ChannelId  int
	CategoryId string
	CreatedAt  int
	CreatedOn  time.Time
}

type ChannelStruct struct{}

var CH ChannelStruct

// soft delete check
func IsDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = 0")
}

/*channel list*/
func (Ch ChannelStruct) Channellist(limit, offset int, filter Filter, activestatus bool, DB *gorm.DB) (chn []tblchannel, chcount int64, err error) {

	query := DB.Table("tbl_channels").Scopes(IsDeleted).Order("id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(channel_name)) ILIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if activestatus {

		query = query.Where("is_active=1")

	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Order("id asc").Find(&chn)

	} else {

		query.Find(&chn).Count(&chcount)

		return chn, chcount, nil
	}

	return chn, 0, nil
}

func (Ch ChannelStruct) GetChannelByChannelName(ch *tblchannel, name string, DB *gorm.DB) error {

	if err := DB.Table("tbl_channels").Where("channel_name=? and is_deleted=0", name).First(&ch).Error; err != nil {

		return err
	}

	return nil
}
