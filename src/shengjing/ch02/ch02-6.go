package main

import (
	"fmt"
	"sort"
)

func main(){
	ages := map[string]int{
    "alice":   31,
    "charlie": 34,
	}
	fmt.Println(ages)
	fmt.Println(ages["alice"])
	delete(ages, "alice")
	fmt.Println(ages)
	for name,age :=range ages{
		fmt.Printf("%s\t%d\n",name,age)
	}

	var names []string
	for name := range ages{
		names = append(names,name)
	}
	sort.Strings(names)
	for _,name := range  names{
		fmt.Printf("%s\t%d\n",name,ages[name])
	}

	var resultV1 bool = equal(map[string]int{"A": 0},
		map[string]int{"A": 0})
	fmt.Println(resultV1)
}

func equal(x,y map[string]int) bool{
	if len(x)!=len(y){
		return false
	}
	for k,xv :=range  x{
		if yv,ok := y[k];!ok||yv!=xv{
			return false
		}
	}
	return true
}
