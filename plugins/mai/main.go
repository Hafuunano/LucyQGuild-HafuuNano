package mai

import (
	"encoding/json"
	"github.com/FloatTech/gg"
	nano "github.com/fumiama/NanoBot"
	"github.com/moyoez/HafuuNano/utils"
	"strconv"
)

func init() {
	nano.OnMessageCommand("mai", nano.OnlyToMe).SetBlock(true).Limit(utils.LimitByUser).Handle(func(ctx *nano.Ctx) {
		MaimaiRenderBase(ctx)
	})
}

// MaimaiRenderBase Render Base Maimai B50.
func MaimaiRenderBase(ctx *nano.Ctx) {
	// check the user using.
	getUserID := ctx.UserID()
	getUsername := utils.QueryUserGeneralInfo(ctx.UserID()).UserName
	if getUsername == "" {
		ctx.SendPlainMessage(true, "你还没有绑定呢！使用/bind <UserName | UserEmail > 以绑定水鱼查分器")
		return
	}
	getUserData, err := QueryMaiBotDataFromUserName(getUsername)
	if err != nil {
		ctx.SendPlainMessage(true, err)
		return
	}
	var data player
	_ = json.Unmarshal(getUserData, &data)
	renderImg := FullPageRender(data, ctx)
	_ = gg.NewContextForImage(renderImg).SavePNG(utils.ResLoader("mai") + "save/" + strconv.Itoa(int(getUserID)) + ".png")
	ctx.SendImage("file:///"+utils.ResLoader("mai")+"save/"+strconv.Itoa(int(getUserID))+".png", true)
	/*
		if GetUserSwitcherInfoFromDatabase(getUserID) == tru									e {
			// use lxns checker service.
			// check bind first, get user friend id.
			getFriendID := GetUserMaiFriendID(getUserID)
			if getFriendID.MaimaiID == 0 {
				ctx.SendPlainMessage(true, "你还没有绑定呢！使用/mai lxbind <friendcode> 以绑定")
				return
			}
			getUserData := RequestBasicDataFromLxns(getFriendID.MaimaiID)
			if getUserData.Code != 200 {
				ctx.SendPlainMessage(true, "aw 出现了一点小错误~：\n - 请检查你是否有上传过数据\n - 请检查你的设置是否允许了第三方查看")
				return
			}
			getGameUserData := RequestB50DataByFriendCode(getUserData.Data.FriendCode)
			if getGameUserData.Code != 200 {
				ctx.SendPlainMessage(true, "aw 出现了一点小错误~：\n - 请检查你是否有上传过数据\n - 请检查你的设置是否允许了第三方查看")
				return
			}
			getImager, _ := ReFullPageRender(getGameUserData, getUserData, ctx)
			_ = gg.NewContextForImage(getImager).SavePNG(engine.DataFolder() + "save/" + "LXNS_" + strconv.Itoa(int(getUserID)) + ".png")
			if israw {
				getDocumentType := &tgba.DocumentConfig{
					BaseFile: tgba.BaseFile{BaseChat: tgba.BaseChat{
						ChatConfig: tgba.ChatConfig{ChatID: ctx.Message.Chat.ID},
					},
						File: tgba.FilePath(engine.DataFolder() + "save/" + "LXNS_" + strconv.Itoa(int(getUserID)) + ".png")},
					Caption:         "",
					CaptionEntities: nil,
				}
				ctx.Send(true, getDocumentType)
			} else {
				ctx.SendPhoto(tgba.FilePath(engine.DataFolder()+"save/"+"LXNS_"+strconv.Itoa(int(getUserID))+".png"), true, "")
			}
		} else {
			// diving fish checker:
			getUsername := GetUserInfoNameFromDatabase(getUserID)
			if getUsername == "" {
				ctx.SendPlainMessage(true, "你还没有绑定呢！使用/mai bind <UserName> 以绑定")
				return
			}
			getUserData, err := QueryMaiBotDataFromUserName(getUsername)
			if err != nil {
				ctx.SendPlainMessage(true, err)
				return
			}
			var data player
			_ = json.Unmarshal(getUserData, &data)
			renderImg := FullPageRender(data, ctx)
			_ = gg.NewContextForImage(renderImg).SavePNG(engine.DataFolder() + "save/" + strconv.Itoa(int(getUserID)) + ".png")

			if israw {
				getDocumentType := &tgba.DocumentConfig{
					BaseFile: tgba.BaseFile{BaseChat: tgba.BaseChat{
						ChatConfig: tgba.ChatConfig{ChatID: ctx.Message.Chat.ID},
					},
						File: tgba.FilePath(engine.DataFolder() + "save/" + strconv.Itoa(int(getUserID)) + ".png")},
					Caption:         "",
					CaptionEntities: nil,
				}
				ctx.Send(true, getDocumentType)
			} else {
				ctx.SendPhoto(tgba.FilePath(engine.DataFolder()+"save/"+strconv.Itoa(int(getUserID))+".png"), true, "")
			}
		}

	*/
}
