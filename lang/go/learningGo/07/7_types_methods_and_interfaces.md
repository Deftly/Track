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
type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5)) // 15
  // method value
	f1 := myAdder.AddTo // Assign method to a variable
	fmt.Println(f1(10)) // 20
  // method expression
	f2 := Adder.AddTo            // Create a function from the type itself
	fmt.Println(f2(myAdder, 15)) // 25
}
```

We'll see how we can use method values and method expressions when we look at dependency injection later in this chapter.

### Functions Versus Methods
Anytime your logic depends on values that are configured at startup or changed while your program is running, those values should be stored in a struct, and that logic should be implemented as a method. If your logic depends only on the input parameters, it should be a function.

### Type Declarations Aren't Inheritance
In addition to declaring types based on built-in Go types and struct literals, you can also declare a user-defined type based on another user defined type.
```go
type HighScore Score
type Employee Person
```

This might look like inheritance but it isn't. The two types have the same underlying type, but that's all. In Go you can't assign an instance of type `HighScore` to a variable of type `Score`, or vice versa, without a type conversion, nor can you assign either of them to a variable of type `int` without a type conversion. Also, any methods defined on `Score` aren't defined on `HighScore`:
```go
// assigning untyped constnats is valid
var i int = 300
var s Score = 100
var hs HighScore = 200

hs = s            // compilation error!
s = i             // compilation error!
s = Score(i)      // ok
hs = HighScore(s) // ok
```

User-defined types whose underlying types are built-in types can be assigned literals and constants compatible with the underlying type.

### Types Are Executable Documentation
Types act as documentation, they make code clearer by providing a name for a concept and describing the kind of data that is expected. It's clearer when a method has a parameter of type `Percentage` than of type `int`, and it's harder for it to be invoked with an invalid value.

The same logic applied when declaring one user-defined type based on another user defined type. When you have the same underlying data, but different sets of operations to perform, make two types. Declaring one as being based on the other avoid some repetition and make it clear that the two types are related.

## iota Is for Enumerations - Sometimes
Many languages have the concept of enumerations which allow you to specify that a type can have only a limited set of values. Go doesn't have an enumeration type. It has `iota` which lets you assign an increasing value to a set of constants.

When using `iota`, the best practice is to first define a type based on `int` that will represent all the valid values. Then we use a `const` block to define a set of values for our type:
```go
type MailCategory int

const (
  Uncategorized MailCategory = iota
  Personal
  Spam
  Social
  Advertisements
)
```

The first constant in the `const` block has the type specified and its value set to `iota`. When the Go compiler sees this it repeats the type and assignment to all subsequent constants in the block. The value of `iota` increments for each constant defined in the constant block, starting with `0`. When a new `const` block is created `iota` is set back to `0`.

The value of `iota` increments for each constant in the `const` block, whether or not `iota` is used to define the value of a constant: 
```go
const (
	Field1 = 0        // 0
	Field2 = 1 + iota // 2
	Field3 = 20       // 20
	Field4            // 20
	Field5 = iota     // 4
)
```

It's important to remember that nothing in Go will stop you(or anyone else) from creating additional values of your type. When you insert a new identifier in the middle of your list of literals, all subsequent ones will be renumbered. This will break your application if those constants represented values in another system or database. Because of these limitations it only makes sense to use `iota`-based enumerations when you want to be able to differentiate between a set of values and don't care what the value is behind the scenes.

## Use Embedding for Composition
Go encourages code reuse via built-in support for composition and promotion:
```go
type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// do business logic
}
```

Note that `Manager` contains a field of type `Employee`, but no name is assigned to that field. This makes `Employee` and *embedded field*. Any fields or methods declared on an embedded field are *promoted* to the containing struct and can be invoke directly:
```go
func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)            // 12345
	fmt.Println(m.Description()) // Bob Bobson (12345)
}
```

If the containing struct has fields or methods with the same name as an embedded field, you need to use the embedded field's type to refer to the obscured fields or methods:
```go
type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

