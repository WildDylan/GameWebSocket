package main

import (
	. "banner"
	. "configuration"
	. "scheduler"
	"time"
	. "web"
)

func main() {
	LoadConfig()
	LoadBanner()

	InitScheduler(1 * time.Second)

	StartWebServer()
	StartWebSocketServer()
}
