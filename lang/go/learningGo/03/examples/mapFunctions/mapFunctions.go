package main

import (
	"fmt"
	"maps"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Example 1: Basic map comparison with maps.Equal
	m1 := map[string]int{
		"hello": 5,
		"world": 10,
	}

	m2 := map[string]int{
		"world": 10,
		"hello": 5,
	}

	m3 := map[string]int{
		"hello": 5,
		"world": 11,
	}

	fmt.Printf("m1 equal to m2: %v\n", maps.Equal(m1, m2))
	fmt.Printf("m1 equal to m3: %v\n", maps.Equal(m1, m3))

	// Example 2: Custom struct comparison with maps.EqualFunc
	p1 := map[string]Person{
		"employee1": {"Alice", 30},
		"employee2": {"Bob", 25},
	}
	p2 := map[string]Person{
		"employee1": {"Alice", 31},
		"employee2": {"Bob", 26},
	}

	equalPerson := maps.EqualFunc(p1, p2, func(v1, v2 Person) bool {
		return v1.Name == v2.Name // Compare only by name
	})
	fmt.Printf("Person map equality by name: %v\n", equalPerson)
}
