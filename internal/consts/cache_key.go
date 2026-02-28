package consts

import "fmt"

// Redis key 模板
const (
	CacheKeyProductDetail = "product:detail:%d"
	CacheKeyCategoryList  = "category:list"
	CacheKeyCartPrefix    = "cart:%d"
	CacheKeyIdempotent    = "idempotent:%s"
	CacheKeyRateLimit     = "rate_limit:%s:%d"
	CacheKeyProductStock  = "stock:product:%d"
	CacheKeyUserToken     = "user:token:%d"
)

func ProductDetailKey(id int64) string {
	return fmt.Sprintf(CacheKeyProductDetail, id)
}

func CartKey(userId int64) string {
	return fmt.Sprintf(CacheKeyCartPrefix, userId)
}

func IdempotentKey(key string) string {
	return fmt.Sprintf(CacheKeyIdempotent, key)
}

func RateLimitKey(ip string, userId int64) string {
	return fmt.Sprintf(CacheKeyRateLimit, ip, userId)
}
