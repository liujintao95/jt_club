package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"jt_chat/internal/consts"
)

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code    int         `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.Exit()
}

func dataReturn(r *ghttp.Request, code int, req ...interface{}) *JsonRes {
	var msg string
	var data interface{}
	if len(req) > 0 {
		msg = gconv.String(req[0])
	}
	if len(req) > 1 {
		data = req[1]
	}
	response := &JsonRes{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	r.SetParam("apiReturnRes", response)
	return response
}

func Auth(r *ghttp.Request) {
	res := dataReturn(r, consts.CodeUnauthorized, "用户验证失败，请登录！")
	r.Response.Status = 401
	r.Response.WriteJsonExit(res)
}
