package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"jt_chat/internal/controller/chat"
	"jt_chat/internal/controller/user"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareCORS)
			s.Use(ghttp.MiddlewareHandlerResponse)
			token, err := StartGToken()
			if err != nil {
				return err
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				err = token.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.NewV1(),
					)
				})
				group.Group("/chat", func(group *ghttp.RouterGroup) {
					group.Bind(
						chat.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
