package entity

import "github.com/gogf/gf/v2/os/gtime"

type User struct {
	Id          int64       `json:"id"`
	Openid      string      `json:"openid"`
	UnionId     string      `json:"unionId"`
	Nickname    string      `json:"nickname"`
	AvatarUrl   string      `json:"avatarUrl"`
	Phone       string      `json:"phone"`
	Gender      int         `json:"gender"`
	Points      int         `json:"points"`
	TotalPoints int         `json:"totalPoints"`
	MemberLevel int         `json:"memberLevel"`
	Status      int         `json:"status"`
	CreatedAt   *gtime.Time `json:"createdAt"`
	UpdatedAt   *gtime.Time `json:"updatedAt"`
}

type Address struct {
	Id        int64       `json:"id"`
	UserId    int64       `json:"userId"`
	Name      string      `json:"name"`
	Phone     string      `json:"phone"`
	Province  string      `json:"province"`
	City      string      `json:"city"`
	District  string      `json:"district"`
	Detail    string      `json:"detail"`
	IsDefault int         `json:"isDefault"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}
