package main

import (
	"fmt"
	"crypto/sha256"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main(){
	symbol := [...]string{USD:"$",EUR:"€",GBP:"￡",RMB:"￥"}
	fmt.Println(USD,symbol[USD])

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(&a[1])
	fmt.Println(&b)
	fmt.Println(&c)

	fmt.Println(a == b, a == c, b == c) // "true false false"



	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Println(c1,"-------->",&c1)
	fmt.Println(c2,"-------->",&c2)
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}