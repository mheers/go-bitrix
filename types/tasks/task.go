package tasks

import "time"

type AllTasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID                   int           `json:"ID"`
	ParentID             int           `json:"PARENT_ID"`
	Title                string        `json:"TITLE"`
	Description          string        `json:"DESCRIPTION"`
	Mark                 string        `json:"MARK"`
	Priority             string        `json:"PRIORITY"`
	Status               string        `json:"STATUS,omitempty"`
	Multitask            string        `json:"MULTITASK"`
	NotViewed            string        `json:"NOT_VIEWED"`
	Replicate            string        `json:"REPLICATE"`
	GroupID              string        `json:"GROUP_ID"`
	StageID              string        `json:"STAGE_ID"`
	CreatedBy            string        `json:"CREATED_BY"`
	CreatedDate          time.Time     `json:"CREATED_DATE"`
	ResponsibleID        string        `json:"RESPONSIBLE_ID,omitempty"`
	ChangedBy            string        `json:"CHANGED_BY"`
	ChangedDate          time.Time     `json:"CHANGED_DATE"`
	StatusChangedBy      string        `json:"STATUS_CHANGED_BY"`
	StatusChangedDate    time.Time     `json:"STATUS_CHANGED_DATE"`
	ClosedBy             string        `json:"CLOSED_BY"`
	ClosedDate           time.Time     `json:"CLOSED_DATE"`
	ActivityDate         time.Time     `json:"ACTIVITY_DATE"`
	DateStart            time.Time     `json:"DATE_START"`
	Deadline             time.Time     `json:"DEADLINE"`
	StartDatePlan        time.Time     `json:"START_DATE_PLAN"`
	EndDatePlan          time.Time     `json:"END_DATE_PLAN"`
	GUID                 string        `json:"GUID"`
	XMLID                int           `json:"XML_ID"`
	CommentsCount        string        `json:"COMMENTS_COUNT"`
	ServiceCommentsCount string        `json:"SERVICE_COMMENTS_COUNT"`
	AllowChangeDeadline  string        `json:"ALLOW_CHANGE_DEADLINE"`
	AllowTimeTracking    string        `json:"ALLOW_TIME_TRACKING"`
	TaskControl          string        `json:"TASK_CONTROL"`
	AddInReport          string        `json:"ADD_IN_REPORT"`
	ForkedByTemplateID   int           `json:"FORKED_BY_TEMPLATE_ID"`
	TimeEstimate         string        `json:"TIME_ESTIMATE"`
	TimeSpentInLogs      string        `json:"TIME_SPENT_IN_LOGS"`
	MatchWorkTime        string        `json:"MATCH_WORK_TIME"`
	ForumTopicID         string        `json:"FORUM_TOPIC_ID"`
	ForumID              string        `json:"FORUM_ID"`
	SiteID               string        `json:"SITE_ID"`
	Subordinate          string        `json:"SUBORDINATE"`
	Favorite             string        `json:"FAVORITE"`
	ExchangeModified     interface{}   `json:"EXCHANGE_MODIFIED"`
	ExchangeID           int           `json:"EXCHANGE_ID"`
	OutlookVersion       string        `json:"OUTLOOK_VERSION"`
	ViewedDate           time.Time     `json:"VIEWED_DATE"`
	Sorting              interface{}   `json:"SORTING"`
	DurationPlan         interface{}   `json:"DURATION_PLAN"`
	DurationFact         string        `json:"DURATION_FACT"`
	DurationType         string        `json:"DURATION_TYPE"`
	IsMuted              string        `json:"IS_MUTED"`
	IsPinned             string        `json:"IS_PINNED"`
	IsPinnedInGroup      string        `json:"IS_PINNED_IN_GROUP"`
	DescriptionInBbcode  string        `json:"DESCRIPTION_IN_BBCODE"`
	Auditors             []interface{} `json:"AUDITORS"`
	Accomplices          []interface{} `json:"ACCOMPLICES"`
	NewCommentsCount     int           `json:"NEW_COMMENTS_COUNT"`
	Group                interface{}   `json:"GROUP"` //in all tasks this field is like `GroupData struct`, but in one and simgle it is empty [], so -- interface{}
	// Creator              UserData      `json:"CREATOR"`
	// Responsible          UserData      `json:"RESPONSIBLE"`
	SubStatus int `json:"SUB_STATUS"`
}

type GroupData struct {
	ID             int           `json:"id,string"`
	Name           string        `json:"name"`
	Opened         bool          `json:"opened"`
	MembersCount   int           `json:"membersCount"`
	Image          string        `json:"image"`
	AdditionalData []interface{} `json:"additionalData"`
}

type UserData struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
	Link string `json:"link,omitempty"`
	Icon string `json:"icon,omitempty"`
}
