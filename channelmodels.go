package channels

import (
	"time"

	"github.com/spurtcms/categories"
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
	ChannelEntries     []tblchannelentries `gorm:"-"`
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

type Section struct {
	SectionId     int    `json:"SectionId"`
	SectionNewId  int    `json:"SectionNewId"`
	SectionName   string `json:"SectionName"`
	MasterFieldId int    `json:"MasterFieldId"`
	OrderIndex    int    `json:"OrderIndex"`
}

type Fiedlvalue struct {
	MasterFieldId    int            `json:"MasterFieldId"`
	FieldId          int            `json:"FieldId"`
	NewFieldId       int            `json:"NewFieldId"`
	SectionId        int            `json:"SectionId"`
	SectionNewId     int            `json:"SectionNewId"`
	FieldName        string         `json:"FieldName"`
	DateFormat       string         `json:"DateFormat"`
	TimeFormat       string         `json:"TimeFormat"`
	OptionValue      []OptionValues `json:"OptionValue"`
	CharacterAllowed int            `json:"CharacterAllowed"`
	IconPath         string         `json:"IconPath"`
	Url              string         `json:"Url"`
	OrderIndex       int            `json:"OrderIndex"`
	Mandatory        int            `json:"Mandatory"`
}

type OptionValues struct {
	Id         int    `json:"Id"`
	NewId      int    `json:"NewId"`
	FieldId    int    `json:"FieldId"`
	NewFieldId int    `json:"NewFieldId"`
	Value      string `json:"Value"`
}
type tblfieldoption struct {
	Id          int
	OptionName  string
	OptionValue string
	FieldId     int
	CreatedOn   time.Time
	CreatedBy   int
	ModifiedOn  time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted   int       `gorm:"DEFAULT:0"`
	DeletedOn   time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
	Idstring    string    `gorm:"-"`
}

type tblfield struct {
	Id               int
	FieldName        string
	FieldDesc        string
	FieldTypeId      int
	MandatoryField   int
	OptionExist      int
	InitialValue     string
	Placeholder      string
	CreatedOn        time.Time
	CreatedBy        int
	ModifiedOn       time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	IsDeleted        int       `gorm:"DEFAULT:0"`
	DeletedOn        time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	OrderIndex       int
	ImagePath        string
	TblFieldOption   []tblfieldoption `gorm:"<-:false; foreignKey:FieldId"`
	DatetimeFormat   string
	TimeFormat       string
	Url              string
	Values           string         `gorm:"-"`
	CheckBoxValue    []FieldValueId `gorm:"-"`
	SectionParentId  int
	FieldTypeName    string `gorm:"-;column:type_name"`
	CharacterAllowed int
	FieldOptions     []tblfieldoption     `gorm:"-"`
	FieldValue       TblChannelEntryField `gorm:"-"`
}

type FieldValueId struct {
	Id     int
	CValue int
}

type ChannelModel struct{}

var CH ChannelModel

// soft delete check
func IsDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = 0")
}

/*channel list*/
func (Ch ChannelModel) Channellist(limit, offset int, filter Filter, activestatus bool, DB *gorm.DB) (chn []tblchannel, chcount int64, err error) {

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

func (Ch ChannelModel) GetChannelByChannelName(name string, DB *gorm.DB) (ch tblchannel, err error) {

	if err := DB.Table("tbl_channels").Where("channel_name=? and is_deleted=0", name).First(&ch).Error; err != nil {

		return tblchannel{}, err
	}

	return ch, nil
}

/*Get Channel*/
func (Ch ChannelModel) GetChannelById(id int, DB *gorm.DB) (ch tblchannel, err error) {

	if err := DB.Table("tbl_channels").Where("id=?", id).First(&ch).Error; err != nil {

		return tblchannel{}, err
	}

	return ch, nil
}

/*Getfieldid using fieldgroupid*/
func (Ch ChannelModel) GetFieldIdByGroupId(id int, DB *gorm.DB) (grpfield []TblGroupField, err error) {

	if err := DB.Table("tbl_group_fields").Where("channel_id=?", id).Find(&grpfield).Error; err != nil {

		return []TblGroupField{}, err
	}

	return grpfield, nil
}

/*Get optionvalue*/
func (Ch ChannelModel) GetFieldAndOptionValue(id []int, DB *gorm.DB) (fld []tblfield, err error) {

	if err := DB.Table("tbl_fields").Where("id in (?) and is_deleted != 1", id).Preload("TblFieldOption", func(db *gorm.DB) *gorm.DB {
		return DB.Where("is_deleted!=1")
	}).Order("order_index asc").Find(&fld).Error; err != nil {

		return []tblfield{}, err
	}

	return fld, nil
}

func (Ch ChannelModel) GetSelectedCategoryChannelById(id int, DB *gorm.DB) (ChannelCategory []tblchannelcategory, err error) {

	if err := DB.Table("tbl_channel_categories").Where("channel_id=?", id).Find(&ChannelCategory).Error; err != nil {

		return []tblchannelcategory{}, err
	}

	return ChannelCategory, nil

}

func (Ch ChannelModel) GetCategoriseById(id []int, DB *gorm.DB) (category []categories.TblCategories, err error) {

	if err := DB.Table("tbl_categories").Where("id in (?)", id).Order("id asc").Find(&category).Error; err != nil {

		return category, err
	}

	return category, nil

}

func (Ch ChannelModel) DeleteEntryByChannelId(id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_channel_entries").Where("channel_id=?", id).UpdateColumns(map[string]interface{}{"is_deleted": 1}).Error; err != nil {

		return err
	}

	return nil

}

/*Delete Channel*/
func (Ch ChannelModel) DeleteChannelById(id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_channels").Where("id=?", id).UpdateColumns(map[string]interface{}{"is_deleted": 1}).Error; err != nil {

		return err
	}

	return nil
}

/*Delete Channel*/
func (Ch ChannelModel) DeleteFieldGroupById(tblfieldgrp *TblFieldGroup, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_field_groups").Where("id=?", id).UpdateColumns(map[string]interface{}{"is_deleted": tblfieldgrp.IsDeleted, "deleted_by": tblfieldgrp.DeletedBy, "deleted_on": tblfieldgrp.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*Isactive channel*/
func (Ch ChannelModel) ChannelIsActive(tblch *TblChannel, id, val int, DB *gorm.DB) error {

	if err := DB.Table("tbl_channels").Where("id=?", id).UpdateColumns(map[string]interface{}{"is_active": val, "modified_on": tblch.ModifiedOn, "modified_by": tblch.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil
}