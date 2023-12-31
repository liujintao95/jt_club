package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserInfoBase struct {
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type UserGroupItem struct {
	Gid    string `json:"gid"     `
	Name   string `json:"name"    `
	Avatar string `json:"avatar"  `
}

type RegisterReq struct {
	g.Meta   `path:"/register" method:"post" summary:"用户注册"`
	Name     string `json:"name" v:"required"`
	Email    string `json:"email" v:"required"`
	Password string `json:"password" v:"required"`
}

type RegisterRes struct {
	Email string `json:"email"`
}

type UpdateReq struct {
	g.Meta `path:"/update" method:"post" summary:"更新用户信息"`
	Name   string `json:"name" v:"required"`
	Avatar string `json:"avatar" v:"required"`
}

type UpdateRes struct {
	UserInfoBase
}

type GetListReq struct {
	g.Meta   `path:"/list" method:"post" summary:"查询用户列表"`
	NameOrId string `json:"name_or_id" description:"查询条件" v:"required"`
	Page     int    `json:"page" description:"分页码" v:"required"`
	Size     int    `json:"size" description:"分页数量" v:"required"`
}

type GetListRes struct {
	Users interface{} `json:"users" description:"用户列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type GetContactListReq struct {
	g.Meta   `path:"/contact/list" method:"post" summary:"查询联系人列表"`
	NameOrId string `json:"name_or_id"`
}

type GetContactListRes struct {
	Contacts interface{} `json:"contacts" description:"联系人列表"`
	Total    int         `json:"total" description:"数据总数"`
}

type DeleteContactReq struct {
	g.Meta `path:"/contact/delete" method:"post" summary:"删除联系人"`
	Uid    string `json:"uid" description:"联系人Uid" v:"required"`
}

type DeleteContactRes struct{}

type GetContactApplicationListReq struct {
	g.Meta      `path:"/contact/application/list" method:"post" summary:"申请添加联系人"`
	ContactType uint `json:"contact_type" description:"申请类型" v:"required"`
	Page        int  `json:"page" description:"分页码" v:"required"`
	Size        int  `json:"size" description:"分页数量" v:"required"`
}

type GetContactApplicationListRes struct {
	Applications interface{} `json:"applications" description:"联系人请求列表"`
	Page         int         `json:"page" description:"分页码"`
	Size         int         `json:"size" description:"分页数量"`
	Total        int         `json:"total" description:"数据总数"`
}

type CreateContactApplicationReq struct {
	g.Meta      `path:"/contact/application/create" method:"post" summary:"申请添加联系人"`
	ContactId   string `json:"contact_id" description:"申请对象ID" v:"required"`
	ContactType uint   `json:"contact_type" description:"申请类型" v:"required"`
	Notice      string `json:"notice" description:"申请备注" v:"required"`
}

type CreateContactApplicationRes struct{}

type UpdateContactApplicationReq struct {
	g.Meta `path:"/contact/confirm/update" method:"post" summary:"审批添加联系人信息"`
	AppId  string `json:"app_id" description:"申请ID" v:"required"`
	Status uint   `json:"status" description:"状态" v:"required"`
	Notice string `json:"notice" description:"备注"`
}

type UpdateContactApplicationRes struct{}

type CreateUserGroupReq struct {
	g.Meta `path:"/group/create" method:"post" summary:"新建群"`
	Name   string `json:"name" description:"群名" v:"required"`
	Avatar string `json:"avatar" description:"群头像" v:"required"`
	Notice string `json:"notice" description:"群备注"`
}

type CreateUserGroupRes struct{}

type UpdateUserGroupReq struct {
	g.Meta  `path:"/group/update" method:"post" summary:"修改群信息"`
	Gid     string `json:"gid" description:"群ID" v:"required"`
	Name    string `json:"name" description:"群名"`
	Avatar  string `json:"avatar" description:"群头像"`
	Notice  string `json:"notice" description:"群备注"`
	AdminId string `json:"admin_id" description:"群管理"`
}

type UpdateUserGroupRes struct{}

type DeleteUserGroupMapReq struct {
	g.Meta `path:"/group/exit" method:"post" summary:"退出群聊"`
	Gid    string `json:"gid" description:"群ID" v:"required"`
}

type DeleteUserGroupMapRes struct{}
