package main

import (
	"awesomeProject1/user"
	"awesomeProject1/user/delivery/echo"
	"awesomeProject1/user/repository/sqlite3"
	"flag"
	"github.com/labstack/gommon/log"
)

func main() {
	addr := flag.String("addr", ":8080", "サーバーのアドレス")
	flag.Parse()
	userDelivery := echo.NewUserDelivery(user.NewUserUsecase(sqlite3.NewUserRepository()))
	err := userDelivery.Start(*addr)
	if err != nil {
		log.Fatal("Start:", err)
	}
}