package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net/http"
	"sscmgroup.com/app/common"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func ServerInfo(c *gin.Context) {
	info, _ := load.Avg()
	fmt.Printf("ServerInfo load info %v", info)
	dis, _ := disk.Usage("/")
	diskTotalGB := int(dis.Total) / GB
	diskFreeGB := int(dis.Free) / GB
	diskDic := make(map[string]interface{}, 0)
	diskDic["total"] = diskTotalGB
	diskDic["free"] = diskFreeGB

	mem, _ := mem.VirtualMemory()
	memUsedMB := int(mem.Used) / GB
	memTotalMB := int(mem.Total) / GB
	memFreeMB := int(mem.Free) / GB
	memUsedPercent := int(mem.UsedPercent)
	memDic := make(map[string]interface{}, 0)
	memDic["total"] = memTotalMB
	memDic["used"] = memUsedMB
	memDic["free"] = memFreeMB
	memDic["usage"] = memUsedPercent

	cpuDic := make(map[string]interface{}, 0)
	cpuDic["cpuInfo"], _ = cpu.Info()
	percent, _ := cpu.Percent(0, false)
	cpuDic["Percent"] = common.Round(percent[0], 2)
	cpuDic["cpuNum"], _ = cpu.Counts(false)

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]interface{}{
			"mem":  memDic,
			"disk": diskDic,
			"cpu":  cpuDic,
		},
		"msg": "ok",
	})
	return
}
