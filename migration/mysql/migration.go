package mysql

import (
	"time"

	"gorm.io/gorm"
)

type TblChannel struct {
	Id                 int       `gorm:"primaryKey;auto_increment"`
	ChannelName        string    `gorm:"type:varchar(255)"`
	ChannelDescription string    `gorm:"type:varchar(255)"`
	SlugName           string    `gorm:"type:varchar(255)"`
	FieldGroupId       int       `gorm:"type:int"`
	IsActive           int       `gorm:"type:int"`
	IsDeleted          int       `gorm:"type:int"`
	CreatedOn          time.Time `gorm:"type:datetime"`
	CreatedBy          int       `gorm:"type:int"`
	ModifiedOn         time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId           int       `gorm:"type:int"`
}

type TblChannelCategory struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	ChannelId  int       `gorm:"type:int"`
	CategoryId string    `gorm:"type:varchar(255)"`
	CreatedAt  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime"`
	TenantId   int       `gorm:"type:int"`
}

type TblGroupField struct {
	Id        int `gorm:"primaryKey;auto_increment"`
	ChannelId int `gorm:"type:int"`
	FieldId   int `gorm:"type:int"`
	TenantId  int `gorm:"type:int"`
}

type TblChannelEntry struct {
	Id              int       `gorm:"primaryKey;auto_increment"`
	Title           string    `gorm:"type:varchar(255)"`
	Slug            string    `gorm:"type:varchar(255)"`
	Description     string    `gorm:"type:text"`
	UserId          int       `gorm:"type:int"`
	ChannelId       int       `gorm:"type:int"`
	Status          int       `gorm:"type:int"` //0-draft 1-publish 2-unpublish
	CoverImage      string    `gorm:"type:varchar(255)"`
	ThumbnailImage  string    `gorm:"type:varchar(255)"`
	MetaTitle       string    `gorm:"type:varchar(255)"`
	MetaDescription string    `gorm:"type:varchar(255)"`
	Keyword         string    `gorm:"type:varchar(255)"`
	CategoriesId    string    `gorm:"type:varchar(255)"`
	RelatedArticles string    `gorm:"type:varchar(255)"`
	Feature         int       `gorm:"DEFAULT:0"`
	ViewCount       int       `gorm:"DEFAULT:0"`
	CreateTime      time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	PublishedTime   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ImageAltTag     string    `gorm:"type:varchar(255)"`
	Author          string    `gorm:"type:varchar(255)"`
	SortOrder       int       `gorm:"type:int"`
	Excerpt         string    `gorm:"type:varchar(255)"`
	ReadingTime     int       `gorm:"type:int"`
	Tags            string    `gorm:"type:varchar(255)"`
	CreatedOn       time.Time `gorm:"type:datetime"`
	CreatedBy       int       `gorm:"type:int"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn      time.Time `gorm:"DEFAULT:NULL;type:int"`
	IsActive        int       `gorm:"type:int"`
	IsDeleted       int       `gorm:"DEFAULT:0"`
	DeletedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId        int       `gorm:"type:int"`
}

type TblChannelEntryField struct {
	Id             int       `gorm:"primaryKey;auto_increment"`
	FieldName      string    `gorm:"type:varchar(255)"`
	FieldValue     string    `gorm:"type:varchar(255)"`
	ChannelEntryId int       `gorm:"type:int"`
	FieldId        int       `gorm:"type:int"`
	CreatedOn      time.Time `gorm:"type:datetime"`
	CreatedBy      int       `gorm:"type:int"`
	ModifiedOn     time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedBy      int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedOn      time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId       int       `gorm:"type:int"`
}

type TblField struct {
	Id               int       `gorm:"primaryKey;auto_increment"`
	FieldName        string    `gorm:"type:varchar(255)"`
	FieldDesc        string    `gorm:"type:varchar(255)"`
	FieldTypeId      int       `gorm:"type:int"`
	MandatoryField   int       `gorm:"type:int"`
	OptionExist      int       `gorm:"type:int"`
	InitialValue     string    `gorm:"type:varchar(255)"`
	Placeholder      string    `gorm:"type:varchar(255)"`
	OrderIndex       int       `gorm:"type:int"`
	ImagePath        string    `gorm:"type:varchar(255)"`
	DatetimeFormat   string    `gorm:"type:varchar(255)"`
	TimeFormat       string    `gorm:"type:varchar(255)"`
	Url              string    `gorm:"type:varchar(255)"`
	SectionParentId  int       `gorm:"type:int"`
	CharacterAllowed int       `gorm:"type:int"`
	CreatedOn        time.Time `gorm:"type:datetime"`
	CreatedBy        int       `gorm:"type:int"`
	ModifiedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted        int       `gorm:"DEFAULT:0"`
	DeletedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId         int       `gorm:"type:int"`
}

type TblFieldGroup struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	GroupName  string    `gorm:"type:varchar(255)"`
	CreatedOn  time.Time `gorm:"type:datetime"`
	CreatedBy  int       `gorm:"type:int"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted  int       `gorm:"DEFAULT:0"`
	DeletedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId   int       `gorm:"type:int"`
}

