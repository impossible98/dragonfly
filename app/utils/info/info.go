package info

import (
	// import built-in packages
	"fmt"
	"strconv"
	"time"
	// import local packages
	"dragonfly/app/config"
	"dragonfly/app/utils/color"
	"dragonfly/app/utils/ip"
	"dragonfly/global"
)

func ShowInfo(startTime int64) {
	fmt.Print("\033[H\033[2J")
	endTime := time.Now().UnixNano()
	difference := (endTime - startTime) / 1000 / 1000
	// control flow
	if difference == 0 {
		fmt.Printf("  %s %s  %s %s ms\n", color.Bold(color.Green(global.APP_NAME)), color.Green("v"+global.VERSION), color.Dim("ready in"), color.Bold(strconv.Itoa(1)))
	} else {
		fmt.Printf("  %s %s  %s %s ms\n", color.Bold(color.Green(global.APP_NAME)), color.Green("v"+global.VERSION), color.Dim("ready in"), color.Bold(strconv.Itoa(int(difference))))
	}
	fmt.Println()
	fmt.Printf("  %s  %s   %s\n", color.Green("➜"), color.Bold("Local:"), color.Cyan("http://127.0.0.1:"+strconv.Itoa(config.ServerPort)+"/"))
	fmt.Printf("  %s  %s %s\n", color.Green("➜"), color.Bold("Network:"), color.Cyan("http://"+ip.Ip()+":"+strconv.Itoa(config.ServerPort)+"/"))
	fmt.Println()
}
