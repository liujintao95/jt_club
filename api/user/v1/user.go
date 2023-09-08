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
	Users interface{}
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type GetContactListReq struct {
	g.Meta   `path:"/contact/list" method:"post" summary:"查询联系人列表"`
	NameOrId string `json:"name_or_id"`
	Page     int    `json:"page" description:"分页码" v:"required"`
	Size     int    `json:"size" description:"分页数量" v:"required"`
}

type GetContactListRes struct {
	Contacts interface{}
	Page     int `json:"page" description:"分页码"`
	Size     int `json:"size" description:"分页数量"`
	Total    int `json:"total" description:"数据总数"`
}

type GetContactApplicationListReq struct {
	g.Meta `path:"/contact/application/list" method:"post" summary:"申请添加联系人"`
	Page   int `json:"page" description:"分页码" v:"required"`
	Size   int `json:"size" description:"分页数量" v:"required"`
}

type GetContactApplicationListRes struct {
	Applications interface{}
	Page         int `json:"page" description:"分页码" v:"required"`
	Size         int `json:"size" description:"分页数量" v:"required"`
	Total        int `json:"total" description:"数据总数" v:"required"`
}

type SetContactApplicationReq struct {
	g.Meta      `path:"/contact/application" method:"post" summary:"申请添加联系人"`
	ContactId   string `json:"contact_id" v:"required"`
	ContactType uint   `json:"contact_type" v:"required"`
	Notice      string `json:"notice" v:"required"`
}

type SetContactApplicationRes struct{}

type UpdateContactApplicationReq struct {
	g.Meta `path:"/contact/confirm" method:"post" summary:"审批添加联系人信息"`
	AppId  string `json:"app_id" v:"required"`
	Status uint   `json:"status" v:"required"`
}

type UpdateContactApplicationRes struct{}