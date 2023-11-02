package model

import "github.com/gogf/gf/v2/frame/g"

type GetHistoryMessageInput struct {
	MessageId string `json:"message_id" description:"起始ID"`
	ContactId string `json:"contact_id" description:"聊天对象ID"`
	Page      int    `json:"page" description:"分页码"`
	Size      int    `json:"size" description:"分页数量"`
}

type MessageInfoItem struct {
	g.Meta       `orm:"table:message"`
	MessageId    string `json:"message_id"`
	CreatedAt    string `json:"created_at"`
	Avatar       string `json:"avatar"`
	FromUsername string `json:"from_username"`
	From         string `json:"from"`
	To           string `json:"to"`
	Content      string `json:"content"`
	ContentType  int    `json:"contentType"`
	Type         string `json:"type"`
	MessageType  int    `json:"messageType"`
	Url          string `json:"url"`
	FileSuffix   string `json:"fileSuffix"`
	FilePath     string `json:"filePath"`
}

type GetHistoryMessageOutput struct {
	Messages []MessageInfoItem
	Total    int `json:"total" description:"数据总数"`
}
