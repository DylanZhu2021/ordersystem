package service

import (
	"context"

	v1 "ordersystem/api/v1"
)

type IAuth interface {
	WechatLogin(ctx context.Context, req *v1.WechatLoginReq) (*v1.WechatLoginRes, error)
	GetUserInfo(ctx context.Context, userId int64) (*v1.UserInfoRes, error)
	UpdateUser(ctx context.Context, userId int64, req *v1.UserUpdateReq) error
}

type IAddress interface {
	List(ctx context.Context, userId int64) (*v1.AddressListRes, error)
	Create(ctx context.Context, userId int64, req *v1.AddressCreateReq) (int64, error)
	Update(ctx context.Context, userId int64, req *v1.AddressUpdateReq) error
	Delete(ctx context.Context, userId int64, id int64) error
}

type ICategory interface {
	List(ctx context.Context) (*v1.CategoryListRes, error)
}

type IProduct interface {
	List(ctx context.Context, req *v1.ProductListReq) (*v1.ProductListRes, error)
	Detail(ctx context.Context, id int64) (*v1.ProductDetailRes, error)
	Search(ctx context.Context, req *v1.ProductSearchReq) (*v1.ProductListRes, error)
}

type ICart interface {
	Add(ctx context.Context, userId int64, req *v1.CartAddReq) error
	List(ctx context.Context, userId int64) (*v1.CartListRes, error)
	Update(ctx context.Context, userId int64, req *v1.CartUpdateReq) error
	Delete(ctx context.Context, userId int64, req *v1.CartDeleteReq) error
	Clear(ctx context.Context, userId int64) error
}

type IOrder interface {
	Create(ctx context.Context, userId int64, req *v1.OrderCreateReq) (*v1.OrderCreateRes, error)
	List(ctx context.Context, userId int64, req *v1.OrderListReq) (*v1.OrderListRes, error)
	Detail(ctx context.Context, userId int64, id int64) (*v1.OrderDetailRes, error)
	Cancel(ctx context.Context, userId int64, id int64) error
	Refund(ctx context.Context, userId int64, req *v1.OrderRefundReq) error
}

type IPayment interface {
	SimulatePay(ctx context.Context, userId int64, orderId int64) (*v1.PaymentSimulateRes, error)
}

type ICoupon interface {
	List(ctx context.Context, userId int64, req *v1.CouponListReq) (*v1.CouponListRes, error)
	Claim(ctx context.Context, userId int64, couponId int64) error
	MyCoupons(ctx context.Context, userId int64, req *v1.MyCouponListReq) (*v1.MyCouponListRes, error)
}

type IReview interface {
	Create(ctx context.Context, userId int64, req *v1.ReviewCreateReq) error
	List(ctx context.Context, req *v1.ReviewListReq) (*v1.ReviewListRes, error)
}

type IFavorite interface {
	Add(ctx context.Context, userId int64, productId int64) error
	Delete(ctx context.Context, userId int64, productId int64) error
	List(ctx context.Context, userId int64, req *v1.FavoriteListReq) (*v1.FavoriteListRes, error)
}

var (
	localAuth     IAuth
	localAddress  IAddress
	localCategory ICategory
	localProduct  IProduct
	localCart     ICart
	localOrder    IOrder
	localPayment  IPayment
	localCoupon   ICoupon
	localReview   IReview
	localFavorite IFavorite
)

func Auth() IAuth         { return localAuth }
func Address() IAddress   { return localAddress }
func Category() ICategory { return localCategory }
func Product() IProduct   { return localProduct }
func Cart() ICart         { return localCart }
func Order() IOrder       { return localOrder }
func Payment() IPayment   { return localPayment }
func Coupon() ICoupon     { return localCoupon }
func Review() IReview     { return localReview }
func Favorite() IFavorite { return localFavorite }

func SetAuth(s IAuth)         { localAuth = s }
func SetAddress(s IAddress)   { localAddress = s }
func SetCategory(s ICategory) { localCategory = s }
func SetProduct(s IProduct)   { localProduct = s }
func SetCart(s ICart)         { localCart = s }
func SetOrder(s IOrder)       { localOrder = s }
func SetPayment(s IPayment)   { localPayment = s }
func SetCoupon(s ICoupon)     { localCoupon = s }
func SetReview(s IReview)     { localReview = s }
func SetFavorite(s IFavorite) { localFavorite = s }
