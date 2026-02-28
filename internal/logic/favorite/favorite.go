package favorite

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
)

type sFavorite struct{}

func init() {
	service.SetFavorite(&sFavorite{})
}

func (s *sFavorite) Add(ctx context.Context, userId int64, productId int64) error {
	_, err := g.DB().Exec(ctx,
		`INSERT INTO favorite (user_id, product_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`,
		userId, productId,
	)
	return err
}

func (s *sFavorite) Delete(ctx context.Context, userId int64, productId int64) error {
	_, err := g.DB().Ctx(ctx).Model("favorite").
		Where("user_id", userId).Where("product_id", productId).Delete()
	return err
}

func (s *sFavorite) List(ctx context.Context, userId int64, req *v1.FavoriteListReq) (*v1.FavoriteListRes, error) {
	m := g.DB().Ctx(ctx).Model("favorite f").
		LeftJoin("product p", "f.product_id = p.id").
		Where("f.user_id", userId).
		Where("p.status", 1)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var list []v1.ProductItem
	err = m.Fields("p.id, p.category_id, p.name, p.price, p.original_price, p.image_url, p.stock, p.sales, p.is_hot, p.is_recommend").
		OrderDesc("f.id").Page(req.Page, req.PageSize).Scan(&list)
	if err != nil {
		return nil, err
	}

	return &v1.FavoriteListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
