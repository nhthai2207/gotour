package main

import (
	"fmt"
	"time"
)

func main(){
	defer fmt.Println("In the end")
	sum := 0
	for i := 0; i < 10 ; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Tuesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
