package v1

import "github.com/gogf/gf/v2/frame/g"

type WebsocketReq struct {
	g.Meta `path:"/ws" method:"get" summary:"链接websocket"`
}

type WebsocketRes struct {
}

type GetHistoryMessageReq struct {
	g.Meta    `path:"/message/list" method:"post" summary:"获取联系人的历史聊天信息"`
	MessageId string `json:"message_id" description:"起始ID"`
	ContactId string `json:"contact_id" description:"聊天对象ID" v:"required"`
	Page      int    `json:"page" description:"分页码" v:"required"`
	Size      int    `json:"size" description:"分页数量" v:"required"`
}

type GetHistoryMessageRes struct {
	Messages interface{} `json:"messages" description:"消息列表"`
	Page     int         `json:"page" description:"分页码"`
	Size     int         `json:"size" description:"分页数量"`
	Total    int         `json:"total" description:"数据总数"`
}
