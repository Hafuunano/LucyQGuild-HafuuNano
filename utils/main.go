package utils

import (
	"bytes"
	"fmt"
	"github.com/FloatTech/floatbox/binary"
	"github.com/FloatTech/floatbox/file"
	"github.com/FloatTech/floatbox/web"
	"github.com/FloatTech/gg"
	"github.com/FloatTech/imgfactory"
	nano "github.com/fumiama/NanoBot"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"hash/crc64"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
	"unicode/utf8"
	"unsafe"
)

var loaderHere = file.Pwd() + "/data/res/"

// ReturnFontLocation Return Font Location.
func ReturnFontLocation() string {
	return loaderHere + "fonts/"
}

// ResLoader Only refer To loader Path here, no reader.
func ResLoader(plugin string) string {
	if _, err := os.Stat(loaderHere + plugin); os.IsNotExist(err) {
		err := os.MkdirAll(loaderHere+plugin+"/", 0755)
		if err != nil {
			panic(err)
		}
	}
	return loaderHere + plugin + "/"
}

// ResReader Read refer Files here, return status and bytes
func ResReader(plugin string, insertDataFile string) (err bool, dataByte []byte) {
	if _, err := os.Stat(loaderHere + plugin); os.IsNotExist(err) {
		err := os.MkdirAll(loaderHere+plugin+"/", 0755)
		if err != nil {
			panic(err)
		}
	} else {
		getStat, err := os.ReadFile(loaderHere + plugin + insertDataFile)
		if err != nil {
			return true, nil
		}
		return false, getStat
	}
	return false, nil
}

// GetUserInfoName Get Info Name, for it can satisfy QQ Onebot Code here.
func GetUserInfoName(ctx *nano.Ctx) string {
	getMember, err := ctx.GetGuildMemberOf(ctx.Message.GuildID, strconv.FormatUint(ctx.UserID(), 10))
	if err != nil {
		panic(err)
	}
	return getMember.User.Username
}

// GetUserAvatar Get User's Avatar from QGuild, Return Raw Image.
func GetUserAvatar(ctx *nano.Ctx) image.Image {
	getMember, err := ctx.GetGuildMemberOf(ctx.Message.GuildID, strconv.FormatUint(ctx.UserID(), 10))
	if err != nil {
		panic(err)
	}
	fmt.Print("\n" + getMember.User.Avatar)
	getData, err := web.GetData(getMember.User.Avatar)
	if err != nil {
		panic(err)
	}
	getdataImage := bytes.NewReader(getData)
	getDataImage, _, _ := image.Decode(getdataImage)
	return getDataImage
}

// GetBotSelfImage Get User's Avatar from QGuild, Return Raw Image.
func GetBotSelfImage(ctx *nano.Ctx) image.Image {
	botInfo, err := ctx.GetMyInfo()
	if err != nil {
		panic(err)
	}
	getData, err := web.GetData(botInfo.Avatar)
	if err != nil {
		panic(err)
	}
	getdataImage := bytes.NewReader(getData)
	getDataImage, _, _ := image.Decode(getdataImage)
	return getDataImage
}

// ReturnLucyMainDataIndex Lucy's main Path here.
func ReturnLucyMainDataIndex(pluginName string) string {
	return "/root/Lucy_Project/main/Lucy_zerobot/data/" + pluginName + "/"
}

// RandSenderPerDayN 每个用户每天随机数
func RandSenderPerDayN(uid int64, n int) int {
	sum := crc64.New(crc64.MakeTable(crc64.ISO))
	_, _ = sum.Write(binary.StringToBytes(time.Now().Format("20060102")))
	_, _ = sum.Write((*[8]byte)(unsafe.Pointer(&uid))[:])
	r := rand.New(rand.NewSource(int64(sum.Sum64())))
	return r.Intn(n)
}

// DRAW USAGE, For Render Like Maimai.

// SaveImage Save Cover Chun | Maimai
func SaveImage(img image.Image, path string) {
	files, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(files *os.File) {
		err := files.Close()
		if err != nil {
			return
		}
	}(files)
	err = png.Encode(files, img)
	if err != nil {
		log.Fatal(err)
	}
}

