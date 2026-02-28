package v1

// ==================== 优惠券 ====================

type CouponListReq struct {
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"20"`
}

type CouponListRes struct {
	List     []CouponItem `json:"list"`
	Total    int          `json:"total"`
	Page     int          `json:"page"`
	PageSize int          `json:"pageSize"`
}

type CouponItem struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	Type          int     `json:"type"`
	DiscountValue float64 `json:"discountValue"`
	MinAmount     float64 `json:"minAmount"`
	StartTime     string  `json:"startTime"`
	EndTime       string  `json:"endTime"`
	Claimed       bool    `json:"claimed"`
}

type CouponClaimReq struct {
	CouponId int64 `json:"couponId" v:"required"`
}

type CouponClaimRes struct{}

type MyCouponListReq struct {
	Status   int `json:"status" d:"-1"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"20"`
}

type MyCouponListRes struct {
	List     []MyCouponItem `json:"list"`
	Total    int            `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}

type MyCouponItem struct {
	Id            int64   `json:"id"`
	CouponId      int64   `json:"couponId"`
	Name          string  `json:"name"`
	Type          int     `json:"type"`
	DiscountValue float64 `json:"discountValue"`
	MinAmount     float64 `json:"minAmount"`
	Status        int     `json:"status"`
	EndTime       string  `json:"endTime"`
}

// ==================== 评价 ====================

type ReviewCreateReq struct {
	OrderId   int64    `json:"orderId" v:"required"`
	ProductId int64    `json:"productId" v:"required"`
	Rating    int      `json:"rating" v:"required|between:1,5"`
	Content   string   `json:"content"`
	Images    []string `json:"images"`
}

type ReviewCreateRes struct{}

type ReviewListReq struct {
	ProductId int64 `json:"productId" v:"required"`
	Page      int   `json:"page" d:"1"`
	PageSize  int   `json:"pageSize" d:"20"`
}

type ReviewListRes struct {
	List     []ReviewItem `json:"list"`
	Total    int          `json:"total"`
	Page     int          `json:"page"`
	PageSize int          `json:"pageSize"`
}

type ReviewItem struct {
	Id        int64    `json:"id"`
	Nickname  string   `json:"nickname"`
	AvatarUrl string   `json:"avatarUrl"`
	Rating    int      `json:"rating"`
	Content   string   `json:"content"`
	Images    []string `json:"images"`
	Reply     string   `json:"reply"`
	CreatedAt string   `json:"createdAt"`
}

// ==================== 收藏 ====================

type FavoriteAddReq struct {
	ProductId int64 `json:"productId" v:"required"`
}

type FavoriteAddRes struct{}

type FavoriteDeleteReq struct {
	ProductId int64 `json:"productId" v:"required"`
}

type FavoriteDeleteRes struct{}

type FavoriteListReq struct {
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"20"`
}

type FavoriteListRes struct {
	List     []ProductItem `json:"list"`
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
}
