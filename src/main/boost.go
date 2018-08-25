package main

import (
	. "configuration"
	. "banner"
	. "scheduler"
	. "web"
	"time"
)

func main() {
	LoadConfig()
	LoadBanner()

	InitScheduler(1 * time.Second)

	StartWebServer()
	StartWebSocketServer()
}