func main() {
	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)       // 20
	fmt.Println(o.Inner.X) // 10
}
```

## Embedding Is Not Inheritance
Built-in embedding support is rare in programming languages, and many developers try to understand embedding by treating it as inheritance, don't do this. You cannot assign a variable of type `Manager` to a variable of type `Employee`. To access the `Employee` field in `Manager`, you must do so explicitly:
```go
var eFail Employee = m // compilation error: cannot use m (type Manager) as type Employee in assignment
var eOk Employee = m.Employee // ok
```

Go has no [*dynamic dispatch*](../addendums/dynamicDispatch.md) for concrete types. The methods on the embedded field have no idea that they are embedded. If you have a method on an embedded field that calls another method on the embedded field, and the containing struct has a method of same name, the method on the embedded field is invoked:
```go
type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	S string
	Inner
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func main() {
	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.Double()) // Inner 20
}
```

While embedding once concrete type inside another won't allow you to treat the outer type as the inner type, the methods on an embedded field do count toward the *method set* of the containing struct. This means they can make the containing struct implement an interface.

## A Quick Lesson on Interfaces
We'll start with looking at how to declare interfaces(the only abstract type in Go), like other user defined types you use the `type` keyword. Here's the definition of the `Stringer` interface in the `fmt` package:
```go
type Stringer interface {
  String() string
}
```

In an interface declaration, an interface literal appears which lists the methods that must be implemented by a concrete type to meet the interface. This list of methods is the method set of the interface. We mentioned previously that the method set of a pointer instance contains the methods defined with both pointer and value receivers, while the method set of a value instance contains only the methods with value receivers. Here's an example:
```go
type Counter struct {
	lastUpdated time.Time
	total       int
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type Incrementer interface {
	Increment()
}

func main() {
	var myStringer fmt.Stringer
	var myIncrementer Incrementer
	pointerCounter := &Counter{}
	valueCounter := Counter{}

	myStringer = pointerCounter    // ok
	myStringer = valueCounter      // ok
	myIncrementer = pointerCounter // ok
	myIncrementer = valueCounter   // compile-time error
}
```

Trying to compile this code will results in the error: `cannot use valueCounter (variable of type Counter) as Incrementer value in assignment: Counter does not implement Incrementer (method Increment has pointer receiver)`.

## Interfaces Are Type-Safe Duck Typing
What makes interfaces in Go special is that they are implemented implicitly. A concrete type does not explicitly declare that it implements an interface. Instead, if a concrete type's method set contains all the methods in an interface's method set, the concrete type implicitly implements the interface. This allows the concrete type to be assigned to a variable or field of the interface type, enabling both type safety and decoupling, bridging the functionality found in both static and dynamic languages.

Interfaces allow you to depend on behavior not an implementation, meaning we can swap implementations as needed. This allows your code to evolve over time as requirements inevitably change. Let's take a look at an example:
```go
type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// business logic
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data from somewhere
	c.L.Process(data)
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
```

The Go code provides an interface, but only the caller(`Client`) knows about it. Nothing is declared on `LogicProvider` to indicate that it meets the interface. This allows both a new logic provider to be easily added in the future and provide executable documentation to ensure that any type passed into the client will match the client's needs.

> **_NOTE:_** If an interface in the standard library describes what your code needs, use it. Commonly used interfaces include `io.Reader`, `io.Writer`, and `io.Closer`.

It's fine for a type that meets an interface to specify additional methods that aren't part of the interface. For example, the `io.File` type meets the `io.Reader` and `io.Writer` interfaces. If your code cares only about reading from a file, use the `io.Reader` interface to refer to the file instance and ignore the other methods.

## Embedding and Interfaces
Embedding is not only for structs. You can embed an interface in an interface. For example, the `io.ReadCloser` interface is built out of an `io.Reader` and an `io.Closer`:
```go
type Reader interface {
  Read(p []byte) (n int, err error)
}

type Closer interface {
  Close() error
}

type ReadCloser interface {
  Reader
  Closer
}
```

## Accept Interfaces, Return Structs
Experienced Go developers will often say that your code should "Accept interfaces, return structs". This means that business logic invoked by your functions should be invoked via interfaces, but the output of your functions should be a concrete type. We've seen why functions should accept interfaces: they make your code more flexible and explicitly declare the exact functionality being used.

The main reason your functions should return concrete types is that they make it easier to update your function's return values in new version of your code. When a concrete type is returned by a function, new methods and fields can be added without breaking existing code that calls the function, because the new fields and methods are ignored. Adding a new method to an interface means that existing implementations of that interface must be updated or your code breaks.

Errors are the exception to this rule. We'll see in [chapter 9](../09/9_errors.md) that Go functions and methods can declare a return parameter of the `error` interface type. In the case of `error` it's quite likely that different implementations of the interface could be returned, so you need to use an interface to handle all possible options as interfaces are the only abstract type in Go.

## Interfaces and nil
In the Go runtime, interfaces are implemented as a struct with two pointer fields, one for the value and one for the type of the value. As long as the type field in non-nil, the interface is non-nil. It is only when **both** the type and value pointers are `nil` that the interface itself also `nil`: 
```go
var pointerCounter *Counter
fmt.Println(pointerCounter == nil) // true
var incrementer Incrementer
fmt.Println(incrementer == nil) // true
incrementer = pointerCounter
fmt.Println(incrementer == nil) // false
```

What `nil` indicates for a variable with an interface type is whether you can invoke methods on it. We've seen previously that you can invoke methods on `nil` concrete instances, so it makes sense that you can invoke methods on an interface variable that was assigned a `nil` concrete instance. If an interface variable is `nil`, invoking any methods on it will trigger a panic.

Since an interface instance with a non-nil type is not equal to `nil`, it's not straightforward to tell whether the value associated with the interface is `nil`. You have to use reflection(covered in [chapter 16](../16/16_reflect_unsafe_cgo.md)) to find out.

## Interfaces Are Comparable
In [chatper 3](../03/3_composite_types.md) we covered comparable types, the ones that can be checked for equality with `==`. Interfaces are also comparable. Two instances of an interface type are equal if their types are equal and their values are equal. But what happens if the type isn't comparable? Let's look at an example:
```go
type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func main() {
	var di DoubleInt = 10
	var di2 DoubleInt = 10
	dis := DoubleIntSlice{1, 2, 3}
	dis2 := DoubleIntSlice{1, 2, 3}
	// false because we are comparing pointers, and they point to different values
	DoublerCompare(&di, &di2)
	// false because they have different underlying types
	DoublerCompare(&di, dis)
	// triggers a panic, because the underlying types match but are a
	// non-comparable type
	DoublerCompare(dis, dis2)
}
```

Also, remember that the key of a map must be comparable, so a map can be defined to have an interface as a key: `m := map[Doubler]int{}`

If you were to add a key-value pair to this map and they key isn't comparable, that will also trigger a panic.

Given this behavior, be careful when using `==` or `!=` with interfaces or using an interface as a map key as this can easily generate a panic and crash your program. Even if all your interface implementations are currently comparable, you don't know what will happen if someone else uses or modifies your code, and there's no way to specify that an interface can only be implemented by comparable types. To be extra safe you can use the `Comparable` method on `reflect.Value` to inspect an interface before using it with `==` or `!=`.(We'll cover reflection in [chapter 16](../16/16_reflect_unsafe_cgo.md))

## The Empty Interface Says Nothing
Sometimes in a statically typed language we need a way to say that a variable could store a value of any type. For this Go uses an *empty interface*, `interface{}` to represent this:
```go
var i interface{}
i = 20
i = "hello"
i = struct {
  FirstName string
  LastName  string
}{"Fred", "Fredson"}
```

To improve readability, Go added `any` in Go 1.18 as a type alias for `interface{}`

Because an empty interface doesn't tell you anything about the value it represents, you can't do a lot with it. A common use of `any` is as a placeholder for data of uncertain schema that's read from an external source, like a JSON file:
```go
data := map[string]any{}
contents, err := os.ReadFile("testdata/sample.json")
if err != nil {
  return err
}
json.Unmarshal(contents, &data)
// the contents are now in the data map
```

> **_NOTE:_** User-created data containers written before generics used the empty interface  to store values of any type. Now that generics are part of Go it is recommended that you use them for any newly created data containers. Generics provide type safety, compile-time checking, and eliminate the need for type assertions, resulting in more robust and maintainable code.

In general you should avoid using `any`. Go is a strongly typed language and attempts to work around this are unidiomatic.

If you do end up storing a value into an empty interface you will need to use type assertions and type switches to read the value back again.

## Type Assertions and Type Switches
A *type assertion* names the concrete type that implemented the interface, or names another interface that is also implemented by the concrete type whose value is stored in the interface.
```go
type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt) // i2 is of type MyInt
	fmt.Println(i2 + 1)
}
```

If the type assertion is wrong the code will panic like in this example:
```go
i2 := i.(string)
fmt.Println(i2) // panic: interface conversion: interface {} is main.MyInt, not string
```

Even if two types share an underlying type, a type assertion must match the type of the value stored in the interface:
```go
i2 := i.(int)
fmt.Println(i2 + 1) // panic: interface conversion: interface {} is main.MyInt, not int
```

To avoid crashing we use the comma ok idiom:
```go
i2, ok := i.(int)
if !ok {
  return fmt.Errorf("unexpected type for %v", i)
}
fmt.Println(i2 + 1)
```

When an interface could be one of multiple possible types, use a *type switch* instead:
```go
func doThings(i any) {
	switch j := i.(type) {
	case nil:
		// i is nil, type of j is any
	case int:
		// j is of type int
	case MyInt:
		// j is of type MyInt
	case io.Reader:
		// j is of type io.Reader
	case string:
		// j is a string
	case bool, rune:
		// i is either a bool or rune, so j is of type any
	default:
		// no idea what i is, so j is of type any
	}
}
```

> **_NOTE:_** Since the purpose of a type `switch` is to derive a new variable from an existing one, it is idiomatic to assign the variable being switched on to a variable of the same name(`i := i.(type)`), making this one of the few places where shadowing is a good idea.

While the examples so far have use the `any` interface with type assertions and type switches, you can uncover the concrete type from all interface types.

## Use Type Assertions and Type Switches Sparingly
While being able to extract the concrete implementation from an interface variable can be handy you should use the techniques infrequently. In most cases you should treat a parameter or return value as the type that was supplied and not what else it could be. Otherwise your function's API isn't accurately declaring the types it needs to perform its task.

One common use of a type assertion is to see if the concrete type behind the interface also implements another interface. For example, the standard library uses this technique to allow more efficient copies when the `io.Copy` function is called. This functions has two parameters of types `io.Writer` and `io.Reader` and calls the `io.copyBuffer` function to do its work. If the `io.Writer` parameter also implements `io.WriterTo`, or the `io.Reader` parameter also implements `io.ReaderFrom`, most of the work in the function can be skipped:
```go
// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
  // function continues...
}
```

This optional interface technique has one drawback. It's common for implementations of interfaces to use the decorator pattern to wrap other implementations of the same interface to layer behavior. The problem is that if an optional interface is implemented by one of the wrapped implementations, you cannot detect it with a type assertion or type switch. For example, the standard library includes a `bufio` package that provides a buffered reader. We can buffer any other `io.Reader` implementation by passing it to the `bufio.NewReader` function and using the returned `*bufio.Reader`. If the passed-in `io.Reader` also implemented the `io.ReaderFrom`, wrapping it in a buffered reader prevents the optimization we saw earlier.

Type `switch` statements give us the ability to differentiate between multiple implementations of an interface that require different processing. They are most useful when only certain possible valid types can be supplied for an interface. Be sure to include a `default` case to handle implementations that aren't known at development time. This also protects you if you forget to update the type `switch` statements when new interface implementations are added:
```go
func walkTree(t *treeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("invalid expression")
	case number:
		// we know that t.val is of type number, so return the
		// int value
		return int(val), nil
	case operator:
		// we know that t.val is of type operator, so
		// find the values of the left and right children, then
		// call the process() method on operator to return the
		// result of processing their values.
		left, err := walkTree(t.lchild)
		if err != nil {
			return 0, err
		}
		right, err := walkTree(t.rchild)
		if err != nil {
			return 0, err
		}
		return val.process(left, right), nil
	default:
		// if a new treeVal type is defined, but walkTree wasn't updated
		// to process it, this detects it
		return 0, errors.New("unknown node type")
	}
}
```

## Function Types Are a Bridge to Interfaces


## Implicit Interfaces Make Dependency Injection Easier

## Go Isn't Particularly Object-Oriented(and That's Great)

## Exercises

## Wrapping Up
