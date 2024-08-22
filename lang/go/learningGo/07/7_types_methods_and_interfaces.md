# Types, Methods, and Interfaces

<!--toc:start-->
- [Types, Methods, and Interfaces](#types-methods-and-interfaces)
  - [Types in Go](#types-in-go)
  - [Methods](#methods)
    - [Pointer Receivers and Value Receivers](#pointer-receivers-and-value-receivers)
    - [Code Your Methods for nil Instances](#code-your-methods-for-nil-instances)
    - [Methods Are Functions Too](#methods-are-functions-too)
    - [Functions Versus Methods](#functions-versus-methods)
    - [Type Declarations Aren't Inheritance](#type-declarations-arent-inheritance)
    - [Types Are Executable Documentation](#types-are-executable-documentation)
  - [iota Is for Enumerations - Sometimes](#iota-is-for-enumerations-sometimes)
  - [Use Embedding for Composition](#use-embedding-for-composition)
  - [Embedding Is Not Inheritance](#embedding-is-not-inheritance)
  - [A Quick Lesson on Interfaces](#a-quick-lesson-on-interfaces)
  - [Interfaces Are Type-Safe Duck Typing](#interfaces-are-type-safe-duck-typing)
  - [Embedding and Interfaces](#embedding-and-interfaces)
  - [Accept Interfaces, Return Structs](#accept-interfaces-return-structs)
  - [Interfaces and nil](#interfaces-and-nil)
  - [Interfaces Are Comparable](#interfaces-are-comparable)
  - [The Empty Interface Says Nothing](#the-empty-interface-says-nothing)
  - [Type Assertions and Type Switches](#type-assertions-and-type-switches)
  - [Use Type Assertions and Type Switches Sparingly](#use-type-assertions-and-type-switches-sparingly)
  - [Function Types Are a Bridge to Interfaces](#function-types-are-a-bridge-to-interfaces)
  - [Implicit Interfaces Make Dependency Injection Easier](#implicit-interfaces-make-dependency-injection-easier)
  - [Go Isn't Particularly Object-Oriented(and That's Great)](#go-isnt-particularly-object-orientedand-thats-great)
  - [Exercises](#exercises)
  - [Wrapping Up](#wrapping-up)
<!--toc:end-->

## Types in Go
```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
}
```

The above code should be read as declaring a user defined type with the name `Person` to have the *underlying type* of the struct literal that follows. You can also use any primitive type or compound type literal to define a concrete type:
```go
type Score int
type Converter func(string) Score
type TeamScore map[string]Score
```

Go allows you to declare a type at any block level, from the package block down, but you can only access a type from within its scope. The exceptions are types exported from other packages, which we'll cover in [chapter 10](../10/10_modules_packages_and_imports.md)

## Methods
The methods for a type are defined at the package block level:
```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
```

Method declarations look like function declarations with the addition of the *receiver* specification. By convention, the receiver name is a short abbreviation of the type's name, usually the first letter.

Another key difference between declaring methods and functions is that methods can *only* be defined at the package block level, while functions can be defined inside any block.

### Pointer Receivers and Value Receivers
Method receivers can be *pointer receivers*(the type is a pointer) or *value receivers*(the type is a value type). Use the following rules to determine what type of receiver to use:
- If your method modifies the receiver, you *must* use a pointer receiver.
- If your method needs to handle `nil` instances, then it *must* use a pointer receiver.
- If your method doesn't modify the receiver, you *can* use a value receiver.

Here's some simple code to demonstrate pointer and value receivers:
```go
type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c.String()) // total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
	c.Increment()
	fmt.Println(c.String()) // total: 1, last updated: 2024-08-21 10:53:34.450049506 -0700 PDT m=+0.000060371

	d := &Counter{}
	fmt.Println(d.String()) // total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
	d.Increment()
	fmt.Println(d.String()) // total: 1, last updated: 2024-08-21 10:53:34.450152827 -0700 PDT m=+0.000163682
}
```

You might have noticed in the example that we are able to call pointer receiver methods on a value type as well as value receiver methods on pointer types. This is because Go does some automatic conversions to make this happen: `c.Increment()` is converted to `(&c).Increment()` and `d.String()` is converted to `(*d).String()`.

In general do not write getter and setter methods for Go structs unless you need them to meet an interface. Go encourages you to directly access a field and to reserve methods for business logic. The exception to this is when you need to update multiple fields as a single operation or when the update isn't a straightforward assignment of a new value.

### Code Your Methods for nil Instances
In Go you are allowed to invoke methods on a `nil` instance. If the method has a value receiver, you'll get a panic. However, methods with a pointer receiver can work if the method is written to handle a `nil` instance. Here's an implementation of a binary tree that takes advantage of `nil` values for the receiver:
```go
type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}

	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func main() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
}
```

Pointer receivers work like pointer function parameters, it's a copy of the pointer that's passed into the method. Just like `nil` parameters passed to functions, if you change the copy of the pointer you haven't changed the original. This means you can't write a pointer receiver method that handles `nil` and makes the original pointer non-nil.

### Methods Are Functions Too
Methods in Go are so much like functions that you can use a method in place of a function anytime there's a variable or parameter of a function type. Here's an example:
```go

```

### Functions Versus Methods

### Type Declarations Aren't Inheritance

### Types Are Executable Documentation

## iota Is for Enumerations - Sometimes

## Use Embedding for Composition

## Embedding Is Not Inheritance

## A Quick Lesson on Interfaces

## Interfaces Are Type-Safe Duck Typing

## Embedding and Interfaces

## Accept Interfaces, Return Structs

## Interfaces and nil

## Interfaces Are Comparable

## The Empty Interface Says Nothing

## Type Assertions and Type Switches

## Use Type Assertions and Type Switches Sparingly

## Function Types Are a Bridge to Interfaces

## Implicit Interfaces Make Dependency Injection Easier

## Go Isn't Particularly Object-Oriented(and That's Great)

## Exercises

## Wrapping Up
