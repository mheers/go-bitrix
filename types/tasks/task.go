package tasks

import "time"

type AllTasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID                   int           `json:"ID,string"`
	ParentID             int           `json:"parentId,string"`
	Title                string        `json:"title"`
	Description          string        `json:"description"`
	Mark                 string        `json:"mark"`
	Priority             string        `json:"priority"`
	Status               string        `json:"status"`
	Multitask            string        `json:"multitask"`
	NotViewed            string        `json:"notViewed"`
	Replicate            string        `json:"replicate"`
	GroupID              string        `json:"groupId"`
	StageID              string        `json:"stageId"`
	CreatedBy            string        `json:"createdBy"`
	CreatedDate          time.Time     `json:"createdDate"`
	ResponsibleID        string        `json:"responsibleId"`
	ChangedBy            string        `json:"changedBy"`
	ChangedDate          time.Time     `json:"changedDate"`
	StatusChangedBy      string        `json:"statusChangedBy"`
	StatusChangedDate    time.Time     `json:"statusChangedDate"`
	ClosedBy             string        `json:"closedBy"`
	ClosedDate           time.Time     `json:"closedDate"`
	ActivityDate         time.Time     `json:"activityDate"`
	DateStart            time.Time     `json:"dateStart"`
	Deadline             time.Time     `json:"deadline"`
	StartDatePlan        time.Time     `json:"startDatePlan"`
	EndDatePlan          time.Time     `json:"endDatePlan"`
	GUID                 string        `json:"guid"`
	XMLID                int           `json:"xmlId,string"`
	CommentsCount        string        `json:"commentsCount"`
	ServiceCommentsCount string        `json:"serviceCommentsCount"`
	AllowChangeDeadline  string        `json:"allowChangeDeadline"`
	AllowTimeTracking    string        `json:"allowTimeTracking"`
	TaskControl          string        `json:"taskControl"`
	AddInReport          string        `json:"addInReport"`
	ForkedByTemplateID   int           `json:"forkedByTemplateId,string"`
	TimeEstimate         string        `json:"timeEstimate"`
	TimeSpentInLogs      string        `json:"timeSpentInLogs"`
	MatchWorkTime        string        `json:"matchWorkTime"`
	ForumTopicID         string        `json:"forumTopicId"`
	ForumID              string        `json:"forumId"`
	SiteID               string        `json:"siteId"`
	Subordinate          string        `json:"subordinate"`
	Favorite             string        `json:"favorite"`
	ExchangeModified     interface{}   `json:"exchangeModified"`
	ExchangeID           int           `json:"exchangeId,string"`
	OutlookVersion       string        `json:"outlookVersion"`
	ViewedDate           time.Time     `json:"viewedDate"`
	Sorting              interface{}   `json:"sorting"`
	DurationPlan         interface{}   `json:"durationPlan"`
	DurationFact         string        `json:"durationFact"`
	DurationType         string        `json:"durationType"`
	IsMuted              string        `json:"isMuted"`
	IsPinned             string        `json:"isPinned"`
	IsPinnedInGroup      string        `json:"isPinnedInGroup"`
	DescriptionInBbcode  string        `json:"descriptionInBbcode"`
	Auditors             []interface{} `json:"auditors"`
	Accomplices          []interface{} `json:"accomplices"`
	NewCommentsCount     int           `json:"newCommentsCount"`
	Group                interface{}   `json:"group"` //in all tasks this field is like `GroupData struct`, but in one and simgle it is empty [], so -- interface{}
	Creator              UserData      `json:"creator"`
	Responsible          UserData      `json:"responsible"`
	SubStatus            int           `json:"subStatus,string"`
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
	Link string `json:"link"`
	Icon string `json:"icon"`
}
