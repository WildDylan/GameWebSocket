package main

import (
	. "banner"
	. "configuration"
	. "socket"
)

func main() {
	LoadConfig()
	LoadBanner()

	Start()
}
