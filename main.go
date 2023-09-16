package main

import (
	"fmt"

	"github.com/ZhangMuran/GinMall/conf"
)

func main() {
	fmt.Println("hello world")
	loadInfo()

}

func loadInfo() {
	conf.Init()
}