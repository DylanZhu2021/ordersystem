package admin

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	adminApi "ordersystem/api/v1/admin"
	"ordersystem/internal/service"
)

type sAdminProduct struct{}

func init() {
	service.SetAdminProduct(&sAdminProduct{})
}

func (s *sAdminProduct) Create(ctx context.Context, req *adminApi.ProductCreateReq) (int64, error) {
	images, _ := json.Marshal(req.Images)
	id, err := g.DB().Ctx(ctx).Model("product").InsertAndGetId(g.Map{
		"category_id":    req.CategoryId,
		"name":           req.Name,
		"description":    req.Description,
		"price":          req.Price,
		"original_price": req.OriginalPrice,
		"image_url":      req.ImageUrl,
		"images":         string(images),
		"stock":          req.Stock,
		"is_hot":         req.IsHot,
		"is_recommend":   req.IsRecommend,
		"sort_order":     req.SortOrder,
		"status":         1,
	})
	return id, err
}

func (s *sAdminProduct) Update(ctx context.Context, req *adminApi.ProductUpdateReq) error {
	data := g.Map{}
	if req.CategoryId > 0 {
		data["category_id"] = req.CategoryId
	}
	if req.Name != "" {
		data["name"] = req.Name
	}
	if req.Description != "" {
		data["description"] = req.Description
	}
	if req.Price > 0 {
		data["price"] = req.Price
	}
	if req.OriginalPrice > 0 {
		data["original_price"] = req.OriginalPrice
	}
	if req.ImageUrl != "" {
		data["image_url"] = req.ImageUrl
	}
	if req.Images != nil {
		images, _ := json.Marshal(req.Images)
		data["images"] = string(images)
	}
	if req.Stock >= 0 {
		data["stock"] = req.Stock
	}
	data["is_hot"] = req.IsHot
	data["is_recommend"] = req.IsRecommend
	data["sort_order"] = req.SortOrder

	_, err := g.DB().Ctx(ctx).Model("product").Where("id", req.Id).Data(data).Update()
	return err
}

func (s *sAdminProduct) Delete(ctx context.Context, id int64) error {
	result, err := g.DB().Ctx(ctx).Model("product").Where("id", id).Delete()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return gerror.New("商品不存在")
	}
	return nil
}

func (s *sAdminProduct) UpdateStatus(ctx context.Context, id int64, status int) error {
	_, err := g.DB().Ctx(ctx).Model("product").Where("id", id).Data(g.Map{"status": status}).Update()
	return err
}
