package admin

// ==================== 管理员认证 ====================

type AdminLoginReq struct {
	Username string `json:"username" v:"required"`
	Password string `json:"password" v:"required"`
}

type AdminLoginRes struct {
	Token    string `json:"token"`
	AdminId  int64  `json:"adminId"`
	RealName string `json:"realName"`
	RoleName string `json:"roleName"`
}

type AdminInfoReq struct{}

type AdminInfoRes struct {
	Id          int64    `json:"id"`
	Username    string   `json:"username"`
	RealName    string   `json:"realName"`
	RoleId      int      `json:"roleId"`
	RoleName    string   `json:"roleName"`
	Permissions []string `json:"permissions"`
}

// ==================== 商品管理 ====================

type ProductCreateReq struct {
	CategoryId    int      `json:"categoryId" v:"required"`
	Name          string   `json:"name" v:"required"`
	Description   string   `json:"description"`
	Price         float64  `json:"price" v:"required|min:0.01"`
	OriginalPrice float64  `json:"originalPrice"`
	ImageUrl      string   `json:"imageUrl"`
	Images        []string `json:"images"`
	Stock         int      `json:"stock" v:"required|min:0"`
	IsHot         int      `json:"isHot"`
	IsRecommend   int      `json:"isRecommend"`
	SortOrder     int      `json:"sortOrder"`
}

type ProductCreateRes struct {
	Id int64 `json:"id"`
}

type ProductUpdateReq struct {
	Id            int64    `json:"id" v:"required"`
	CategoryId    int      `json:"categoryId"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Price         float64  `json:"price"`
	OriginalPrice float64  `json:"originalPrice"`
	ImageUrl      string   `json:"imageUrl"`
	Images        []string `json:"images"`
	Stock         int      `json:"stock"`
	IsHot         int      `json:"isHot"`
	IsRecommend   int      `json:"isRecommend"`
	SortOrder     int      `json:"sortOrder"`
}

type ProductUpdateRes struct{}

type ProductDeleteReq struct {
	Id int64 `json:"id" v:"required"`
}

type ProductDeleteRes struct{}

type ProductStatusReq struct {
	Id     int64 `json:"id" v:"required"`
	Status int   `json:"status" v:"required|in:0,1"`
}

type ProductStatusRes struct{}
