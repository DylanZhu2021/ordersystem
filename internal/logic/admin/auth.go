package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/crypto/bcrypt"

	adminApi "ordersystem/api/v1/admin"
	"ordersystem/internal/service"
	"ordersystem/utility"
)

type sAdminAuth struct{}

func init() {
	service.SetAdminAuth(&sAdminAuth{})
}

func (s *sAdminAuth) Login(ctx context.Context, req *adminApi.AdminLoginReq) (*adminApi.AdminLoginRes, error) {
	admin, err := g.DB().Ctx(ctx).Model("admin_user").Where("username", req.Username).Where("status", 1).One()
	if err != nil {
		return nil, err
	}
	if admin.IsEmpty() {
		return nil, gerror.New("用户名或密码错误")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(admin["password_hash"].String()), []byte(req.Password))
	if err != nil {
		return nil, gerror.New("用户名或密码错误")
	}

	adminId := admin["id"].Int64()

	// 更新最后登录时间
	_, _ = g.DB().Exec(ctx, `UPDATE admin_user SET last_login_at = NOW() WHERE id = $1`, adminId)

	// 查询角色
	roleName := ""
	role, _ := g.DB().Ctx(ctx).Model("role").Where("id", admin["role_id"].Int()).One()
	if !role.IsEmpty() {
		roleName = role["name"].String()
	}

	token, err := utility.GenerateToken(ctx, adminId, "admin")
	if err != nil {
		return nil, err
	}

	return &adminApi.AdminLoginRes{
		Token:    token,
		AdminId:  adminId,
		RealName: admin["real_name"].String(),
		RoleName: roleName,
	}, nil
}

func (s *sAdminAuth) GetInfo(ctx context.Context, adminId int64) (*adminApi.AdminInfoRes, error) {
	admin, err := g.DB().Ctx(ctx).Model("admin_user").Where("id", adminId).One()
	if err != nil || admin.IsEmpty() {
		return nil, gerror.New("管理员不存在")
	}

	role, _ := g.DB().Ctx(ctx).Model("role").Where("id", admin["role_id"].Int()).One()

	return &adminApi.AdminInfoRes{
		Id:       adminId,
		Username: admin["username"].String(),
		RealName: admin["real_name"].String(),
		RoleId:   admin["role_id"].Int(),
		RoleName: role["name"].String(),
	}, nil
}
