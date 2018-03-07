package main

import "fmt"

func compute(fn func(float32, float32) float32) int{
	return int(fn(3.2,4.3))
}

func fibonacci() func() int {
	x0 := 0
	x1 := 1
	return func() int{
		x0, x1 = x1, x0 + x1
		return x1
	}
}


func main(){
	testfn := func(x, y float32) float32 {
		return x*y
	}

	fmt.Println(testfn(3,4))
	fmt.Println(compute(testfn))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
