package entity

import "github.com/gogf/gf/v2/os/gtime"

type Category struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	IconUrl   string      `json:"iconUrl"`
	SortOrder int         `json:"sortOrder"`
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

type Product struct {
	Id            int64       `json:"id"`
	CategoryId    int         `json:"categoryId"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Price         float64     `json:"price"`
	OriginalPrice float64     `json:"originalPrice"`
	ImageUrl      string      `json:"imageUrl"`
	Images        string      `json:"images"`
	Stock         int         `json:"stock"`
	Sales         int         `json:"sales"`
	IsHot         int         `json:"isHot"`
	IsRecommend   int         `json:"isRecommend"`
	Status        int         `json:"status"`
	SortOrder     int         `json:"sortOrder"`
	CreatedAt     *gtime.Time `json:"createdAt"`
	UpdatedAt     *gtime.Time `json:"updatedAt"`
}

type ProductSpec struct {
	Id        int64       `json:"id"`
	ProductId int64       `json:"productId"`
	SpecName  string      `json:"specName"`
	SpecValue string      `json:"specValue"`
	PriceDiff float64     `json:"priceDiff"`
	Stock     int         `json:"stock"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}
