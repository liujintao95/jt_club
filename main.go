package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"jt_chat/internal/chatserver"
	_ "jt_chat/internal/packed"

	_ "jt_chat/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"jt_chat/internal/cmd"
)

func main() {
	go chatserver.ChatServer.Start()
	cmd.Main.Run(gctx.GetInitCtx())
}
