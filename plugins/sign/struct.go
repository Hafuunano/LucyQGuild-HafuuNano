package sign

import (
	"image/color"
	"strings"
	"time"
)

// used For Time Switcher.
var chineseDigits = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
var chineseTens = []string{"", "十", "二十", "三十", "四十", "五十", "六十", "七十", "八十", "九十"}
var chineseHundreds = []string{"", "一百", "二百", "三百", "四百", "五百", "六百", "七百", "八百", "九百"}

func getGreeting() string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	currentTime := time.Now().In(loc)
	currentHour := currentTime.Hour()
	var greeting string
	if currentHour >= 6 && currentHour < 12 {
		greeting = "早上好ya~"
	} else if currentHour >= 12 && currentHour < 18 {
		greeting = "下午好ww"
	} else {
		greeting = "晚上好~"
	}

	return greeting
}

// NumberToChinese To make it AOSP Style.
func NumberToChinese(number int) string {
	var result strings.Builder
	if number >= 100 {
		result.WriteString(chineseHundreds[number/100])
		number %= 100
	}
	if number >= 10 {
		if number == 10 {
			result.WriteString(chineseTens[1])
		} else {
			result.WriteString(chineseTens[number/10])
			number %= 10
		}
	}
	if number > 0 {
		result.WriteString(chineseDigits[number])
	}
	return result.String()
}

func MixColorWithWhite(colorToMix color.RGBA, weight float64) color.RGBA {
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	mixedColor := color.RGBA{
		R: uint8(float64(white.R)*(1-weight) + float64(colorToMix.R)*weight),
		G: uint8(float64(white.G)*(1-weight) + float64(colorToMix.G)*weight),
		B: uint8(float64(white.B)*(1-weight) + float64(colorToMix.B)*weight),
		A: uint8(float64(white.A)*(1-weight) + float64(colorToMix.A)*weight),
	}
	return mixedColor
}

func GetHourWord(t time.Time) (reply string) {
	switch {
	case 5 <= t.Hour() && t.Hour() < 12:
		reply = "起的很早嘛w! 是不懒睡觉的好孩子w"
	case 12 <= t.Hour() && t.Hour() < 14:
		reply = "吃饭了嘛w~如果没有快去吃饭哦w"
	case 14 <= t.Hour() && t.Hour() < 19:
		reply = "下午的话~记得多去补点水呢~ww,辛苦了哦w"
	case 19 <= t.Hour() && t.Hour() < 24:
		reply = "今天过的开心嘛ww ╰(￣ω￣ｏ)w"
	case 0 <= t.Hour() && t.Hour() < 5:
		reply = "快去睡觉~已经很晚了w"
	default:
	}
	return
}
