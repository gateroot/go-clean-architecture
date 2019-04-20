package main

import (
	"awesomeProject1/di"
	"flag"

	"github.com/labstack/gommon/log"
)

func main() {
	addr := flag.String("addr", ":8080", "サーバーのアドレス")
	flag.Parse()
	userDelivery := di.InjectUserDelivery()
	err := userDelivery.Start(*addr)
	if err != nil {
		log.Fatal("Start:", err)
	}
}