package activity_stream

import "time"

//import "time"

type ActivityStream struct {
	ID               int           `json:"ID,string"`
	BlogID           int           `json:"BLOG_ID,string"`
	AuthorID         int           `json:"AUTHOR_ID,string"`
	PublishStatus    string        `json:"PUBLISH_STATUS"`
	Title            string        `json:"TITLE"`
	EnableComments   string        `json:"ENABLE_COMMENTS"`
	NumComments      string        `json:"NUM_COMMENTS"`
	Code             interface{}   `json:"CODE"`
	Micro            string        `json:"MICRO"`
	DetailText       string        `json:"DETAIL_TEXT"`
	DatePublish      time.Time     `json:"DATE_PUBLISH"`
	CategoryID       string        `json:"CATEGORY_ID"`
	HasSocnetAll     string        `json:"HAS_SOCNET_ALL"`
	HasTags          string        `json:"HAS_TAGS"`
	HasImages        string        `json:"HAS_IMAGES"`
	HasProps         string        `json:"HAS_PROPS"`
	HasCommentImages string        `json:"HAS_COMMENT_IMAGES"`
	Files            []interface{} `json:"FILES,omitempty"`

	UfGratitude       Uf_Gratitude       `json:"UF_GRATITUDE,omitempty"`
	UfImprtantDateEnd Uf_ImprtantDateEnd `json:"UF_IMPRTANT_DATE_END,omitempty"`
	UfMailMessage     Uf_MailMessage     `json:"UF_MAIL_MESSAGE,omitempty"`

	UfBlogPostDoc     Uf_BlogPost_Doc     `json:"UF_BLOG_POST_DOC,omitempty"`
	UfBlogPostImprtnt Uf_BlogPost_Imprtnt `json:"UF_BLOG_POST_IMPRTNT,omitempty"`
	UfBlogPostVote    Uf_BlogPost_Vote    `json:"UF_BLOG_POST_VOTE,omitempty"`
	UfBlogPostFile    Uf_BlogPost_File    `json:"UF_BLOG_POST_FILE,omitempty"`
	UfBlogPostFEdit   Uf_BlogPost_FEdit   `json:"UF_BLOG_POST_F_EDIT,omitempty"`
	UfBlogPostURLPrv  Uf_BlogPost_URLPrv  `json:"UF_BLOG_POST_URL_PRV,omitempty"`
}

type Uf_BlogPost_Doc struct {
	ID              int         `json:"ID,string"`
	EntityID        string      `json:"ENTITY_ID"`
	EntityValueID   int         `json:"ENTITY_VALUE_ID"`
	UserTypeID      string      `json:"USER_TYPE_ID"`
	XMLID           string      `json:"XML_ID"`
	FieldName       string      `json:"FIELD_NAME"`
	Sort            string      `json:"SORT"`
	Multiple        string      `json:"MULTIPLE"`
	Mandatory       string      `json:"MANDATORY"`
	ShowFilter      string      `json:"SHOW_FILTER"`
	ShowInList      string      `json:"SHOW_IN_LIST"`
	EditInList      string      `json:"EDIT_IN_LIST"`
	IsSearchable    string      `json:"IS_SEARCHABLE"`
	EditFormLabel   interface{} `json:"EDIT_FORM_LABEL"`
	ListColumnLabel interface{} `json:"LIST_COLUMN_LABEL"`
	ListFilterLabel interface{} `json:"LIST_FILTER_LABEL"`
	ErrorMessage    interface{} `json:"ERROR_MESSAGE"`
	HelpMessage     interface{} `json:"HELP_MESSAGE"`
	UserType        UserType    `json:"USER_TYPE"`

	Value bool `json:"VALUE"`

	Settings struct {
		Size           int           `json:"SIZE"`
		ListWidth      int           `json:"LIST_WIDTH"`
		ListHeight     int           `json:"LIST_HEIGHT"`
		MaxShowSize    int           `json:"MAX_SHOW_SIZE"`
		MaxAllowedSize int           `json:"MAX_ALLOWED_SIZE"`
		Extensions     []interface{} `json:"EXTENSIONS"`
		TargetBlank    string        `json:"TARGET_BLANK"`
	} `json:"SETTINGS"`
}

