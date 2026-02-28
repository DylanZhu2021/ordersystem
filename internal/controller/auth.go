package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
	"ordersystem/utility"
)

var Auth = &cAuth{}

type cAuth struct{}

func (c *cAuth) WechatLogin(r *ghttp.Request) {
	var req v1.WechatLoginReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	res, err := service.Auth().WechatLogin(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cAuth) GetUserInfo(r *ghttp.Request) {
	userId := r.GetCtxVar("userId").Int64()
	res, err := service.Auth().GetUserInfo(r.Context(), userId)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cAuth) UpdateUser(r *ghttp.Request) {
	var req v1.UserUpdateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Auth().UpdateUser(r.Context(), userId, &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}
