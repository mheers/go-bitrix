package task_comments

import (
	"time"
)

type TaskComments struct {
	ID              int                   `json:"ID,string"`
	AuthorID        int                   `json:"AUTHOR_ID,string"`
	AuthorName      string                `json:"AUTHOR_NAME"`
	AuthorEmail     string                `json:"AUTHOR_EMAIL"`
	PostDate        time.Time             `json:"POST_DATE"`
	PostMessage     string                `json:"POST_MESSAGE"`
	PostMessageHTML string                `json:"POST_MESSAGE_HTML"`
	AttachedObjects map[string]Attachment `json:"ATTACHED_OBJECTS,omitempty"`
}

type Attachment struct {
	AttachmentID int    `json:"ATTACHMENT_ID,string"`
	Name         string `json:"NAME"`
	Size         int    `json:"SIZE,string"`
	FileID       int    `json:"FILE_ID,string"`
	DownloadURL  string `json:"DOWNLOAD_URL"`
	ViewURL      string `json:"VIEW_URL"`
}
