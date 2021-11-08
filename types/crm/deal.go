package crm

import (
	"time"
)

type Deal struct {
	Id                  float64   `json:"ID"`
	Title               string    `json:"TITLE"`
	TypeId              string    `json:"TYPE_ID"`
	CategoryId          string    `json:"CATEGORY_ID"`
	StageId             string    `json:"STAGE_ID"`
	StageSemanticId     string    `json:"STAGE_SEMANTIC_ID"`
	IsNew               string    `json:"IS_NEW"`
	IsRecurring         string    `json:"IS_RECURRING"`
	IsReturnCustomer    string    `json:"IS_RETURN_CUSTOMER"`
	IsRepeatedApproach  string    `json:"IS_REPEATED_APPROACH"`
	Probability         int       `json:"PROBABILITY"`
	CurrencyId          string    `json:"CURRENCY_ID"`
	Opportunity         float64   `json:"OPPORTUNITY"`
	IsManualOpportunity string    `json:"IS_MANUAL_OPPORTUNITY"`
	TaxValue            float64   `json:"TAX_VALUE"`
	ExchRate            float64   `json:"EXCH_RATE"`
	AccountCurrencyId   string    `json:"ACCOUNT_CURRENCY_ID"`
	OpportunityAccount  float64   `json:"OPPORTUNITY_ACCOUNT"`
	TaxValueAccount     float64   `json:"TAX_VALUE_ACCOUNT"`
	CompanyId           string    `json:"COMPANY_ID"`
	ContactId           string    `json:"CONTACT_ID"`
	ContactIds          string    `json:"CONTACT_IDS"`
	QuoteId             string    `json:"QUOTE_ID"`
	Begindate           time.Time `json:"BEGINDATE"`
	Closedate           time.Time `json:"CLOSEDATE"`
	Opened              string    `json:"OPENED"`
	Closed              string    `json:"CLOSED"`
	Comments            string    `json:"COMMENTS"`
	AssignedById        string    `json:"ASSIGNED_BY_ID"`
	CreatedById         string    `json:"CREATED_BY_ID"`
	ModifyById          string    `json:"MODIFY_BY_ID"`
	DateCreate          time.Time `json:"DATE_CREATE"`
	DateModify          time.Time `json:"DATE_MODIFY"`
	SourceId            string    `json:"SOURCE_ID"`
	SourceDescription   string    `json:"SOURCE_DESCRIPTION"`
	LeadId              string    `json:"LEAD_ID"`
	AdditionalInfo      string    `json:"ADDITIONAL_INFO"`
	LocationId          string    `json:"LOCATION_ID"`
	OriginatorId        string    `json:"ORIGINATOR_ID"`
	OriginId            string    `json:"ORIGIN_ID"`
	UtmSource           string    `json:"UTM_SOURCE"`
	UtmMedium           string    `json:"UTM_MEDIUM"`
	UtmCampaign         string    `json:"UTM_CAMPAIGN"`
	UtmContent          string    `json:"UTM_CONTENT"`
	UtmTerm             string    `json:"UTM_TERM"`
}
