package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "ordersystem/api/v1"
	"ordersystem/internal/service"
	"ordersystem/utility"
)

type sAuth struct{}

func init() {
	service.SetAuth(&sAuth{})
}

func (s *sAuth) WechatLogin(ctx context.Context, req *v1.WechatLoginReq) (*v1.WechatLoginRes, error) {
	// 调用微信 code2session
	session, err := utility.Code2Session(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	// 查询或创建用户
	var userId int64
	isNew := false

	result, err := g.DB().Ctx(ctx).Raw(
		`INSERT INTO "user" (openid, union_id) VALUES (?, ?)
		 ON CONFLICT (openid) DO UPDATE SET updated_at = NOW()
		 RETURNING id, (xmax = 0) AS is_new`,
		session.OpenId, session.UnionId,
	).One()
	if err != nil {
		return nil, err
	}

	userId = result["id"].Int64()
	isNew = result["is_new"].Bool()

	// 生成 JWT
	token, err := utility.GenerateToken(ctx, userId, "user")
	if err != nil {
		return nil, err
	}

	return &v1.WechatLoginRes{
		Token:  token,
		UserId: userId,
		IsNew:  isNew,
	}, nil
}

func (s *sAuth) GetUserInfo(ctx context.Context, userId int64) (*v1.UserInfoRes, error) {
	one, err := g.DB().Ctx(ctx).Model("user").Where("id", userId).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, gerror.New("用户不存在")
	}

	return &v1.UserInfoRes{
		Id:          one["id"].Int64(),
		Nickname:    one["nickname"].String(),
		AvatarUrl:   one["avatar_url"].String(),
		Phone:       one["phone"].String(),
		Points:      one["points"].Int(),
		MemberLevel: one["member_level"].Int(),
	}, nil
}

func (s *sAuth) UpdateUser(ctx context.Context, userId int64, req *v1.UserUpdateReq) error {
	data := g.Map{}
	if req.Nickname != "" {
		data["nickname"] = req.Nickname
	}
	if req.AvatarUrl != "" {
		data["avatar_url"] = req.AvatarUrl
	}
	if req.Phone != "" {
		data["phone"] = req.Phone
	}
	if req.Gender > 0 {
		data["gender"] = req.Gender
	}
	if len(data) == 0 {
		return nil
	}

	_, err := g.DB().Ctx(ctx).Model("user").Where("id", userId).Data(data).Update()
	return err
}
