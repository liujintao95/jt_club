package main

import (
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
