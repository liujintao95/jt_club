package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
	"jt_chat/internal/cmd"
	"jt_chat/internal/consts"
	"testing"
)

const (
	Prefix   = "http://127.0.0.1:5000"
	Account  = "admin@jt.com"
	Password = "123456"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type LoginResp struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data cmd.LoginRes `json:"data"`
}

func TestLogin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var resp Resp
		client := g.Client()
		client.SetPrefix(Prefix)
		data := g.Map{
			"account":  Account,
			"password": Password,
		}
		content := client.GetContent(context.Background(), "/login", data)
		fmt.Println(content)
		err := json.Unmarshal([]byte(content), &resp)
		t.AssertNil(err)
		t.Assert(resp.Code, 0)
	})
}

func GetLoginClient() (map[string]string, error) {
	var (
		resp LoginResp
	)
	headerMap := make(map[string]string)
	client := g.Client().SetPrefix(Prefix)
	data := g.Map{
		"account":  Account,
		"password": Password,
	}
	content := client.GetContent(context.Background(), "/login", data)
	err := gconv.Struct(content, &resp)
	if err != nil {
		return nil, err
	}
	headerMap["Authorization"] = fmt.Sprintf("%s %s", resp.Data.Type, resp.Data.Token)
	return headerMap, err
}

func TestGetList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var resp Resp
		headerMap, err := GetLoginClient()
		t.AssertNil(err)
		client := g.Client().SetPrefix(Prefix).Header(headerMap)
		data := g.Map{
			"page":     1,
			"size":     10,
			"NameOrId": "忆夕",
		}
		content := client.PostContent(context.Background(), "/user/list", data)
		fmt.Println(content)
		err = gconv.Struct(content, &resp)
		t.AssertNil(err)
		t.Assert(resp.Code, 0)
	})
}

func TestSetContactApplication(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var resp Resp
		headerMap, err := GetLoginClient()
		t.AssertNil(err)
		client := g.Client().SetPrefix(Prefix).Header(headerMap)
		data := g.Map{
			"ContactId":   100001,
			"ContactType": consts.ContactsUserType,
			"Notice":      "加加加",
		}
		content := client.PostContent(context.Background(), "/user/contact/application", data)
		fmt.Println(content)
		err = gconv.Struct(content, &resp)
		t.AssertNil(err)
		t.Assert(resp.Code, 0)
	})
}

func TestGetContactApplicationList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var resp Resp
		headerMap, err := GetLoginClient()
		t.AssertNil(err)
		client := g.Client().SetPrefix(Prefix).Header(headerMap)
		data := g.Map{
			"Page": 1,
			"Size": 10,
		}
		content := client.PostContent(context.Background(), "/user/contact/application/list", data)
		fmt.Println(content)
		err = gconv.Struct(content, &resp)
		t.AssertNil(err)
		t.Assert(resp.Code, 0)
	})
}

func TestUpdateContactApplication(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var resp Resp
		headerMap, err := GetLoginClient()
		t.AssertNil(err)
		client := g.Client().SetPrefix(Prefix).Header(headerMap)
		data := g.Map{
			"app_id": "865d0269-b1da-4880-9d62-aef37f74a4c5",
			"status": 1,
		}
		content := client.PostContent(context.Background(), "/user/contact/confirm", data)
		fmt.Println(content)
		err = gconv.Struct(content, &resp)
		t.AssertNil(err)
		t.Assert(resp.Code, 0)
	})
}

func TestGetContactList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var resp Resp
		headerMap, err := GetLoginClient()
		t.AssertNil(err)
		client := g.Client().SetPrefix(Prefix).Header(headerMap)
		data := g.Map{
			"Page":       1,
			"Size":       10,
			"name_or_id": "01",
		}
		content := client.PostContent(context.Background(), "/user/contact/list", data)
		fmt.Println(content)
		err = gconv.Struct(content, &resp)
		t.AssertNil(err)
		t.Assert(resp.Code, 0)
	})
}
