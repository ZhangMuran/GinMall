package main

import (
	"fmt"

	"github.com/ZhangMuran/GinMall/conf"
	"github.com/ZhangMuran/GinMall/routes"
)

func main() {
	fmt.Println("hello world")
	loadInfo()
	r := routes.SetupRouter()
	if err := r.Run(conf.HttpPort); err != nil {
		panic(err)
	}

}

func loadInfo() {
	conf.Init()
}