package channels

import (
	"time"

	"github.com/spurtcms/categories"
	"github.com/spurtcms/team"
	permission "github.com/spurtcms/team-roles"
	"gorm.io/gorm"
)

type Filter struct {
	Keyword    string
	Status     bool
	CreateOnly bool
}

type Channels struct {
	Id             int
	Slug           string
	Limit          int
	Offset         int
	Keyword        string
	IsActive       bool
	SortBy         string
	SortingOrder   int
	TenantId       int
	EntriesCount   bool
	ChannelEntries bool
	AuthorDetail   bool
	CreateOnly     bool
	Count          bool
}

type Tblchannel struct {
	Id                 int                 `gorm:"column:id"`
	ChannelName        string              `gorm:"column:channel_name"`
	ChannelDescription string              `gorm:"column:channel_description"`
	SlugName           string              `gorm:"column:slug_name"`
	FieldGroupId       int                 `gorm:"column:field_group_id"`
	IsActive           int                 `gorm:"column:is_active"`
	IsDeleted          int                 `gorm:"column:is_deleted"`
	CreatedOn          time.Time           `gorm:"column:created_on"`
	CreatedBy          int                 `gorm:"column:created_by"`
	ModifiedOn         time.Time           `gorm:"column:modified_on;DEFAULT:NULL"`
	ModifiedBy         int                 `gorm:"column:modified_by;DEFAULT:NULL"`
	DateString         string              `gorm:"-"`
	EntriesCount       int                 `gorm:"-"`
	ChannelEntries     []TblChannelEntries `gorm:"foreignKey:ChannelId"`
	ProfileImagePath   string              `gorm:"<-:false"`
	AuthorDetails      team.TblUser        `gorm:"foreignKey:Id;references:CreatedBy"`
	ChannelType        string              `gorm:"column:channel_type"`
	TenantId           int                 `gorm:"column:tenant_id"`
	Username           string              `gorm:"<-:false"`
	FirstName          string              `gorm:"<-:false"`
	LastName           string              `gorm:"<-:false"`
}

type tblchannelcategory struct {
	Id         int       `gorm:"column:id"`
	ChannelId  int       `gorm:"column:channel_id"`
	CategoryId string    `gorm:"column:category_id"`
	CreatedAt  int       `gorm:"column:created_at"`
	CreatedOn  time.Time `gorm:"column:created_on"`
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
	OrderIndex int    `json:"OrderIndex"`
}
type TblFieldOption struct {
	Id          int       `gorm:"column:id"`
	OptionName  string    `gorm:"column:option_name"`
	OptionValue string    `gorm:"column:option_value"`
	FieldId     int       `gorm:"column:field_id"`
	CreatedOn   time.Time `gorm:"column:created_on"`
	CreatedBy   int       `gorm:"column:created_by"`
	ModifiedOn  time.Time `gorm:"column:modified_on;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"column:modified_by;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"column:is_deleted;DEFAULT:0"`
	DeletedOn   time.Time `gorm:"column:deleted_on;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"column:deleted_by;DEFAULT:NULL"`
	Idstring    string    `gorm:"-:migration;<-:false"`
	TenantId    int       `gorm:"type:integer"`
	OrderIndex  int       `gorm:"type:integer"`
}

