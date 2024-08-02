package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(100))
		fmt.Print(arr[i], " ")
		switch v := arr[i]; {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
