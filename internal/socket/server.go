package socket

import (
	"JT_CLUB/conf"
	"JT_CLUB/internal/constant"
	"JT_CLUB/internal/models"
	"JT_CLUB/pkg/db"
	"JT_CLUB/pkg/log"
	"JT_CLUB/pkg/protocol"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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
}

func NewSocketServer() *Server {
	return &Server{
		mutex:        &sync.Mutex{},
		Clients:      make(map[string]*Client),
		Send:         make(chan []byte),
		Register:     make(chan *Client),
		Cancellation: make(chan *Client),
	}
}

func (s *Server) Start() {
	log.Logger.Info("start socket server")
	for {
		select {
		case client := <-s.Register:
			s.registerClient(client)
			log.Logger.Info(fmt.Sprintf("%s(%s) login jt", client.Name, client.Uid))
		case client := <-s.Cancellation:
			s.cancellationClient(client)
			log.Logger.Info(fmt.Sprintf("%s(%s) logout jt", client.Name, client.Uid))
		case message := <-s.Send:
			// 信息发送
			msg := &protocol.Message{}
			err := proto.Unmarshal(message, msg)
			if err != nil {
				log.Logger.Error("proto unmarshal:" + err.Error())
				continue
			}
			_, exits := s.Clients[msg.From]
			if !exits {
				continue
			}
			s.saveMsg(msg)
			if msg.MessageType == constant.SingleChat {
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
		FromUsername: constant.AdminName,
		From:         constant.AdminUid,
		To:           client.Uid,
		Content:      constant.WelcomeMsg,
		ContentType:  constant.TextType,
		Type:         constant.NormalType,
		MessageType:  constant.SingleChat,
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		log.Logger.Error("proto marshal:" + err.Error())
		return
	}
	err = client.Conn.WriteMessage(websocket.BinaryMessage, msgBytes)
	if err != nil {
		log.Logger.Error("register welcome message:" + err.Error())
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
	var users []*models.Users
	sql := "SELECT `user`.* FROM `user` INNER JOIN user_group_map WHERE user_group_map.gid = ?"
	err := db.Conn.Get(&users, sql, msg.To)
	if err != nil {
		log.Logger.Error("socket get users:" + err.Error())
		return
	}
	for _, user := range users {
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
	var filePath string
	messageId := uuid.New().String()
	if msg.File != nil {
		filePath = filepath.Join(conf.FilePath, messageId)
		err := os.WriteFile(filePath, msg.File, os.ModePerm)
		if err != nil {
			log.Logger.Error("write file to local:" + err.Error())
		}
	}
	sql := `
		insert into message(
			message_id, avatar, from_username, 'from', 'to', content, content_type,
            type, message_type, url, file_suffix, file_path
        ) values (?,?,?,?,?,?,?,?,?,?,?,?)
	`
	_, err := db.Conn.Exec(sql,
		messageId, msg.Avatar, msg.FromUsername, msg.From, msg.To, msg.Content,
		msg.ContentType, msg.Type, msg.MessageType, msg.Url, msg.FileSuffix, filePath,
	)
	if err != nil {
		log.Logger.Error("save msg:" + err.Error())
	}
}
