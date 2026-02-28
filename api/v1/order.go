package v1

// ==================== 订单 ====================

type OrderCreateReq struct {
	DeliveryType   int    `json:"deliveryType" v:"required|in:1,2,3"`
	AddressId      int64  `json:"addressId"`
	TableNo        string `json:"tableNo"`
	CouponId       int64  `json:"couponId"`
	Remark         string `json:"remark"`
	IdempotencyKey string `json:"idempotencyKey" v:"required"`
}

type OrderCreateRes struct {
	OrderId     int64   `json:"orderId"`
	OrderNo     string  `json:"orderNo"`
	TotalAmount float64 `json:"totalAmount"`
	PayAmount   float64 `json:"payAmount"`
}

type OrderListReq struct {
	Status   int `json:"status" d:"-1"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"20"`
}

type OrderListRes struct {
	List     []OrderListItem `json:"list"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}

type OrderListItem struct {
	Id           int64           `json:"id"`
	OrderNo      string          `json:"orderNo"`
	TotalAmount  float64         `json:"totalAmount"`
	PayAmount    float64         `json:"payAmount"`
	Status       int             `json:"status"`
	DeliveryType int             `json:"deliveryType"`
	ItemCount    int             `json:"itemCount"`
	Items        []OrderItemBrief `json:"items"`
	CreatedAt    string          `json:"createdAt"`
}

type OrderItemBrief struct {
	ProductName  string  `json:"productName"`
	ProductImage string  `json:"productImage"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
}

type OrderDetailReq struct {
	Id int64 `json:"id" v:"required"`
}

type OrderDetailRes struct {
	Id             int64             `json:"id"`
	OrderNo        string            `json:"orderNo"`
	TotalAmount    float64           `json:"totalAmount"`
	DiscountAmount float64           `json:"discountAmount"`
	PayAmount      float64           `json:"payAmount"`
	Status         int               `json:"status"`
	DeliveryType   int               `json:"deliveryType"`
	TableNo        string            `json:"tableNo"`
	ContactName    string            `json:"contactName"`
	ContactPhone   string            `json:"contactPhone"`
	Remark         string            `json:"remark"`
	Items          []OrderDetailItem `json:"items"`
	CreatedAt      string            `json:"createdAt"`
	PaidAt         string            `json:"paidAt"`
}

type OrderDetailItem struct {
	ProductId    int64   `json:"productId"`
	ProductName  string  `json:"productName"`
	ProductImage string  `json:"productImage"`
	SpecInfo     string  `json:"specInfo"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	TotalAmount  float64 `json:"totalAmount"`
}

type OrderCancelReq struct {
	Id int64 `json:"id" v:"required"`
}

type OrderCancelRes struct{}

type OrderRefundReq struct {
	Id     int64  `json:"id" v:"required"`
	Reason string `json:"reason"`
}

type OrderRefundRes struct{}

// ==================== 支付 ====================

type PaymentSimulateReq struct {
	OrderId int64 `json:"orderId" v:"required"`
}

type PaymentSimulateRes struct {
	TransactionNo string `json:"transactionNo"`
	Status        int    `json:"status"`
}
