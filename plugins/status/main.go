package status

import (
	"fmt"
	"github.com/moyoez/HafuuNano/utils"
	"math"
	"time"

	nano "github.com/fumiama/NanoBot"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func init() {
	nano.OnMessageCommand("status", nano.OnlyToMe).SetBlock(true).Limit(utils.LimitByUser).Handle(func(ctx *nano.Ctx) {
		ctx.SendPlainMessage(false, "* Hosted On ?(HiMoYo ServerVPS).\n",
			"* CPU Usage: ", cpuPercent(), "%\n",
			"* RAM Usage: ", memPercent(), "%\n",
			"* DiskInfo Usage Check: ", diskPercent(), "\n",
			"  Lucyは、高性能ですから！")
	})
	nano.OnMessageCommand("help", nano.OnlyToMe).SetBlock(true).Limit(utils.LimitByUser).Handle(func(ctx *nano.Ctx) {
		// TODO : 画个图吧受不了傻逼腾讯了
		//	ctx.SendPlainMessage(true,"")
	})

}

func cpuPercent() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return -1
	}
	return math.Round(percent[0])
}

func memPercent() float64 {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return -1
	}
	return math.Round(memInfo.UsedPercent)
}

func diskPercent() string {
	parts, err := disk.Partitions(true)
	if err != nil {
		return err.Error()
	}
	msg := ""
	for _, p := range parts {
		diskInfo, err := disk.Usage(p.Mountpoint)
		if err != nil {
			msg += "\n  - " + err.Error()
			continue
		}
		pc := uint(math.Round(diskInfo.UsedPercent))
		if pc > 0 {
			msg += fmt.Sprintf("\n  - %s(%dM) %d%%", p.Mountpoint, diskInfo.Total/1024/1024, pc)
		}
	}
	return msg
}
