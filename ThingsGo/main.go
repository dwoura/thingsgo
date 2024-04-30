package main

import (
	_ "IOT/initialize/log"

	_ "IOT/modules/dataService"

	_ "IOT/initialize/cache"
	_ "IOT/initialize/psql"

	_ "IOT/initialize/cron"
	_ "IOT/initialize/send_message"
	_ "IOT/initialize/session"
	_ "IOT/initialize/validate"
	_ "IOT/routers"

	services "IOT/services"
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	_ "IOT/cron"
)

var Ticker *time.Ticker

func main() {
	// 读取服务器信息
	Ticker = time.NewTicker(time.Millisecond * 5000)
	go func() {
		for t := range Ticker.C {
			fmt.Println(t)
			var ResourcesService services.ResourcesService
			percent, _ := cpu.Percent(time.Second, false)
			cpu_str := fmt.Sprintf("%.2f", percent[0])
			memInfo, _ := mem.VirtualMemory()
			mem_str := fmt.Sprintf("%.2f", memInfo.UsedPercent)
			currentTime := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))
			ResourcesService.Add(cpu_str, mem_str, currentTime)
		}
	}()
	beego.SetStaticPath("/extensions", "extensions")
	beego.SetStaticPath("/files", "files")
	beego.Run()
}
