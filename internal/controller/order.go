package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
	"ordersystem/utility"
)

var Order = &cOrder{}

type cOrder struct{}

func (c *cOrder) Create(r *ghttp.Request) {
	var req v1.OrderCreateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	res, err := service.Order().Create(r.Context(), userId, &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cOrder) List(r *ghttp.Request) {
	var req v1.OrderListReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	res, err := service.Order().List(r.Context(), userId, &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cOrder) Detail(r *ghttp.Request) {
	id := r.Get("id").Int64()
	userId := r.GetCtxVar("userId").Int64()
	res, err := service.Order().Detail(r.Context(), userId, id)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cOrder) Cancel(r *ghttp.Request) {
	id := r.Get("id").Int64()
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Order().Cancel(r.Context(), userId, id); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cOrder) Refund(r *ghttp.Request) {
	var req v1.OrderRefundReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Order().Refund(r.Context(), userId, &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

var Payment = &cPayment{}

type cPayment struct{}

func (c *cPayment) SimulatePay(r *ghttp.Request) {
	var req v1.PaymentSimulateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	res, err := service.Payment().SimulatePay(r.Context(), userId, req.OrderId)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}
