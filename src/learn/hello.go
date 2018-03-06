package main 

import (
	"fmt"
	"learn/stringutil"
	"math/rand"
	"math"
)

func add(a int, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func main(){
  fmt.Println(stringutil.Reverse("Hello world!"))
  fmt.Println("random number: " , rand.Intn(20))
  fmt.Printf("Test %g number.\n" , math.Sqrt(9))
  fmt.Println(add(5,7))
  fmt.Println(swap("Thai", "Nguyen"))

  var a, b, c = true, false, "Yes"
  fmt.Println(a,b,c)

  var d float32 = 10.5;
  fmt.Printf("%T, %T \n", d, int(d))
  fmt.Println(d, int(d))
}
