package chat

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"jt_chat/internal/chatserver"
	"jt_chat/internal/consts"

	"jt_chat/api/chat/v1"
)

func (c *ControllerV1) Websocket(ctx context.Context, req *v1.WebsocketReq) (res *v1.WebsocketRes, err error) {
	var (
		r      *ghttp.Request
		ws     *ghttp.WebSocket
		client *chatserver.Client
	)
	r = g.RequestFromCtx(ctx)
	if ws, err = r.WebSocket(); err != nil {
		return nil, err
	}
	client = &chatserver.Client{
		Conn: ws.Conn,
		Uid:  gconv.String(ctx.Value(consts.CtxUserId)),
		Name: gconv.String(ctx.Value(consts.CtxUserName)),
		Send: make(chan []byte),
	}
	chatserver.ChatServer.Register <- client
	go client.Read()
	go client.Write()
	return nil, gerror.NewCode(gcode.CodeOK)
}
