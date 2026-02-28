package address

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
)

type sAddress struct{}

func init() {
	service.SetAddress(&sAddress{})
}

func (s *sAddress) List(ctx context.Context, userId int64) (*v1.AddressListRes, error) {
	var list []v1.AddressItem
	err := g.DB().Ctx(ctx).Model("address").
		Where("user_id", userId).
		OrderAsc("is_default").OrderDesc("id").
		Scan(&list)
	if err != nil {
		return nil, err
	}
	return &v1.AddressListRes{List: list}, nil
}

func (s *sAddress) Create(ctx context.Context, userId int64, req *v1.AddressCreateReq) (int64, error) {
	// 如果设为默认，先取消其他默认
	if req.IsDefault == 1 {
		_, _ = g.DB().Ctx(ctx).Model("address").
			Where("user_id", userId).Data(g.Map{"is_default": 0}).Update()
	}

	result, err := g.DB().Ctx(ctx).Model("address").Data(g.Map{
		"user_id":    userId,
		"name":       req.Name,
		"phone":      req.Phone,
		"province":   req.Province,
		"city":       req.City,
		"district":   req.District,
		"detail":     req.Detail,
		"is_default": req.IsDefault,
	}).InsertAndGetId()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *sAddress) Update(ctx context.Context, userId int64, req *v1.AddressUpdateReq) error {
	if req.IsDefault == 1 {
		_, _ = g.DB().Ctx(ctx).Model("address").
			Where("user_id", userId).Data(g.Map{"is_default": 0}).Update()
	}

	_, err := g.DB().Ctx(ctx).Model("address").
		Where("id", req.Id).Where("user_id", userId).
		Data(g.Map{
			"name":       req.Name,
			"phone":      req.Phone,
			"province":   req.Province,
			"city":       req.City,
			"district":   req.District,
			"detail":     req.Detail,
			"is_default": req.IsDefault,
		}).Update()
	return err
}

func (s *sAddress) Delete(ctx context.Context, userId int64, id int64) error {
	result, err := g.DB().Ctx(ctx).Model("address").
		Where("id", id).Where("user_id", userId).Delete()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return gerror.New("地址不存在")
	}
	return nil
}
