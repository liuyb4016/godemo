package test

import "fmt"

var list1 = [5]float32{3,4,6,7,8}
var list2 = [...]float32{3,4,6,7,8,6,7,8,9,3}

var sv1 float32 = list2[3]

func Testlist(){
	var nv1 [10]int
	var i,j int

	for i=0;i<10;i++{
		nv1[i] = i+100
	}

	for j=0;j<10;j++{
		fmt.Printf("Element[%d] = %d\n", j, nv1[j] )
	}
}
