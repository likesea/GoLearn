package main

import (
	"fmt"
	"goconfigsrc"
)

func main() {
	value, _ := goconfig.LoadConfigFile("conf.ini")
	fmt.Println(value.GetValue("Demo", "key:1"))
}
