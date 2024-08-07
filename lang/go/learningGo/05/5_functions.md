# Functions

<!--toc:start-->
- [Functions](#functions)
  - [Declaring and Calling Functions](#declaring-and-calling-functions)
    - [Simulating Named and Optional Parameters](#simulating-named-and-optional-parameters)
    - [Variadic Input Parameters and Slices](#variadic-input-parameters-and-slices)
    - [Multiple Return Values](#multiple-return-values)
    - [Ignoring Returned Values](#ignoring-returned-values)
    - [Named Return Values](#named-return-values)
  - [Functions Are Values](#functions-are-values)
    - [Function Type Declarations](#function-type-declarations)
    - [Anonymous Functions](#anonymous-functions)
  - [Closures](#closures)
    - [Passing Functions as Parameters](#passing-functions-as-parameters)
    - [Returning Functions from Functions](#returning-functions-from-functions)
  - [defer](#defer)
  - [Go Is Call By Value](#go-is-call-by-value)
  - [Exercises](#exercises)
  - [Wrapping Up](#wrapping-up)
<!--toc:end-->

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
Since functions are values and we can specify the type of a function we can pass functions as parameters into functions. Take a moment to think about the implications of creating a closure that references local variables and then passing that closure to another function.

Let's look at an example of how we can use closures to sort the same data in different ways:
```go
func main() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	// sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
```

The closure that's passed to `sort.Slice` has the parameters `i` and `j`, but within the closure, `people` is used. In computer science terms, people is *captured* by the closure. This code outputs the following:
```
[{Pat Patterson 37} {Tracy Bobdaughter 23} {Fred Fredson 18}]
[{Tracy Bobdaughter 23} {Fred Fredson 18} {Pat Patterson 37}]
[{Fred Fredson 18} {Tracy Bobdaughter 23} {Pat Patterson 37}]
```

### Returning Functions from Functions
We can also return a closure from a function, let's look at an example:
```go
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}
```

This has the following output:
```
0 0
2 3
4 6
```

## defer
Programs often create temporary resources, like files or network connections, that need to be cleaned up. This cleanup must occur no matter how the function exits or whether it completed successfully or not. In Go, this cleanup code is attached to the function with the `defer` keyword.

As an example we'll write a simple version of `cat` the Unix utility for printing the contents of a file:
```go
func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
```

You would build a run this program like so:
```
$ go build
$ ./simpleCat simpleCat.go
package main

import (
        "io"
        "log"
        "os"
)
... // the rest of the file
```

In this example we attempt to open a valid file handle, once we have that and we are done using it we will need to close it no matter how the function exits. To ensure the cleanup code runs we use the `defer` keyword, followed by a function or method call. In this case, we use the `Close` method on the file variable(we'll cover methods in [Chapter 7](../07/7_types_methods_and_interfaces.md)). The `defer` will delay the invocation of the function until the surrounding function exits. 

Some important things to know about `defer`:
  - You can use a function, method, or closure with `defer`. 
  - When we say function with `defer`, mentally expand that to "functions, methods, or closures".
  - You can `defer` multiple functions in a Go function. 
    - They run in last-in, first-out(LIFO) order.
  - The code within `defer` functions run after the return statement.
  - The values you supply a function with input parameters to a `defer` are evaluated immediately and stored until the function runs

Here's a quick example to demonstrate:
```go
func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}
```

Running this code gives the following output:
```
exiting: 30
second: 20
first: 10
```

The best reason to used named return values is to allow a deferred function to examine or modify the return values of its surrounding function. This allows your code to take actions based on an error. In [chapter 9](../09/9_errors.md) we'll discuss a pattern that uses `defer` to add contextual information to an error returned from a function.

This next example shows a way to handle database transaction cleanup using named return values and `defer`:
```go
func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
	if err != nil {
		return err
	}
	// use tx to do more database inserts here
	return nil
}
```

In this example  function we create a transaction to do a series of database inserts. If any of them fail, you want to roll back(not modify the database). If all of them succeed, you want to commit(store the database changes). You can use a closure with `defer` to check whether `err` has been assigned a value. If it hasn't we run `tx.Commit`, which could also return an error. If it does, the value `err` is modified. If any database interaction returned an error, we call `tx.Rollback`.

## Go Is Call By Value
Go is a *call-by-value* language meaning that when you supply a variable for a parameter to a function, Go *always* makes a copy of the value of the variable. Here are some examples that demonstrate this:
```go
type person struct {
	name string
	age  int
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func main() {
	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p) // 2 Hello { 0}
}
```

As the function name indicates, the function won't change the values of the parameters passed into it. This behavior is a little bit different for maps and slices:
```go
func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m) // map[2:hello 3:goodbye]

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s) // [2 4 6]
}
```

For map parameters, any changes are reflected in the variable passed to the function. For slices, you can modify any element in the slice, but you can't lengthen the slice. This behavior applies both to maps and slices that are passed directly into functions as well as map and slice fields in structs.

The reason for this difference in behavior is because maps and slices are both implemented with pointers, which we'll cover in more detail in the [next chapter](../06/6_pointers.md).

## Exercises
1. The simple calculator program doesn't handle one error case: division by zero. Change the function signature for the math operations to return both an `int` and an `error`. In the `div` function, if the divisor is `0`, return `errors.New("division by zero")` for the error. In all other cases, return `nil`. Adjust the `main` function to check for this error. [Solution](./exercises/ex1/main.go)
2. Write a function called `fileLen` that has an input parameter of type `string` and returns an `int` and an `error`. The function takes in a filename and returns the number of bytes in the file. If there is an error reading the file, return the error. Use `defer` to make sure the file is closed properly. [Solution](./exercises/ex2/main.go)
3. Write a function called `prefixer` that has an input parameter of type `string` and returns a function that has an input parameter of type `string` and returns a `string`. The returned function should prefix its input with the string passed into `prefixer`. Use the following `main` function to test `prefixer`:
```go
func main() {
  helloPrefix := prefixer("Hello")
  fmt.Println(helloPrefix("Bob")) // Should print "Hello Bob"
  fmt.Println(helloPrefix("Maria")) // Should print "Hello Maria"
}
```
[Solution](./exercises/ex3/main.go)

## Wrapping Up
This chapter covered functions in Go and their unique features. The [next chapter](../06/6_pointers.md) will cover pointers and how to take advantage of them to write efficient programs.
