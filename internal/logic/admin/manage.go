package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	adminApi "ordersystem/api/v1/admin"
	"ordersystem/internal/service"
)

type sAdminCoupon struct{}
type sAdminCategory struct{}
type sAdminStats struct{}
type sAdminUser struct{}

func init() {
	service.SetAdminCoupon(&sAdminCoupon{})
	service.SetAdminCategory(&sAdminCategory{})
	service.SetAdminStats(&sAdminStats{})
	service.SetAdminUser(&sAdminUser{})
}

// ==================== 优惠券管理 ====================

func (s *sAdminCoupon) Create(ctx context.Context, req *adminApi.CouponCreateReq) (int64, error) {
	id, err := g.DB().Ctx(ctx).Model("coupon").InsertAndGetId(g.Map{
		"name":           req.Name,
		"type":           req.Type,
		"discount_value": req.DiscountValue,
		"min_amount":     req.MinAmount,
		"total_count":    req.TotalCount,
		"per_user_limit": req.PerUserLimit,
		"start_time":     req.StartTime,
		"end_time":       req.EndTime,
		"status":         1,
	})
	return id, err
}

func (s *sAdminCoupon) Update(ctx context.Context, req *adminApi.CouponUpdateReq) error {
	data := g.Map{}
	if req.Name != "" {
		data["name"] = req.Name
	}
	if req.TotalCount > 0 {
		data["total_count"] = req.TotalCount
	}
	if req.PerUserLimit > 0 {
		data["per_user_limit"] = req.PerUserLimit
	}
	if req.StartTime != "" {
		data["start_time"] = req.StartTime
	}
	if req.EndTime != "" {
		data["end_time"] = req.EndTime
	}
	if req.Status >= 0 {
		data["status"] = req.Status
	}
	_, err := g.DB().Ctx(ctx).Model("coupon").Where("id", req.Id).Data(data).Update()
	return err
}

func (s *sAdminCoupon) Delete(ctx context.Context, id int64) error {
	_, err := g.DB().Ctx(ctx).Model("coupon").Where("id", id).Delete()
	return err
}

// ==================== 分类管理 ====================

func (s *sAdminCategory) Create(ctx context.Context, req *adminApi.CategoryCreateReq) (int, error) {
	id, err := g.DB().Ctx(ctx).Model("category").InsertAndGetId(g.Map{
		"name":       req.Name,
		"icon_url":   req.IconUrl,
		"sort_order": req.SortOrder,
		"status":     1,
	})
	return int(id), err
}

func (s *sAdminCategory) Update(ctx context.Context, req *adminApi.CategoryUpdateReq) error {
	data := g.Map{}
	if req.Name != "" {
		data["name"] = req.Name
	}
	if req.IconUrl != "" {
		data["icon_url"] = req.IconUrl
	}
	data["sort_order"] = req.SortOrder
	data["status"] = req.Status
	_, err := g.DB().Ctx(ctx).Model("category").Where("id", req.Id).Data(data).Update()
	return err
}

func (s *sAdminCategory) Delete(ctx context.Context, id int) error {
	// 检查是否有商品关联
	count, _ := g.DB().Ctx(ctx).Model("product").Where("category_id", id).Count()
	if count > 0 {
		return gerror.New("该分类下有商品，无法删除")
	}
	_, err := g.DB().Ctx(ctx).Model("category").Where("id", id).Delete()
	return err
}

// ==================== 数据统计 ====================

func (s *sAdminStats) Dashboard(ctx context.Context) (*adminApi.DashboardRes, error) {
	// 今日销售额
	todaySales, _ := g.DB().Ctx(ctx).Raw(
		`SELECT COALESCE(SUM(pay_amount), 0) FROM "order" WHERE status >= 1 AND created_at >= CURRENT_DATE`,
	).Value()
	// 今日订单数
	todayOrders, _ := g.DB().Ctx(ctx).Raw(
		`SELECT COUNT(*) FROM "order" WHERE created_at >= CURRENT_DATE`,
	).Value()
	// 总用户数
	totalUsers, _ := g.DB().Ctx(ctx).Raw(
		`SELECT COUNT(*) FROM "user"`,
	).Value()
	// 平均客单价
	avgAmount, _ := g.DB().Ctx(ctx).Raw(
		`SELECT COALESCE(AVG(pay_amount), 0) FROM "order" WHERE status >= 1 AND created_at >= CURRENT_DATE`,
	).Value()

	return &adminApi.DashboardRes{
		TodaySales:     todaySales.Float64(),
		TodayOrders:    todayOrders.Int(),
		TotalUsers:     totalUsers.Int(),
		AvgOrderAmount: avgAmount.Float64(),
	}, nil
}

func (s *sAdminStats) SalesStats(ctx context.Context, req *adminApi.SalesStatsReq) (*adminApi.SalesStatsRes, error) {
	records, err := g.DB().Ctx(ctx).Raw(
		`SELECT DATE(created_at) as date, COALESCE(SUM(pay_amount), 0) as sales, COUNT(*) as order_count
		 FROM "order" WHERE status >= 1 AND created_at >= $1 AND created_at <= $2
		 GROUP BY DATE(created_at) ORDER BY date`,
		req.StartDate, req.EndDate,
	).All()
	if err != nil {
		return nil, err
	}

	var list []adminApi.SalesStatItem
	for _, r := range records {
		list = append(list, adminApi.SalesStatItem{
			Date:       r["date"].String(),
			Sales:      r["sales"].Float64(),
			OrderCount: r["order_count"].Int(),
		})
	}

	return &adminApi.SalesStatsRes{List: list}, nil
}

func (s *sAdminStats) HotProducts(ctx context.Context, limit int) (*adminApi.HotProductsRes, error) {
	records, err := g.DB().Ctx(ctx).Raw(
		`SELECT p.id as product_id, p.name as product_name, p.sales,
		 COALESCE(SUM(oi.total_amount), 0) as revenue
		 FROM product p LEFT JOIN order_item oi ON p.id = oi.product_id
		 GROUP BY p.id, p.name, p.sales ORDER BY p.sales DESC LIMIT $1`,
		limit,
	).All()
	if err != nil {
		return nil, err
	}

	var list []adminApi.HotProductItem
	for _, r := range records {
		list = append(list, adminApi.HotProductItem{
			ProductId:   r["product_id"].Int64(),
			ProductName: r["product_name"].String(),
			Sales:       r["sales"].Int(),
			Revenue:     r["revenue"].Float64(),
		})
	}

	return &adminApi.HotProductsRes{List: list}, nil
}

// ==================== 用户管理 ====================

func (s *sAdminUser) List(ctx context.Context, req *adminApi.UserListReq) (*adminApi.UserListRes, error) {
	m := g.DB().Ctx(ctx).Model("user")
	if req.Keyword != "" {
		m = m.Where("nickname LIKE ? OR phone LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	records, err := m.OrderDesc("id").Page(req.Page, req.PageSize).All()
	if err != nil {
		return nil, err
	}

	return &adminApi.UserListRes{
		List:     records,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
