package consts

// 订单状态
const (
	OrderStatusPending    = 0 // 待支付
	OrderStatusPaid       = 1 // 已支付
	OrderStatusPreparing  = 2 // 制作中
	OrderStatusReady      = 3 // 待取餐
	OrderStatusCompleted  = 4 // 已完成
	OrderStatusCancelled  = 5 // 已取消
	OrderStatusRefunded   = 6 // 已退款
)

// 配送方式
const (
	DeliveryDineIn   = 1 // 堂食
	DeliveryTakeout  = 2 // 外卖
	DeliverySelfPick = 3 // 自提
)

// 优惠券类型
const (
	CouponTypeFixed    = 1 // 满减券
	CouponTypePercent  = 2 // 折扣券
	CouponTypeFree     = 3 // 无门槛券
)

// 优惠券状态
const (
	UserCouponUnused  = 0
	UserCouponUsed    = 1
	UserCouponExpired = 2
)

// 支付状态
const (
	PaymentPending  = 0
	PaymentSuccess  = 1
	PaymentFailed   = 2
	PaymentRefunded = 3
)

// 会员等级
const (
	MemberNormal   = 0
	MemberSilver   = 1
	MemberGold     = 2
	MemberPlatinum = 3
)

// 积分类型
const (
	PointsTypeOrderEarn  = 1
	PointsTypeReviewEarn = 2
	PointsTypeRedeem     = 3
	PointsTypeAdminAdjust = 4
)

// 通用状态
const (
	StatusDisabled = 0
	StatusEnabled  = 1
)
