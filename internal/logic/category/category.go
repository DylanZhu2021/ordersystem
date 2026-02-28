package category

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
)

type sCategory struct{}

func init() {
	service.SetCategory(&sCategory{})
}

func (s *sCategory) List(ctx context.Context) (*v1.CategoryListRes, error) {
	var list []v1.CategoryItem
	err := g.DB().Ctx(ctx).Model("category").
		Where("status", 1).
		OrderAsc("sort_order").
		Scan(&list)
	if err != nil {
		return nil, err
	}
	return &v1.CategoryListRes{List: list}, nil
}
