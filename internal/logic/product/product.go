package product

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
)

type sProduct struct{}

func init() {
	service.SetProduct(&sProduct{})
}

func (s *sProduct) List(ctx context.Context, req *v1.ProductListReq) (*v1.ProductListRes, error) {
	m := g.DB().Ctx(ctx).Model("product").Where("status", 1)

	if req.CategoryId > 0 {
		m = m.Where("category_id", req.CategoryId)
	}
	if req.Keyword != "" {
		m = m.WhereLike("name", "%"+req.Keyword+"%")
	}

	switch req.Sort {
	case "sales":
		m = m.OrderDesc("sales")
	case "price_asc":
		m = m.OrderAsc("price")
	case "price_desc":
		m = m.OrderDesc("price")
	default:
		m = m.OrderDesc("sort_order").OrderDesc("id")
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var list []v1.ProductItem
	err = m.Page(req.Page, req.PageSize).Scan(&list)
	if err != nil {
		return nil, err
	}

	return &v1.ProductListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *sProduct) Detail(ctx context.Context, id int64) (*v1.ProductDetailRes, error) {
	one, err := g.DB().Ctx(ctx).Model("product").Where("id", id).Where("status", 1).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, gerror.New("商品不存在")
	}

	res := &v1.ProductDetailRes{
		Id:            one["id"].Int64(),
		CategoryId:    one["category_id"].Int(),
		Name:          one["name"].String(),
		Description:   one["description"].String(),
		Price:         one["price"].Float64(),
		OriginalPrice: one["original_price"].Float64(),
		ImageUrl:      one["image_url"].String(),
		Stock:         one["stock"].Int(),
		Sales:         one["sales"].Int(),
		IsHot:         one["is_hot"].Int(),
		IsRecommend:   one["is_recommend"].Int(),
	}

	// 解析图片 JSON
	var images []string
	_ = json.Unmarshal(one["images"].Bytes(), &images)
	res.Images = images

	// 查询规格
	var specs []v1.ProductSpecItem
	err = g.DB().Ctx(ctx).Model("product_spec").
		Where("product_id", id).Scan(&specs)
	if err != nil {
		return nil, err
	}
	res.Specs = specs

	return res, nil
}

func (s *sProduct) Search(ctx context.Context, req *v1.ProductSearchReq) (*v1.ProductListRes, error) {
	return s.List(ctx, &v1.ProductListReq{
		Keyword:  req.Keyword,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
}
