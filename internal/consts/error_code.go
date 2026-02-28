package consts

import "github.com/gogf/gf/v2/errors/gcode"

// 业务错误码
var (
	CodeOK              = gcode.New(0, "success", nil)
	CodeParamError      = gcode.New(1001, "参数错误", nil)
	CodeUnauthorized    = gcode.New(1002, "未登录或登录已过期", nil)
	CodeForbidden       = gcode.New(1003, "无权限", nil)
	CodeNotFound        = gcode.New(1004, "资源不存在", nil)
	CodeDuplicate       = gcode.New(1005, "重复操作", nil)
	CodeStockInsufficient = gcode.New(2001, "库存不足", nil)
	CodeOrderCreateFail = gcode.New(2002, "创建订单失败", nil)
	CodePaymentFail     = gcode.New(2003, "支付失败", nil)
	CodeCouponInvalid   = gcode.New(2004, "优惠券不可用", nil)
	CodeCouponClaimed   = gcode.New(2005, "优惠券已领取", nil)
	CodeRateLimited     = gcode.New(3001, "请求过于频繁", nil)
	CodeWechatAuthFail  = gcode.New(4001, "微信授权失败", nil)
	CodeInternalError   = gcode.New(5000, "服务器内部错误", nil)
)
