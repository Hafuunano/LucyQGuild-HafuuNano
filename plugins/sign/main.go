package sign

import (
	"fmt"
	"github.com/FloatTech/floatbox/file"
	"github.com/FloatTech/gg"
	nano "github.com/fumiama/NanoBot"
	"github.com/moyoez/HafuuNano/utils"
	"image/color"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

func init() {
	picDir, err := os.ReadDir(utils.ReturnLucyMainDataIndex("funwork") + "randpic")
	if err != nil {
		return
	}
	picDirNum := len(picDir)
	reg := regexp.MustCompile(`[^.]+`)

	nano.OnMessageCommand("sign", nano.OnlyToMe).SetBlock(true).Limit(utils.LimitByUser).Handle(func(ctx *nano.Ctx) {
		getUserName := utils.GetUserInfoName(ctx)
		getUserID := ctx.UserID()
		userPic := strconv.FormatUint(getUserID, 10) + time.Now().Format("20060102") + ".png" // On Save
		GetUserAwakenDays := utils.QueryUserAwakenFunction(getUserID).AwakenTime + 1
		GetPeopleCountDays := utils.QueryTotalAwakenFunction().Count + 1
		GetSignStatus := utils.QueryUserAwakenFunction(getUserID)

		// Modify Users.
		if GetSignStatus.SignDay == time.Now().Format("20060102") {
			ctx.SendPlainMessage(true, "今天你已经签到过了哦w")
			if file.IsExist(utils.ResLoader("sign") + "userpic/" + userPic) {
				ctx.SendImage("file:///"+utils.ResLoader("sign")+"userpic/"+userPic, true)
			}
			return
		}

		getReturnData := utils.QueryUserCoinsFunction(getUserID)
		CoinsFull := getReturnData.Coin
		// Query get Coins.
		CoinsGet := int64(rand.Intn(300) + 250)
		// Modify Coins

		utils.ModifyUserCoinFunction(utils.CoinFunction{
			UserID: getUserID,
			Coin:   CoinsFull + CoinsGet,
			Status: true,
		})
		// DEBUG MODE, USER NOT INSERT DATA
		/*
			utils.ModifyUserAwakenFunction(utils.AwakenFunction{
				UserID:     getUserID,
				AwakenTime: GetUserAwakenDays + 1,
				SignDay:    time.Now().Format("20060102"),
			})
		*/
		utils.ModifyTotalAwakenFunction(utils.TotalCountFunction{
			TimeDay: time.Now().Format("20060102"),
			Count:   GetPeopleCountDays + 1,
		})

		usersRandPic := utils.RandSenderPerDayN(int64(getUserID)+time.Now().Unix(), picDirNum)
		picDirName := picDir[usersRandPic].Name()
		list := reg.FindAllString(picDirName, -1)
		// remove Emoji
		emojiRegex := regexp.MustCompile(`[\x{1F600}-\x{1F64F}|[\x{1F300}-\x{1F5FF}]|[\x{1F680}-\x{1F6FF}]|[\x{1F700}-\x{1F77F}]|[\x{1F780}-\x{1F7FF}]|[\x{1F800}-\x{1F8FF}]|[\x{1F900}-\x{1F9FF}]|[\x{1FA00}-\x{1FA6F}]|[\x{1FA70}-\x{1FAFF}]|[\x{1FB00}-\x{1FBFF}]|[\x{1F170}-\x{1F251}]|[\x{1F300}-\x{1F5FF}]|[\x{1F600}-\x{1F64F}]|[\x{1FC00}-\x{1FCFF}]|[\x{1F004}-\x{1F0CF}]|[\x{1F170}-\x{1F251}]]+`)
		getUserName = emojiRegex.ReplaceAllString(getUserName, "")
		chooseBackground, err := gg.LoadImage(utils.ReturnLucyMainDataIndex("funwork") + "randpic" + "/" + list[0] + ".png")
		if err != nil {
			panic(err)
		}
		getBGContent := gg.NewContextForImage(chooseBackground)
		getVarColorR, getVarColorG, GetVarColorB := utils.GetAverageColorAndMakeAdjust(chooseBackground)
		getFontPath := utils.LoadFontFace(utils.ReturnFontLocation()+"SourceHanSansCN-Regular.otf", 50)
		getBGContent.SetFontFace(getFontPath)
		currentBGColor := color.NRGBA{R: uint8(getVarColorR), G: uint8(getVarColorG), B: uint8(GetVarColorB), A: 255}
		utils.DrawBorderSimple(getBGContent, getGreeting()+" "+getUserName, float64(getBGContent.W()/32), float64(getBGContent.H()/12), currentBGColor, false)
		getFontPathBigger := utils.LoadFontFace(utils.ReturnFontLocation()+"SourceHanSansCN-Regular.otf", 80)
		getBGContent.SetFontFace(getFontPathBigger)
		utils.DrawBorderSimple(getBGContent, "现在是", float64(getBGContent.W()/32), float64(getBGContent.H()/12)+150, currentBGColor, false)
		getCurrentTime := time.Now()
		utils.DrawBorderSimple(getBGContent, NumberToChinese(getCurrentTime.Hour())+"时", float64(getBGContent.W()/32), float64(getBGContent.H()/12)+300, currentBGColor, false)
		utils.DrawBorderSimple(getBGContent, NumberToChinese(getCurrentTime.Minute())+"分", float64(getBGContent.W()/32), float64(getBGContent.H()/12)+450, currentBGColor, false)
		getFontPathSmaller := utils.LoadFontFace(utils.ReturnFontLocation()+"SourceHanSansCN-Regular.otf", 35)
		getBGContent.SetFontFace(getFontPathSmaller)
		utils.DrawBorderSimple(getBGContent, fmt.Sprintf("柠檬片:  ( %d + %d ) w ", CoinsFull, CoinsGet), float64(getBGContent.W()/32), float64(getBGContent.H()/12)+550, currentBGColor, false)
		getBGContent.Fill()

		// Draw Year Process
		getFontPathSmallest := utils.LoadFontFace(utils.ReturnFontLocation()+"SourceHanSansCN-Regular.otf", 20)
		getBGContent.SetFontFace(getFontPathSmallest)
		now := time.Now()
		thisYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		daysPassed := now.Sub(thisYear).Hours() / 24
		isLeapYear := now.Year()%4 == 0 && (now.Year()%100 != 0 || now.Year()%400 == 0)
		var getProgress float64
		if isLeapYear {
			getProgress = daysPassed / 366
		} else {
			getProgress = daysPassed / 365
		}
		var getNeedDraw float64
		getNeedDraw = float64(getBGContent.W()) * getProgress
		getBGContent.SetColor(currentBGColor)
		getBGContent.SetLineWidth(5)
		getBGContent.DrawRectangle(0, float64(getBGContent.H())-50, getNeedDraw, 50)
		getBGContent.Fill()
		colors := currentBGColor
		R, G, B, A := colors.RGBA()
		getBGDARK := utils.IsDark(color.RGBA{
			R: uint8(R),
			G: uint8(G),
			B: uint8(B),
			A: uint8(A),
		})
		if getBGDARK {
			getBGContent.SetColor(color.Black)
		} else {
			getBGContent.SetColor(color.White)
		}
		getBGContent.DrawLine(0, float64(getBGContent.H())-50, getNeedDraw, float64(getBGContent.H())-50)
		getBGContent.Stroke()
		getBGContent.SetColor(currentBGColor)
		if daysPassed > 18 {
			utils.DrawBorderSimple(getBGContent, "今年已经过去了"+strconv.Itoa(int(daysPassed))+"天~", getNeedDraw/3, float64(getBGContent.H())-15, currentBGColor, false)
			getBGContent.Fill()
		}
		getBGContent.Fill()
		// get Sign Tips Table
		getBGContent.SetColor(MixColorWithWhite(color.RGBA(currentBGColor), 0.2))
		getBGContent.DrawRoundedRectangle(float64(getBGContent.W()-450), float64(getBGContent.H())/12, 500, 250, 30)
		getBGContent.FillPreserve()
		getBGContent.SetColor(color.Black)
		getBGContent.SetLineWidth(3)
		getBGContent.Stroke()
		getBGContent.SetFontFace(getFontPathSmallest)
		utils.DrawBorderSimple(getBGContent, "今天你是"+strconv.Itoa(int(GetPeopleCountDays))+"个签到的~", float64(getBGContent.W()-435), float64(getBGContent.H())/12+75, color.Black, false)
		utils.DrawBorderSimple(getBGContent, " qwq? ", float64(getBGContent.W()-435), float64(getBGContent.H())/12+25, color.Black, false)
		utils.DrawBorderSimple(getBGContent, " 一共签到了 "+strconv.Itoa(int(GetUserAwakenDays))+"次~好欸w", float64(getBGContent.W()-435), float64(getBGContent.H())/12+125, color.Black, false)
		utils.DrawBorderSimple(getBGContent, GetHourWord(time.Now()), float64(getBGContent.W()-435), float64(getBGContent.H())/12+175, color.Black, false)
		// get sub image here
		getFontPathSmallestUpper := utils.LoadFontFace(utils.ReturnFontLocation()+"SourceHanSansCN-Regular.otf", 20)
		getBGContent.SetFontFace(getFontPathSmallestUpper)
		getLengthWidth, getLengthHeight := getBGContent.MeasureString("Generated By Lucy  (  HafuuNano & 2 Paradise Ver. ),  Designed By MoeMagicMango. ")
		utils.DrawBorderSimple(getBGContent, "Generated By Lucy  (  HafuuNano & 2 Paradise Ver. ),  Designed By MoeMagicMango. ", float64(getBGContent.W())-80-getLengthWidth, float64(getBGContent.H())-40-getLengthHeight, currentBGColor, false)
		getBGContent.Stroke()
		getBGContent.SaveJPG(utils.ResLoader("sign")+"userpic/"+userPic, 90)
		ctx.SendImage("file:///"+utils.ResLoader("sign")+"userpic/"+userPic, true)
	})
}
