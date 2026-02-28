package service

import (
	"context"

	adminApi "ordersystem/api/v1/admin"
)

type IAdminAuth interface {
	Login(ctx context.Context, req *adminApi.AdminLoginReq) (*adminApi.AdminLoginRes, error)
	GetInfo(ctx context.Context, adminId int64) (*adminApi.AdminInfoRes, error)
}

type IAdminProduct interface {
	Create(ctx context.Context, req *adminApi.ProductCreateReq) (int64, error)
	Update(ctx context.Context, req *adminApi.ProductUpdateReq) error
	Delete(ctx context.Context, id int64) error
	UpdateStatus(ctx context.Context, id int64, status int) error
}

type IAdminOrder interface {
	List(ctx context.Context, req *adminApi.OrderListReq) (*adminApi.OrderListRes, error)
	UpdateStatus(ctx context.Context, req *adminApi.OrderStatusUpdateReq) error
	Refund(ctx context.Context, req *adminApi.OrderRefundReq) error
}

type IAdminCoupon interface {
	Create(ctx context.Context, req *adminApi.CouponCreateReq) (int64, error)
	Update(ctx context.Context, req *adminApi.CouponUpdateReq) error
	Delete(ctx context.Context, id int64) error
}

type IAdminCategory interface {
	Create(ctx context.Context, req *adminApi.CategoryCreateReq) (int, error)
	Update(ctx context.Context, req *adminApi.CategoryUpdateReq) error
	Delete(ctx context.Context, id int) error
}

type IAdminStats interface {
	Dashboard(ctx context.Context) (*adminApi.DashboardRes, error)
	SalesStats(ctx context.Context, req *adminApi.SalesStatsReq) (*adminApi.SalesStatsRes, error)
	HotProducts(ctx context.Context, limit int) (*adminApi.HotProductsRes, error)
}

type IAdminUser interface {
	List(ctx context.Context, req *adminApi.UserListReq) (*adminApi.UserListRes, error)
}

var (
	localAdminAuth     IAdminAuth
	localAdminProduct  IAdminProduct
	localAdminOrder    IAdminOrder
	localAdminCoupon   IAdminCoupon
	localAdminCategory IAdminCategory
	localAdminStats    IAdminStats
	localAdminUser     IAdminUser
)

func AdminAuth() IAdminAuth         { return localAdminAuth }
func AdminProduct() IAdminProduct   { return localAdminProduct }
func AdminOrder() IAdminOrder       { return localAdminOrder }
func AdminCoupon() IAdminCoupon     { return localAdminCoupon }
func AdminCategory() IAdminCategory { return localAdminCategory }
func AdminStats() IAdminStats       { return localAdminStats }
func AdminUser() IAdminUser         { return localAdminUser }

func SetAdminAuth(s IAdminAuth)         { localAdminAuth = s }
func SetAdminProduct(s IAdminProduct)   { localAdminProduct = s }
func SetAdminOrder(s IAdminOrder)       { localAdminOrder = s }
func SetAdminCoupon(s IAdminCoupon)     { localAdminCoupon = s }
func SetAdminCategory(s IAdminCategory) { localAdminCategory = s }
func SetAdminStats(s IAdminStats)       { localAdminStats = s }
func SetAdminUser(s IAdminUser)         { localAdminUser = s }
