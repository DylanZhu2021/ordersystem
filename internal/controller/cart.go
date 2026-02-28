package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
	"ordersystem/utility"
)

var Cart = &cCart{}

type cCart struct{}

func (c *cCart) Add(r *ghttp.Request) {
	var req v1.CartAddReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Cart().Add(r.Context(), userId, &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cCart) List(r *ghttp.Request) {
	userId := r.GetCtxVar("userId").Int64()
	res, err := service.Cart().List(r.Context(), userId)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cCart) Update(r *ghttp.Request) {
	var req v1.CartUpdateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Cart().Update(r.Context(), userId, &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cCart) Delete(r *ghttp.Request) {
	var req v1.CartDeleteReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Cart().Delete(r.Context(), userId, &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cCart) Clear(r *ghttp.Request) {
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Cart().Clear(r.Context(), userId); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}
