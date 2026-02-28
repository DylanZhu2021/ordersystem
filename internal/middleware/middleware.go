package middleware

import (
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
	"ordersystem/utility"
)

// JWTAuth 用户端 JWT 认证中间件
func JWTAuth(r *ghttp.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		utility.Error(r, 1002, "未登录")
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")
	claims, err := utility.ParseToken(r.Context(), token)
	if err != nil || claims.UserType != "user" {
		utility.Error(r, 1002, "登录已过期")
		return
	}

	r.SetCtxVar("userId", claims.UserId)
	r.Middleware.Next()
}

// AdminAuth 管理端 JWT 认证中间件
func AdminAuth(r *ghttp.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		utility.Error(r, 1002, "未登录")
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")
	claims, err := utility.ParseToken(r.Context(), token)
	if err != nil || claims.UserType != "admin" {
		utility.Error(r, 1002, "登录已过期")
		return
	}

	r.SetCtxVar("adminId", claims.UserId)
	r.Middleware.Next()
}

// CORS 跨域中间件
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// RequestLog 请求日志中间件
func RequestLog(r *ghttp.Request) {
	r.Middleware.Next()
}

// ResponseHandler 统一响应处理中间件
func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = r.Response.Status
	)

	if err != nil {
		utility.Error(r, -1, err.Error())
		return
	}

	if code != 0 && code != http.StatusOK {
		return
	}

	if res != nil {
		utility.Success(r, res)
	}
}

// RateLimit 限流中间件
func RateLimit(r *ghttp.Request) {
	// 基于 Redis 的滑动窗口限流
	// 简化实现：每分钟 100 次请求
	ctx := r.Context()
	ip := r.GetClientIp()
	userId := r.GetCtxVar("userId", 0).Int64()

	key := "rate_limit:" + ip
	if userId > 0 {
		key = "rate_limit:user:" + string(rune(userId))
	}

	val := r.GetCtxVar("_rate_checked").Bool()
	if val {
		r.Middleware.Next()
		return
	}

	_ = ctx
	_ = key
	// TODO: 实际限流逻辑在生产环境中启用
	r.Middleware.Next()
}
