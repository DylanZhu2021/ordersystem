package v1

// ==================== 购物车 ====================

type CartAddReq struct {
	ProductId int64 `json:"productId" v:"required"`
	SpecId    int64 `json:"specId"`
	Quantity  int   `json:"quantity" v:"required|min:1"`
}

type CartAddRes struct{}

type CartListReq struct{}

type CartListRes struct {
	List       []CartItem `json:"list"`
	TotalPrice float64    `json:"totalPrice"`
}

type CartItem struct {
	ProductId    int64   `json:"productId"`
	ProductName  string  `json:"productName"`
	ProductImage string  `json:"productImage"`
	SpecId       int64   `json:"specId"`
	SpecInfo     string  `json:"specInfo"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	Stock        int     `json:"stock"`
}

type CartUpdateReq struct {
	ProductId int64 `json:"productId" v:"required"`
	SpecId    int64 `json:"specId"`
	Quantity  int   `json:"quantity" v:"required|min:0"`
}

type CartUpdateRes struct{}

type CartDeleteReq struct {
	ProductId int64 `json:"productId" v:"required"`
	SpecId    int64 `json:"specId"`
}

type CartDeleteRes struct{}

type CartClearReq struct{}

type CartClearRes struct{}
