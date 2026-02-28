package payment

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/consts"
	"ordersystem/internal/service"
	"ordersystem/utility"
)

type sPayment struct{}

func init() {
	service.SetPayment(&sPayment{})
}

func (s *sPayment) SimulatePay(ctx context.Context, userId int64, orderId int64) (*v1.PaymentSimulateRes, error) {
	// 查询订单
	order, err := g.DB().Ctx(ctx).Model("order").
		Where("id", orderId).Where("user_id", userId).One()
	if err != nil {
		return nil, err
	}
	if order.IsEmpty() {
		return nil, gerror.New("订单不存在")
	}
	if order["status"].Int() != consts.OrderStatusPending {
		return nil, gerror.New("订单状态不允许支付")
	}

	transactionNo := utility.GeneratePaymentNo()

	// 创建支付记录
	_, err = g.DB().Ctx(ctx).Model("payment_log").Data(g.Map{
		"order_id":       orderId,
		"transaction_no": transactionNo,
		"amount":         order["pay_amount"].Float64(),
		"method":         "wechat_simulated",
		"status":         consts.PaymentSuccess,
	}).Insert()
	if err != nil {
		return nil, err
	}

	// 更新订单状态为已支付
	_, err = g.DB().Exec(ctx,
		`UPDATE "order" SET status = $1, paid_at = $2, updated_at = NOW() WHERE id = $3`,
		consts.OrderStatusPaid, time.Now(), orderId,
	)
	if err != nil {
		return nil, err
	}

	// 增加用户积分
	payAmount := order["pay_amount"].Float64()
	points := int(payAmount) // 1元=1积分
	if points > 0 {
		_, _ = g.DB().Exec(ctx,
			`UPDATE "user" SET points = points + $1, total_points = total_points + $1 WHERE id = $2`,
			points, userId,
		)
		_, _ = g.DB().Ctx(ctx).Model("points_log").Data(g.Map{
			"user_id": userId,
			"change":  points,
			"balance": 0, // 简化处理
			"type":    consts.PointsTypeOrderEarn,
			"ref_id":  orderId,
			"remark":  "订单消费积分",
		}).Insert()
	}

	return &v1.PaymentSimulateRes{
		TransactionNo: transactionNo,
		Status:        consts.PaymentSuccess,
	}, nil
}
