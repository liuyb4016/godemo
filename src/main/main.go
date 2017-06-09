package main

import (
	"fmt"
	"test"
	"time"
)
func main(){
	fmt.Println("Hello,World!!!!")
	var addval int = test.Add(10,300)
	fmt.Println("加法结果：",addval)
	var mutival int = test.Muti(10,300)
	fmt.Println("乘法结果：",mutival)
	fmt.Println(test.DEMO_STR)
	test.Dotestif("C",93)
	fmt.Println(time.Now().Weekday())
	test.Switchtype(true)
	test.Selecttest()
	test.Xunhuantest()
	test.Sushutest(1000)
}




