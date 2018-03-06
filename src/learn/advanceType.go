package main

import "fmt"

func main(){
	var i = 100
	p := &i
	fmt.Println(*p)
	*p = 300
	fmt.Println(i)

}