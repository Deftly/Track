package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(word string) string {
		return prefix + " " + word
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // Should print "Hello Bob"
	fmt.Println(helloPrefix("Maria")) // Should print "Hello Maria"
}
