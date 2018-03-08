package main

import (
	"time"
	"fmt"
)

func say(s string) {
	for i := 0; i< 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main(){
	go say("worldddddddddddddd")
	say("hello")


	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
	channelBuffer()


}

func channelBuffer(){
	ch := make(chan int,2 )
	go deadLock(ch, 1)
	go deadLock(ch, 2)
}

func deadLock(ch chan int, k int){
	ch <- k
	fmt.Println(<-ch)
}