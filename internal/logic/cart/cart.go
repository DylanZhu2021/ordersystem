package cart

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/consts"
	"ordersystem/internal/service"
)

type sCart struct{}

func init() {
	service.SetCart(&sCart{})
}

type cartEntry struct {
	ProductId int64 `json:"productId"`
	SpecId    int64 `json:"specId"`
	Quantity  int   `json:"quantity"`
}

func cartField(productId, specId int64) string {
	return fmt.Sprintf("%d:%d", productId, specId)
}

func (s *sCart) Add(ctx context.Context, userId int64, req *v1.CartAddReq) error {
	key := consts.CartKey(userId)
	field := cartField(req.ProductId, req.SpecId)

	// 检查商品是否存在
	count, err := g.DB().Ctx(ctx).Model("product").Where("id", req.ProductId).Where("status", 1).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("商品不存在或已下架")
	}

	// 获取当前数量
	existing, err := g.Redis().HGet(ctx, key, field)
	if err != nil {
		return err
	}

	newQty := req.Quantity
	if !existing.IsNil() && !existing.IsEmpty() {
		var entry cartEntry
		if err := json.Unmarshal(existing.Bytes(), &entry); err == nil {
			newQty += entry.Quantity
		}
	}

	entry := cartEntry{
		ProductId: req.ProductId,
		SpecId:    req.SpecId,
		Quantity:  newQty,
	}
	data, _ := json.Marshal(entry)

	_, err = g.Redis().HSet(ctx, key, map[string]interface{}{field: string(data)})
	if err != nil {
		return err
	}

	// 设置 30 天过期
	_, _ = g.Redis().Expire(ctx, key, 30*24*3600)
	return nil
}

func (s *sCart) List(ctx context.Context, userId int64) (*v1.CartListRes, error) {
	key := consts.CartKey(userId)
	all, err := g.Redis().HGetAll(ctx, key)
	if err != nil {
		return nil, err
	}

	res := &v1.CartListRes{List: make([]v1.CartItem, 0)}
	totalPrice := 0.0

	allMap := all.MapStrStr()
	for _, val := range allMap {
		var entry cartEntry
		if err := json.Unmarshal([]byte(val), &entry); err != nil {
			continue
		}

		// 查询商品信息
		product, err := g.DB().Ctx(ctx).Model("product").
			Where("id", entry.ProductId).Where("status", 1).One()
		if err != nil || product.IsEmpty() {
			continue
		}

		price := product["price"].Float64()
		specInfo := ""

		// 查询规格信息
		if entry.SpecId > 0 {
			spec, _ := g.DB().Ctx(ctx).Model("product_spec").
				Where("id", entry.SpecId).One()
			if !spec.IsEmpty() {
				price += spec["price_diff"].Float64()
				specInfo = spec["spec_name"].String() + ": " + spec["spec_value"].String()
			}
		}

		item := v1.CartItem{
			ProductId:    entry.ProductId,
			ProductName:  product["name"].String(),
			ProductImage: product["image_url"].String(),
			SpecId:       entry.SpecId,
			SpecInfo:     specInfo,
			Price:        price,
			Quantity:     entry.Quantity,
			Stock:        product["stock"].Int(),
		}
		res.List = append(res.List, item)
		totalPrice += price * float64(entry.Quantity)
	}

	res.TotalPrice = totalPrice
	return res, nil
}

func (s *sCart) Update(ctx context.Context, userId int64, req *v1.CartUpdateReq) error {
	key := consts.CartKey(userId)
	field := cartField(req.ProductId, req.SpecId)

	if req.Quantity <= 0 {
		_, err := g.Redis().HDel(ctx, key, field)
		return err
	}

	entry := cartEntry{
		ProductId: req.ProductId,
		SpecId:    req.SpecId,
		Quantity:  req.Quantity,
	}
	data, _ := json.Marshal(entry)
	_, err := g.Redis().HSet(ctx, key, map[string]interface{}{field: string(data)})
	return err
}

func (s *sCart) Delete(ctx context.Context, userId int64, req *v1.CartDeleteReq) error {
	key := consts.CartKey(userId)
	field := cartField(req.ProductId, req.SpecId)
	_, err := g.Redis().HDel(ctx, key, field)
	return err
}

func (s *sCart) Clear(ctx context.Context, userId int64) error {
	key := consts.CartKey(userId)
	_, err := g.Redis().Del(ctx, key)
	return err
}

// GetCartEntries 内部方法，供订单创建使用
func GetCartEntries(ctx context.Context, userId int64) ([]cartEntry, error) {
	key := consts.CartKey(userId)
	all, err := g.Redis().HGetAll(ctx, key)
	if err != nil {
		return nil, err
	}

	var entries []cartEntry
	allMap := all.MapStrStr()
	for _, val := range allMap {
		var entry cartEntry
		if err := json.Unmarshal([]byte(val), &entry); err != nil {
			continue
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
