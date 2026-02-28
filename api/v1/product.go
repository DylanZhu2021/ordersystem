package v1

// ==================== 分类 ====================

type CategoryListReq struct{}

type CategoryListRes struct {
	List []CategoryItem `json:"list"`
}

type CategoryItem struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	IconUrl  string `json:"iconUrl"`
}

// ==================== 商品 ====================

type ProductListReq struct {
	CategoryId int    `json:"categoryId"`
	Keyword    string `json:"keyword"`
	Sort       string `json:"sort"` // sales, price_asc, price_desc
	Page       int    `json:"page" d:"1"`
	PageSize   int    `json:"pageSize" d:"20"`
}

type ProductListRes struct {
	List     []ProductItem `json:"list"`
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
}

type ProductItem struct {
	Id            int64   `json:"id"`
	CategoryId    int     `json:"categoryId"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"originalPrice"`
	ImageUrl      string  `json:"imageUrl"`
	Stock         int     `json:"stock"`
	Sales         int     `json:"sales"`
	IsHot         int     `json:"isHot"`
	IsRecommend   int     `json:"isRecommend"`
}

type ProductDetailReq struct {
	Id int64 `json:"id" v:"required"`
}

type ProductDetailRes struct {
	Id            int64         `json:"id"`
	CategoryId    int           `json:"categoryId"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Price         float64       `json:"price"`
	OriginalPrice float64       `json:"originalPrice"`
	ImageUrl      string        `json:"imageUrl"`
	Images        []string      `json:"images"`
	Stock         int           `json:"stock"`
	Sales         int           `json:"sales"`
	IsHot         int           `json:"isHot"`
	IsRecommend   int           `json:"isRecommend"`
	Specs         []ProductSpecItem `json:"specs"`
}

type ProductSpecItem struct {
	Id        int64   `json:"id"`
	SpecName  string  `json:"specName"`
	SpecValue string  `json:"specValue"`
	PriceDiff float64 `json:"priceDiff"`
	Stock     int     `json:"stock"`
}

type ProductSearchReq struct {
	Keyword  string `json:"keyword" v:"required"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20"`
}

type ProductSearchRes = ProductListRes
