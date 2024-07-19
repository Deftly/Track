package main

import "fmt"

func main() {
	i := 1
	// while loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// usage of range
	for i := range 3 {
		fmt.Println("range", i)
	}

	// infinite loop with break
	for {
		fmt.Println("loop")
		break
	}

	// range with continue
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