// DownloadImage Simple Downloader.
func DownloadImage(url string) (image.Image, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// LoadPictureWithResize Load Picture
func LoadPictureWithResize(link string, w int, h int) image.Image {
	getImage, err := gg.LoadImage(link)
	if err != nil {
		return nil
	}
	return Resize(getImage, w, h)
}

// Resize Image width height
func Resize(image image.Image, w int, h int) image.Image {
	return imgfactory.Size(image, w, h).Image()
}

// GetAverageColorAndMakeAdjust different from k-means algorithm,it uses origin plugin's algorithm.(Reduce the cost of average color usage.)
func GetAverageColorAndMakeAdjust(image image.Image) (int, int, int) {
	var RList []int
	var GList []int
	var BList []int
	width, height := image.Bounds().Size().X, image.Bounds().Size().Y
	// use the center of the bg, to make it more quickly and save memory and usage.
	for x := int(math.Round(float64(width) / 1.5)); x < int(math.Round(float64(width))); x++ {
		for y := height / 10; y < height/2; y++ {
			r, g, b, _ := image.At(x, y).RGBA()
			RList = append(RList, int(r>>8))
			GList = append(GList, int(g>>8))
			BList = append(BList, int(b>>8))
		}
	}
	RAverage := int(Average(RList))
	GAverage := int(Average(GList))
	BAverage := int(Average(BList))
	return RAverage, GAverage, BAverage
}

// Average sum all the numbers and divide by the length of the list.
func Average(numbers []int) float64 {
	var sum float64
	for _, num := range numbers {
		sum += float64(num)
	}
	return math.Round(sum / float64(len(numbers)))
}

// DrawBorderString GG Package Not support The string render, so I write this (^^)
func DrawBorderString(page *gg.Context, s string, size int, x float64, y float64, ax float64, ay float64, inlineRGB color.Color, outlineRGB color.Color) {
	page.SetColor(outlineRGB)
	n := size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				continue
			}
			renderX := x + float64(dx)
			renderY := y + float64(dy)
			page.DrawStringAnchored(s, renderX, renderY, ax, ay)
		}
	}
	page.SetColor(inlineRGB)
	page.DrawStringAnchored(s, x, y, ax, ay)
}

// SplitChineseString Split Chinese type chart, Used In maimai | CHUN Render.
func SplitChineseString(s string, length int) []string {
	results := make([]string, 0)
	runes := []rune(s)
	start := 0
	for i := 0; i < len(runes); i++ {
		size := utf8.RuneLen(runes[i])
		if start+size > length {
			results = append(results, string(runes[0:i]))
			runes = runes[i:]
			i, start = 0, 0
		}
		start += size
	}
	if len(runes) > 0 {
		results = append(results, string(runes))
	}
	return results
}

// IsDark judge which font color I prefer to use (black or white)
func IsDark(rgb color.RGBA) bool {
	var (
		r = float64(rgb.R) * 0.299
		g = float64(rgb.G) * 0.587
		b = float64(rgb.B) * 0.114
	)
	lum := r + g + b
	return lum < 192
}

// DrawBorderSimple Simple Usage of draw Border
func DrawBorderSimple(page *gg.Context, s string, x float64, y float64, inlineRGB color.Color, reverse bool) {
	colors := inlineRGB
	R, G, B, A := colors.RGBA()
	if IsDark(color.RGBA{
		R: uint8(R),
		G: uint8(G),
		B: uint8(B),
		A: uint8(A),
	}) {
		DrawBorderString(page, s, 5, x, y, 0, 0, inlineRGB, color.White)
	} else {
		if reverse {
			DrawBorderString(page, s, 5, x, y, 0, 0, inlineRGB, color.White)
		} else {
			DrawBorderString(page, s, 5, x, y, 0, 0, inlineRGB, color.Black)
		}
	}
}

// LoadFontFace load font face once before running, to work it quickly and save memory.
func LoadFontFace(filePath string, size float64) font.Face {
	fontFile, _ := os.ReadFile(filePath)
	fontFileParse, _ := opentype.Parse(fontFile)
	fontFace, _ := opentype.NewFace(fontFileParse, &opentype.FaceOptions{Size: size, DPI: 70, Hinting: font.HintingFull})
	return fontFace
}

// BreakWords Reduce the length of strings, if out of range, use ".." instead.
func BreakWords(getSongName string, breakerCount float64) string {
	charCount := 0.0
	setBreaker := false
	var truncated string
	var charFloatNum float64
	for _, runeValue := range getSongName {
		charWidth := utf8.RuneLen(runeValue)
		if charWidth == 3 {
			charFloatNum = 2
		} else {
			charFloatNum = float64(charWidth)
		}
		if charCount+charFloatNum > breakerCount {
			setBreaker = true
			break
		}
		truncated += string(runeValue)
		charCount += charFloatNum
	}
	if setBreaker {
		getSongName = truncated + ".."
	} else {
		getSongName = truncated
	}
	return getSongName
}
