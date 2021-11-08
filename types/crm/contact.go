package crm

import (
	"time"
)

type CrmMultifield struct {
	Value     string `json:"VALUE"`
	ValueType string `json:"VALUE_TYPE"`
	TypeId    string `json:"TYPE_ID"`
}

type Contact struct {
	Id                 int             `json:"ID"`
	Honorific          string          `json:"HONORIFIC"`
	Name               string          `json:"NAME"`
	SecondName         string          `json:"SECOND_NAME"`
	LastName           string          `json:"LAST_NAME"`
	FullName           string          `json:"FULL_NAME"`
	Photo              int             `json:"PHOTO"`
	Birthdate          time.Time       `json:"BIRTHDATE"`
	BirthdaySort       int             `json:"BIRTHDAY_SORT"`
	TypeId             string          `json:"TYPE_ID"`
	SourceId           string          `json:"SOURCE_ID"`
	SourceDescription  string          `json:"SOURCE_DESCRIPTION"`
	Post               string          `json:"POST"`
	Address            string          `json:"ADDRESS"`
	Address2           string          `json:"ADDRESS_2"`
	AddressCity        string          `json:"ADDRESS_CITY"`
	AddressPostalCode  string          `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion      string          `json:"ADDRESS_REGION"`
	AddressProvince    string          `json:"ADDRESS_PROVINCE"`
	AddressCountry     string          `json:"ADDRESS_COUNTRY"`
	AddressCountryCode string          `json:"ADDRESS_COUNTRY_CODE"`
	Comments           string          `json:"COMMENTS"`
	Opened             string          `json:"OPENED"`
	Export             string          `json:"EXPORT"`
	HasPhone           bool            `json:"HAS_PHONE"`
	HasEmail           bool            `json:"HAS_EMAIL"`
	HasImol            bool            `json:"HAS_IMOL"`
	AssignedById       string          `json:"ASSIGNED_BY_ID"`
	CreatedById        string          `json:"CREATED_BY_ID"`
	ModifyById         string          `json:"MODIFY_BY_ID"`
	DateCreate         time.Time       `json:"DATE_CREATE"`
	DateModify         time.Time       `json:"DATE_MODIFY"`
	CompanyId          string          `json:"COMPANY_ID"`
	CompanyIds         string          `json:"COMPANY_IDS"`
	LeadId             string          `json:"LEAD_ID"`
	OriginatorId       string          `json:"ORIGINATOR_ID"`
	OriginId           string          `json:"ORIGIN_ID"`
	OriginVersion      string          `json:"ORIGIN_VERSION"`
	FaceId             int             `json:"FACE_ID"`
	UtmSource          string          `json:"UTM_SOURCE"`
	UtmMedium          string          `json:"UTM_MEDIUM"`
	UtmCampaign        string          `json:"UTM_CAMPAIGN"`
	UtmContent         string          `json:"UTM_CONTENT"`
	UtmTerm            string          `json:"UTM_TERM"`
	Phone              []CrmMultifield `json:"PHONE"`
	Email              []CrmMultifield `json:"EMAIL"`
	Web                string          `json:"WEB"`
	Im                 string          `json:"IM"`
}
