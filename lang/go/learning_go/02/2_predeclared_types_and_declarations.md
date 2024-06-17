# Predeclared Types And Declarations

<!--toc:start-->
- [Predeclared Types And Declarations](#predeclared-types-and-declarations)
  - [The Predeclared Types](#the-predeclared-types)
    - [The Zero Value](#the-zero-value)
    - [Literals](#literals)
    - [Booleans](#booleans)
    - [Numeric Types](#numeric-types)
      - [Integer types](#integer-types)
      - [The special integer types](#the-special-integer-types)
      - [Choosing which integer to use](#choosing-which-integer-to-use)
      - [Integer operators](#integer-operators)
      - [Floating-point types](#floating-point-types)
      - [Complex types(probably won't ever use them)](#complex-typesprobably-wont-ever-use-them)
    - [A Taste of Strings and Runes](#a-taste-of-strings-and-runes)
    - [Explicit Type Conversion](#explicit-type-conversion)
    - [Literals Are Untyped](#literals-are-untyped)
  - [var Versus :=](#var-versus)
  - [Using const](#using-const)
  - [Typed and Untyped Constants](#typed-and-untyped-constants)
  - [Exercises](#exercises)
  - [Wrapping Up](#wrapping-up)
<!--toc:end-->

## The Predeclared Types
Go has many built-in types, also called *predeclared types*, similar to types found in other languages: booleans, integers, floats, and strings. This chapter will cover how to best use these types in Go. Before that we'll cover some concepts that apply to all types.

### The Zero Value
Go assigns a default *zero value* to any variable that is declared but not assigned a value. This removes a source of bugs found in C and C++. As we cover each type we'll cover the zero value for the type.

### Literals
A Go *literal* is an explicitly specified number, character, or string. Go programs have four common kinds of literals(the fifth is rare and is used for complex numbers).

*Integer literals* are base 10 by default, but different prefixes are used to indicate other bases:
  - `0b` for binary
  - `0o` for octal
  - `0x` for hexadecimal

*Floating point literals* have a decimal point to indicate the fractional portion of the value. The can also have an exponent specified with letter `e` and a positive or negative number(such as 6.03e23).

A *rune literal* represents a character and is surrounded by single quotes. Unlike many other languages, single quotes and double quotes are not interchangeable in Go. Runes can be written:
  - Single Unicode characters `'a'`
  - 8-bit hexadecimal numbers `'\x61'`
  - 16-bit hexadecimal numbers `'\u0061'`
  - 32-bit Unicode numbers `'\U00000061'`
There are also several backslash escaped rune literals(`'\n', '\t', '\'', '\\'`)

There are two ways to indicate string literals. Most of the time you'll use double quotes to create an *interpreted string literal*: `"Go is a statically typed language\nWith excellent concurrency features"`

The other option is to use a *raw string literal*. These are delimited with backquotes(`) and can contain any character except a backquote. There is no escape character in a raw string literal:
```go
`Go is a statically typed language
With excellent concurrency features`
```

Literals are considered *untyped*. There are situations in Go where the type isn't explicitly declared, in those cases Go uses the *default type* for a literal.

### Booleans
The `bool` type represents boolean variables which can have one of two values: `true` or `false`. The zero value is `false`.
```go
var flag bool // no value assigned, set to false
var isAwesome = true
```

### Numeric Types
Go has 12 numeric types(and a few special names) that are grouped into three categories. Some of them are used quite frequently while others are more obscure.

#### Integer types
| Type | Value range |
|------|-------------|
|int8| -128 to 127|
|int16| -32768 to 32767|
|int32| –2147483648 to 2147483647|
|int64| –9223372036854775808 to 9223372036854775807|
|uint8| 0 to 255|
|uint16| 0 to 65535|
|uint32| 0 to 4294967295|
|uint64| 0 to 18446744073709551615|

#### The special integer types
A `byte` is an alias for `uint8`; it is legal to assign, compare, or perform mathematical operations between a `byte` and a `uint8`. You'll rarely see `uint8`, just use `byte` instead.

The second special name is `int`. On a 32-bit CPU, `int` is a 32-bit signed integer like an `int32`. On most 64-bit CPUs, `int` is a 64-bit singed integer like `int64`. Because `int` isn't consistent from platform to platform it is a compile time error to assign, compare, or perform mathematical operations between an `int` and `int32` or `int64` without a type conversion.

Integer literals default to being of `int` type.

`uint` is also a special name and follows the same rules as `int` only it is unsigned.

There are two other special names for integer types, `rune` and `uintptr`. We saw rune literals earlier and we'll cover the `rune` type [later in this chapter](#a-taste-of-strings-and-runes) and `uintptr` in [Chapter 16](../16/16_reflect_unsafe_cgo.md).

#### Choosing which integer to use
There are three simple rules to choosing what integer type to use:
- If you are working with a binary file format or network protocol that has an integer of a specific size or sign, use the corresponding integer type.
- If you are writing a library function that should work with any integer type, use a generic type parameter([more on generics in chapter 8](../08/8_generics.md)) to represent any integer type.
- In all other cases just use `int`.

#### Integer operators
Go integers support all the usual arithmetic operators: `+`, `-`, `*`, `/`, and `%` for modulus. Any of them can be combined with `=` to modify a variable:
```go
var x int = 10
x *= 2  // 20
```

Comparisons are done with: `==`, `!=`, `>`, `>=`, `<`, and `<=`.

Go also has bit manipulation operators for integers:
  - Bit shift left `<<`
  - Bit shift right `>>`
  - Bitwise AND `&`
  - Bitwise OR `^`
  - Bitwise AND NOT `&^`
Like with arithmetic operators you can also combine them with `=` to modify a variable.

#### Floating-point types
| Type | Largest absolute value |
|------|-------------|
|float32|3.40282346638528859811704183484516925440e+38|
|float64|1.797693134862315708145274237317043567981e+308|

Like the integer types, the zero value for floating-point types is 0.

Picking which floating-point type to use is easy, always use `float64` unless you have to be compatible with an existing format that uses `float32`.

In most cases you should avoid using floating point numbers. Like in other languages, floating-point numbers in Go have a huge range but can't store every value in that range. They store the nearest approximation, this limits their use to cases where inexact values are acceptable.

> **_NOTE:_** Go(and most other programming languages) store floating point numbers using a specification called IEEE 754

You can use all the standard mathematical and comparison operators with floats except `%`. When dividing a nonzero floating-point variable by 0 it returns `+Inf` or `-Inf` depending on the sign of the number. Dividing a floating-point variable set to 0 by 0 returns `NaN`(Not a number).

#### Complex types(probably won't ever use them)
Feel free to skip this section. Go defines two complex number types:
  - `complex64` uses `float32` values to represent the real and imaginary part
  - `complex128` uses `float54` values
Both are declared using the `complex` built-in function:
```go
var complexNum = complex(20.3, 10.2)
```
All the standard floating-point arithmetic operators work on complex numbers and they have same precision limitations. You can extract the real and imaginary portions using the `real` and `imag` built-in functions. The `math/cmplx` package has additional functions for manipulating `complex128` values.

The zero value for both types of complex numbers has 0 assigned to both the real and imaginary portions of the number.

### A Taste of Strings and Runes
You can put any Unicode character into a string. Like integers and floats, strings are compared for equality using `==`, difference with `!=`, or ordering with `>`, `>=`, `<`, or `<=`. They can be concatenated by using the `+` operator. The zero value for a string is the empty string.

Strings in Go are immutable, you can reassign the value of a string variable but you cannot change the value of the string that is assigned to it.

The *rune* type is an alias for `int32`. If you are referring to a character, use the rune type, not `int32`. While they are the same to the compiler you want to use the type that clarifies the intent of your code.

The [next chapter](../03/3_composite_types.md) will cover a lot more about strings, including some implementation details, relationships with bytes and runes, as well as some advanced features and pitfalls.

### Explicit Type Conversion
Go doesn't allow automatic type promotion between variables, you must use a *type conversion* when variable types do not match. Even different sized integers and floats must be converted to the same type to interact. This make sit clear exactly what type you want without having to memorize any type conversion rules.

Since all type conversions in Go are explicit you cannot treat another Go type as a boolean. Many other languages allow nonzero number or nonempty string to be interpreted as a boolean `true`. This is not the case in Go. No other type can be converted to a `bool`, implicitly or explicitly. To convert from another data type to boolean, you must use one of the comparison operators.

### Literals Are Untyped
While you can't add two integer variables together if they are declared to be of different types of integers, Go lets you use an integer literal in floating-point expressions or even assign an integer literal to a floating-point variable:
```go
var x flaot64 = 10
var y float64 = 200.3 * 5
```
This is allowed because literals in Go are untyped, meaning they can be used with any variable whose type is compatible with the literal. In [chapter 7](../07/7_types_methods_and_interfaces.md) we'll see that we can even use literals with user defined types based on predefined types.

## var Versus :=
Here's a list of many of the ways to declare a variable in Go using the `var` keyword:
```go
var x int = 10
var y = 10
var z int
var x2, y2 int = 10, 20
var x3, y3 int
var a, b = 10, "hello"
```

Go also supports a short declaration and assignment format. When within a function you can use the `:=` operator to replace a `var` declaration that uses type inference. The following statements are equivalent:
```go
var x = 10
x := 10
```
Just like with `var`, you can declare multiple variables at once using `:=`. The `:=` operator also allow s you to assign values to existing variables as long as at least one new variable is on the left hand side of `:=`, this cannot be done with `var`.

The one limitation of `:=` is that it cannot be used to declare a variable at the package level, it is only legal within the scope of a function.

## Using const
Most languages have a way to declare that a value is immutable. In Go, this is done using the `cosnt` keyword:
```go
package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 1 // this will not compile!
	y = "bye" // this will not compile!
	fmt.Println(x)
	fmt.Println(y)
}
```
This code will fail to compile with the following error:
```
./test.go:20:2: cannot assign to x (neither addressable nor a map index expression)
./test.go:21:2: cannot assign to y (neither addressable nor a map index expression)
```

Constants in Go are a way to give names to literals. They can only hold values that the compiler can figure out at compile time:
  - Numeric literals
  - `true` and `false`
  - Strings
  - Runes
  - Values returned by the built-in functions `complex`, `real`, `imag`, `len`, and `cap`
  - Expressions that consist of operators and the preceding values

> **_NOTE:_** We'll cover the `len` and `cap` functions in the [next chapter](../03/3_composite_types.md). Another value that can be used with `const` is `iota`, this will be covered in [chapter 7](../07/7_types_methods_and_interfaces.md).

In Go there is no way to specify that a value calculated at runtime is immutable. There are no immutable arrays, slices, maps, or structs. This isn't as limiting as it sounds. Within a function, it is clear if a variable is being modified so immutability is less important. Later we'll discuss how "Go Is Call by Value" and how that prevents modifications to variables that are passed as parameters to functions.

## Typed and Untyped Constants
An untyped constant works exactly like a literal, it has no type of its own but does have a default type that is used when no other type can be inferred. A typed constant can be directly assigned only to a variable of that type.

In general, leaving a constant untyped gives you more flexibility, however, there are some situations where you'll want a constant to enforce a type.
```go
const x = 10

var y int = x
var z float64 = x
var d byte = x

fmt.Println(y, z, d)

const typedX int = 10 // can only be assigned directly to an int
```

## Exercises
1. Write a program that declares an integer variable called `i` with the value 20. Assign `i` to a floating-point variable named `f`. Print out `i` and `f`. [Solution](./exercises/ex1/ex1.go)
2. Write a program that declares a constant called `value` that can be assigned to both an integer and a floating-point variable. Assign it to an integer called `i` and a floating-point variable called `f`. Print out `i` and `f`. [Solution](./exercises/ex2/ex2.go)
3. Write a program with three variables, on named `b` of type `byte`, one named `smallI` of type `int32`, and one named `bigI` of type `uint64`. Assign each variable the maximum legal value for its type; then add `1` to each variable. Print out their values. [Solution](./exercises/ex3/ex3.go)

## Wrapping Up
This chapter covered how to use the predeclared types, declare variables, and work with assignments and operators. The [next chapter](../03/3_composite_types.md) will look at the composite types in Go: arrays, slices, maps, and structs.
