package utility

import (
	"fmt"
	"sync/atomic"
	"time"
)

var orderSeq int64

// GenerateOrderNo 生成订单号: ORD + 年月日时分秒(14位) + 序列号(6位)
func GenerateOrderNo() string {
	now := time.Now()
	timestamp := now.Format("20060102150405")
	seq := atomic.AddInt64(&orderSeq, 1) % 1000000
	return fmt.Sprintf("ORD%s%06d", timestamp, seq)
}

// GeneratePaymentNo 生成支付流水号
func GeneratePaymentNo() string {
	now := time.Now()
	timestamp := now.Format("20060102150405")
	seq := atomic.AddInt64(&orderSeq, 1) % 1000000
	return fmt.Sprintf("PAY%s%06d", timestamp, seq)
}
