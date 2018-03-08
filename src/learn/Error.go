package main

import (
	"math"
	"fmt"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot do sqrt for %f", err)
}

func Sqrt(a ErrNegativeSqrt) (float64, error) {
	if(a >= 0){
		return math.Sqrt(float64(a)), nil
	}
	return 0, a
}

func main(){
	if k, err := Sqrt(-100); err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(k)
	}
}