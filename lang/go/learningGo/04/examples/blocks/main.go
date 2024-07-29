package main // Package block begins

import "fmt" // File block for "fmt"

var globalVar = 10 // In package block

func main() { // Function block begins
	localVar := 20 // In function block

	{
		innerVar := 30        // In an inner block
		fmt.Println(innerVar) // Can access innerVar here
	}
	// fmt.Println(innerVar) // This would case an error

	if x := 5; x > 0 { // Control structure block
		fmt.Println(x) // x is accessible here
	}
	// fmt.Println(x) // This would cause an error

	fmt.Println(localVar)
}
