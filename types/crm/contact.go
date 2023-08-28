package crm

import "time"

type CrmMultifield struct {
	Value     string `json:"VALUE"`
	ValueType string `json:"VALUE_TYPE"`
	TypeId    string `json:"TYPE_ID"`
}

type Contact struct {
	Id                 int        `json:"ID,string"`
	Honorific          string     `json:"HONORIFIC"`
	Name               string     `json:"NAME"`
	SecondName         string     `json:"SECOND_NAME"`
	LastName           string     `json:"LAST_NAME"`
	FullName           string     `json:"FULL_NAME"`
	Photo              Photo      `json:"PHOTO,omitempty"`
	Birthdate          *time.Time `json:"BIRTHDATE,omitempty"`
	BirthdaySort       int        `json:"BIRTHDAY_SORT"`
	TypeId             string     `json:"TYPE_ID"`
	SourceId           string     `json:"SOURCE_ID"`
	SourceDescription  string     `json:"SOURCE_DESCRIPTION"`
	Post               string     `json:"POST"`
	Address            string     `json:"ADDRESS"`
	Address2           string     `json:"ADDRESS_2"`
	AddressCity        string     `json:"ADDRESS_CITY"`
	AddressPostalCode  string     `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion      string     `json:"ADDRESS_REGION"`
	AddressProvince    string     `json:"ADDRESS_PROVINCE"`
	AddressCountry     string     `json:"ADDRESS_COUNTRY"`
	AddressCountryCode string     `json:"ADDRESS_COUNTRY_CODE"`
	Comments           string     `json:"COMMENTS"`
	Opened             string     `json:"OPENED"`
	Export             string     `json:"EXPORT"`
	HasPhone           bool       `json:"HAS_PHONE"`
	HasEmail           bool       `json:"HAS_EMAIL"`
	HasImol            bool       `json:"HAS_IMOL"`

	AssignedById  int             `json:"ASSIGNED_BY_ID,string"`
	CreatedById   int             `json:"CREATED_BY_ID,string"`
	ModifyById    int             `json:"MODIFY_BY_ID,string"`
	DateCreate    *time.Time      `json:"DATE_CREATE,omitempty"`
	DateModify    *time.Time      `json:"DATE_MODIFY,omitempty"`
	CompanyId     string          `json:"COMPANY_ID"`
	CompanyIds    string          `json:"COMPANY_IDS"`
	LeadId        int             `json:"LEAD_ID,string"`
	OriginatorId  string          `json:"ORIGINATOR_ID"`
	OriginId      string          `json:"ORIGIN_ID"`
	OriginVersion string          `json:"ORIGIN_VERSION"`
	FaceId        int             `json:"FACE_ID,string"`
	UtmSource     string          `json:"UTM_SOURCE"`
	UtmMedium     string          `json:"UTM_MEDIUM"`
	UtmCampaign   string          `json:"UTM_CAMPAIGN"`
	UtmContent    string          `json:"UTM_CONTENT"`
	UtmTerm       string          `json:"UTM_TERM"`
	Phone         []CrmMultifield `json:"PHONE,omitempty"`
	Email         []CrmMultifield `json:"EMAIL,omitempty"`
	Web           []CrmMultifield `json:"WEB,omitempty"`
	Im            string          `json:"IM"`
	Value         string          `json:"VALUE"`
	ValueType     string          `json:"VALUE_TYPE"`

	Userfield map[string]interface{}
}

type Photo struct {
	ID          int    `json:"id"`
	ShowURL     string `json:"showUrl"`
	DownloadURL string `json:"downloadUrl"`
}
