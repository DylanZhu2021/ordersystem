package order

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/consts"
	"ordersystem/internal/logic/cart"
	"ordersystem/internal/service"
	"ordersystem/utility"
)

type sOrder struct{}

func init() {
	service.SetOrder(&sOrder{})
}

func (s *sOrder) Create(ctx context.Context, userId int64, req *v1.OrderCreateReq) (*v1.OrderCreateRes, error) {
	// 1. 幂等性检查
	if !utility.IdempotentCheck(ctx, req.IdempotencyKey) {
		return nil, gerror.New("请勿重复提交")
	}

	// 2. 获取购物车数据
	entries, err := cart.GetCartEntries(ctx, userId)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, gerror.New("购物车为空")
	}

	// 3. 在事务中创建订单
	var orderRes *v1.OrderCreateRes
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		totalAmount := 0.0
		orderNo := utility.GenerateOrderNo()

		type itemData struct {
			ProductId    int64
			ProductName  string
			ProductImage string
			SpecId       int64
			SpecInfo     string
			Price        float64
			Quantity     int
			TotalAmount  float64
		}
		var items []itemData

		// 4. 校验商品并扣减库存
		for _, entry := range entries {
			// 原子扣减库存: UPDATE ... WHERE stock >= ?
			result, err := tx.Exec(
				`UPDATE product SET stock = stock - $1, sales = sales + $1, updated_at = NOW()
				 WHERE id = $2 AND stock >= $1 AND status = 1`,
				entry.Quantity, entry.ProductId,
			)
			if err != nil {
				return gerror.Wrapf(err, "扣减库存失败")
			}
			affected, _ := result.RowsAffected()
			if affected == 0 {
				return gerror.Newf("商品库存不足 (productId=%d)", entry.ProductId)
			}

			// 查询商品信息（快照）
			product, err := tx.GetOne(
				`SELECT name, image_url, price FROM product WHERE id = $1`, entry.ProductId,
			)
			if err != nil {
				return err
			}

			price := product["price"].Float64()
			specInfo := ""

			if entry.SpecId > 0 {
				spec, _ := tx.GetOne(
					`SELECT spec_name, spec_value, price_diff FROM product_spec WHERE id = $1`,
					entry.SpecId,
				)
				if !spec.IsEmpty() {
					price += spec["price_diff"].Float64()
					specInfo = spec["spec_name"].String() + ": " + spec["spec_value"].String()
				}
			}

			itemTotal := price * float64(entry.Quantity)
			totalAmount += itemTotal

			items = append(items, itemData{
				ProductId:    entry.ProductId,
				ProductName:  product["name"].String(),
				ProductImage: product["image_url"].String(),
				SpecId:       entry.SpecId,
				SpecInfo:     specInfo,
				Price:        price,
				Quantity:     entry.Quantity,
				TotalAmount:  itemTotal,
			})
		}

		// 5. 计算优惠
		discountAmount := 0.0
		if req.CouponId > 0 {
			discount, err := applyCoupon(ctx, tx, userId, req.CouponId, totalAmount)
			if err != nil {
				return err
			}
			discountAmount = discount
		}

		payAmount := totalAmount - discountAmount
		if payAmount < 0 {
			payAmount = 0
		}

		// 6. 创建订单
		orderId, err := tx.InsertAndGetId("order", g.Map{
			"order_no":        orderNo,
			"user_id":         userId,
			"total_amount":    totalAmount,
			"discount_amount": discountAmount,
			"pay_amount":      payAmount,
			"coupon_id":       req.CouponId,
			"delivery_type":   req.DeliveryType,
			"table_no":        req.TableNo,
			"address_id":      req.AddressId,
			"remark":          req.Remark,
			"status":          consts.OrderStatusPending,
			"idempotency_key": req.IdempotencyKey,
		})
		if err != nil {
			return gerror.Wrapf(err, "创建订单失败")
		}

		// 7. 创建订单明细
		for _, item := range items {
			_, err := tx.Insert("order_item", g.Map{
				"order_id":      orderId,
				"product_id":    item.ProductId,
				"product_name":  item.ProductName,
				"product_image": item.ProductImage,
				"spec_id":       item.SpecId,
				"spec_info":     item.SpecInfo,
				"price":         item.Price,
				"quantity":      item.Quantity,
				"total_amount":  item.TotalAmount,
			})
			if err != nil {
				return err
			}
		}

		orderRes = &v1.OrderCreateRes{
			OrderId:     orderId,
			OrderNo:     orderNo,
			TotalAmount: totalAmount,
			PayAmount:   payAmount,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// 8. 清空购物车
	_ = service.Cart().Clear(ctx, userId)

	return orderRes, nil
}

