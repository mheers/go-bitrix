package crm

import (
	"time"
)

type Activity struct {
	Id                 int       `json:"ID,string"`
	OwnerId            int       `json:"OWNER_ID,string"`
	OwnerTypeId        string    `json:"OWNER_TYPE_ID"`
	TypeId             string    `json:"TYPE_ID"`
	ProviderId         string    `json:"PROVIDER_ID"`
	ProviderTypeId     string    `json:"PROVIDER_TYPE_ID"`
	ProviderGroupId    string    `json:"PROVIDER_GROUP_ID"`
	AssociatedEntityId int       `json:"ASSOCIATED_ENTITY_ID,string"`
	Subject            string    `json:"SUBJECT"`
	StartTime          string    `json:"START_TIME"`
	EndTime            string    `json:"END_TIME"`
	Deadline           time.Time `json:"DEADLINE"`
	Completed          string    `json:"COMPLETED"`
	Status             string    `json:"STATUS"`
	ResponsibleId      string    `json:"RESPONSIBLE_ID"`
	Priority           string    `json:"PRIORITY"`
	NotifyType         string    `json:"NOTIFY_TYPE"`
	NotifyValue        int       `json:"NOTIFY_VALUE,string"`
	Description        string    `json:"DESCRIPTION"`
	DescriptionType    string    `json:"DESCRIPTION_TYPE"`
	Direction          string    `json:"DIRECTION"`
	Location           string    `json:"LOCATION"`
	Created            time.Time `json:"CREATED"`
	AuthorId           string    `json:"AUTHOR_ID"`
	LastUpdated        time.Time `json:"LAST_UPDATED"`
	EditorId           string    `json:"EDITOR_ID"`
	OriginId           string    `json:"ORIGIN_ID"`
	OriginatorId       string    `json:"ORIGINATOR_ID"`
	ResultStatus       int       `json:"RESULT_STATUS,string"`
	ResultStream       int       `json:"RESULT_STREAM,string"`
	ResultSourceId     string    `json:"RESULT_SOURCE_ID"`
	ProviderData       string    `json:"PROVIDER_DATA"`
	ResultMark         int       `json:"RESULT_MARK,string"`
	ResultValue        float64   `json:"RESULT_VALUE,omitempty"`
	ResultSum          float64   `json:"RESULT_SUM,omitempty"`
	ResultCurrencyId   string    `json:"RESULT_CURRENCY_ID"`
	AutocompleteRule   int       `json:"AUTOCOMPLETE_RULE,string"`
}