type tblfield struct {
	Id               int                  `gorm:"column:id"`
	FieldName        string               `gorm:"column:field_name"`
	FieldDesc        string               `gorm:"column:field_desc"`
	FieldTypeId      int                  `gorm:"column:field_type_id"`
	MandatoryField   int                  `gorm:"column:mandatory_field"`
	OptionExist      int                  `gorm:"column:option_exist"`
	InitialValue     string               `gorm:"column:initial_value"`
	Placeholder      string               `gorm:"column:place_holder"`
	CreatedOn        time.Time            `gorm:"column:created_on"`
	CreatedBy        int                  `gorm:"column:created_by"`
	ModifiedOn       time.Time            `gorm:"column:modified_on;DEFAULT:NULL"`
	ModifiedBy       int                  `gorm:"column:modified_by;DEFAULT:NULL"`
	IsDeleted        int                  `gorm:"column:is_deleted;DEFAULT:0"`
	DeletedOn        time.Time            `gorm:"column:deleted_on;DEFAULT:NULL"`
	DeletedBy        int                  `gorm:"column:deleted_by;DEFAULT:NULL"`
	OrderIndex       int                  `gorm:"column:order_index"`
	ImagePath        string               `gorm:"column:image_path"`
	TblFieldOption   []TblFieldOption     `gorm:"<-:false; foreignKey:FieldId"`
	DatetimeFormat   string               `gorm:"column:datetime_format"`
	TimeFormat       string               `gorm:"column:time_format"`
	Url              string               `gorm:"column:url"`
	Values           string               `gorm:"-"`
	CheckBoxValue    []FieldValueId       `gorm:"-"`
	SectionParentId  int                  `gorm:"column:section_parent_id"`
	FieldTypeName    string               `gorm:"-;column:type_name"`
	CharacterAllowed int                  `gorm:"column:character_allowed"`
	FieldOptions     []TblFieldOption     `gorm:"-"`
	FieldValue       TblChannelEntryField `gorm:"-"`
	TenantId         int
}

type TblChannel struct {
	Id                 int       `gorm:"primaryKey;auto_increment;type:serial"`
	ChannelName        string    `gorm:"type:character varying"`
	ChannelDescription string    `gorm:"type:character varying"`
	SlugName           string    `gorm:"type:character varying"`
	FieldGroupId       int       `gorm:"type:integer"`
	IsActive           int       `gorm:"type:integer"`
	IsDeleted          int       `gorm:"type:integer"`
	CreatedOn          time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy          int       `gorm:"type:integer"`
	ModifiedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL"`
	TenantId           int       `gorm:"type:integer"`
	ChannelType        string    `gorm:"type:character varying"`
}

type TblChannelCategorie struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	ChannelId  int       `gorm:"type:integer"`
	CategoryId string    `gorm:"type:character varying"`
	CreatedAt  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
	TenantId   int       `gorm:"type:integer"`
}

type TblGroupField struct {
	Id        int `gorm:"primaryKey;auto_increment;type:serial"`
	ChannelId int `gorm:"type:integer"`
	FieldId   int `gorm:"type:integer"`
	TenantId  int `gorm:"type:integer"`
}

type TblChannelEntries struct {
	Id              int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title           string    `gorm:"type:character varying"`
	Slug            string    `gorm:"type:character varying"`
	Description     string    `gorm:"type:character varying"`
	UserId          int       `gorm:"type:integer"`
	ChannelId       int       `gorm:"type:integer"`
	Status          int       `gorm:"type:integer"` //0-draft 1-publish 2-unpublish
	CoverImage      string    `gorm:"type:character varying"`
	ThumbnailImage  string    `gorm:"type:character varying"`
	MetaTitle       string    `gorm:"type:character varying"`
	MetaDescription string    `gorm:"type:character varying"`
	Keyword         string    `gorm:"type:character varying"`
	CategoriesId    string    `gorm:"type:character varying"`
	RelatedArticles string    `gorm:"type:character varying"`
	Feature         int       `gorm:"DEFAULT:0"`
	ViewCount       int       `gorm:"DEFAULT:0"`
	Author          string    `gorm:"type:character varying"`
	SortOrder       int       `gorm:"type:integer"`
	CreateTime      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	PublishedTime   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ReadingTime     int       `gorm:"type:integer;DEFAULT:0"`
	Tags            string    `gorm:"type:character varying"`
	Excerpt         string    `gorm:"type:character varying"`
	ImageAltTag     string    `gorm:"type:character varying"`
	CreatedOn       time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy       int       `gorm:"type:integer"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL"`
	ModifiedOn      time.Time `gorm:"DEFAULT:NULL"`
	IsActive        int       `gorm:"type:integer"`
	IsDeleted       int       `gorm:"DEFAULT:0"`
	DeletedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"DEFAULT:NULL"`
	Uuid            string    `gorm:"character varying"`
	OrderIndex      int       `gorm:"type:integer"`
}

type TblChannelEntryField struct {
	Id             int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldName      string    `gorm:"type:character varying"`
	FieldValue     string    `gorm:"type:character varying"`
	FieldTypeId    int       `gorm:"-:migration;<-:false"`
	ChannelEntryId int       `gorm:"type:integer"`
	FieldId        int       `gorm:"type:integer"`
	CreatedOn      time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy      int       `gorm:"type:integer"`
	ModifiedOn     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"DEFAULT:NULL"`
	DeletedBy      int       `gorm:"DEFAULT:NULL"`
	DeletedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId       int       `gorm:"type:integer"`
}

