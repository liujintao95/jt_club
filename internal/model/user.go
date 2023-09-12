package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type RegisterInput struct {
	Name     string `json:"name" description:"用户名"`
	Email    string `json:"email" description:"邮箱"`
	Password string `json:"password" description:"密码"`
}

type RegisterOutput struct {
	Id string
}

type UpdateInput struct {
	Name   string `json:"name" description:"用户名"`
	Avatar string `json:"avatar" description:"头像"`
}

type UpdateOutput struct {
	Uid    string `json:"uid" description:"UID"`
	Name   string `json:"name" description:"用户名"`
	Avatar string `json:"avatar" description:"头像"`
}

type GetListInput struct {
	NameOrId string `json:"name_or_id" description:"名称或ID"`
	Page     int    `json:"page" description:"分页码"`
	Size     int    `json:"size" description:"分页数量"`
}

type UserInfoItem struct {
	g.Meta `orm:"table:user"`
	Uid    string
	Name   string
	Email  string
	Avatar string
}

type GetListOutput struct {
	Users []UserInfoItem
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type GetContactListInput struct {
	NameOrId string `json:"name_or_id" description:"名称或ID"`
	Page     int    `json:"page" description:"分页码"`
	Size     int    `json:"size" description:"分页数量"`
}

type UserGroupItem struct {
	g.Meta  `orm:"table:user_group"`
	Gid     string
	Name    string
	AdminId string
	Notice  string
	Avatar  string
}

type ContactInfoItem struct {
	g.Meta       `orm:"table:user_contacts"`
	ContactId    string
	ContactType  int
	ContactNotes string
	LastMsg      string
	LastTime     time.Time
	User         UserInfoItem  `json:"user" orm:"with:uid=contact_id"`
	Group        UserGroupItem `json:"user_group" orm:"with:gid=contact_id"`
}

type GetContactListOutput struct {
	Contacts []ContactInfoItem
	Page     int `json:"page" description:"分页码"`
	Size     int `json:"size" description:"分页数量"`
	Total    int `json:"total" description:"数据总数"`
}

type GetContactApplicationListInput struct {
	ContactType uint `json:"contact_type" description:"申请类型"`
	Page        int  `json:"page" description:"分页码"`
	Size        int  `json:"size" description:"分页数量"`
}

type ContactApplicationItem struct {
	g.Meta      `orm:"table:contact_application"`
	AppId       string
	Uid         string
	ContactId   string
	ContactType uint
	Status      uint
	Notice      string
	User        UserInfoItem `orm:"with:uid=uid"`
}

type GetContactApplicationListOutput struct {
	Applications []ContactApplicationItem
	Page         int `json:"page" description:"分页码"`
	Size         int `json:"size" description:"分页数量"`
	Total        int `json:"total" description:"数据总数"`
}

type CreateContactApplicationInput struct {
	ContactId   string `json:"contact_id" description:"请求对象ID"`
	ContactType uint   `json:"contact_type" description:"请求对象类型"`
	Notice      string `json:"notice" description:"备注"`
}

type CreateContactApplicationOutput struct {
	AppID string `json:"app_id" description:"请求ID"`
}

type UpdateContactApplicationInput struct {
	AppId  string `json:"app_id" description:"请求ID"`
	Status uint   `json:"status" description:"状态信息"`
	Notice string `json:"notice" description:"备注"`
}

type UpdateContactApplicationOutput struct {
}

type UserGroupMapItem struct {
	g.Meta `orm:"table:user_group_map"`
	MapId  string
	Uid    string
	Gid    string
	User   UserInfoItem  `json:"user" orm:"with:uid=uid"`
	Group  UserGroupItem `json:"user_group" orm:"with:gid=git"`
}

type CreateUserGroupInput struct {
	Name   string `json:"name" description:"群名"`
	Avatar string `json:"avatar" description:"群头像"`
	Notice string `json:"notice" description:"群备注"`
}

type CreateUserGroupOutput struct {
	Gid string `json:"gid" description:"群ID"`
}

type UpdateUserGroupInput struct {
	Gid     string `json:"gid" description:"群ID" v:"required"`
	Name    string `json:"name" description:"群名"`
	Avatar  string `json:"avatar" description:"群头像"`
	Notice  string `json:"notice" description:"群备注"`
	AdminId string `json:"admin_id" description:"群管理"`
}

type UpdateUserGroupOutput struct {
}
