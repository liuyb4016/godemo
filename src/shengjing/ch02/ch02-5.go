package main

import "fmt"

func reverse(s []int){
	 for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i],s[j] = s[j],s[i]
	}
}

func main(){
	a := [...]int{0,1,2,3,4,5,6}

	reverse(a[:2])
	reverse(a[2:])
	reverse(a[:])
	fmt.Println(a)
	appendstring()
	appendIntV2()
}

func appendstring(){
	var runes []rune //针对中文的切片
	for _,r := range "Hello,世界"{
		runes = append(runes,r)
	}
	fmt.Printf("%q\n",runes)
}

func appendInt(x []int,y int) []int {
	var z []int
	zlen :=len(x)+1
	if zlen>cap(x){
		z = x[:zlen]
	}else{
		zcap := zlen
		if zcap<2*len(x){
			zcap=2*len(x)
		}
		z = make([]int,zlen,zcap)
		copy(z,x)
	}
	z[len(x)] = y
	return z
}
func appendIntV2(){
	var x, y []int
    for i := 0; i < 10; i++ {
        y = appendInt(x, i)
        fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
        x = y
    }
}