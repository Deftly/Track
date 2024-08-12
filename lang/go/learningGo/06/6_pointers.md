# Pointers

<!--toc:start-->
- [Pointers](#pointers)
  - [Pointers Indicate Mutable Parameters](#pointers-indicate-mutable-parameters)
  - [Pointers Are a Last Resort](#pointers-are-a-last-resort)
  - [The Difference Between Maps and Slices](#the-difference-between-maps-and-slices)
  - [Slices as Buffers](#slices-as-buffers)
  - [Reducing the Garbage Collector's Workload](#reducing-the-garbage-collectors-workload)
  - [Tuning the Garbage Collector](#tuning-the-garbage-collector)
  - [Exercises](#exercises)
  - [Wrapping Up](#wrapping-up)
<!--toc:end-->

A pointer is a variable that contains the address where another variable is stored:
```go
var x int32 = 10
var y bool = true
pointerX := &x
pointerY := &y
var pointerZ *string
```
![pointersInMemory](./assets/pointersInMemory.png)

While different types of variables can take up different numbers of memory locations, every pointer, no matter what type it is pointing to, is always the same number of memory locations as seen in the example above.

The zero value for a pointer is `nil`. We've seen this before as the zero value for slices, maps, and functions. All these types are implemented with pointers(channels and interfaces are also implemented with pointers). Unlike `NULL` in C, `nil` is not another name for 0; `nil` is an untyped identifier that represents the lack of a value.

The `&` is the *address* operator. It precedes a value type and returns the address where the value is stored:
```go
x := hello
pointerToX := &x
```

The `*` is the *indirection* operator. It precedes a variable of pointer type and returns the pointed-to value. This is called *dereferencing*:
```go
x := hello
pointerToX := &x
fmt.Println(pointerToX)  // A memory address
fmt.Println(*pointerToX) // 10
z := 5 + *pointerToX
fmt.Println(z) // 15
```

Before dereferencing a pointer, you must make sure that the pointer is non-nil. Your program will panic if you attempt to dereference a `nil` pointer:
```go
var x *int
fmt.Println(x == nil) // true
fmt.Println(*x)       // panics
```

A *pointer type* is a type that represents a pointer. It is written with a `*` before a type name.

For structs, use an `&` before a struct literal to create a pointer instance. You can't use an `&` before a primitive literal(numbers, booleans, and strings) or a constant because they don't have memory addresses and they exist only at compile time:
- When we say primitive values don't have memory addresses, we're referring to how these values are handled by the compiler and at runtime. In many programming languages, including Go, primitive values are often treated differently from complex types like structs or slices.
  - Complex types typically have a definite location in memory where their data is stored. 
  - Primitive values might be stored in CPU registers, inlined into the machine code, or handled in other optimized ways.
- This optimization is done because primitive values are small and frequently used so treating them like full-fledged objects with memory addresses would be inefficient.

When you need a pointer to a primitive type, declare a variable and point to it:
```go
x := &Foo{}
var y string
z := &y
```

If you have a struct with a field of a pointer to a primitive type, you can't assign a literal directly to the field for the reasons discussed above:
```go
type person struct {
  Firstname  string
  MiddleName *string
  LastName   string
}

p := person{
  Firstname:  "Pat",
  MiddleName: "Perry", // cannot use "Perry" (untyped string constant) as *string value in struct literal
  LastName:   "Peterson",
}
```

If you try putting an `&` before `"Perry"`, you'll get the following error: `invalid operation: cannot take address of "Perry" (untyped string constant)`

There are two ways around this problem. The first is to do what we've seen before which is to introduce a variable to hold the constant value. The second is to write a generic helper function that takes in a parameter of any type and returns a pointer to that type:
```go
func makePointer[T any](t T) *T {
	return &t
}
...
p := person{
  Firstname:  "Pat",
  MiddleName: makePointer("Perry"), // This works
  LastName:   "Peterson",
}
```

This works because passing a constant to a function creates a copy and assigns it to the parameter which is a variable. And since it's a variable, it has an address in memory which we return from the function.

## Pointers Indicate Mutable Parameters
Since Go is a call-by-value language, the values passed to functions are copies. For non-pointer types like primitives, structs, and arrays, this means that the called function cannot modify the original.

When a pointer is passed to a function, the function gets a copy of the pointer. This still points to the original data, which means that the original data can be modified by the called function.

The first implication of this is that when you pass a `nil` pointer to a function, you cannot make the value non-nil. This might be confusing at first, but it makes sense. Since the memory location was passed to the function via call-by-value, you can't change the memory address, any more than you could change the value of an `int` parameter. Here is an example:
```go
func failedUpdate(g *int) {
	x := 10
	g = &x
}

func main() {
	var f *int // f is nil
	failedUpdate(f)
	fmt.Println(f) // nil
}
```

The flow of this code is shown below:
![failNilPointerUpdate](./assets/failNilPointerUpdate.png)

The second implication of copying a pointer is that if you want the value assigned to a pointer parameter to still be there when you exit the function, you must dereference the pointer and set the value. If you change the pointer, you have changed the copy, not the original. De-referencing puts the new value in the memory location pointed to by both the original and the copy:
```go
func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func main() {
	x := 10
	failedUpdate(&x)
	fmt.Println(x) // 10
	update(&x)
	fmt.Println(x) // 20
}
```

The flow of this code is shown below
![updatingPointers](./assets/updatingPointers.png)

## Pointers Are a Last Resort
We should be careful about when we use pointers in Go. They make it harder to understand data flow and can create extra work for the garbage collector. For example, rather than populating a struct by passing a pointer to it into a function, we have the function instantiate and return the struct:
```go
// Don't do this
func MakeFoo1(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

// Do this
func MakeFoo2() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}
```

The only time you should use pointer parameters to modify a variable is when the function expects an interface.

When returning values from a function you should favor value types. Use a pointer type as a return type only if there is state within the data type that needs to be modified.

## The Difference Between Maps and Slices
Now that we know a little about pointers we can understand why modifications made to a map that's passed to a function are reflected in the original variable. Within the Go runtime, a map is implemented as a pointer to a struct.

Passing a slice to a function has more a complicated behavior: any modification to the slice's contents is reflected in the original variable, but using `append` to change the length isn't reflected in the original variable even if the slice has a capacity greater than its length. This is because a slice is implemented as a struct with three fields: two `int` fields for length and capacity and a pointer to a block of memory.

![memLayoutOfSlice](./assets/memLayoutOfSlice.png)

When a slice is copied to a different variable or passed to a function, a copy is made of the length, capacity, and the pointer:

![sliceCopy](./assets/sliceCopy.png)

Changing the values in the slice changes the memory that the pointer points to, so the changes are seen in both the copy and original:

![modifySlice](./assets/modifySlice.png)

If the slice copy is appended to and there is enough capacity in the slice for new values, the length changes in the copy, and the new values are stored in the block of memory that's shared. However, the length in the original slice remains unchanged. The Go runtime prevents the original slice from seeing those values since they are beyond the length of the original slice:

![appendSlice](./assets/appendSlice.png)

If the slice copy is appended to and there isn't enough capacity a new bigger block of memory is allocated, values are copied over, and the length, capacity, and pointer fields in the copy are updated. These changes are not reflected in the original:

![sliceCapacityChange](./assets/sliceCapacityChange.png)

Slices are frequently passed around in Go programs and by default you should assume that a slice is not modified by a function. If your function does modify a slice this should be specified in your function's documentation.

The ability to modify the contents(but not the size) of a slice input parameter makes them ideal for use as a reusable buffer.

## Slices as Buffers
When reading data from an external resource(like a file or a network connection), man languages use code like this:
```
r = open_resource()
while r.has_data() {
  data_chunk = r.next_chunk()
  process(data_chunk)
}
```

The problem with this pattern is that every time you iterate through the `while` loop you allocate another `data_chunk` even though each one is used only once, creating lots of unnecessary memory allocations.

Writing idiomatic Go avoids means avoiding unneeded allocations. Rather than returning a new allocation each time we read from a data source, we create a slice of bytes once and use it as a buffer to read data from the data source:
```go
f, err := os.Open(fileName)
if err != nil {
  return err
}
defer f.Close()
data := make([]byte, 100)
for {
  count, err := f.Read(data)
  process(data[:count])
  if err != nil {
    if errors.Is(err, io.EOF) {
      return nil
    }
    return err
  }
}
```

## Reducing the Garbage Collector's Workload
If you've spent time learning how programming languages are implemented, you've probably learned about the *heap* and the stack. A *stack* is a consecutive block of memory. Every function call in a thread of execution shared the same stack. Allocating memory on the stack is fast and simple. A *stack pointer* tracks the last location where memory was allocated. Allocating additional memory is done by changing the value of the stack pointer. When a function is invoked, a new *stack frame* is created for the function's data. Local variables are stored on the stack, along with parameters passed into a function. Each new variable moves the stack pointer by the size of the value. When a function exits, its return values are copied back to the beginning of the stack frame from the exited function, deallocated all the stack memory that was used by that function's local variables and parameters.

To store something on the stack, you have to know exactly how big it is at compile time. All the value types in Go(primitive values, arrays, and structs) have a known size at compile time and can be stored on the stack.

The size of a pointer type is also know, and can be stored on the stack. The rules are more complicated when it comes to the data that the pointer points to. For Go to allocate the data the pointer points to on the stack, several conditions must be true:
  - The data must be a local variable whose data size is known at compile time.
  - The pointer cannot be returned from the function. 
  - If the pointer is passed into a function, the compiler must be able to ensure that these conditions still hold

When the compiler determines that the data can't be stored on the stack, we say that the data the pointer points to *escapes* the stack, and the compiler stores the data on the heap.

The heap is the memory that's managed by the garbage collector(GC). Any data that is stored on the heap is valid as long as it can be tracked back to a pointer type variable on a stack. Once there are no more variables on the stack pointing to the data, the data becomes *garbage*, and the GC needs to clear it out.

There are two main performance related problems with storing things on the heap. The first is that the GC takes time to do its work. It isn't trivial to track all available chunks of free memory on the heap or to track which used blocks of memory still have valid pointers. This is time that's taken away from doing the processing your program is written to do. The Go runtime has a GC that favors low latency(finish the garbage scan as quickly as possible instead of finding the most garbage possible). Each garbage-collection cycle is designed to "stop the world" for fewer than 500 microseconds. However, if your Go program creates lots of garbage, the GC won't be able to gind all the garbage during the cycle, slowing the collector and increasing memory usage.

The second problem deals with the nature of computer hardware. While RAM means "random access memory" the fastest way to read from memory is to read it sequentially. A slice of structs in Go has all the data laid out sequentially in memory. This makes it fast to load and process. A slice of pointers to structs (or structs whose fields are pointers) has its data scattered across RAM, making it far slower to read and process.

For these reasons Go encourages us to use pointers sparingly. This reduces the GC's workload and when the GC does do work, it is optimized to return quickly rather than gather the most garbage. The key to making this approach work is to create less garbage in the first place.

To learn more about heap versus stack allocation and escape analysis in Go check out this [addendum](../addendums/stack_allocation_escape_analysis.md)

## Tuning the Garbage Collector
A GC doesn't immediately reclaim memory as soon as it is no longer referenced, doing so would seriously impact performance. Instead, it lets garbage pile up for a bit. The heap almost always contains both live data and memory that's no longer needed. The Go runtime provides users a couple of settings to control the heap's size.

The first is the `GOGC` environment variable. The GC looks at the heap size at the end of a garbage collection cycle and uses the formula `CURRENT_HEAP_SIZE + CURRENT_HEAP_SIZE*GOGC/100` to calculate the heap size that needs to be reached to trigger the next GC cycle. By default, `GOGC` is set to `100`, which means that the heap size that triggers the next collection is roughly double the heap size at the end of the current collection.

The second GC setting specifies a limit on the total amount of memory your Go program is allowed to use. By default, it is disabled(technically set to `math.MaxInt64`, but it's unlikely that your computer has that much memory). The value for `GOMEMLIMIT` is specified in bytes, but you can optionally use the suffixes `B`, `KiB`, `MiB`, `GiB`, and `TiB`. For example, `GOMEMLIMIT=3GiB` sets the memory limit to 3 gigabytes(3,221,225,472 bytes).


It might seem counter intuitive that limiting the maximum amount of memory could improve a program's performance but there's a good reason why this flag was added. Computers don't have infinite memory. So, if a sudden, temporary spike occurs in memory usage, relying on `GOGC` alone might result in the maximum heap size exceeding the amount of available memory. This can cause memory to swap to disk, which is very slow, and in some cases may even cause your program to crash. Specifying a maximum memory limit prevents the heap from growing beyond the computer's resources.

`GOMEMLIMIT` is a *soft* limit that can be exceeded in certain circumstances. A common problem that can occur is when the collector is unable to free enough memory to get within a memory limit or the garbage-collection cycles are rapidly being triggered because a program is repeatedly hitting the limit. This is called *thrashing*, and results in a program that does nothing other than run the garbage collector. If the Go runtime detects thrashing it chooses to end the current garbage-collection cycle and exceed the limit. This does mean that you should set `GOMEMLIMIT` below the absolute maximum amount of available memory so you have spare capacity.

To learn more about how to use `GOGC` and `GOMEMLIMIT` read [A Guide to the Go Garbage Collector](https://go.dev/doc/gc-guide) from Go's development team.

## Exercises
1. Create a struct named `Person` with three fields: `FirstName` and `Lastname` of type `string` and `Age` of type `int`. Write a function called `MakePerson` that takes in `firstName`, `lastName`, and `age` and returns a `Person`. Write a second function `MakePersonPointer` that takes in `firstName`, `lastName`, and `age` and returns a `*Person`. Call both from `main`. Compile the program with `go build -gcflags="-m"`. This both compiles the code and prints out which values escape to the heap. Take not of the results. [Solution](./exercises/ex1/main.go)
2. Write two functions. The `UpdateSlice` function takes in a `[]string` and a `string`. It sets the last position in in the passed-in slice to the passed-in `string`. At the end of `UpdatedSlice`, print the slice after making the change. The `GrowSlice` function also takes in a `[]string` and a `string`. It appends the `string` onto the slice. At the end of `GrowSlice`, print the slice after making the change. Call these functions from `main`. Print out the slice before each function is called and after each function is called. Do you understand why some changes are visible in `main` and why some changes are not? [Solution](./exercises/ex2/main.go)
3. Write a program that builds `[]Person` with 10,000,000 entries(they could all be the same names and ages). See how long it takes to run. Change the value of `GOGC` and see how that affects the time it takes for the program to complete. Set the environment variable `GODEBUG=gctrace=1` to see when garbage collections happen and see how changing `GOGC` changes the number of garbage collections. What happens if you create the slice with a capacity of 10,000,000? [Solution](./exercises/ex3/main.go)

## Wrapping Up
This chapter covered pointers and when to use them. In the [next chapter](../07/7_types_methods_and_interfaces.md) we'll cover Go's implementation of methods, interfaces and types. 
