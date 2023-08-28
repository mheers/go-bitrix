package crm

import "time"

type Deal struct {
	Id                  int         `json:"ID,string"`
	Title               string      `json:"TITLE"`
	TypeId              string      `json:"TYPE_ID"`
	CategoryId          int         `json:"CATEGORY_ID,string"`
	StageId             string      `json:"STAGE_ID"`
	StageSemanticId     string      `json:"STAGE_SEMANTIC_ID"`
	IsNew               string      `json:"IS_NEW"`
	IsRecurring         string      `json:"IS_RECURRING"`
	IsReturnCustomer    string      `json:"IS_RETURN_CUSTOMER"`
	IsRepeatedApproach  string      `json:"IS_REPEATED_APPROACH"`
	Probability         int         `json:"PROBABILITY"`
	CurrencyId          string      `json:"CURRENCY_ID"`
	Opportunity         float64     `json:"OPPORTUNITY,string"`
	IsManualOpportunity string      `json:"IS_MANUAL_OPPORTUNITY"`
	TaxValue            float64     `json:"TAX_VALUE,string"`
	CompanyID           int         `json:"COMPANY_ID,string"`
	ContactId           int         `json:"CONTACT_ID,string"`
	QuoteId             string      `json:"QUOTE_ID"`
	Begindate           time.Time   `json:"BEGINDATE"`
	Closedate           time.Time   `json:"CLOSEDATE"`
	Opened              string      `json:"OPENED"`
	Closed              string      `json:"CLOSED"`
	Comments            string      `json:"COMMENTS"`
	AssignedByID        int         `json:"ASSIGNED_BY_ID,string"`
	CreatedByID         int         `json:"CREATED_BY_ID,string"`
	ModifyByID          int         `json:"MODIFY_BY_ID,string"`
	DateCreate          time.Time   `json:"DATE_CREATE"`
	DateModify          time.Time   `json:"DATE_MODIFY"`
	SourceId            string      `json:"SOURCE_ID"`
	SourceDescription   string      `json:"SOURCE_DESCRIPTION"`
	LeadId              int         `json:"LEAD_ID,string"`
	AdditionalInfo      string      `json:"ADDITIONAL_INFO"`
	LocationId          string      `json:"LOCATION_ID"`
	OriginatorId        string      `json:"ORIGINATOR_ID"`
	OriginId            string      `json:"ORIGIN_ID"`
	UtmSource           string      `json:"UTM_SOURCE"`
	UtmMedium           string      `json:"UTM_MEDIUM"`
	UtmCampaign         string      `json:"UTM_CAMPAIGN"`
	UtmContent          string      `json:"UTM_CONTENT"`
	UtmTerm             string      `json:"UTM_TERM"`
	MovedByID           int         `json:"MOVED_BY_ID,string"`
	MovedTime           time.Time   `json:"MOVED_TIME"`
	ParentID133         interface{} `json:"PARENT_ID_133"`
	ExchRate            float64     `json:"EXCH_RATE"`
	AccountCurrencyId   string      `json:"ACCOUNT_CURRENCY_ID"`
	OpportunityAccount  float64     `json:"OPPORTUNITY_ACCOUNT"`
	TaxValueAccount     float64     `json:"TAX_VALUE_ACCOUNT"`
	ContactIds          string      `json:"CONTACT_IDS"`

	Userfield map[string]interface{}
}
