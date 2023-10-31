package chatserver

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"jt_chat/internal/consts"
	protocol "jt_chat/manifest/protobuf"
)

type Client struct {
	Conn *websocket.Conn
	Uid  string
	Name string
	Send chan []byte
	Ctx  context.Context
	Stop chan bool
}

func (c *Client) Write() {
	for {
		select {
		case <-c.Stop:
			return
		case message := <-c.Send:
			err := c.Conn.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				g.Log().Error(c.Ctx, "client write message:"+err.Error())
			}
		}
	}
}

func (c *Client) Read() {
	defer func() {
		ChatServer.Cancellation <- c
		c.Stop <- true
		close(c.Send)
		close(c.Stop)
		_ = c.Conn.Close()
	}()
	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			g.Log().Error(c.Ctx, "client read:"+err.Error())
			break
		}
		msg := &protocol.Message{}
		err = proto.Unmarshal(message, msg)
		if err != nil {
			g.Log().Error(c.Ctx, "proto unmarshal:"+err.Error())
			continue
		}
		if msg.Type == consts.TransportTypeHeartBeat {
			pong := &protocol.Message{
				Content: consts.MsgPong,
				Type:    consts.TransportTypeHeartBeat,
			}
			pongBytes, err := proto.Marshal(pong)
			if err != nil {
				g.Log().Error(c.Ctx, "proto marshal:"+err.Error())
				continue
			}
			err = c.Conn.WriteMessage(websocket.BinaryMessage, pongBytes)
			if err != nil {
				g.Log().Error(c.Ctx, "client pong message:"+err.Error())
			}
		} else {
			ChatServer.Send <- message
		}

	}
}
