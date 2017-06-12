package main

import (
	"fmt"
	"test"
	"time"
)
func main() {
	fmt.Println("Hello,World!!!!")
	var addval int = test.Add(10, 300)
	fmt.Println("加法结果：", addval)
	var mutival int = test.Muti(10, 300)
	fmt.Println("乘法结果：", mutival)
	fmt.Println(test.DEMO_STR)
	test.Dotestif("C", 93)
	fmt.Println(time.Now().Weekday())
	test.Switchtype(true)
	test.Selecttest()
	test.Xunhuantest()
	test.Sushutest(1000)
	test.Testlist()
	test.TestPointerV1()
	test.TestPointerV2()
	test.TestPointerV3()
	test.TestPointerV4()
	test.TestStructV1()
	test.TestSlicelV1()
	test.TestRange()
	test.TestMapV1()
	test.TestMapV2()
	fmt.Printf("%d 的阶乘是 %d\n", 15, test.Factorial(15))
	test.TestFibonaci()
	test.TestInterface()
	test.TestError()
	test.Testosargs()
}




