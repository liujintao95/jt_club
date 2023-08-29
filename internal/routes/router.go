package routes

import (
	"JT_CLUB/internal/api"
	"JT_CLUB/internal/constant"
	"JT_CLUB/internal/models"
	"JT_CLUB/pkg/cache"
	"JT_CLUB/pkg/log"
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(Cors())
	server.Use(ZapLogger(log.Logger))
	server.Use(ZapRecovery(log.Logger, true))

	server.POST("/sign_in", api.SignIn)
	server.POST("/sign_up", api.SignUp)
	user := server.Group("/user", LoginRequired())
	{
		user.POST("/edit", api.UserEdit)
		user.POST("/select", api.UserSelect)
		user.POST("/contact/application", api.ContactApplication)
		user.POST("/contact/confirm", api.ContactConfirm)
		user.POST("/contact/confirm/list", api.ContactConfirmList)
		user.POST("/contact/list", api.ContactList)
	}
	msg := server.Group("/msg", LoginRequired())
	{
		msg.POST("/socket", api.SocketClient)
	}
	return server
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			ctx.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Logger.Error("HttpError", zap.Any("HttpError", err))
			}
		}()

		ctx.Next()
	}
}

func ZapLogger(lg *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		post := ""

		if ctx.Request.Method == "POST" {
			// 把request的内容读取出来
			bodyBytes, _ := io.ReadAll(ctx.Request.Body)
			_ = ctx.Request.Body.Close()
			// 把刚刚读出来的再写进去
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			switch ctx.ContentType() {
			case "application/json":
				var result map[string]interface{}
				d := jsoniter.NewDecoder(bytes.NewReader(bodyBytes))
				d.UseNumber()
				if err := d.Decode(&result); err == nil {
					bt, _ := jsoniter.Marshal(result)
					post = string(bt)
				}
			default:
				post = string(bodyBytes)
			}
		}

		ctx.Next()

		cost := time.Since(start)
		lg.Info(path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("post", post),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

func ZapRecovery(lg *zap.Logger, stack bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					lg.Error(ctx.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = ctx.Error(err.(error)) // nolint: err check
					ctx.Abort()
					return
				}

				if stack {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}

func LoginRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var currentUser *models.User
		token, err := ctx.Cookie(constant.TokenKey)
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
		}
		err = cache.Cache.Get(token).Scan(currentUser)
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
		}
		ctx.Set(constant.CurrentUserKey, currentUser)
		ctx.Next()
	}
}
