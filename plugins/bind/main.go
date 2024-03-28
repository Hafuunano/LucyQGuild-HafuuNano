// Package bind Bind CHUN & MAI ACCOUNT | DIVING FISH SERVICE.
package bind

import (
	nano "github.com/fumiama/NanoBot"
	"github.com/moyoez/HafuuNano/utils"
	"regexp"
)

func init() {
	nano.OnMessageCommand("bind", nano.OnlyToMe).SetBlock(true).Limit(utils.LimitByUser).Handle(func(ctx *nano.Ctx) {
		_, getSplit := utils.SplitCommandTo(ctx.MessageString(), 2)
		// to check if user input in a wrong way.
		patterns, err := regexp.Compile("<[^<>]+>")
		if err != nil {
			panic(err)
		}
		switch {
		case len(getSplit) < 2:
			ctx.SendPlainMessage(true, "你还没有填完整! /bind <UserName | UserEmail> 进行绑定 水鱼查分器 ")
			return
		case patterns.MatchString(getSplit[1]):
			ctx.SendPlainMessage(true, "笨蛋~! 为什么要带上 '<>'  "+" ( ｡ •̀ ᴖ •́ ｡)💢 ")
			return
		}
		utils.BindUserGenaralInfo(ctx.UserID(), getSplit[1])
		ctx.SendPlainMessage(true, "绑定成功了哦~ 绑定ID: "+getSplit[1])
	})
}
