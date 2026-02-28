package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
	"ordersystem/utility"
)

var Coupon = &cCoupon{}

type cCoupon struct{}

func (c *cCoupon) List(r *ghttp.Request) {
	var req v1.CouponListReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId", 0).Int64()
	res, err := service.Coupon().List(r.Context(), userId, &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cCoupon) Claim(r *ghttp.Request) {
	var req v1.CouponClaimReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Coupon().Claim(r.Context(), userId, req.CouponId); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cCoupon) MyCoupons(r *ghttp.Request) {
	var req v1.MyCouponListReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	res, err := service.Coupon().MyCoupons(r.Context(), userId, &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

var Review = &cReview{}

type cReview struct{}

func (c *cReview) Create(r *ghttp.Request) {
	var req v1.ReviewCreateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Review().Create(r.Context(), userId, &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cReview) List(r *ghttp.Request) {
	var req v1.ReviewListReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	res, err := service.Review().List(r.Context(), &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

var Favorite = &cFavorite{}

type cFavorite struct{}

func (c *cFavorite) Add(r *ghttp.Request) {
	var req v1.FavoriteAddReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Favorite().Add(r.Context(), userId, req.ProductId); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cFavorite) Delete(r *ghttp.Request) {
	var req v1.FavoriteDeleteReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Favorite().Delete(r.Context(), userId, req.ProductId); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cFavorite) List(r *ghttp.Request) {
	var req v1.FavoriteListReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	res, err := service.Favorite().List(r.Context(), userId, &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

var Address = &cAddress{}

type cAddress struct{}

func (c *cAddress) List(r *ghttp.Request) {
	userId := r.GetCtxVar("userId").Int64()
	res, err := service.Address().List(r.Context(), userId)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, res)
}

func (c *cAddress) Create(r *ghttp.Request) {
	var req v1.AddressCreateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	id, err := service.Address().Create(r.Context(), userId, &req)
	if err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r, v1.AddressCreateRes{Id: id})
}

func (c *cAddress) Update(r *ghttp.Request) {
	var req v1.AddressUpdateReq
	if err := r.Parse(&req); err != nil {
		utility.Error(r, 1001, err.Error())
		return
	}
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Address().Update(r.Context(), userId, &req); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}

func (c *cAddress) Delete(r *ghttp.Request) {
	id := r.Get("id").Int64()
	userId := r.GetCtxVar("userId").Int64()
	if err := service.Address().Delete(r.Context(), userId, id); err != nil {
		utility.ErrorMsg(r, err.Error())
		return
	}
	utility.Success(r)
}
