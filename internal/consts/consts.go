package consts

const (
	Salt              = "jt"
	TokenType         = "Bearer"
	CacheModeRedis    = 2
	BackendServerName = "锦涛社"
	MultiLogin        = false
	GTokenExpireIn    = 10 * 24 * 60 * 60
)

const (
	ErrCodeMissingParameterMsg = "请检查是否缺少参数"
	ErrLoginMsg                = "登录失败，账号或密码错误"
	ErrSecretAnswerMsg         = "密保问题不正确"
	ErrResourcePermissionFail  = "没有权限操作"
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
	ContactsUserType = iota
	ContactsGroupType
)

const (
	ApplicationWaitStatus = iota
	ApplicationAgreeStatus
	ApplicationRefuseStatus
)

const (
	TransportTypeHeartBeat = "heartbeat"
	TransportTypeWebRTC    = "webrtc"
	TransportTypeNormal    = "normal"
)

const (
	MsgPong    = "pong"
	MsgWelcome = "欢迎登录锦涛社!"
)

const (
	AdminName = "admin"
	AdminUid  = "0000001"
)

const (
	ContentTypeText = iota
	ContentTypeFile
	ContentTypePicture
	ContentTypeAudio
	ContentTypeVideo
)

const (
	ChatSingle = ContactsUserType
	ChatGroup  = ContactsGroupType
)

const DefaultDuration = 60 * 60 * 8

const FilePath = "./files"

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 1024
)

const (
	EmailRegex = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
)
