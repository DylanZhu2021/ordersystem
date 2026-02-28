package review

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
)

type sReview struct{}

func init() {
	service.SetReview(&sReview{})
}

func (s *sReview) Create(ctx context.Context, userId int64, req *v1.ReviewCreateReq) error {
	images, _ := json.Marshal(req.Images)
	_, err := g.DB().Ctx(ctx).Model("review").Data(g.Map{
		"user_id":    userId,
		"product_id": req.ProductId,
		"order_id":   req.OrderId,
		"rating":     req.Rating,
		"content":    req.Content,
		"images":     string(images),
	}).Insert()
	return err
}

func (s *sReview) List(ctx context.Context, req *v1.ReviewListReq) (*v1.ReviewListRes, error) {
	m := g.DB().Ctx(ctx).Model("review r").
		LeftJoin(`"user" u`, "r.user_id = u.id").
		Where("r.product_id", req.ProductId).
		Where("r.status", 1)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	records, err := m.Fields("r.id, u.nickname, u.avatar_url, r.rating, r.content, r.images, r.reply, r.created_at").
		OrderDesc("r.id").Page(req.Page, req.PageSize).All()
	if err != nil {
		return nil, err
	}

	var list []v1.ReviewItem
	for _, r := range records {
		var images []string
		_ = json.Unmarshal(r["images"].Bytes(), &images)

		list = append(list, v1.ReviewItem{
			Id:        r["id"].Int64(),
			Nickname:  r["nickname"].String(),
			AvatarUrl: r["avatar_url"].String(),
			Rating:    r["rating"].Int(),
			Content:   r["content"].String(),
			Images:    images,
			Reply:     r["reply"].String(),
			CreatedAt: r["created_at"].String(),
		})
	}

	return &v1.ReviewListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
