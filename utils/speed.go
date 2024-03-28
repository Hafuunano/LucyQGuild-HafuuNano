package utils

import (
	nano "github.com/fumiama/NanoBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"time"
)

var defaultLimiterManager = rate.NewManager[int64](time.Second*1, 1)

// LimitByUser 默认限速器 每 10s 5次触发
//
//	按 发送者 限制
func LimitByUser(ctx *nano.Ctx) *rate.Limiter {
	if _, ok := ctx.Value.(*nano.Message); ok {
		return defaultLimiterManager.Load(int64(ctx.UserID()))
	}
	return defaultLimiterManager.Load(0)
}
