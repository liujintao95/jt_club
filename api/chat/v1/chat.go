package v1

import "github.com/gogf/gf/v2/frame/g"

type WebsocketReq struct {
	g.Meta `path:"/ws" method:"get" summary:"链接websocket"`
}

type WebsocketRes struct {
}
