package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func main(){
	var i = 100
	p := &i
	fmt.Println(*p)
	*p = 300
	fmt.Println(i)

	fmt.Println(Vertex{1,2})

	var a [2] string
	a[0] = "Hello"
	a[1] = "Thai"
	fmt.Println(a)
	array()
	testRange()
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	list := strings.Split(s, " ")
	for _,ss := range list {
		ret[ss] += 1
	}
	return ret
}


func testRange(){
	list := [] int {4,7,10,20}
	for i, v := range list {
		fmt.Printf("index = %d, value = %d \n", i, v)
	}
}

func array() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:1]
	printSlice(s)

	// Notice that the high bounds use the capacity instead of length
	s = s[:5]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}