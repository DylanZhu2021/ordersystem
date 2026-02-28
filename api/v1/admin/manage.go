package admin

// ==================== 订单管理 ====================

type OrderListReq struct {
	Status   int    `json:"status" d:"-1"`
	Keyword  string `json:"keyword"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20"`
}

type OrderListRes struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type OrderStatusUpdateReq struct {
	Id     int64 `json:"id" v:"required"`
	Status int   `json:"status" v:"required"`
}

type OrderStatusUpdateRes struct{}

type OrderRefundReq struct {
	Id     int64  `json:"id" v:"required"`
	Reason string `json:"reason"`
}

type OrderRefundRes struct{}

type OrderExportReq struct {
	Status    int    `json:"status" d:"-1"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// ==================== 优惠券管理 ====================

type CouponCreateReq struct {
	Name          string  `json:"name" v:"required"`
	Type          int     `json:"type" v:"required|in:1,2,3"`
	DiscountValue float64 `json:"discountValue" v:"required|min:0.01"`
	MinAmount     float64 `json:"minAmount"`
	TotalCount    int     `json:"totalCount" v:"required|min:1"`
	PerUserLimit  int     `json:"perUserLimit" d:"1"`
	StartTime     string  `json:"startTime" v:"required"`
	EndTime       string  `json:"endTime" v:"required"`
}

type CouponCreateRes struct {
	Id int64 `json:"id"`
}

type CouponUpdateReq struct {
	Id            int64   `json:"id" v:"required"`
	Name          string  `json:"name"`
	TotalCount    int     `json:"totalCount"`
	PerUserLimit  int     `json:"perUserLimit"`
	StartTime     string  `json:"startTime"`
	EndTime       string  `json:"endTime"`
	Status        int     `json:"status"`
}

type CouponUpdateRes struct{}

type CouponDeleteReq struct {
	Id int64 `json:"id" v:"required"`
}

type CouponDeleteRes struct{}

// ==================== 数据统计 ====================

type DashboardReq struct{}

type DashboardRes struct {
	TodaySales     float64 `json:"todaySales"`
	TodayOrders    int     `json:"todayOrders"`
	TotalUsers     int     `json:"totalUsers"`
	AvgOrderAmount float64 `json:"avgOrderAmount"`
}

type SalesStatsReq struct {
	StartDate string `json:"startDate" v:"required"`
	EndDate   string `json:"endDate" v:"required"`
}

type SalesStatsRes struct {
	List []SalesStatItem `json:"list"`
}

type SalesStatItem struct {
	Date       string  `json:"date"`
	Sales      float64 `json:"sales"`
	OrderCount int     `json:"orderCount"`
}

type HotProductsReq struct {
	Limit int `json:"limit" d:"10"`
}

type HotProductsRes struct {
	List []HotProductItem `json:"list"`
}

type HotProductItem struct {
	ProductId   int64   `json:"productId"`
	ProductName string  `json:"productName"`
	Sales       int     `json:"sales"`
	Revenue     float64 `json:"revenue"`
}

// ==================== 分类管理 ====================

type CategoryCreateReq struct {
	Name      string `json:"name" v:"required"`
	IconUrl   string `json:"iconUrl"`
	SortOrder int    `json:"sortOrder"`
}

type CategoryCreateRes struct {
	Id int `json:"id"`
}

type CategoryUpdateReq struct {
	Id        int    `json:"id" v:"required"`
	Name      string `json:"name"`
	IconUrl   string `json:"iconUrl"`
	SortOrder int    `json:"sortOrder"`
	Status    int    `json:"status"`
}

type CategoryUpdateRes struct{}

type CategoryDeleteReq struct {
	Id int `json:"id" v:"required"`
}

type CategoryDeleteRes struct{}

// ==================== 用户管理 ====================

type UserListReq struct {
	Keyword  string `json:"keyword"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20"`
}

type UserListRes struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