type Uf_Gratitude struct {
	ID              int         `json:"ID,string"`
	EntityID        string      `json:"ENTITY_ID"`
	EntityValueID   int         `json:"ENTITY_VALUE_ID"`
	UserTypeID      string      `json:"USER_TYPE_ID"`
	XMLID           string      `json:"XML_ID"`
	FieldName       string      `json:"FIELD_NAME"`
	Sort            string      `json:"SORT"`
	Multiple        string      `json:"MULTIPLE"`
	Mandatory       string      `json:"MANDATORY"`
	ShowFilter      string      `json:"SHOW_FILTER"`
	ShowInList      string      `json:"SHOW_IN_LIST"`
	EditInList      string      `json:"EDIT_IN_LIST"`
	IsSearchable    string      `json:"IS_SEARCHABLE"`
	EditFormLabel   interface{} `json:"EDIT_FORM_LABEL"`
	ListColumnLabel interface{} `json:"LIST_COLUMN_LABEL"`
	ListFilterLabel interface{} `json:"LIST_FILTER_LABEL"`
	ErrorMessage    interface{} `json:"ERROR_MESSAGE"`
	HelpMessage     interface{} `json:"HELP_MESSAGE"`
	UserType        UserType    `json:"USER_TYPE"`

	Value interface{} `json:"VALUE"`

	Settings Settings `json:"SETTINGS"`
}

type Uf_BlogPost_Imprtnt struct {
	ID              int         `json:"ID,string"`
	EntityID        string      `json:"ENTITY_ID"`
	EntityValueID   int         `json:"ENTITY_VALUE_ID"`
	UserTypeID      string      `json:"USER_TYPE_ID"`
	XMLID           string      `json:"XML_ID"`
	FieldName       string      `json:"FIELD_NAME"`
	Sort            string      `json:"SORT"`
	Multiple        string      `json:"MULTIPLE"`
	Mandatory       string      `json:"MANDATORY"`
	ShowFilter      string      `json:"SHOW_FILTER"`
	ShowInList      string      `json:"SHOW_IN_LIST"`
	EditInList      string      `json:"EDIT_IN_LIST"`
	IsSearchable    string      `json:"IS_SEARCHABLE"`
	EditFormLabel   interface{} `json:"EDIT_FORM_LABEL"`
	ListColumnLabel interface{} `json:"LIST_COLUMN_LABEL"`
	ListFilterLabel interface{} `json:"LIST_FILTER_LABEL"`
	ErrorMessage    interface{} `json:"ERROR_MESSAGE"`
	HelpMessage     interface{} `json:"HELP_MESSAGE"`
	UserType        UserType    `json:"USER_TYPE"`

	Value string `json:"VALUE"`

	Settings Settings `json:"SETTINGS"`
}

type Uf_BlogPost_Vote struct {
	ID              int         `json:"ID,string"`
	EntityID        string      `json:"ENTITY_ID"`
	EntityValueID   int         `json:"ENTITY_VALUE_ID"`
	UserTypeID      string      `json:"USER_TYPE_ID"`
	XMLID           string      `json:"XML_ID"`
	FieldName       string      `json:"FIELD_NAME"`
	Sort            string      `json:"SORT"`
	Multiple        string      `json:"MULTIPLE"`
	Mandatory       string      `json:"MANDATORY"`
	ShowFilter      string      `json:"SHOW_FILTER"`
	ShowInList      string      `json:"SHOW_IN_LIST"`
	EditInList      string      `json:"EDIT_IN_LIST"`
	IsSearchable    string      `json:"IS_SEARCHABLE"`
	EditFormLabel   interface{} `json:"EDIT_FORM_LABEL"`
	ListColumnLabel interface{} `json:"LIST_COLUMN_LABEL"`
	ListFilterLabel interface{} `json:"LIST_FILTER_LABEL"`
	ErrorMessage    interface{} `json:"ERROR_MESSAGE"`
	HelpMessage     interface{} `json:"HELP_MESSAGE"`
	UserType        UserType    `json:"USER_TYPE"`

	Value interface{} `json:"VALUE"`

	Settings struct {
		ChannelID     int `json:"CHANNEL_ID"`
		Unique        int `json:"UNIQUE"`
		UniqueIPDelay struct {
			Delay     string `json:"DELAY"`
			DelayType string `json:"DELAY_TYPE"`
		} `json:"UNIQUE_IP_DELAY"`
		Notify    string `json:"NOTIFY"`
		Converted string `json:"converted"`
	} `json:"SETTINGS"`
}

