package main

import (
	// import built-in packages
	"time"
	// import local packages
	"dragonfly/app/server"
	"dragonfly/app/utils/info"
)

func main() {
	startTime := time.Now().UnixNano()
	info.ShowInfo(startTime)
	server.SetupServer()
}
