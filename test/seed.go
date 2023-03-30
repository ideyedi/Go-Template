package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rate int64 = 
	rand.Seed(rate.UnixNano())
	fmt.Println("integer : " + rand.intn(50))
	fmt.Println("float : " + rand.Float64())
}
