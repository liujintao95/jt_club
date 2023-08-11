package socket

import (
	"JT_CLUB/internal/constant"
	"JT_CLUB/pkg/log"
	"JT_CLUB/pkg/protocol"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Uid  string
	Name string
	Send chan []byte
}

func (c *Client) Write() {
	defer func() {
		_ = c.Conn.Close()
	}()
	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			log.Logger.Error("client write message:" + err.Error())
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
			log.Logger.Error("client read:" + err.Error())
			break
		}
		msg := &protocol.Message{}
		err = proto.Unmarshal(message, msg)
		if err != nil {
			log.Logger.Error("proto unmarshal:" + err.Error())
			continue
		}
		if msg.Type == constant.HeatBeatType {
			pong := &protocol.Message{
				Content: constant.PongMsg,
				Type:    constant.HeatBeatType,
			}
			pongBytes, err := proto.Marshal(pong)
			if err != nil {
				log.Logger.Error("proto marshal:" + err.Error())
				continue
			}
			err = c.Conn.WriteMessage(websocket.BinaryMessage, pongBytes)
			if err != nil {
				log.Logger.Error("client pong message:" + err.Error())
			}
		} else {
			ChatServer.Send <- message
		}

	}
}
