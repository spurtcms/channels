package postgres

import (
	"time"

	"gorm.io/gorm"
)

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
}

type TblChannelCategory struct {
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

type TblChannelEntry struct {
	Id              int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title           string    `gorm:"type:character varying"`
	Slug            string    `gorm:"type:character varying"`
	Description     string    `gorm:"type:text"`
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
	TenantId        int       `gorm:"type:integer"`
}

type TblChannelEntryField struct {
	Id             int       `gorm:"primaryKey;auto_increment;type:serial"`
	FieldName      string    `gorm:"type:character varying"`
	FieldValue     string    `gorm:"type:character varying"`
	FieldTypeId    int       `gorm:"-:migration;<-false"`
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
	TenantId   int       `gorm:"type:integer"`
}

type TblFieldOption struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	OptionName  string    `gorm:"type:character varying"`
	OptionValue string    `gorm:"type:character varying"`
	FieldId     int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted   int       `gorm:"DEFAULT:0"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
	TenantId    int       `gorm:"type:integer"`
}

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
	TenantId   int       `gorm:"type:integer"`
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
