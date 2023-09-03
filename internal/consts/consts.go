package consts

const (
	Salt              = "jt"
	TokenType         = "Bearer"
	CacheModeRedis    = 2
	BackendServerName = "锦涛社"
	MultiLogin        = true
	GTokenExpireIn    = 10 * 24 * 60 * 60
)

const (
	CodeMissingParameterMsg = "请检查是否缺少参数"
	ErrLoginMsg             = "登录失败，账号或密码错误"
	ErrSecretAnswerMsg      = "密保问题不正确"
	ResourcePermissionFail  = "没有权限操作"
)

const (
	CtxUserId     = "CtxUserId"
	CtxUserName   = "CtxUserName"
	CtxUserAvatar = "CtxUserAvatar"
	CtxUserEmail  = "CtxUserEmail"
	CtxUserSex    = "CtxUserSex"
	CtxUserSign   = "CtxUserSign"
	CtxUserStatus = "CtxUserStatus"
)

const (
	ContactsUserType  = iota
	ContactsGroupType = iota
)

const (
	ApplicationWaitStatus   = iota
	ApplicationAgreeStatus  = iota
	ApplicationRefuseStatus = iota
)

const (
	HeatBeatType = "heatbeat"
	WebRTCType   = "webrtc"
	NormalType   = "normal"
)

const (
	PongMsg    = "pong"
	WelcomeMsg = "welcome to jt club!"
)

const (
	AdminName = "admin"
	AdminUid  = "0000001"
)

const (
	TextType = iota
	FileType
	PictureType
)

const (
	SingleChat = iota
	GroupChat
)

const DefaultDuration = 60 * 60 * 8

const FilePath = "./files"

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 1024
)
