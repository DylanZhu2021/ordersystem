package utility

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"ordersystem/internal/consts"
)

// IdempotentCheck 幂等性检查，返回 true 表示首次请求
func IdempotentCheck(ctx context.Context, key string) bool {
	redisKey := consts.IdempotentKey(key)
	ok, err := g.Redis().SetNX(ctx, redisKey, 1)
	if err != nil || !ok {
		return false
	}
	_, _ = g.Redis().Expire(ctx, redisKey, 300) // 5分钟过期
	return true
}
