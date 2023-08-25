package bll

import (
	"JT_CLUB/internal/constant"
	"JT_CLUB/internal/models"
	"JT_CLUB/internal/socket"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func RunSocketClient(ctx *gin.Context, user *models.User) error {
	var (
		client   *socket.Client
		ws       *websocket.Conn
		err      error
		upGrader = websocket.Upgrader{
			ReadBufferSize:  constant.ReadBufferSize,
			WriteBufferSize: constant.WriteBufferSize,
		}
	)
	ws, err = upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return err
	}
	client = &socket.Client{
		Conn: ws,
		Uid:  user.Uid,
		Name: user.Name,
		Send: make(chan []byte),
	}
	socket.ChatServer.Register <- client
	go client.Read()
	go client.Write()
	return nil
}
