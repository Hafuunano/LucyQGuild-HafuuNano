// Package utils ARK Message model, Now work yet.
package utils

import (
	"bytes"
	nano "github.com/fumiama/NanoBot"
	"image"
	"image/png"
)

// MessageArkSender User Should return ARK MESSAGE TYPE.
func MessageArkSender(ctx *nano.Ctx, Tosender bool, ark *nano.MessageArk) {
	ctx.Post(Tosender, &nano.MessagePost{Ark: ark})
}

func SendArkCoverWithText(ctx *nano.Ctx, Tosender bool, image string, title string, content string, url ...string) {
	getArkTemple := nano.MessageArk{
		TemplateID: 37,
		KV: []nano.MessageArkKV{
			{Key: "#METACOVER#", Value: image},
			{Key: "#METATITLE#", Value: title},
			{Key: "#PROMPT#", Value: content},
			{Key: "#METAURL#", Value: url[0]},
		},
	}
	MessageArkSender(ctx, Tosender, &getArkTemple)
}

func SendArkTextWithSmallerImage(ctx *nano.Ctx, toSender bool, image2 string, title string, content string, url ...string) {
	getArkTemple := nano.MessageArk{
		TemplateID: 24,
		KV: []nano.MessageArkKV{
			{Key: "#IMG#", Value: image2},
			{Key: "#TITLE#", Value: title},
			{Key: "#METADESC#", Value: content},
			{Key: "#LINK#", Value: url[0]},
			{Key: "#PROMPT#", Value: ""},
			{Key: "#METADESC#", Value: ""},
			{Key: "#SUBTITLE#", Value: ""},
		},
	}
	MessageArkSender(ctx, toSender, &getArkTemple)
}

func encodePNG(imageBytes *[]byte, img image.Image) error {
	var err error
	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, img)
	if err != nil {
		return err
	}
	*imageBytes = buffer.Bytes()
	return nil
}
