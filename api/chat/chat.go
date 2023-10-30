// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package chat

import (
	"context"
	
	"jt_chat/api/chat/v1"
)

type IChatV1 interface {
	Websocket(ctx context.Context, req *v1.WebsocketReq) (res *v1.WebsocketRes, err error)
	GetHistoryMessage(ctx context.Context, req *v1.GetHistoryMessageReq) (res *v1.GetHistoryMessageRes, err error)
}


