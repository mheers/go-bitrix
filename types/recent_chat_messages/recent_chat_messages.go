package recent_chat_messages

import "time"

type MessageList struct {
	ChatID   int           `json:"chat_id"`
	Messages []MessageData `json:"messages"`
	Users    []UserData    `json:"users"`
	Files    []interface{} `json:"files"`
}

type MessageData struct {
	ID       int         `json:"id"`
	ChatID   int         `json:"chat_id"`
	AuthorID int         `json:"author_id"`
	Date     time.Time   `json:"date"`
	Text     string      `json:"text"`
	Unread   bool        `json:"unread"`
	UUID     interface{} `json:"uuid"`
	Params   interface{} `json:"params,omitempty"` // map[string]interface{}
}

type UserData struct {
	ID               int               `json:"id"`
	Active           bool              `json:"active"`
	Name             string            `json:"name"`
	FirstName        string            `json:"first_name"`
	LastName         string            `json:"last_name"`
	WorkPosition     string            `json:"work_position"`
	Color            string            `json:"color"`
	Avatar           string            `json:"avatar"`
	Gender           string            `json:"gender"`
	Birthday         string            `json:"birthday"`
	Extranet         bool              `json:"extranet"`
	Network          bool              `json:"network"`
	Bot              bool              `json:"bot"`
	Connector        bool              `json:"connector"`
	ExternalAuthID   string            `json:"external_auth_id"`
	Status           string            `json:"status"`
	Idle             bool              `json:"idle"`
	LastActivityDate time.Time         `json:"last_activity_date"`
	MobileLastDate   interface{}       `json:"mobile_last_date"` //time.Time
	Absent           bool              `json:"absent"`
	Departments      []int             `json:"departments"`
	Phones           map[string]string `json:"phones,omitempty"`
}
