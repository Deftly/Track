# Functions

## Declaring and Calling Functions
Go has functions similar to those found in other languages but it adds its own twist on function features.

We've already seen functions being declared and used. Every Go program starts from a `main` function, and we've been calling the `fmt.Println` function. Now let's take a look at a function that takes in parameters and returns values:
```go
func main() {
	result := div(5, 2)
	fmt.Println(result)
}

func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}
```

A function declaration has four parts: the keyword `func`, the name of the function, the input parameters, and the return type. Go has a `return` keyword for returning values from a function. If a function returns a value, you *must* supply a `return`. If a function returns nothing, `return` statement is not needed unless you want to exit the function before the last line.

> **_NOTE:_** When you have two or more consecutive input parameters of the same type, you can specify the type once for all of them like this: `func div(num, denom int) int {`

### Simulating Named and Optional Parameters
Two features that Go functions *don't* have are: named and optional input parameters. With one exception that we'll cover in the next section, you must supply all parameters for a function. To emulate named and optional parameters, define a struct that has fields that match the desired parameters, and pass the struct to your function:
```go
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	// do something here
}

func main() {
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})

	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}
```

In practice not having named and optional parameters isn't a limitation. A function shouldn't have more than a few parameters, and named and optional parameters are mostly useful when a function has many inputs. If you find yourself in that situation your function might be too complicated.

### Variadic Input Parameters and Slices
You may have noticed that we can use `fmt.Println` with any number of input parameters, it can do this because Go supports *variadic parameters*. The variadic parameters *must* be the last parameter in the input parameter list and is indicated with three dots before the type. The variable that is created within the function is a slice of the specified type:
```go
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func main() {
	fmt.Println(addTo(3))             // []
	fmt.Println(addTo(3, 2))          // [5]
	fmt.Println(addTo(3, 2, 4, 6, 8)) // [5 7 9 11]
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))                    // [7 6]
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // [4 5 6 7 8]
}
```

### Multiple Return Values
A unique feature of Go is that it allows for multiple return values:
```go
func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
```

We'll cover more about errors in [Chapter 9](../09/9_errors.md). For now, remember that you use Go's multiple return value support to return an `error` if something goes wrong in a function. If the function completes successfully, return `nil` for the error's value. By convention, the `error` is always the last value returned from a function.

### Ignoring Returned Values
We know that go does not allow unused variables, so what do we do when a function return values that we don't need? In these situations we assign the unused variables to the name `_`. For example, if you don't want to use `remainder` from the previous example, you would write the assignment as `result, _, err := divAndRemainder(5, 2)`.

### Named Return Values
Go allows you to specify *names* for your return values. Let's rewrite our previous example:
```go
func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}
```

When you supply names to your return values, what you are doing is pre-declaring variables that you use within the function to hold your return values. Named return values are initialized to their zero values when created, this means that you can return them before any explicit use or assignment.

While named return values can help clarify code they have some problems too. First is the problem of shadowing. Just like with any other variable, you can shadow a named return value so take care that you are assigning to the return value and not to a shadow of it.

The other problem is that you don't have to return them:
```go
func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	result, remainder = 20, 30
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder) // 2 1
}
```

Notice that the values from the `return` statement were returned even though they were never assigned to the named return parameters. The named return parameters give a way to declare an *intent* to use variables to hold the return values, but don't *require* you to use them.

Because of these potential pitfalls named return values provide limited value. However, there is one situation in which they are essential which we will cover when going over `defer` later in the [chapter](#defer).

## Functions Are Values
The type of a function is built out of the keyword `func` and the types of the parameters and return values. This combination is called the *signature* of the function. Any function that has the exact same number and types of parameters and return values meets the type signature.

Since functions are values, you can declare a function variable: `var myFuncVariable func(string) int`

`myFuncVariable` can be assigned any function that has a single parameter of type `string` and returns a single value of type `int`:
```go
func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		fmt.Println(string(v), ":", int(v))
		total += int(v)
	}
	return total
}

func main() {
	var myFuncVariable func(string) int
	myFuncVariable = f1
	result := myFuncVariable("Hello")
	fmt.Println(result) // 5

	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println(result) // 500
}
```

The default zero value for a function variable is `nil`, and attempting to run a function variable with a `nil` value results in a panic.

Having functions as values allows us to do some clever things, such as build a simple calculator using functions as values in a map. First we'll create a set of functions that have the same signature:
```go
func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }
```

Next, we'll create a map to associate a math operator with each function:
```go
var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}
```

Finally we'll test it out on a few different expressions:
```go
func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
```

With out set of test expressions this program has the following output:
```
5
-1
6
0
unsupported operator: %
strconv.Atoi: parsing "two": invalid syntax
invalid expression: [5]
```

> **_NOTE:_** Don't write fragile programs. The majority of the lines in the above program are for error checking and data validation. Failing to do these things leads to unstable and unmaintainable code. Error handling is what separates the professionals from the amateurs.

### Function Type Declarations
We can use the `type` keyword to define a function type the same we use it define a `struct`:
```go
type opFuncType func(int, int) int
```

We can then rewrite the `opMap` declaration:
```go
var opMap = map[string]opFuncType{
  // same as before
}
```

Function types can serve as documentation and it can be useful to give something a name if you're going to refer to it multiple times. We'll cover another use when we look at interfaces in a [later chapter](../07/7_types_methods_and_interfaces.md)

### Anonymous Functions
You can define new functions within a function and assign them to variables:
```go
func main() {
	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}
}
```

Inner functions are *anonymous*; they don't have a name. You declare an anonymous function with the keyword `func` immediately followed by the input parameters, the return values,a nd the opening brace. The above example has the following output:
```
printing 0 from inside of an anonymous function
printing 1 from inside of an anonymous function
printing 2 from inside of an anonymous function
printing 3 from inside of an anonymous function
printing 4 from inside of an anonymous function
```

You don't have to assign an anonymous function to a variable. You can write them inline and call them immediately too:
```go
func main() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside an anonymous function")
		}(i)
	}
}
```

This isn't something you would normally do. If you declare an anonymous function and execute it immediately, you might as well get rid of the function and just call the code. However, declaring anonymous functions without assigning them to variables is useful in two situations: `defer` statements and launching goroutines(covered in [chapter 12](../12/12_concurrency.md)).

## Closures
Closure is a computer science term that means that a functions declared inside functions are able to access and modify variables declared in the outer function:
```go
func main() {
	a := 20
	f := func() {
		fmt.Println(a)
		a = 30
	}
	f()
	fmt.Println(a)
}
```

Running this program gives the following output:
```
20
30
```

The anonymous function assigned to `f` can read and write `a`, even though `a` is not passed in to the function.

One thing that closures allow us to do is to limit a function's scope. If a function is going to be called from only one other function, but it is called multiple times, you can use an inner function to "hide" the called function, reducing the number of package level declarations.

Closures really become interesting when they are passed to other functions or returned from a function. This allows us to take the variables within a function and use those values *outside* of the function.

### Passing Functions as Parameters


## defer


## Go Is Call By Value


## Exercises


## Wrapping Up

