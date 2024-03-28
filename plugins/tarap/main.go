// Package tarap Inspired By Tarnhelm Project, make tracker Stop here.
/*

<This Function Is Forbidden By Tencent Policy.>

2. 基本信息规范

"""
（3） QQ 机器人的名称、简介、头像包含如下内容，将无法通过审核：

▶️ 含有 QQ 、微信、微博、网址、邮箱、手机号等联系信息。

"""

*/

package tarap

import (
	"github.com/FloatTech/AnimeAPI/bilibili"
	nano "github.com/fumiama/NanoBot"
	"github.com/moyoez/HafuuNano/utils"
	"regexp"
)

func init() {
	nano.OnMessageCommand("fix", nano.OnlyToMe).SetBlock(true).Limit(utils.LimitByUser).Handle(func(ctx *nano.Ctx) {
		getLength, content := utils.SplitCommandTo(ctx.Message.Content, 2)
		if getLength < 2 {
			ctx.SendPlainMessage(true, "未获取到链接~")
			return
		}
		getRealLink := CleanerFixedLink(content[1])
		if getRealLink != "" {
			//	utils.SendWithoutFixURL(ctx, true, "已经移除相关跟踪参数~: \n"+getRealLink)
			ctx.SendPlainMessage(true, "已经移除相关跟踪参数~: \n"+getRealLink)
		} else {
			ctx.SendPlainMessage(true, "暂时不支持对应连接参数的移除~")
		}
	})
}

func CleanerFixedLink(link string) string {
	switch {
	case RegexPattern("(?:https?:\\/\\/)?music\\.apple\\.com.*[^\\/?]+\\/\\d+", link) != "":
		// APPLE LIST
		return RegexPattern("(?:https?:\\/\\/)?music\\.apple\\.com.*[^\\/?]+\\/\\d+", link)
	case RegexPattern("(?:https?:\\/\\/)?www\\.bilibili\\.com\\/video\\/[^\\/?]+\\/", link) != "":
		return RegexPattern("(?:https?:\\/\\/)?www\\.bilibili\\.com\\/video\\/[^\\/?]+\\/", link)
	case RegexPattern("(?:https?:\\/\\/)?y\\.music\\.163\\.com\\/m\\/song\\?id=\\d+", link) != "":
		return RegexPattern("(?:https?:\\/\\/)?y\\.music\\.163\\.com\\/m\\/song\\?id=\\d+", link)
	case RegexPattern("((b23|acg).tv|bili2233.cn)/[0-9a-zA-Z]+", link) != "":
		return BilibiliFixedLink(link)
	default:
		return ""
	}
}

func RegexPattern(regex string, needClear string) (returnClear string) {
	getRegex := regexp.MustCompile(regex)
	if len(getRegex.FindAllString(needClear, 1)) == 1 {
		// Matched.
		return getRegex.FindAllString(needClear, 1)[0]
	} else {
		return ""
	}

}

func BilibiliFixedLink(link string) string {
	// query for the last link here.
	getRealLink, err := bilibili.GetRealURL(link)
	if err != nil {
		return ""
	}
	getReq, err := regexp.Compile("bilibili.com\\\\?/video\\\\?/(?:av(\\d+)|([bB][vV][0-9a-zA-Z]+))")
	if err != nil {
		return ""
	}
	return getReq.FindAllString(getRealLink, 1)[0]
}
