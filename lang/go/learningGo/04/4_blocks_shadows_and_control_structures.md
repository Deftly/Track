# Blocks, Shadows, and Control Structures

## Blocks
A block in Go is a place where a declarations occur. It's a scope where variables, constants, types and functions are defined and made accessible:
  - Package block: This is the outermost block in a Go program. Declarations made outside of any function are placed in the package block and they can be accessed throughout the package.
  - File block: This block is specific to a single Go source file. When you import a package, the names from that package become available within the file where the import statement appears.
  - Function block: All variables defined at the top level of a function, including its parameters, are in the function block.
  - Inner blocks: Within a function, every set of braces(`{}`) defines a new block, creating nested scopes within the function.
  - Control structure blocks: Go's control structures(like `if`, `for`, and `switch`) also define their own blocks.

```go
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
```

You can access an identifier defined in any outer block from within any inner block. So what happens when you have a declaration with the same name as an identifier in a containing block? This is referred to as shadowing.

## Shadowing Variables

## if

## for, Four Ways

### The Complete for Statement

### The Condition-Only for Statement

### The Infinite for Statement

### break and continue

### The for-range Statement

### Labeling Your for Statements

### Choosing the Right for Statement

## switch

### Blank Switches

### Choosing Between if and switch

## goto - Yes, goto

## Exercises

## Wrapping Up
