package main

import (
	"fmt"
)

type T struct {
	name string //	name of the object
	val  int    //	integer value
}

func main() {
	//var arr [5]int = {1, 2, 3, 4, 5}
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr : %v", arr)

	for v, num := range arr {
		fmt.Println(v, num)
	}

	t := T{
		name: "EunSeok Ji",
		val:  31,
	}

	fmt.Println(t.name)
}
