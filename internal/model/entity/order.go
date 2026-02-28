package entity

import "github.com/gogf/gf/v2/os/gtime"

type Order struct {
	Id             int64       `json:"id"`
	OrderNo        string      `json:"orderNo"`
	UserId         int64       `json:"userId"`
	TotalAmount    float64     `json:"totalAmount"`
	DiscountAmount float64     `json:"discountAmount"`
	PayAmount      float64     `json:"payAmount"`
	CouponId       int64       `json:"couponId"`
	DeliveryType   int         `json:"deliveryType"`
	TableNo        string      `json:"tableNo"`
	AddressId      int64       `json:"addressId"`
	ContactName    string      `json:"contactName"`
	ContactPhone   string      `json:"contactPhone"`
	Remark         string      `json:"remark"`
	Status         int         `json:"status"`
	IdempotencyKey string      `json:"idempotencyKey"`
	PaidAt         *gtime.Time `json:"paidAt"`
	CompletedAt    *gtime.Time `json:"completedAt"`
	CancelledAt    *gtime.Time `json:"cancelledAt"`
	CreatedAt      *gtime.Time `json:"createdAt"`
	UpdatedAt      *gtime.Time `json:"updatedAt"`
}

type OrderItem struct {
	Id           int64       `json:"id"`
	OrderId      int64       `json:"orderId"`
	ProductId    int64       `json:"productId"`
	ProductName  string      `json:"productName"`
	ProductImage string      `json:"productImage"`
	SpecId       int64       `json:"specId"`
	SpecInfo     string      `json:"specInfo"`
	Price        float64     `json:"price"`
	Quantity     int         `json:"quantity"`
	TotalAmount  float64     `json:"totalAmount"`
	CreatedAt    *gtime.Time `json:"createdAt"`
}

type PaymentLog struct {
	Id            int64       `json:"id"`
	OrderId       int64       `json:"orderId"`
	TransactionNo string      `json:"transactionNo"`
	Amount        float64     `json:"amount"`
	Method        string      `json:"method"`
	Status        int         `json:"status"`
	NotifyData    string      `json:"notifyData"`
	CreatedAt     *gtime.Time `json:"createdAt"`
	UpdatedAt     *gtime.Time `json:"updatedAt"`
}
