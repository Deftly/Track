# Blocks, Shadows, and Control Structures

<!--toc:start-->
- [Blocks, Shadows, and Control Structures](#blocks-shadows-and-control-structures)
  - [Blocks](#blocks)
  - [Shadowing Variables](#shadowing-variables)
  - [if](#if)
  - [for, Four Ways](#for-four-ways)
    - [The Complete for Statement](#the-complete-for-statement)
    - [The Condition-Only for Statement](#the-condition-only-for-statement)
    - [The Infinite for Statement](#the-infinite-for-statement)
    - [break and continue](#break-and-continue)
    - [The for-range Statement](#the-for-range-statement)
      - [The for-range value is a copy](#the-for-range-value-is-a-copy)
    - [Labeling Your for Statements](#labeling-your-for-statements)
  - [switch](#switch)
    - [Blank Switches](#blank-switches)
  - [goto - Yes, goto](#goto-yes-goto)
  - [Exercises](#exercises)
  - [Wrapping Up](#wrapping-up)
<!--toc:end-->

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
```go
func main() {
	x := 10
	if true {
		fmt.Println(x) // 10
		x := 20        // new variable declaration
		fmt.Println(x) // 20
	}
	fmt.Println(x) // 10
}
```

A shadowing variable is a variable that has the same name as a variable in a containing block. As long as the shadowing variable exists, you cannot access a shadowed variable. When using `:=`, make sure that you don't have any variables from an outer scope on the left hand side unless you intend to shadow them.

You also need to be careful to not shadow package imports:
```go
func main() {
  x := 10
  fmt.Println(x)
  fmt := "oops"
  fmt.Println(fmt) // fmt.Println undefined (type string has no field or method Println)
}
```

There is one more block that is a little strange called the universe block. Go only has 25 keywords, but what's interesting is that the built-in types like(`int` and `string`), constants(`true`, `false` and `iota`), and functions like `make` or `close` aren't keywords. Neither is `nil`.

Rather than make them keywords, Go considers these to be *predeclared identifiers* and defines them in the universe block, which contains all other blocks. This means they can be shadowed in other scopes:
```go
fmt.Println(true) // true
true := 10
fmt.Println(true) // 10
```

Be extremely careful to never redefine any identifiers in the universe block.

## if
The `if` statement in Go is very similar to those found in most other programming languages:
```go
n := rand.Intn(10)
if n == 0 {
  fmt.Println("That's too low")
} else if n > 5 {
  fmt.Println("That's too big:", n)
} else {
  fmt.Println("That's a good number:", n)
}
```

Go `if` statements don't put parentheses around the condition and it adds the ability to declare variables that are scoped to the condition and to both the `if` and `else` blocks:
```go
if n := rand.Intn(10); n == 0 {
  fmt.Println("That's too low")
} else if n > 5 {
  fmt.Println("That's too big:", n)
} else {
  fmt.Println("That's a good number:", n)
}
```

This special scope lets you create variables that are available only where they are needed. Once the series of `if/else` statements ends, `n` is undefined.

## for, Four Ways
In Go `for` is the *only* looping keyword in the language and it can be used four different formats:
- A complete, C-style `for`
- A condition-only `for`
- An infinite `for`
- `for-range`

### The Complete for Statement
This `for` statement has three parts, separated by semicolons:
- The initialization - sets one or more variables before the loop begins
- The comparison - This is an expression that must evaluate to a `bool` and is checked immediately *before* each iteration of the loop.
- The increment - Usually something like `i++` but any assignment is valid and it will run immediately after each iteration of the loop, before the condition is evaluated
```go
for i := 0; i < 10; i++ {
  fmt.Println(i)
}

// Go allows you to leave out one or more of the three parts of the for statement
i := 0
for ; i < 10; i++ {
  fmt.Println(i)
}

for i := 0; i < 10; {
  fmt.Println(i)
  if i%2 == 0 {
    i++
  } else {
    i += 2
  }
}
```

### The Condition-Only for Statement
If you leave off both the initialization and increment in a `for` statement, don't include the semicolons. This `for` statement will function like a `while` statement found in other languages:
```go
i := 1
for i < 100 {
  fmt.Println(i)
  i = i * 2
}
```

### The Infinite for Statement
This version of `for` does away with the condition too:
```go
// Infinite loop
for {
  fmt.Println("Hello")
}
```

Press Ctrl-C to interrupt the process and stop the loop

### break and continue
To break out of an infinite loop using code you can use the `break` statement, this will exit the loop immediately and of course it can be used with any `for` statement not just the infinite `for`.

Go also includes the `continue` keyword, this skips over the rest of the `for` loop's body and proceeds directly to the next iteration. You can achieve the same effect as `continue` using nested conditions:
```go
for i := 0; i < 100; i++ {
  if i%3 == 0 {
    if i%5 == 0 {
      fmt.Println("FizzBuzz")
    } else {
      fmt.Println("Fizz")
    }
  } else if i%5 == 0 {
    fmt.Println("Buzz")
  } else {
    fmt.Println(i)
  }
}
```

The code above is not idiomatic and can be confusing to read because of the nested code. Using `continue` makes it easier to understand:
```go
for i := 0; i < 100; i++ {
  if i%3 == 0 && i%5 == 0 {
    fmt.Println("FizzBuzz")
    continue
  }
  if i%3 == 0 {
    fmt.Println("Fizz")
    continue
  }
  if i%5 == 0 {
    fmt.Println("Buzz")
    continue
  }
  fmt.Println(i)
}
```

### The for-range Statement
The fourth `for` statement format is for iterating over elements in some of Go's built-in types. It is called a `for-range` loop and resembles iterators found in other languages. For now we'll cover how to use a `for-range` loop with strings, arrays, slices, and maps. When we cover channels in [Chapter 12](../12/12_concurrency.md) we'll see how to use them with `for-range` loops.

```go
evenVals := []int{2, 4, 6, 8, 10, 12}
for i, v := range evenVals {
  fmt.Println(i, v)
}
```

This produces the following output:
```
0 2
1 4
2 6
3 8
4 10
5 12
```

The `for-range` loop gives you two loop variables. The first variable is the position/key in the data structure being iterated over, while the second is the value at the corresponding position/key. If you don't need to access the key variable you can use an underscore(`_`) as the variable's name which tells Go to ignore the value:
```go
evenVals := []int{2, 4, 6, 8, 10, 12}
for _, v := range evenVals {
  fmt.Println(v)
}
```

If you want the key but don't want the value you can just leave off the second variable:
```go
uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
for k := range uniqueNames {
  fmt.Println(k)
}
```

The most common reason for iterating over the key is when a map is being used as a set, in those cases the value is unimportant.

In this next example we'll iterate over strings using `for-range`:
```go
samples := []string{"hello", "apple_π!"}
for _, sample := range samples {
  for i, r := range sample {
    fmt.Println(i, r, string(r))
  }
  fmt.Println()
}
```

The output for the word "hello" has no surprises:
```
0 104 h
1 101 e
2 108 l
3 108 l
4 111 o
```

The first column is the index; in the second is the numeric value of the letter; and in the third is the numeric value of the letter type converted to a string. Next we'll look at the results for "apple_π!":
```
0 97 a
1 112 p
2 112 p
3 108 l
4 101 e
5 95 _
6 960 π
8 33 !
```

The two things to notice is that the first column skips the number 7 and the value at position 6 is 960 which is far larger than what can fit in a byte. But we saw in the [previous chapter](../03/3_composite_types.md) that strings are made out of bytes.

This is a special behavior from iterating over a string with a `for-range` loop. It iterates over the *runes* not the *bytes*. When it encounters a multi-byte rune in a string it converts the UTF-8 representation into a single 32-bit number and assigns it to the value. The offset is incremented by the number of bytes in the rune.

#### The for-range value is a copy
Be aware that each time the `for-range` loop iterates over your compound type, it *copies* the value from the compound type to the value variable. Modifying the value variable will not modify the value in the compound type:
```go
evenVals := []int{2, 4, 6, 8, 10, 12}
for _, v := range evenVals {
  v *= 2
}
fmt.Println(evenVals) // [2 4 6 8 10 12]
```

In version of Go before 1.22, the value variable is created once and is reused on each iteration through the `for` loop. Since Go 1.22, the default behavior is to create a new index and value variable on each iteration through the `for` loop, this prevents a common bug that has to do with goroutines which we'll cover in [Chapter 12](../12/12_concurrency.md).

Because this is a backward-breaking change, you can control whether to enable this behavior by specifying the Go version in the `go` directive in the *go.mod* file for your module, we'll cover this in greater detail in [Chapter 10](../10/10_modules_packages_and_imports.md).

### Labeling Your for Statements
By default, the `break` and `continue` keywords apply to the `for` loop that directly contains them. To exit or skip over an iteration of an outer loop we can use labels:
```go
func main() {
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}
```

This program gives us the following output:
```
0 104 h
1 101 e
2 108 l
0 97 a
1 112 p
2 112 p
3 108 l
```

Nested `for` loops with labels are rare. They are most commonly used to implement algorithms similar to the following pseudocode:
```go
outer:
	for _, outerVal := range outerValues {
		for _, innerVal := range outerVal {
			// process innerVal
			if invalidSituation(innerVal) {
				continue outer
			}
		}
		// here we have code that runs only when all of the
		// innerVal values were successfully processed
	}
```

## switch
Let's take a look at a basic switch in Go:
```go
words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
for _, word := range words {
  switch size := len(word); size {
  case 1, 2, 3, 4:
    fmt.Println(word, "is a short word!")
  case 5:
    wordLen := len(word)
    fmt.Println(word, "is exactly the right length:", wordLen)
  case 6, 7, 8, 9:
  default:
    fmt.Println(word, "is a long word!")
  }
}
```

This has the following output:
```
a is a short word!
cow is a short word!
smile is exactly the right length: 5
anthropologist is a long word!
```

In this sample program we are switching on the value of an integer, but you can switch on any type that can be compared with `==`, which includes all built-in types except slices, maps, channels, functions and structs that contain fields of these types.

### Blank Switches
A regular switch only allows you to check for equality. A blank switch allows you to use any boolean comparison for each case:
```go
words := []string{"hi", "salutations", "hello"}
for _, word := range words {
  switch wordLen := len(word); {
  case wordLen < 5:
    fmt.Println(word, "is a short word!")
  case wordLen > 10:
    fmt.Println(word, "is a long word!")
  default:
    fmt.Println(word, "is exactly the right length")
  }
}
```

This outputs the following:
```
hi is a short word!
salutations is a long word!
hello is exactly the right length
```

## goto - Yes, goto
Most modern languages don't include `goto` as it is potentially dangerous but it has some use cases and the version of `goto` in Go has limitations that make it safer to use.

In Go, a `goto` statement specifies a labeled line of code, and execution jumps to it. However, you can't jump anywhere. Go forbids jumps that skip over variable declarations and jumps that go into an inner or parallel block.

The following shows two illegal `goto` statements:
```go
func main() {
	a := 10
	goto skip // goto skip jumps over variable declaration at line 8
	b := 20
skip:
	c := 30
	fmt.Println(a, b, c)
	if c > a {
		goto inner // goto inner jumps into block
	}
	if a < b {
	inner:
		fmt.Println("a is less than b")
	}
}
```

The following program demonstrates one of the few valid use cases of `goto`:
```go
func main() {
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}
```

In this example, there is some logic that you don't want to run in the middle of the function, but rather at the end of the function. You could do this without `goto` using boolean flags or duplicating the complicated code after the `for` loop but this is arguable the same functionality as `goto` but more verbose. And duplicating code is problematic because it makes your code harder to maintain. In situations like this `goto` actually improves the code. 

## Exercises
1. Write a `for` loop that puts 100 random numbers between 0 and 100 into an `int` slice. [Solution](./exercises/ex1/main.go)
2. Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
  - If the value is divisible by 2, print "Two!"
  - If the value is divisible by 3, print "Three!"
  - If the value is divisible by 2 and 3, print "Six!". Don't print anything else.
  - Otherwise, print "Never mind".

[Solution](./exercises/ex2/main.go)

3. Start a new program. In `main`, declare an `int` variable called `total`. Write a `for` loop that uses a variable named `i` to iterate from 0(inclusive) to 10(exclusive). The body of the `for` loop should be as follows:
```go
total := total + 1
fmt.Println(total)
```
After the `for` loop, print out the value of `total`. What is printed out? What is the likely bug in this code? [Solution](./exercises/ex3/main.go)

## Wrapping Up
