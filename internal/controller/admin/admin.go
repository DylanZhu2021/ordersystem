package admin

import (
	"github.com/gogf/gf/v2/net/ghttp"
	adminApi "ordersystem/api/v1/admin"
	"ordersystem/internal/service"
	"ordersystem/utility"
)

var Auth = &cAdminAuth{}

type cAdminAuth struct{}

func (c *cAdminAuth) Login(r *ghttp.Request) {
	var req adminApi.AdminLoginReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	res, err := service.AdminAuth().Login(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cAdminAuth) Info(r *ghttp.Request) {
	adminId := r.GetCtxVar("adminId").Int64()
	res, err := service.AdminAuth().GetInfo(r.Context(), adminId)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

var Product = &cAdminProduct{}

type cAdminProduct struct{}

func (c *cAdminProduct) Create(r *ghttp.Request) {
	var req adminApi.ProductCreateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	id, err := service.AdminProduct().Create(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, adminApi.ProductCreateRes{Id: id})
}

func (c *cAdminProduct) Update(r *ghttp.Request) {
	var req adminApi.ProductUpdateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	if err := service.AdminProduct().Update(r.Context(), &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cAdminProduct) Delete(r *ghttp.Request) {
	id := r.Get("id").Int64()
	if err := service.AdminProduct().Delete(r.Context(), id); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cAdminProduct) UpdateStatus(r *ghttp.Request) {
	var req adminApi.ProductStatusReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	if err := service.AdminProduct().UpdateStatus(r.Context(), req.Id, req.Status); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

var Order = &cAdminOrder{}

type cAdminOrder struct{}

func (c *cAdminOrder) List(r *ghttp.Request) {
	var req adminApi.OrderListReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	res, err := service.AdminOrder().List(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cAdminOrder) UpdateStatus(r *ghttp.Request) {
	var req adminApi.OrderStatusUpdateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	if err := service.AdminOrder().UpdateStatus(r.Context(), &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cAdminOrder) Refund(r *ghttp.Request) {
	var req adminApi.OrderRefundReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	if err := service.AdminOrder().Refund(r.Context(), &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

var Coupon = &cAdminCoupon{}

type cAdminCoupon struct{}

func (c *cAdminCoupon) Create(r *ghttp.Request) {
	var req adminApi.CouponCreateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	id, err := service.AdminCoupon().Create(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, adminApi.CouponCreateRes{Id: id})
}

func (c *cAdminCoupon) Update(r *ghttp.Request) {
	var req adminApi.CouponUpdateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	if err := service.AdminCoupon().Update(r.Context(), &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cAdminCoupon) Delete(r *ghttp.Request) {
	id := r.Get("id").Int64()
	if err := service.AdminCoupon().Delete(r.Context(), id); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

var Category = &cAdminCategory{}

type cAdminCategory struct{}

func (c *cAdminCategory) Create(r *ghttp.Request) {
	var req adminApi.CategoryCreateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	id, err := service.AdminCategory().Create(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, adminApi.CategoryCreateRes{Id: id})
}

func (c *cAdminCategory) Update(r *ghttp.Request) {
	var req adminApi.CategoryUpdateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	if err := service.AdminCategory().Update(r.Context(), &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cAdminCategory) Delete(r *ghttp.Request) {
	id := r.Get("id").Int()
	if err := service.AdminCategory().Delete(r.Context(), id); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

var Stats = &cAdminStats{}

type cAdminStats struct{}

func (c *cAdminStats) Dashboard(r *ghttp.Request) {
	res, err := service.AdminStats().Dashboard(r.Context())
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cAdminStats) SalesStats(r *ghttp.Request) {
	var req adminApi.SalesStatsReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	res, err := service.AdminStats().SalesStats(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cAdminStats) HotProducts(r *ghttp.Request) {
	var req adminApi.HotProductsReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	res, err := service.AdminStats().HotProducts(r.Context(), req.Limit)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

var User = &cAdminUser{}

type cAdminUser struct{}

func (c *cAdminUser) List(r *ghttp.Request) {
	var req adminApi.UserListReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	res, err := service.AdminUser().List(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}
