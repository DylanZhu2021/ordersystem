package v1

// ==================== 用户认证 ====================

type WechatLoginReq struct {
	Code string `json:"code" v:"required#请提供微信授权码"`
}

type WechatLoginRes struct {
	Token  string `json:"token"`
	UserId int64  `json:"userId"`
	IsNew  bool   `json:"isNew"`
}

type UserInfoReq struct{}

type UserInfoRes struct {
	Id          int64  `json:"id"`
	Nickname    string `json:"nickname"`
	AvatarUrl   string `json:"avatarUrl"`
	Phone       string `json:"phone"`
	Points      int    `json:"points"`
	MemberLevel int    `json:"memberLevel"`
}

type UserUpdateReq struct {
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
	Phone     string `json:"phone"`
	Gender    int    `json:"gender"`
}

type UserUpdateRes struct{}

// ==================== 地址 ====================

type AddressListReq struct{}

type AddressListRes struct {
	List []AddressItem `json:"list"`
}

type AddressItem struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"district"`
	Detail    string `json:"detail"`
	IsDefault int    `json:"isDefault"`
}

type AddressCreateReq struct {
	Name      string `json:"name" v:"required"`
	Phone     string `json:"phone" v:"required|phone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"district"`
	Detail    string `json:"detail" v:"required"`
	IsDefault int    `json:"isDefault"`
}

type AddressCreateRes struct {
	Id int64 `json:"id"`
}

type AddressUpdateReq struct {
	Id        int64  `json:"id" v:"required"`
	Name      string `json:"name" v:"required"`
	Phone     string `json:"phone" v:"required"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"district"`
	Detail    string `json:"detail" v:"required"`
	IsDefault int    `json:"isDefault"`
}

type AddressUpdateRes struct{}

type AddressDeleteReq struct {
	Id int64 `json:"id" v:"required"`
}

type AddressDeleteRes struct{}
