package fortune

import (
	"encoding/json"
	"fmt"
	"github.com/FloatTech/gg"
	"github.com/FloatTech/imgfactory"
	nano "github.com/fumiama/NanoBot"
	"github.com/moyoez/HafuuNano/utils"
	"image/color"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"
)

type card struct {
	Name string `json:"name"`
	Info struct {
		Description        string `json:"description"`
		ReverseDescription string `json:"reverseDescription"`
		ImgURL             string `json:"imgUrl"`
	} `json:"info"`
}

type cardset = map[string]card

var (
	info     string
	cardMap  = make(cardset, 256)
	position = []string{"正位", "逆位"}
	result   map[int64]int
	signTF   map[string]int
)

func init() {
	signTF = make(map[string]int)
	result = make(map[int64]int)
	// onload fortune mapset.
	data, err := os.ReadFile(utils.ReturnLucyMainDataIndex("funwork") + "tarots.json")
	_ = json.Unmarshal(data, &cardMap)
	picDir, err := os.ReadDir(utils.ReturnLucyMainDataIndex("funwork") + "randpic")
	if err != nil {
		return
	}
	picDirNum := len(picDir)
	reg := regexp.MustCompile(`[^.]+`)

	nano.OnMessageCommand("fortune", nano.OnlyToMe).SetBlock(true).Limit(utils.LimitByUser).Handle(func(ctx *nano.Ctx) {
		getUserName := utils.GetUserInfoName(ctx)
		getUserID := ctx.UserID()
		userPic := strconv.FormatUint(getUserID, 10) + time.Now().Format("20060102") + ".png"
		usersRandPic := utils.RandSenderPerDayN(int64(getUserID), picDirNum)
		picDirName := picDir[usersRandPic].Name()
		list := reg.FindAllString(picDirName, -1)
		p := rand.Intn(2)
		is := rand.Intn(77)
		i := is + 1
		card := cardMap[(strconv.Itoa(i))]
		if p == 0 {
			info = card.Info.Description
		} else {
			info = card.Info.ReverseDescription
		}
		userS := strconv.FormatUint(getUserID, 10)
		now := time.Now().Format("20060102")
		// modify this possibility to 40-100, don't be to low.
		randEveryone := utils.RandSenderPerDayN(int64(getUserID), 50)
		var si = now + userS // use map to store.
		loadNotoSans := utils.ReturnLucyMainDataIndex("funwork") + "NotoSansCJKsc-Regular.otf"
		if signTF[si] == 0 {
			result[int64(getUserID)] = randEveryone + 50
			// background
			img, err := gg.LoadImage(utils.ReturnLucyMainDataIndex("funwork") + "randpic" + "/" + list[0] + ".png")
			if err != nil {
				panic(err)
			}
			bgFormat := imgfactory.Limit(img, 1920, 1080)
			getBackGroundMainColorR, getBackGroundMainColorG, getBackGroundMainColorB := utils.GetAverageColorAndMakeAdjust(bgFormat)
			mainContext := gg.NewContext(bgFormat.Bounds().Dx(), bgFormat.Bounds().Dy())
			mainContextWidth := mainContext.Width()
			mainContextHight := mainContext.Height()
			mainContext.DrawImage(bgFormat, 0, 0)
			// draw Round rectangle
			mainContext.SetFontFace(utils.LoadFontFace(loadNotoSans, 50))
			if err != nil {
				_, _ = ctx.SendPlainMessage(false, "Something wrong while rendering pic? font")
				return
			}
			// shade mode || not bugs(
			mainContext.SetLineWidth(4)
			mainContext.SetRGBA255(255, 255, 255, 255)
			mainContext.DrawRoundedRectangle(0, float64(mainContextHight-150), float64(mainContextWidth), 150, 16)
			mainContext.Stroke()
			mainContext.SetRGBA255(255, 224, 216, 215)
			mainContext.DrawRoundedRectangle(0, float64(mainContextHight-150), float64(mainContextWidth), 150, 16)
			mainContext.Fill()
			// draw third round rectangle
			mainContext.SetRGBA255(91, 57, 83, 255)
			mainContext.SetFontFace(utils.LoadFontFace(loadNotoSans, 25))
			charCount := 0.0
			setBreaker := false
			emojiRegex := regexp.MustCompile(`[\x{1F600}-\x{1F64F}|[\x{1F300}-\x{1F5FF}]|[\x{1F680}-\x{1F6FF}]|[\x{1F700}-\x{1F77F}]|[\x{1F780}-\x{1F7FF}]|[\x{1F800}-\x{1F8FF}]|[\x{1F900}-\x{1F9FF}]|[\x{1FA00}-\x{1FA6F}]|[\x{1FA70}-\x{1FAFF}]|[\x{1FB00}-\x{1FBFF}]|[\x{1F170}-\x{1F251}]|[\x{1F300}-\x{1F5FF}]|[\x{1F600}-\x{1F64F}]|[\x{1FC00}-\x{1FCFF}]|[\x{1F004}-\x{1F0CF}]|[\x{1F170}-\x{1F251}]]+`)
			getUserName = emojiRegex.ReplaceAllString(getUserName, "")
			var truncated string
			var UserFloatNum float64
			// set rune count
			for _, runeValue := range getUserName {
				charWidth := utf8.RuneLen(runeValue)
				if charWidth == 3 {
					UserFloatNum = 1.5
				} else {
					UserFloatNum = float64(charWidth)
				}
				if charCount+UserFloatNum > 24 {
					setBreaker = true
					break
				}
				truncated += string(runeValue)
				charCount += UserFloatNum
			}
			if setBreaker {
				getUserName = truncated + "..."
			} else {
				getUserName = truncated
			}
			nameLength, _ := mainContext.MeasureString(getUserName)
			var renderLength float64
			renderLength = nameLength + 160
			if nameLength+160 <= 450 {
				renderLength = 450
			}
			mainContext.DrawRoundedRectangle(50, float64(mainContextHight-175), renderLength, 250, 20)
			mainContext.Fill()
			// avatar draw end.
			avatarFormatRaw := utils.GetUserAvatar(ctx)
			if avatarFormatRaw != nil {
				mainContext.DrawImage(imgfactory.Size(avatarFormatRaw, 100, 100).Circle(0).Image(), 60, int(float64(mainContextHight-150)+25))
			}
			mainContext.SetRGBA255(255, 255, 255, 255)
			mainContext.DrawString("User Info", 60, float64(mainContextHight-150)+10) // basic ui
			mainContext.SetRGBA255(155, 121, 147, 255)
			mainContext.DrawString(getUserName, 180, float64(mainContextHight-150)+50)
			mainContext.DrawString(fmt.Sprintf("今日人品值: %d", randEveryone+50), 180, float64(mainContextHight-150)+100)
			mainContext.Fill()
			// AOSP time and date
			setInlineColor := color.NRGBA{R: uint8(getBackGroundMainColorR), G: uint8(getBackGroundMainColorG), B: uint8(getBackGroundMainColorB), A: 255}
			if err != nil {
				_, _ = ctx.SendPlainMessage(false, "ERROR: 渲染时发生了一点错误")
				return
			}
			formatTimeDate := time.Now().Format("2006 / 01 / 02")
			formatTimeCurrent := time.Now().Format("15 : 04 : 05")
			formatTimeWeek := time.Now().Weekday().String()
			mainContext.SetFontFace(utils.LoadFontFace(loadNotoSans, 35))
			setOutlineColor := color.White
			utils.DrawBorderString(mainContext, formatTimeCurrent, 5, float64(mainContextWidth-80), 50, 1, 0.5, setInlineColor, setOutlineColor)
			utils.DrawBorderString(mainContext, formatTimeDate, 5, float64(mainContextWidth-80), 100, 1, 0.5, setInlineColor, setOutlineColor)
			utils.DrawBorderString(mainContext, formatTimeWeek, 5, float64(mainContextWidth-80), 150, 1, 0.5, setInlineColor, setOutlineColor)
			mainContext.FillPreserve()
			if err != nil {
				return
			}
			mainContext.SetFontFace(utils.LoadFontFace(loadNotoSans, 140))
			utils.DrawBorderString(mainContext, "|", 5, float64(mainContextWidth-30), 65, 1, 0.5, setInlineColor, setOutlineColor)
			// throw tarot card
			mainContext.SetFontFace(utils.LoadFontFace(loadNotoSans, 20))
			if err != nil {
				_, _ = ctx.SendPlainMessage(false, "Something wrong while rendering pic?")
				return
			}
			mainContext.SetRGBA255(91, 57, 83, 255)
			mainContext.DrawRoundedRectangle(float64(mainContextWidth-300), float64(mainContextHight-350), 450, 300, 20)
			mainContext.Fill()
			mainContext.SetRGBA255(255, 255, 255, 255)
			mainContext.SetLineWidth(3)
			mainContext.DrawString("今日塔罗牌", float64(mainContextWidth-300)+10, float64(mainContextHight-350)+30)
			mainContext.SetRGBA255(155, 121, 147, 255)
			mainContext.DrawString(card.Name, float64(mainContextWidth-300)+10, float64(mainContextHight-350)+60)
			mainContext.DrawString(fmt.Sprintf("- %s", position[p]), float64(mainContextWidth-300)+10, float64(mainContextHight-350)+280)
			placedList := utils.SplitChineseString(info, 44)
			for ist, v := range placedList {
				mainContext.DrawString(v, float64(mainContextWidth-300)+10, float64(mainContextHight-350)+90+float64(ist*30))
			}
			// output
			mainContext.SetFontFace(utils.LoadFontFace(loadNotoSans, 20))
			mainContext.SetRGBA255(186, 163, 157, 255)
			mainContext.DrawStringAnchored("Generated By Lucy ( HafuuNano & 2 Paradise Ver. ), Design By MoeMagicMango", float64(mainContextWidth-15), float64(mainContextHight-30), 1, 1)
			mainContext.Fill()
			_ = mainContext.SavePNG(utils.ResLoader("fortune") + "jrrp/" + userPic)
			ctx.SendImage("file:///"+utils.ResLoader("fortune")+"jrrp/"+userPic, true, "")
			signTF[si] = 1
		} else {
			_, _ = ctx.SendImage("file:///"+(utils.ResLoader("fortune")+"jrrp/"+userPic), true, "今天已经测试过了哦w")
		}
	})
}
