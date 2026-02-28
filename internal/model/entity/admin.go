package entity

import "github.com/gogf/gf/v2/os/gtime"

type AdminUser struct {
	Id           int64       `json:"id"`
	Username     string      `json:"username"`
	PasswordHash string      `json:"-"`
	RealName     string      `json:"realName"`
	Phone        string      `json:"phone"`
	RoleId       int         `json:"roleId"`
	Status       int         `json:"status"`
	LastLoginAt  *gtime.Time `json:"lastLoginAt"`
	CreatedAt    *gtime.Time `json:"createdAt"`
	UpdatedAt    *gtime.Time `json:"updatedAt"`
}

type Role struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Permissions string      `json:"permissions"`
	CreatedAt   *gtime.Time `json:"createdAt"`
	UpdatedAt   *gtime.Time `json:"updatedAt"`
}
