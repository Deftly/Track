package main

import "fmt"

func main() {
	// The type of elements and length are both part of the array's type.
	// By default an array is zero-valued.
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get", a[4])

	fmt.Println("len:", len(a))

	// This syntax declares and initializes an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// You can have the compiler count the number of elements with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// I fyou specify the index with :, the elements in between will be zeroed.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Arrays are one dimensional but you can compose them together to build
	// multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[0]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d:", twoD)
}
