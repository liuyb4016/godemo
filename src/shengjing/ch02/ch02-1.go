package main

import (
	"fmt"
	"flag"
	"strings"
)

func incr(p *int) int{
	*p++
	return *p
}

var n = flag.Bool("n",false,"omit trailing new line")
var sep = flag.String("s"," ","separator")
func main(){
	v := 1
	fmt.Println(v)
	v2 := incr(&v)
	fmt.Println(v2)
	fmt.Println(n)
	fmt.Println(sep)

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n{
		fmt.Println()
	}

	p:=new(int)
	fmt.Println(*p)
	fmt.Println(p)
	*p = 2
	fmt.Println(*p)
	fmt.Println(p)


}


