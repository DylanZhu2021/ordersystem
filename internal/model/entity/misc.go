package entity

import "github.com/gogf/gf/v2/os/gtime"

type Coupon struct {
	Id            int64       `json:"id"`
	Name          string      `json:"name"`
	Type          int         `json:"type"`
	DiscountValue float64     `json:"discountValue"`
	MinAmount     float64     `json:"minAmount"`
	TotalCount    int         `json:"totalCount"`
	ClaimedCount  int         `json:"claimedCount"`
	PerUserLimit  int         `json:"perUserLimit"`
	StartTime     *gtime.Time `json:"startTime"`
	EndTime       *gtime.Time `json:"endTime"`
	Status        int         `json:"status"`
	CreatedAt     *gtime.Time `json:"createdAt"`
	UpdatedAt     *gtime.Time `json:"updatedAt"`
}

type UserCoupon struct {
	Id        int64       `json:"id"`
	UserId    int64       `json:"userId"`
	CouponId  int64       `json:"couponId"`
	OrderId   int64       `json:"orderId"`
	Status    int         `json:"status"`
	ClaimedAt *gtime.Time `json:"claimedAt"`
	UsedAt    *gtime.Time `json:"usedAt"`
}

type Review struct {
	Id        int64       `json:"id"`
	UserId    int64       `json:"userId"`
	ProductId int64       `json:"productId"`
	OrderId   int64       `json:"orderId"`
	Rating    int         `json:"rating"`
	Content   string      `json:"content"`
	Images    string      `json:"images"`
	Reply     string      `json:"reply"`
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

type Favorite struct {
	Id        int64       `json:"id"`
	UserId    int64       `json:"userId"`
	ProductId int64       `json:"productId"`
	CreatedAt *gtime.Time `json:"createdAt"`
}

type PointsLog struct {
	Id        int64       `json:"id"`
	UserId    int64       `json:"userId"`
	Change    int         `json:"change"`
	Balance   int         `json:"balance"`
	Type      int         `json:"type"`
	RefId     int64       `json:"refId"`
	Remark    string      `json:"remark"`
	CreatedAt *gtime.Time `json:"createdAt"`
}

type OperationLog struct {
	Id        int64       `json:"id"`
	AdminId   int64       `json:"adminId"`
	Module    string      `json:"module"`
	Action    string      `json:"action"`
	Content   string      `json:"content"`
	Ip        string      `json:"ip"`
	UserAgent string      `json:"userAgent"`
	CreatedAt *gtime.Time `json:"createdAt"`
}
