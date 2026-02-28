package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
	"ordersystem/utility"
)

var Product = &cProduct{}

type cProduct struct{}

func (c *cProduct) List(r *ghttp.Request) {
	var req v1.ProductListReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	res, err := service.Product().List(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cProduct) Detail(r *ghttp.Request) {
	id := r.Get("id").Int64()
	res, err := service.Product().Detail(r.Context(), id)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cProduct) Search(r *ghttp.Request) {
	var req v1.ProductSearchReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	res, err := service.Product().Search(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

var Category = &cCategory{}

type cCategory struct{}

func (c *cCategory) List(r *ghttp.Request) {
	res, err := service.Category().List(r.Context())
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}
