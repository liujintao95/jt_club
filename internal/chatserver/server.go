package chatserver

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"jt_chat/internal/consts"
	"jt_chat/internal/dao"
	"jt_chat/internal/model/entity"
	"jt_chat/manifest/protobuf"
	"os"
	"path/filepath"
	"sync"
)

var ChatServer = NewSocketServer()

type Server struct {
	mutex        *sync.Mutex
	Clients      map[string]*Client
	Send         chan []byte
	Register     chan *Client
	Cancellation chan *Client
	ctx          context.Context
}

func NewSocketServer() *Server {
	return &Server{
		mutex:        &sync.Mutex{},
		Clients:      make(map[string]*Client),
		Send:         make(chan []byte),
		Register:     make(chan *Client),
		Cancellation: make(chan *Client),
		ctx:          gctx.New(),
	}
}

func (s *Server) Start() {
	g.Log().Info(s.ctx, "start chat server")
	for {
		select {
		case client := <-s.Register:
			s.registerClient(client)
			g.Log().Info(s.ctx, fmt.Sprintf("%s(%s) login jt", client.Name, client.Uid))
		case client := <-s.Cancellation:
			s.cancellationClient(client)
			g.Log().Info(s.ctx, fmt.Sprintf("%s(%s) logout jt", client.Name, client.Uid))
		case message := <-s.Send:
			// 信息发送
			msg := &protocol.Message{}
			err := proto.Unmarshal(message, msg)
			if err != nil {
				g.Log().Error(s.ctx, "proto unmarshal:"+err.Error())
				continue
			}
			_, exits := s.Clients[msg.From]
			if !exits {
				continue
			}
			s.saveMsg(msg)
			if msg.MessageType == consts.SingleChat {
				s.sendSingleMsg(msg)
			} else {
				s.sendGroupMsg(msg)
			}
		}
	}
}

func (s *Server) registerClient(client *Client) {
	s.Clients[client.Uid] = client
	msg := &protocol.Message{
		FromUsername: consts.AdminName,
		From:         consts.AdminUid,
		To:           client.Uid,
		Content:      consts.WelcomeMsg,
		ContentType:  consts.TextType,
		Type:         consts.NormalType,
		MessageType:  consts.SingleChat,
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		g.Log().Error(s.ctx, "proto marshal:"+err.Error())
		return
	}
	err = client.Conn.WriteMessage(websocket.BinaryMessage, msgBytes)
	if err != nil {
		g.Log().Error(s.ctx, "register welcome message:"+err.Error())
	}
}

func (s *Server) cancellationClient(client *Client) {
	if _, ok := s.Clients[client.Uid]; ok {
		delete(s.Clients, client.Uid)
	}
}

func (s *Server) sendSingleMsg(msg *protocol.Message) {
	client, ok := s.Clients[msg.To]
	if ok {
		msgByte, err := proto.Marshal(msg)
		if err == nil {
			client.Send <- msgByte
		}
	}
}

func (s *Server) sendGroupMsg(msg *protocol.Message) {
	var (
		groupUsers []*entity.UserGroupMap
		err        error
	)
	err = dao.UserGroupMap.Ctx(s.ctx).Where(
		dao.UserGroupMap.Columns().Gid, msg.To,
	).Where(
		dao.UserGroupMap.Columns().Deleted, false,
	).Scan(groupUsers)
	if err != nil {
		g.Log().Error(s.ctx, "socket get users:"+err.Error())
		return
	}
	for _, user := range groupUsers {
		if user.Uid == msg.From {
			continue
		}
		sendMsg := msg
		sendMsg.From = msg.To
		sendMsg.To = user.Uid
		client, ok := s.Clients[msg.To]
		if !ok {
			continue
		}
		msgByte, err := proto.Marshal(sendMsg)
		if err != nil {
			client.Send <- msgByte
		}
	}
}

func (s *Server) saveMsg(msg *protocol.Message) {
	var (
		messageId string
		filePath  string
		newMsg    entity.Message
		err       error
	)
	messageId = uuid.New().String()
	if msg.File != nil {
		filePath = filepath.Join(consts.FilePath, messageId)
		err := os.WriteFile(filePath, msg.File, os.ModePerm)
		if err != nil {
			g.Log().Error(s.ctx, "write file to local:"+err.Error())
		}
	}
	newMsg = entity.Message{
		MessageId:    messageId,
		Avatar:       msg.Avatar,
		FromUsername: msg.FromUsername,
		From:         msg.From,
		To:           msg.To,
		Content:      msg.Content,
		ContentType:  int(msg.ContentType),
		Type:         msg.Type,
		MessageType:  int(msg.MessageType),
		Url:          msg.Url,
		FileSuffix:   msg.FileSuffix,
		FilePath:     filePath,
	}
	_, err = dao.Message.Ctx(s.ctx).Data(newMsg).InsertAndGetId()
	if err != nil {
		g.Log().Error(s.ctx, "save msg:"+err.Error())
	}
}
