package test

import "fmt"

func Dotestif(grade string, marks int) {

	switch {
	case marks >= 90 && marks <= 100:
		grade = "A"
	case marks >= 80 && marks < 90:
		grade = "B"
	case marks >= 60 && marks < 79:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n");
	}
	fmt.Printf("你的等级是 %s\n", grade);
}

func Switchtype(x interface{}) {

	switch i := x.(type) {
	case nil:
		fmt.Println(" x 的类型 :%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool:
		fmt.Println("x 是 bool 型")
	case string:
		fmt.Println("x 是 string 型")
	default:
		fmt.Println("未知型")
	}
}

var c1, c2, c3 chan int

func Selecttest() {
	c1 = make(chan  int)
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case i2 = <-c2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

func Xunhuantest(){
	var b int = 15
	var a int

	numbers :=[6]int{1,3,5,6,8}

	for a := 0; a<10 ; a++{
		fmt.Printf("a的值为：%d\n",a)
	}

	for a<b{
		a++
		fmt.Printf("a的值为：%d\n",a)
	}

	for i,x:=range  numbers{
		fmt.Printf("第 %d 位的值：%d\n",i,x)
	}
}

func Sushutest(num1 int){
	var i,j int
	for i=2;i<num1;i++{
		for j=2;j<=(i/j);j++{
			if(i%j==0){
				break;
			}
		}
		if (j>(i/j)){
			fmt.Printf("%d 是素数\n",i)
		}
	}
}