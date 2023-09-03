package chatserver

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"jt_chat/internal/consts"
	protocol "jt_chat/manifest/protobuf"
	"net/http"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  consts.ReadBufferSize,
	WriteBufferSize: consts.WriteBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn *websocket.Conn
	Uid  string
	Name string
	Send chan []byte
	ctx  context.Context
}

func (c *Client) Write() {
	defer func() {
		_ = c.Conn.Close()
	}()
	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			g.Log().Error(c.ctx, "client write message:"+err.Error())
		}
	}
}

func (c *Client) Read() {
	defer func() {
		ChatServer.Cancellation <- c
		close(c.Send)
		_ = c.Conn.Close()
	}()
	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			g.Log().Error(c.ctx, "client read:"+err.Error())
			break
		}
		msg := &protocol.Message{}
		err = proto.Unmarshal(message, msg)
		if err != nil {
			g.Log().Error(c.ctx, "proto unmarshal:"+err.Error())
			continue
		}
		if msg.Type == consts.HeatBeatType {
			pong := &protocol.Message{
				Content: consts.PongMsg,
				Type:    consts.HeatBeatType,
			}
			pongBytes, err := proto.Marshal(pong)
			if err != nil {
				g.Log().Error(c.ctx, "proto marshal:"+err.Error())
				continue
			}
			err = c.Conn.WriteMessage(websocket.BinaryMessage, pongBytes)
			if err != nil {
				g.Log().Error(c.ctx, "client pong message:"+err.Error())
			}
		} else {
			ChatServer.Send <- message
		}

	}
}
