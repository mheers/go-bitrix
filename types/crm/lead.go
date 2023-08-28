package crm

import "time"

type Lead struct {
	Id                 int             `json:"ID,string"`
	Title              string          `json:"TITLE"`
	Honorific          string          `json:"HONORIFIC"`
	Name               string          `json:"NAME"`
	SecondName         string          `json:"SECOND_NAME"`
	LastName           string          `json:"LAST_NAME"`
	Birthdate          string          `json:"BIRTHDATE"`
	CompanyTitle       string          `json:"COMPANY_TITLE"`
	SourceId           string          `json:"SOURCE_ID"`
	SourceDescription  string          `json:"SOURCE_DESCRIPTION"`
	StatusId           string          `json:"STATUS_ID"`
	StatusDescription  string          `json:"STATUS_DESCRIPTION"`
	StatusSemanticId   string          `json:"STATUS_SEMANTIC_ID"`
	Post               string          `json:"POST"`
	Address            string          `json:"ADDRESS"`
	Address2           string          `json:"ADDRESS_2"`
	AddressCity        string          `json:"ADDRESS_CITY"`
	AddressPostalCode  string          `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion      string          `json:"ADDRESS_REGION"`
	AddressProvince    string          `json:"ADDRESS_PROVINCE"`
	AddressCountry     string          `json:"ADDRESS_COUNTRY"`
	AddressCountryCode string          `json:"ADDRESS_COUNTRY_CODE"`
	CurrencyId         string          `json:"CURRENCY_ID"`
	Opportunity        float64         `json:"OPPORTUNITY,string"`
	Opened             string          `json:"OPENED"`
	Comments           string          `json:"COMMENTS"`
	HasPhone           string          `json:"HAS_PHONE"`
	HasEmail           string          `json:"HAS_EMAIL"`
	HasImol            string          `json:"HAS_IMOL"`
	AssignedById       int             `json:"ASSIGNED_BY_ID,string"`
	CreatedById        int             `json:"CREATED_BY_ID,string"`
	ModifyById         int             `json:"MODIFY_BY_ID,string"`
	DateCreate         time.Time       `json:"DATE_CREATE"`
	DateModify         time.Time       `json:"DATE_MODIFY"`
	CompanyId          int             `json:"COMPANY_ID,string"`
	ContactId          int             `json:"CONTACT_ID,string"`
	IsReturnCustomer   string          `json:"IS_RETURN_CUSTOMER"`
	DateClosed         string          `json:"DATE_CLOSED"`
	OriginatorId       string          `json:"ORIGINATOR_ID"`
	OriginId           string          `json:"ORIGIN_ID"`
	UtmSource          string          `json:"UTM_SOURCE"`
	UtmMedium          string          `json:"UTM_MEDIUM"`
	UtmCampaign        string          `json:"UTM_CAMPAIGN"`
	UtmContent         string          `json:"UTM_CONTENT"`
	UtmTerm            string          `json:"UTM_TERM"`
	Phone              []CrmMultifield `json:"PHONE"`
	Email              []CrmMultifield `json:"EMAIL"`

	Userfield map[string]interface{}
}
