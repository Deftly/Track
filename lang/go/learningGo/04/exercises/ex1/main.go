package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
}