type TblFieldOption struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	OptionName  string    `gorm:"type:varchar(255)"`
	OptionValue string    `gorm:"type:varchar(255)"`
	FieldId     int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	IsDeleted   int       `gorm:"DEFAULT:0"`
	DeletedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId    int       `gorm:"type:int"`
}

type TblFieldType struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	TypeName   string    `gorm:"type:varchar(255)"`
	TypeSlug   string    `gorm:"type:varchar(255)"`
	IsActive   int       `gorm:"type:int"`
	IsDeleted  int       `gorm:"type:int"`
	CreatedBy  int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime"`
	ModifiedBy int       `gorm:"type:int"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId   int       `gorm:"type:int"`
}

// MigrateTable creates this package related tables in your database
func MigrationTables(db *gorm.DB) {

	if err := db.AutoMigrate(
		&TblChannelCategory{},
		&TblChannelEntry{},
		&TblChannelEntryField{},
		&TblChannel{},
		&TblFieldGroup{},
		&TblFieldOption{},
		&TblFieldType{},
		&TblField{},
	); err != nil {

		panic(err)
	}

	db.Exec(`CREATE INDEX IF NOT EXISTS email_unique
    ON public.tbl_members USING btree
    (email COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default
    WHERE is_deleted = 0;`)

	db.Exec(`CREATE INDEX IF NOT EXISTS mobile_no_unique
    ON public.tbl_members USING btree
    (mobile_no COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default
    WHERE is_deleted = 0;`)

	//create default channel
	db.Exec(`INSERT INTO public.tbl_channels(id, channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by, channel_description) VALUES (1, 'Default_Channel', 'default_channel', 0, 1, 0, '2024-03-04 10:49:17', '1', 'default description');`)

	db.Exec(`INSERT INTO public.tbl_channel_categories(id, channel_id, category_id, created_at, created_on) VALUES (1, 1, '1,2', 1, '2024-03-04 10:49:17');`)

	//Channel default fields
	db.Exec(`INSERT INTO public.tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (1, 'Label', 'label', 1,  0, 1, '2023-03-14 11:09:12'), (2, 'Text', 'text', 1,  0, 1, '2023-03-14 11:09:12'),(3, 'Link', 'link', 1,  0, 1, '2023-03-14 11:09:12'),(4, 'Date & Time', 'date&time', 1,  0, 1, '2023-03-14 11:09:12'), (5, 'Select', 'select', 1,  0, 1, '2023-03-14 11:09:12'),(6, 'Date', 'date', 1,  0, 1, '2023-03-14 11:09:12'),(7, 'TextBox', 'textbox', 1,  0, 1, '2023-03-14 11:09:12'),(8, 'TextArea', 'textarea', 1, 0, 1, '2023-03-14 11:09:12'), (9, 'Radio Button', 'radiobutton', 1, 0, 1, '2023-03-14 11:09:12'),(10, 'CheckBox', 'checkbox', 1, 0, 1, '2023-03-14 11:09:12'),(11, 'Text Editor', 'texteditor', 1, 0, 1, '2023-03-14 11:09:12'),(12, 'Section', 'section', 1, 0, 1, '2023-03-14 11:09:12'),(13, 'Section Break', 'sectionbreak', 1, 0, 1, '2023-03-14 11:09:12'),(14, 'Members', 'member', 1,  0, 1, '2023-03-14 11:09:12');
	`)
}