// applyCoupon 在事务中应用优惠券
func applyCoupon(ctx context.Context, tx gdb.TX, userId, couponId int64, totalAmount float64) (float64, error) {
	// 查询用户优惠券
	uc, err := tx.GetOne(
		`SELECT uc.id, c.type, c.discount_value, c.min_amount
		 FROM user_coupon uc JOIN coupon c ON uc.coupon_id = c.id
		 WHERE uc.id = $1 AND uc.user_id = $2 AND uc.status = 0
		 AND c.start_time <= NOW() AND c.end_time >= NOW()`,
		couponId, userId,
	)
	if err != nil {
		return 0, err
	}
	if uc.IsEmpty() {
		return 0, gerror.New("优惠券不可用")
	}

	minAmount := uc["min_amount"].Float64()
	if totalAmount < minAmount {
		return 0, gerror.Newf("未满足最低消费 %.2f 元", minAmount)
	}

	var discount float64
	switch uc["type"].Int() {
	case consts.CouponTypeFixed, consts.CouponTypeFree:
		discount = uc["discount_value"].Float64()
	case consts.CouponTypePercent:
		discount = totalAmount * (1 - uc["discount_value"].Float64()/100)
	}

	if discount > totalAmount {
		discount = totalAmount
	}

	// 标记优惠券已使用
	_, err = tx.Exec(
		`UPDATE user_coupon SET status = 1, used_at = NOW() WHERE id = $1`,
		uc["id"].Int64(),
	)
	return discount, err
}

func (s *sOrder) List(ctx context.Context, userId int64, req *v1.OrderListReq) (*v1.OrderListRes, error) {
	m := g.DB().Ctx(ctx).Model("order").Where("user_id", userId)
	if req.Status >= 0 {
		m = m.Where("status", req.Status)
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var orders []v1.OrderListItem
	records, err := m.OrderDesc("id").Page(req.Page, req.PageSize).All()
	if err != nil {
		return nil, err
	}

	for _, r := range records {
		// 查询订单明细
		var items []v1.OrderItemBrief
		_ = g.DB().Ctx(ctx).Model("order_item").
			Where("order_id", r["id"].Int64()).
			Scan(&items)

		orders = append(orders, v1.OrderListItem{
			Id:           r["id"].Int64(),
			OrderNo:      r["order_no"].String(),
			TotalAmount:  r["total_amount"].Float64(),
			PayAmount:    r["pay_amount"].Float64(),
			Status:       r["status"].Int(),
			DeliveryType: r["delivery_type"].Int(),
			ItemCount:    len(items),
			Items:        items,
			CreatedAt:    r["created_at"].String(),
		})
	}

	return &v1.OrderListRes{
		List:     orders,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *sOrder) Detail(ctx context.Context, userId int64, id int64) (*v1.OrderDetailRes, error) {
	one, err := g.DB().Ctx(ctx).Model("order").
		Where("id", id).Where("user_id", userId).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, gerror.New("订单不存在")
	}

	var items []v1.OrderDetailItem
	_ = g.DB().Ctx(ctx).Model("order_item").
		Where("order_id", id).Scan(&items)

	return &v1.OrderDetailRes{
		Id:             one["id"].Int64(),
		OrderNo:        one["order_no"].String(),
		TotalAmount:    one["total_amount"].Float64(),
		DiscountAmount: one["discount_amount"].Float64(),
		PayAmount:      one["pay_amount"].Float64(),
		Status:         one["status"].Int(),
		DeliveryType:   one["delivery_type"].Int(),
		TableNo:        one["table_no"].String(),
		ContactName:    one["contact_name"].String(),
		ContactPhone:   one["contact_phone"].String(),
		Remark:         one["remark"].String(),
		Items:          items,
		CreatedAt:      one["created_at"].String(),
		PaidAt:         one["paid_at"].String(),
	}, nil
}

func (s *sOrder) Cancel(ctx context.Context, userId int64, id int64) error {
	result, err := g.DB().Exec(ctx,
		`UPDATE "order" SET status = $1, cancelled_at = $2, updated_at = NOW()
		 WHERE id = $3 AND user_id = $4 AND status = $5`,
		consts.OrderStatusCancelled, time.Now(), id, userId, consts.OrderStatusPending,
	)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return gerror.New("订单不可取消")
	}

	// 恢复库存
	items, _ := g.DB().Ctx(ctx).Model("order_item").Where("order_id", id).All()
	for _, item := range items {
		_, _ = g.DB().Exec(ctx,
			`UPDATE product SET stock = stock + $1, sales = sales - $1 WHERE id = $2`,
			item["quantity"].Int(), item["product_id"].Int64(),
		)
	}

	return nil
}

func (s *sOrder) Refund(ctx context.Context, userId int64, req *v1.OrderRefundReq) error {
	result, err := g.DB().Exec(ctx,
		`UPDATE "order" SET status = $1, updated_at = NOW()
		 WHERE id = $2 AND user_id = $3 AND status = $4`,
		consts.OrderStatusRefunded, req.Id, userId, consts.OrderStatusPaid,
	)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return gerror.New("订单不可退款")
	}
	return nil
}
