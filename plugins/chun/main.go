package chun

import (
	"encoding/json"
	"github.com/FloatTech/gg"
	nano "github.com/fumiama/NanoBot"
	"github.com/moyoez/HafuuNano/utils"
	"strconv"
)

func init() {
	nano.OnMessageCommand("chun", nano.OnlyToMe).SetBlock(true).Limit(utils.LimitByUser).Handle(func(ctx *nano.Ctx) {
		checker := utils.QueryUserGeneralInfo(ctx.UserID())
		if checker.UserName == "" {
			// none
			ctx.SendPlainMessage(true, "你还没有绑定呢! 使用 /bind <UserName | UserEmail> 进行绑定 水鱼查分器")
			return
		}
		// query Data
		getData, err := QueryChunDataFromUserName(checker.UserName)
		if err != nil {
			ctx.SendPlainMessage(true, "ERR: ", err)
			return
		}
		var structData ChunData
		json.Unmarshal(getData, &structData)
		// Var Handler
		// result Image here.
		gg.SaveJPG(utils.ResLoader("chun")+"UserRender/"+strconv.FormatUint(ctx.UserID(), 10)+"_chun.png", BaseRender(structData, ctx), 90)
		ctx.SendImage("file:///"+utils.ResLoader("chun")+"UserRender/"+strconv.FormatUint(ctx.UserID(), 10)+"_chun.png", true)
	})

}
