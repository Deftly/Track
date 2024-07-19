package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Example 1: Case-insensitive string comparison
	s1 := []string{"hello", "world"}
	s2 := []string{"HELLO", "WORLD"}
	equal := slices.EqualFunc(s1, s2, func(a, b string) bool {
		return strings.EqualFold(a, b)
	})
	fmt.Printf("Case-insensitive equality: %v\n", equal)

	// Example 2: Custom struct comparison
	p1 := []Person{{"Alice", 30}, {"Bob", 25}}
	p2 := []Person{{"Alice", 30}, {"Bob", 26}}
	equalPerson := slices.EqualFunc(p1, p2, func(a, b Person) bool {
		return a.Name == b.Name // Compare only by name
	})
	fmt.Printf("Person equality by name: %v\n", equalPerson)

	// Example 3: Approximate float comparison
	f1 := []float64{1.0, 2.0, 3.0}
	f2 := []float64{1.00001, 1.99999, 3.00002}
	equalFloat := slices.EqualFunc(f1, f2, func(a, b float64) bool {
		return math.Abs(a-b) < 0.0001
	})
	fmt.Printf("Approximate float equality: %v\n", equalFloat)
}
