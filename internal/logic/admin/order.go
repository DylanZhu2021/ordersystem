package admin

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	adminApi "ordersystem/api/v1/admin"
	"ordersystem/internal/consts"
	"ordersystem/internal/service"
)

type sAdminOrder struct{}

func init() {
	service.SetAdminOrder(&sAdminOrder{})
}

func (s *sAdminOrder) List(ctx context.Context, req *adminApi.OrderListReq) (*adminApi.OrderListRes, error) {
	m := g.DB().Ctx(ctx).Model("order")
	if req.Status >= 0 {
		m = m.Where("status", req.Status)
	}
	if req.Keyword != "" {
		m = m.Where("order_no LIKE ?", "%"+req.Keyword+"%")
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	records, err := m.OrderDesc("id").Page(req.Page, req.PageSize).All()
	if err != nil {
		return nil, err
	}

	return &adminApi.OrderListRes{
		List:     records,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *sAdminOrder) UpdateStatus(ctx context.Context, req *adminApi.OrderStatusUpdateReq) error {
	data := g.Map{"status": req.Status, "updated_at": time.Now()}
	if req.Status == consts.OrderStatusCompleted {
		data["completed_at"] = time.Now()
	}
	if req.Status == consts.OrderStatusCancelled {
		data["cancelled_at"] = time.Now()
	}

	result, err := g.DB().Ctx(ctx).Model("order").Where("id", req.Id).Data(data).Update()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return gerror.New("订单不存在")
	}
	return nil
}

func (s *sAdminOrder) Refund(ctx context.Context, req *adminApi.OrderRefundReq) error {
	result, err := g.DB().Exec(ctx,
		`UPDATE "order" SET status = $1, updated_at = NOW() WHERE id = $2 AND status IN ($3, $4)`,
		consts.OrderStatusRefunded, req.Id, consts.OrderStatusPaid, consts.OrderStatusPreparing,
	)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return gerror.New("订单不可退款")
	}

	// 恢复库存
	items, _ := g.DB().Ctx(ctx).Model("order_item").Where("order_id", req.Id).All()
	for _, item := range items {
		_, _ = g.DB().Exec(ctx,
			`UPDATE product SET stock = stock + $1 WHERE id = $2`,
			item["quantity"].Int(), item["product_id"].Int64(),
		)
	}

	return nil
}
