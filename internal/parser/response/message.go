package response

type MsgResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessMsg(data any) *MsgResponse {
	msg := &MsgResponse{
		Code: 0,
		Msg:  "SUCCESS",
		Data: data,
	}
	return msg
}

func FailMsg(msg string) *MsgResponse {
	msgObj := &MsgResponse{
		Code: -1,
		Msg:  msg,
	}
	return msgObj
}