type TblField struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldName        string    `gorm:"type:character varying"`
	FieldDesc        string    `gorm:"type:character varying"`
	FieldTypeId      int       `gorm:"type:integer"`
	MandatoryField   int       `gorm:"type:integer"`
	OptionExist      int       `gorm:"type:integer"`
	InitialValue     string    `gorm:"type:character varying"`
	Placeholder      string    `gorm:"type:character varying"`
	OrderIndex       int       `gorm:"type:integer"`
	ImagePath        string    `gorm:"type:character varying"`
	DatetimeFormat   string    `gorm:"type:character varying"`
	TimeFormat       string    `gorm:"type:character varying"`
	Url              string    `gorm:"type:character varying"`
	SectionParentId  int       `gorm:"type:integer"`
	CharacterAllowed int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	IsDeleted        int       `gorm:"DEFAULT:0"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	TenantId         int       `gorm:"type:integer"`
}

type TblFieldGroup struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	GroupName  string    `gorm:"type:character varying"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL"`
	IsDeleted  int       `gorm:"DEFAULT:0"`
	DeletedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"DEFAULT:NULL"`
}

// type TblFieldOption struct {
// 	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
// 	OptionName  string    `gorm:"type:character varying"`
// 	OptionValue string    `gorm:"type:character varying"`
// 	FieldId     int       `gorm:"type:integer"`
// 	CreatedOn   time.Time `gorm:"type:timestamp without time zone"`
// 	CreatedBy   int       `gorm:"type:integer"`
// 	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
// 	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
// 	IsDeleted   int       `gorm:"DEFAULT:0"`
// 	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
// 	DeletedBy   int       `gorm:"DEFAULT:NULL"`
// }

type TblFieldType struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	TypeName   string    `gorm:"type:character varying"`
	TypeSlug   string    `gorm:"type:character varying"`
	IsActive   int       `gorm:"type:integer"`
	IsDeleted  int       `gorm:"type:integer"`
	CreatedBy  int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
	ModifiedBy int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type ChannelUpdate struct {
	Sections           []Section
	FieldValues        []Fiedlvalue
	Deletesections     []Section
	DeleteFields       []Fiedlvalue
	DeleteOptionsvalue []OptionValues
	CategoryIds        []string
	ModifiedBy         int
}

type ChannelCreate struct {
	ChannelName        string
	ChannelDescription string
	CategoryIds        []string
	CreatedBy          int
}

type ChannelAddtionalField struct {
	Sections    []Section
	FieldValues []Fiedlvalue
	CreatedBy   int
}

type FieldValueId struct {
	Id     int
	CValue int
}

type ChannelModel struct {
	Userid     int
	Dataaccess int
}

var CH ChannelModel

// soft delete check
func IsDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = 0")
}

/*channel list*/
func (Ch ChannelModel) Channellist(DB *gorm.DB,channel *Channel, inputs Channels, channels *[]Tblchannel, count *int64) (err error) {

	query := DB.Debug().Table("tbl_channels").Where("tbl_channels.is_deleted = 0")

	if channel.PermissionEnable && (channel.Auth.RoleId != 1 && channel.Auth.RoleId != 2) {

		query = query.Debug().Where("channel_name in (select display_name from tbl_module_permissions inner join tbl_modules on tbl_modules.id = tbl_module_permissions.module_id inner join tbl_role_permissions on tbl_role_permissions.permission_id = tbl_module_permissions.id where role_id =(?) and tbl_modules.module_name='Entries') ", channel.Auth.RoleId)

	}

	if inputs.TenantId != -1 {

		query = query.Where("tbl_channels.tenant_id=?", inputs.TenantId)

	}

	if inputs.CreateOnly && Ch.Dataaccess == 1 {

		query = query.Where("tbl_channels.created_by = ?", Ch.Userid)
	}

	if inputs.Keyword != "" {

		query = query.Where("LOWER(TRIM(channel_name)) LIKE LOWER(TRIM(?))", "%"+inputs.Keyword+"%")
	}

	if inputs.IsActive {

		query = query.Where("tbl_channels.is_active=1")

	}

	if inputs.Count {

		err = query.Count(count).Error

		if err != nil {

			return err
		}
	}

	if inputs.AuthorDetail {

		query = query.Preload("AuthorDetails", "is_deleted = ?", 0)
	}

	if inputs.ChannelEntries {

		query = query.Preload("ChannelEntries", "is_deleted = ?", 0)
	}

	if inputs.SortBy != "" {

		if inputs.SortingOrder == 0 {

			query = query.Order(inputs.SortingOrder)

		} else if inputs.SortingOrder == 1 {

			query = query.Order(inputs.SortBy + " desc")

		}

	} else {

		query = query.Order("id desc")
	}

	if inputs.Limit != 0 {

		query = query.Limit(inputs.Limit)
	}

	if inputs.Offset != -1 {

		query = query.Offset(inputs.Offset)
	}

	err = query.Find(&channels).Error

	if err != nil {

		return err
	}

	return nil
}

/*Craete channel */
func (Ch ChannelModel) CreateChannel(chn *TblChannel, DB *gorm.DB) (TblChannel, error) {

	if err := DB.Table("tbl_channels").Create(&chn).Error; err != nil {

		return TblChannel{}, err

	}

	return *chn, nil

}

func (Ch ChannelModel) ChannelDetail(DB *gorm.DB, inputs Channels, channelDetail *Tblchannel) error {

	query := DB.Debug().Model(TblChannel{}).Where("tbl_channels.is_deleted = 0")

	if inputs.Id != 0 {

		query = query.Where("id=?", inputs.Id)
	}

	if inputs.Slug != "" {

		query = query.Where("slug_name=?", inputs.Slug)
	}

	if inputs.TenantId != -1 {

		query = query.Where("tbl_channels.tenant_id=?", inputs.TenantId)
	}

	if inputs.CreateOnly && Ch.Dataaccess == 1 {

		query = query.Where("tbl_channels.created_by = ?", Ch.Userid)
	}

	if inputs.Keyword != "" {

		query = query.Where("LOWER(TRIM(channel_name)) LIKE LOWER(TRIM(?))", "%"+inputs.Keyword+"%")
	}

	if inputs.IsActive {

		query = query.Where("tbl_channels.is_active=?", 1)
	}

	if inputs.AuthorDetail {

		query = query.Preload("AuthorDetails", "is_deleted = ?", 0)
	}

	if inputs.ChannelEntries {

		query = query.Preload("ChannelEntries", "is_deleted = ?", 0)
	}

	err := query.First(&channelDetail).Error

	if err != nil {

		return err
	}

	return nil
}

func (Ch ChannelModel) GetChannelByChannelName(name string, DB *gorm.DB, tenantid int) (ch Tblchannel, err error) {

	if err := DB.Table("tbl_channels").Where("channel_name=? and tenant_id=? and is_deleted=0", name, tenantid).First(&ch).Error; err != nil {

		return Tblchannel{}, err
	}

	return ch, nil
}

/*Get Channel*/
func (Ch ChannelModel) GetChannelById(id int, DB *gorm.DB, tenantid int) (ch Tblchannel, err error) {

	if err := DB.Table("tbl_channels").Where("id=? and tenant_id=?", id, tenantid).First(&ch).Error; err != nil {

		return Tblchannel{}, err
	}

	return ch, nil
}

/*Getfieldid using fieldgroupid*/
func (Ch ChannelModel) GetFieldIdByGroupId(id int, DB *gorm.DB, tenantid int) (grpfield []TblGroupField, err error) {

	if err := DB.Table("tbl_group_fields").Where("channel_id=? and tenant_id=?", id, tenantid).Find(&grpfield).Error; err != nil {

		return []TblGroupField{}, err
	}

	return grpfield, nil
}

/*Get optionvalue*/
func (Ch ChannelModel) GetFieldAndOptionValue(id []int, DB *gorm.DB, tenantid int) (fld []tblfield, err error) {

	if err := DB.Table("tbl_fields").Where("id in (?) and tenant_id=? and is_deleted != 1", id, tenantid).Preload("TblFieldOption", func(db *gorm.DB) *gorm.DB {
		return DB.Where("is_deleted!=1").Order("order_index asc")
	}).Order("order_index asc").Find(&fld).Error; err != nil {

		return []tblfield{}, err
	}

	return fld, nil
}

func (Ch ChannelModel) GetSelectedCategoryChannelById(id int, DB *gorm.DB, tenantid int) (ChannelCategory []tblchannelcategory, err error) {

	if err := DB.Table("tbl_channel_categories").Where("channel_id=? and tenant_id=?", id, tenantid).Find(&ChannelCategory).Error; err != nil {

		return []tblchannelcategory{}, err
	}

	return ChannelCategory, nil

}

func (Ch ChannelModel) GetCategoriseById(id []int, DB *gorm.DB, tenantid int) (category []categories.TblCategories, err error) {

	if err := DB.Table("tbl_categories").Where("id in (?) and tenant_id=?", id, tenantid).Order("id asc").Find(&category).Error; err != nil {

		return category, err
	}

	return category, nil

}

func (Ch ChannelModel) DeleteEntryByChannelId(id int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_channel_entries").Where("channel_id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1}).Error; err != nil {

		return err
	}

	return nil

}

/*Delete Channel*/
func (Ch ChannelModel) DeleteChannelById(id int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_channels").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1}).Error; err != nil {

		return err
	}

	return nil
}

/*Delete Channel*/
func (Ch ChannelModel) DeleteFieldGroupById(tblfieldgrp *TblFieldGroup, id int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_field_groups").Where("id=?  and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": tblfieldgrp.IsDeleted, "deleted_by": tblfieldgrp.DeletedBy, "deleted_on": tblfieldgrp.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*Isactive channel*/
func (Ch ChannelModel) ChannelIsActive(tblch *TblChannel, id, val int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_channels").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_active": val, "modified_on": tblch.ModifiedOn, "modified_by": tblch.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil
}

/*Get all master fields*/
func (Ch ChannelModel) GetAllField(DB *gorm.DB, tenantid int) (channel []TblFieldType, err error) {

	if err := DB.Table("tbl_field_types").Where("is_deleted=0").Find(&channel).Error; err != nil {

		return []TblFieldType{}, err
	}
	return channel, nil
}

/*Update Channel Details*/
func (Ch ChannelModel) UpdateChannelDetails(chn *TblChannel, id int, DB *gorm.DB, TenantId int) error {

	if err := DB.Table("tbl_channels").Where("id=? and tenant_id=?", id, TenantId).UpdateColumns(map[string]interface{}{"channel_name": chn.ChannelName, "channel_description": chn.ChannelDescription, "modified_by": chn.ModifiedBy, "modified_on": chn.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*Update Field Details*/
func (Ch ChannelModel) UpdateFieldDetails(fds *TblField, id int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_fields").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"field_name": fds.FieldName, "field_desc": fds.FieldDesc, "mandatory_field": fds.MandatoryField, "datetime_format": fds.DatetimeFormat, "time_format": fds.TimeFormat, "initial_value": fds.InitialValue, "placeholder": fds.Placeholder, "modified_on": fds.ModifiedOn, "modified_by": fds.ModifiedBy, "order_index": fds.OrderIndex, "url": fds.Url, "character_allowed": fds.CharacterAllowed}).Error; err != nil {

		return err
	}

	return nil
}

/*CheckCategoryId Already Exists*/
func (Ch ChannelModel) CheckChannelCategoryAlreadyExitst(channelid int, categoryids string, DB *gorm.DB, tenantid int) error {

	var category tblchannelcategory

	if err := DB.Table("tbl_channel_categories").Where("channel_id=? and category_id=? and tenant_id=?", channelid, categoryids, tenantid).First(&category).Error; err != nil {

		return err
	}

	return nil

}

/*Create Channel Categories*/
func (Ch ChannelModel) CreateChannelCategory(channelcategory *TblChannelCategorie, DB *gorm.DB) error {

	if err := DB.Model(TblChannelCategorie{}).Create(&channelcategory).Error; err != nil {

		return err
	}

	return nil

}

/*update channel entry permission*/
func (Ch ChannelModel) UpdateChannelNameInEntries(modper *permission.TblModulePermission, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_module_permissions").Where("route_name=? and tenant_id=?", modper.RouteName, tenantid).UpdateColumns(map[string]interface{}{
		"display_name": modper.DisplayName, "slug_name": modper.SlugName}).Error; err != nil {

		return err
	}

	return nil
}

/**/
func (Ch ChannelModel) GetChannelCategoryNotExist(category *[]tblchannelcategory, channelid int, categoryids []string, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_channel_categories").Where("channel_id=? and category_id not in (?) and tenant_id=?", channelid, categoryids, tenantid).Find(&category).Error; err != nil {

		return err
	}

	return nil
}

/*Delete Channel Category*/
func (Ch ChannelModel) DeleteChannelCategoryByValue(category *tblchannelcategory, id int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_channel_categories").Where("id=? and tenant_id=?", id, tenantid).Delete(&category).Error; err != nil {

		return err
	}

	return nil
}

/*Delete Field By Id*/
func (Ch ChannelModel) DeleteFieldById(field *TblField, id []int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_fields").Where("id in(?) and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": field.DeletedBy, "deleted_on": field.DeletedOn}).Error; err != nil {

		return err
	}

	return nil
}

/*Delete FieldOption By fieldid*/
func (Ch ChannelModel) DeleteFieldOptionById(fieldopt *TblFieldOption, id []string, fid int, DB *gorm.DB, tenantid int) error {

	if len(id) > 0 {

		if err := DB.Table("tbl_field_options").Where("option_name not in (?) and field_id=? and tenant_id=?", id, fid, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": fieldopt.DeletedBy, "deleted_on": fieldopt.DeletedOn}).Error; err != nil {

			return err
		}

	} else {

		if err := DB.Table("tbl_field_options").Where("field_id=? and tenant_id=?", fid, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": fieldopt.DeletedBy, "deleted_on": fieldopt.DeletedOn}).Error; err != nil {

			return err
		}
	}

	return nil

}

/*Delete FieldOption By fieldid*/
func (Ch ChannelModel) DeleteOptionById(fieldopt *TblFieldOption, id []int, fid []int, DB *gorm.DB, tenantid int) error {

	if len(id) > 0 {

		if err := DB.Table("tbl_field_options").Where("id in (?) and tenant_id=? ", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": fieldopt.DeletedBy, "deleted_on": fieldopt.DeletedOn}).Error; err != nil {

			return err
		}

	} else {

		if err := DB.Table("tbl_field_options").Where("field_id in (?) and tenant_id=?", fid, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": fieldopt.DeletedBy, "deleted_on": fieldopt.DeletedOn}).Error; err != nil {

			return err
		}
	}

	return nil

}

/*create field*/
func (Ch ChannelModel) CreateFields(flds *TblField, DB *gorm.DB) (*TblField, error) {

	if err := DB.Table("tbl_fields").Create(&flds).Error; err != nil {

		return &TblField{}, err
	}

	return flds, nil
}

func (Ch ChannelModel) CreateGroupField(grpfield *TblGroupField, DB *gorm.DB) error {

	if err := DB.Table("tbl_group_fields").Create(&grpfield).Error; err != nil {

		return err
	}

	return nil

}

/*Update Field Option Details*/
func (Ch ChannelModel) UpdateFieldOption(fdoption *TblFieldOption, id int, DB *gorm.DB, tenantid int) error {

	if err := DB.Debug().Table("tbl_field_options").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"option_name": fdoption.OptionName, "option_value": fdoption.OptionValue, "modified_on": fdoption.ModifiedOn, "modified_by": fdoption.ModifiedBy, "order_index": fdoption.OrderIndex}).Error; err != nil {

		return err
	}

	return nil
}

/*create option value*/
func (Ch ChannelModel) CreateFieldOption(optval *TblFieldOption, DB *gorm.DB) error {

	if err := DB.Table("tbl_field_options").Create(&optval).Error; err != nil {

		return err
	}

	return nil
}

func (ch ChannelModel) GetChannelCount(count *int64, DB *gorm.DB, tenantid int) error {

	if err := DB.Debug().Table("tbl_channels").Distinct("tbl_channels.id").Joins("inner join tbl_channel_entries on tbl_channel_entries.channel_id = tbl_channels.id").
		Joins("inner join tbl_channel_categories on tbl_channel_categories.channel_id = tbl_channels.id").
		Where("tbl_channels.is_deleted = 0 and tbl_channels.is_active = 1 and tbl_channel_entries.status = 1 and tbl_channel_entries.tenant_id=?", tenantid).Count(count).Error; err != nil {

		return err
	}

	return nil
}

func (ch ChannelModel) GetChannels(channels *[]Tblchannel, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_channels").Where("is_deleted = 0 and is_active = 1 and tenant_id=?", tenantid).Find(&channels).Error; err != nil {

		return err
	}

	return nil
}

func (ch ChannelModel) GetPermissionChannel(channels *Channel, DB *gorm.DB, tenantid int) (channel []Tblchannel, err error) {

	query := DB.Debug().Table("tbl_channels").Where("is_deleted=0 and is_active=1 and tenant_id=?", tenantid).Order("id desc")

	if channels.PermissionEnable && (channels.Auth.RoleId != 1 && channels.Auth.RoleId != 2) {

		query = query.Where("channel_name in (select display_name from tbl_module_permissions inner join tbl_modules on tbl_modules.id = tbl_module_permissions.module_id inner join tbl_role_permissions on tbl_role_permissions.permission_id = tbl_module_permissions.id where role_id =(?) and tbl_modules.module_name='Entries') ", channels.Auth.RoleId)

	}

	if err := query.Find(&channel).Error; err != nil {

		return channel, err
	}

	return channel, nil
}

// Channel type change
func (ch ChannelModel) ChangeChanelType(Channels Tblchannel, DB *gorm.DB) (Error error) {

	if err := DB.Debug().Table("tbl_channels").Where("id=?", Channels.Id).Updates(map[string]interface{}{"channel_type": Channels.ChannelType}).Error; err != nil {

		return err

	}

	return nil
}

func (ch ChannelModel) ModulePermissionChannelDelete(routename string, DB *gorm.DB, tenantid int) error {
	result := DB.Table("tbl_module_permissions").Where("route_name = ? AND tenant_id = ?", routename, tenantid).Delete(nil)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Total Channel Count

func (ch ChannelModel) AllChannelCount(DB *gorm.DB, tenantid int) (count int64, err error) {

	if err := DB.Table("tbl_channels").Where("tbl_channels.is_deleted = 0 and  tenant_id = ?", tenantid).Count(&count).Error; err != nil {

		return 0, err
	}

	return count, nil

}

// Last 10 days Channel Count
func (ch ChannelModel) NewChannelCount(DB *gorm.DB, tenantid int) (count int64, err error) {

	if err := DB.Table("tbl_channels").Where("created_on >=? and  tenant_id=? and is_deleted = 0", time.Now().AddDate(0, 0, -10), tenantid).Count(&count).Error; err != nil {

		return 0, err
	}

	return count, nil

}
