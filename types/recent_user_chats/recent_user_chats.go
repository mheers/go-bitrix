package recent_user_chats

import "time"

type RecentUserChats struct {
	ID         interface{} `json:"id"`
	ChatID     int         `json:"chat_id"`
	Type       string      `json:"type"`
	Avatar     AvatarData  `json:"avatar"`
	Title      string      `json:"title"`
	Message    MessageData `json:"message"`
	Counter    int         `json:"counter"`
	Pinned     bool        `json:"pinned"`
	Unread     bool        `json:"unread"`
	DateUpdate time.Time   `json:"date_update"`
	Chat       ChatData    `json:"chat,omitempty"`
	User       UserData    `json:"user,omitempty"`
	Options    interface{} `json:"options"` // must be "map[string]bool" but sometimes responds "options":[], not "options":{"string":bool}}
}

type AvatarData struct {
	URL   string `json:"url"`
	Color string `json:"color"`
}

type MessageData struct {
	ID       int         `json:"id"`
	Text     string      `json:"text"`
	File     bool        `json:"file"`
	AuthorID int         `json:"author_id"`
	Attach   bool        `json:"attach"`
	Date     time.Time   `json:"date"`
	Status   string      `json:"status"`
	UUID     interface{} `json:"uuid"`
}

type ChatData struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Owner       int           `json:"owner"`
	Extranet    bool          `json:"extranet"`
	Avatar      string        `json:"avatar"`
	Color       string        `json:"color"`
	Type        string        `json:"type"`
	EntityType  string        `json:"entity_type"`
	EntityID    string        `json:"entity_id"`
	EntityData1 string        `json:"entity_data_1"`
	EntityData2 string        `json:"entity_data_2"`
	EntityData3 string        `json:"entity_data_3"`
	MuteList    []interface{} `json:"mute_list"`
	ManagerList []interface{} `json:"manager_list"`
	DateCreate  time.Time     `json:"date_create"`
	MessageType string        `json:"message_type"`
}

type UserData struct {
	ID               int         `json:"id"`
	Active           bool        `json:"active"`
	Name             string      `json:"name"`
	FirstName        string      `json:"first_name"`
	LastName         string      `json:"last_name"`
	WorkPosition     string      `json:"work_position"`
	Color            string      `json:"color"`
	Avatar           string      `json:"avatar"`
	Gender           string      `json:"gender"`
	Birthday         string      `json:"birthday"`
	Extranet         bool        `json:"extranet"`
	Network          bool        `json:"network"`
	Bot              bool        `json:"bot"`
	Connector        bool        `json:"connector"`
	ExternalAuthID   string      `json:"external_auth_id"`
	Status           string      `json:"status"`
	Idle             bool        `json:"idle"`
	LastActivityDate time.Time   `json:"last_activity_date"`
	Absent           bool        `json:"absent"`
	Departments      []int       `json:"departments"`
	MobileLastDate   interface{} `json:"mobile_last_date"`  // must be "time.Time" but sometimes responds "false"
	DesktopLastDate  interface{} `json:"desktop_last_date"` // must be "time.Time" but sometimes responds "false"
	Phones           interface{} `json:"phones,omitempty"`  // must be "map[string]string" but sometimes responds "false"
}