type Uf_BlogPost_File struct {
	ID              int         `json:"ID,string"`
	EntityID        string      `json:"ENTITY_ID"`
	EntityValueID   int         `json:"ENTITY_VALUE_ID"`
	UserTypeID      string      `json:"USER_TYPE_ID"`
	XMLID           string      `json:"XML_ID"`
	FieldName       string      `json:"FIELD_NAME"`
	Sort            string      `json:"SORT"`
	Multiple        string      `json:"MULTIPLE"`
	Mandatory       string      `json:"MANDATORY"`
	ShowFilter      string      `json:"SHOW_FILTER"`
	ShowInList      string      `json:"SHOW_IN_LIST"`
	EditInList      string      `json:"EDIT_IN_LIST"`
	IsSearchable    string      `json:"IS_SEARCHABLE"`
	EditFormLabel   interface{} `json:"EDIT_FORM_LABEL"`
	ListColumnLabel interface{} `json:"LIST_COLUMN_LABEL"`
	ListFilterLabel interface{} `json:"LIST_FILTER_LABEL"`
	ErrorMessage    interface{} `json:"ERROR_MESSAGE"`
	HelpMessage     interface{} `json:"HELP_MESSAGE"`
	UserType        UserType    `json:"USER_TYPE"`

	Value []interface{} `json:"VALUE"`

	Settings struct {
		IblockID          int    `json:"IBLOCK_ID"`
		SectionID         int    `json:"SECTION_ID"`
		UfToSaveAllowEdit string `json:"UF_TO_SAVE_ALLOW_EDIT"`
	} `json:"SETTINGS"`
}

type Uf_BlogPost_FEdit struct {
	ID              int         `json:"ID,string"`
	EntityID        string      `json:"ENTITY_ID"`
	EntityValueID   int         `json:"ENTITY_VALUE_ID"`
	UserTypeID      string      `json:"USER_TYPE_ID"`
	XMLID           string      `json:"XML_ID"`
	FieldName       string      `json:"FIELD_NAME"`
	Sort            string      `json:"SORT"`
	Multiple        string      `json:"MULTIPLE"`
	Mandatory       string      `json:"MANDATORY"`
	ShowFilter      string      `json:"SHOW_FILTER"`
	ShowInList      string      `json:"SHOW_IN_LIST"`
	EditInList      string      `json:"EDIT_IN_LIST"`
	IsSearchable    string      `json:"IS_SEARCHABLE"`
	EditFormLabel   interface{} `json:"EDIT_FORM_LABEL"`
	ListColumnLabel interface{} `json:"LIST_COLUMN_LABEL"`
	ListFilterLabel interface{} `json:"LIST_FILTER_LABEL"`
	ErrorMessage    interface{} `json:"ERROR_MESSAGE"`
	HelpMessage     interface{} `json:"HELP_MESSAGE"`
	UserType        UserType    `json:"USER_TYPE"`

	Value interface{} `json:"VALUE"`

	Settings struct {
		DefaultValue  int           `json:"DEFAULT_VALUE"`
		Display       string        `json:"DISPLAY"`
		Label         []interface{} `json:"LABEL"`
		LabelCheckbox interface{}   `json:"LABEL_CHECKBOX"`
	} `json:"SETTINGS"`
}

type Uf_BlogPost_URLPrv struct {
	ID              int         `json:"ID,string"`
	EntityID        string      `json:"ENTITY_ID"`
	EntityValueID   int         `json:"ENTITY_VALUE_ID"`
	UserTypeID      string      `json:"USER_TYPE_ID"`
	XMLID           string      `json:"XML_ID"`
	FieldName       string      `json:"FIELD_NAME"`
	Sort            string      `json:"SORT"`
	Multiple        string      `json:"MULTIPLE"`
	Mandatory       string      `json:"MANDATORY"`
	ShowFilter      string      `json:"SHOW_FILTER"`
	ShowInList      string      `json:"SHOW_IN_LIST"`
	EditInList      string      `json:"EDIT_IN_LIST"`
	IsSearchable    string      `json:"IS_SEARCHABLE"`
	EditFormLabel   interface{} `json:"EDIT_FORM_LABEL"`
	ListColumnLabel interface{} `json:"LIST_COLUMN_LABEL"`
	ListFilterLabel interface{} `json:"LIST_FILTER_LABEL"`
	ErrorMessage    interface{} `json:"ERROR_MESSAGE"`
	HelpMessage     interface{} `json:"HELP_MESSAGE"`
	UserType        UserType    `json:"USER_TYPE"`

	Value interface{} `json:"VALUE"`

	Settings []interface{} `json:"SETTINGS"`
}

