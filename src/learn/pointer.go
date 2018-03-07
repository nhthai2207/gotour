package main

import (
	"math"
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}


func (t Vertex) sqrt() float64 {
	return math.Sqrt(float64(t.X * t.X + t.Y * t.Y))
}

// toString function
func (t Vertex) String () string {
	return fmt.Sprintf("X = %d, Y = %d", t.X, t.Y)
}


func main()  {
	v := Vertex{3,4}
	fmt.Println(v.sqrt())
	fmt.Println(v)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	ss := make([]string, 4)
	for i, b := range ip {
		ss[i] = string(b)
	}
	return strings.Join(ss, ".")
}

