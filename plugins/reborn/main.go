package reborn

import (
	"encoding/json"
	"fmt"
	nano "github.com/fumiama/NanoBot"
	"github.com/moyoez/HafuuNano/utils"
	_ "github.com/moyoez/HafuuNano/utils"
	wr "github.com/mroth/weightedrand"
	"math/rand"
	"os"
)

var (
	areac     *wr.Chooser
	gender, _ = wr.NewChooser(
		wr.Choice{Item: "男孩子", Weight: 33707},
		wr.Choice{Item: "女孩子", Weight: 39292},
		wr.Choice{Item: "雌雄同体", Weight: 1001},
		wr.Choice{Item: "猫猫!", Weight: 10000},
		wr.Choice{Item: "狗狗!", Weight: 10000},
		wr.Choice{Item: "🐉~", Weight: 3000},
		wr.Choice{Item: "龙猫~", Weight: 3000},
	)
)

type ratego []struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

func init() {
	area := make(ratego, 226)
	err := load(&area, utils.ReturnLucyMainDataIndex("funwork")+"ratego.json")
	if err != nil {
		panic(err)
	}
	choices := make([]wr.Choice, len(area))
	for i, a := range area {
		choices[i].Item = a.Name
		choices[i].Weight = uint(a.Weight * 1e9)
	}
	areac, err = wr.NewChooser(choices...)
	if err != nil {
		panic(err)
	}
	nano.OnMessageCommand("reborn", nano.OnlyToMe).SetBlock(true).Limit(utils.LimitByUser).Handle(func(ctx *nano.Ctx) {
		if rand.Int31() > 1<<27 {
			ctx.SendPlainMessage(true, fmt.Sprintf("投胎成功！\n您出生在 %s, 是 %s。", randcoun(), randgen()))
		} else {
			ctx.SendPlainMessage(true, "投胎失败！\n您没能活到出生，希望下次运气好一点呢~！")
		}
	})
}

// load 加载rate数据
func load(area *ratego, jsonfile string) error {
	data, err := os.ReadFile(jsonfile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, area)
}

func randcoun() string {
	return areac.Pick().(string)
}

func randgen() string {
	return gender.Pick().(string)
}