type Uf_ImprtantDateEnd struct {
	ID              int         `json:"ID,string"`
	EntityID        string      `json:"ENTITY_ID"`
	EntityValueID   int         `json:"ENTITY_VALUE_ID"`
	UserTypeID      string      `json:"USER_TYPE_ID"`
	XMLID           string      `json:"XML_ID"`
	FieldName       string      `json:"FIELD_NAME"`
	Sort            string      `json:"SORT"`
	Multiple        string      `json:"MULTIPLE"`
	Mandatory       string      `json:"MANDATORY"`
	ShowFilter      string      `json:"SHOW_FILTER"`
	ShowInList      string      `json:"SHOW_IN_LIST"`
	EditInList      string      `json:"EDIT_IN_LIST"`
	IsSearchable    string      `json:"IS_SEARCHABLE"`
	EditFormLabel   string      `json:"EDIT_FORM_LABEL"`
	ListColumnLabel string      `json:"LIST_COLUMN_LABEL"`
	ListFilterLabel interface{} `json:"LIST_FILTER_LABEL"`
	ErrorMessage    interface{} `json:"ERROR_MESSAGE"`
	HelpMessage     interface{} `json:"HELP_MESSAGE"`
	UserType        UserType    `json:"USER_TYPE"`

	Value string `json:"VALUE"`

	Settings struct {
		DefaultValue struct {
			Type  string `json:"TYPE"`
			Value string `json:"VALUE"`
		} `json:"DEFAULT_VALUE"`
		UseSecond   string `json:"USE_SECOND"`
		UseTimezone string `json:"USE_TIMEZONE"`
	} `json:"SETTINGS"`
}

type Uf_MailMessage struct {
	ID              int         `json:"ID,string"`
	EntityID        string      `json:"ENTITY_ID"`
	EntityValueID   int         `json:"ENTITY_VALUE_ID"`
	UserTypeID      string      `json:"USER_TYPE_ID"`
	XMLID           string      `json:"XML_ID"`
	FieldName       string      `json:"FIELD_NAME"`
	Sort            string      `json:"SORT"`
	Multiple        string      `json:"MULTIPLE"`
	Mandatory       string      `json:"MANDATORY"`
	ShowFilter      string      `json:"SHOW_FILTER"`
	ShowInList      string      `json:"SHOW_IN_LIST"`
	EditInList      string      `json:"EDIT_IN_LIST"`
	IsSearchable    string      `json:"IS_SEARCHABLE"`
	EditFormLabel   interface{} `json:"EDIT_FORM_LABEL"`
	ListColumnLabel interface{} `json:"LIST_COLUMN_LABEL"`
	ListFilterLabel interface{} `json:"LIST_FILTER_LABEL"`
	ErrorMessage    interface{} `json:"ERROR_MESSAGE"`
	HelpMessage     interface{} `json:"HELP_MESSAGE"`
	UserType        UserType    `json:"USER_TYPE"`

	Value interface{} `json:"VALUE"`

	Settings interface{} `json:"SETTINGS"`
}

type UserType struct {
	UserTypeID  string `json:"USER_TYPE_ID"`
	ClassName   string `json:"CLASS_NAME"`
	Description string `json:"DESCRIPTION"`
	BaseType    string `json:"BASE_TYPE"`

	EditCallback []string `json:"EDIT_CALLBACK,omitempty"`
	ViewCallback []string `json:"VIEW_CALLBACK,omitempty"`

	OnBeforeSave []string `json:"onBeforeSave,omitempty"`
	OnDelete     []string `json:"onDelete,omitempty"`

	Tag []string `json:"TAG,omitempty"`

	UseFieldComponent bool `json:"USE_FIELD_COMPONENT,omitempty"`
}

type Settings struct {
	Size         int `json:"SIZE"`
	MinValue     int `json:"MIN_VALUE"`
	MaxValue     int `json:"MAX_VALUE"`
	DefaultValue int `json:"DEFAULT_VALUE"`
}
