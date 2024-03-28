package ping

import (
	nano "github.com/fumiama/NanoBot"
	"github.com/moyoez/HafuuNano/utils"
)

func init() {
	nano.OnMessageCommand("ping", nano.OnlyToMe).SetBlock(true).Limit(utils.LimitByUser).Handle(func(ctx *nano.Ctx) {
		ctx.SendPlainMessage(true, "pong!")
	})
}
