package test

import "fmt"

func Factorial(x int) (result int){
	if x==0{
		result = 1
	}else{
		result = x * Factorial(x-1)
 	}
	return;
}

func fibonaci(n int) int{
	if n<2{
		return n
	}
	return fibonaci(n-2)+fibonaci(n-1)
}

func TestFibonaci(){
	var i int
	 for i = 0; i < 5; i++ {
       fmt.Printf("%d\t", fibonaci(i))
    }
}