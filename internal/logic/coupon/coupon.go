package coupon

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
)

type sCoupon struct{}

func init() {
	service.SetCoupon(&sCoupon{})
}

func (s *sCoupon) List(ctx context.Context, userId int64, req *v1.CouponListReq) (*v1.CouponListRes, error) {
	m := g.DB().Ctx(ctx).Model("coupon").
		Where("status", 1).
		Where("end_time >= NOW()").
		Where("start_time <= NOW()")

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	records, err := m.Page(req.Page, req.PageSize).All()
	if err != nil {
		return nil, err
	}

	var list []v1.CouponItem
	for _, r := range records {
		// 检查用户是否已领取
		claimed := false
		if userId > 0 {
			cnt, _ := g.DB().Ctx(ctx).Model("user_coupon").
				Where("user_id", userId).Where("coupon_id", r["id"].Int64()).Count()
			claimed = cnt > 0
		}

		list = append(list, v1.CouponItem{
			Id:            r["id"].Int64(),
			Name:          r["name"].String(),
			Type:          r["type"].Int(),
			DiscountValue: r["discount_value"].Float64(),
			MinAmount:     r["min_amount"].Float64(),
			StartTime:     r["start_time"].String(),
			EndTime:       r["end_time"].String(),
			Claimed:       claimed,
		})
	}

	return &v1.CouponListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *sCoupon) Claim(ctx context.Context, userId int64, couponId int64) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 原子扣减优惠券数量
		result, err := tx.Exec(
			`UPDATE coupon SET claimed_count = claimed_count + 1
			 WHERE id = $1 AND claimed_count < total_count AND status = 1
			 AND start_time <= NOW() AND end_time >= NOW()`,
			couponId,
		)
		if err != nil {
			return err
		}
		affected, _ := result.RowsAffected()
		if affected == 0 {
			return gerror.New("优惠券已领完或不可用")
		}

		// 检查用户领取限制
		coupon, _ := tx.GetOne(`SELECT per_user_limit FROM coupon WHERE id = $1`, couponId)
		limit := coupon["per_user_limit"].Int()

		cnt, _ := tx.GetCount(
			`SELECT COUNT(*) FROM user_coupon WHERE user_id = $1 AND coupon_id = $2`,
			userId, couponId,
		)
		if int(cnt) >= limit {
			// 回滚扣减
			_, _ = tx.Exec(`UPDATE coupon SET claimed_count = claimed_count - 1 WHERE id = $1`, couponId)
			return gerror.New("已达领取上限")
		}

		// 创建领取记录
		_, err = tx.Insert("user_coupon", g.Map{
			"user_id":   userId,
			"coupon_id": couponId,
			"status":    0,
		})
		return err
	})
}

func (s *sCoupon) MyCoupons(ctx context.Context, userId int64, req *v1.MyCouponListReq) (*v1.MyCouponListRes, error) {
	m := g.DB().Ctx(ctx).Model("user_coupon uc").
		LeftJoin("coupon c", "uc.coupon_id = c.id").
		Where("uc.user_id", userId)

	if req.Status >= 0 {
		m = m.Where("uc.status", req.Status)
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	records, err := m.Fields("uc.id, uc.coupon_id, c.name, c.type, c.discount_value, c.min_amount, uc.status, c.end_time").
		OrderDesc("uc.id").Page(req.Page, req.PageSize).All()
	if err != nil {
		return nil, err
	}

	var list []v1.MyCouponItem
	for _, r := range records {
		list = append(list, v1.MyCouponItem{
			Id:            r["id"].Int64(),
			CouponId:      r["coupon_id"].Int64(),
			Name:          r["name"].String(),
			Type:          r["type"].Int(),
			DiscountValue: r["discount_value"].Float64(),
			MinAmount:     r["min_amount"].Float64(),
			Status:        r["status"].Int(),
			EndTime:       r["end_time"].String(),
		})
	}

	return &v1.MyCouponListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
