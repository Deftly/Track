# Composite Types

<!--toc:start-->
- [Composite Types](#composite-types)
  - [Arrays-Too Rigid to Use Directly](#arrays-too-rigid-to-use-directly)
  - [Slices](#slices)
    - [len](#len)
    - [append](#append)
    - [Capacity](#capacity)
    - [make](#make)
    - [Emptying a Slice](#emptying-a-slice)
    - [Declaring Your Slice](#declaring-your-slice)
    - [Slicing Slices](#slicing-slices)
    - [copy](#copy)
    - [Converting Arrays to Slices](#converting-arrays-to-slices)
    - [Converting Slices to Arrays](#converting-slices-to-arrays)
  - [Strings, Runes, and Bytes](#strings-runes-and-bytes)
  - [Maps](#maps)
    - [Reading and Writing a Map](#reading-and-writing-a-map)
    - [The comma ok Idiom](#the-comma-ok-idiom)
    - [Deleting from Maps](#deleting-from-maps)
    - [Emptying from Maps](#emptying-from-maps)
    - [Comparing Maps](#comparing-maps)
    - [Using Maps as Sets](#using-maps-as-sets)
  - [Structs](#structs)
    - [Anonymous Structs](#anonymous-structs)
    - [Comparing and Converting Structs](#comparing-and-converting-structs)
  - [Exercises](#exercises)
  - [Wrapping Up](#wrapping-up)
<!--toc:end-->

## Arrays-Too Rigid to Use Directly
Go has arrays like most languages, however they are rarely used directly and we'll see why shortly.

All elements in an array must be of the type that's specified and there are a few different declaration styles:
```go
var x [3]int // all elements are initialized to the zero value of the specified type
var y = [3]int{10, 20, 30}
var z = [12]int{1, 5: 4, 6, 10: 100, 15} // [1 0 0 0 0 4 6 0 0 0 100 15]
var a = [...]int{10, 20, 30} // When using an array literal to initialize you can use ...
```

You can use `==` and `!=` to compare two arrays. Arrays are equal if they are the same length and contain the same values:
```go
var x = [...]int{1, 2, 3}
var y = [3]int{1, 2, 3}
fmt.Println(x == y) // true
```

Go has only one-dimensional arrays, but you can simulate multidimensional arrays like so:
```go
var x [2][3]int
```
This declares that `x` is an array of length 2 whose type is array of length 3. Other languages like Julia have true matrix support, Go does not.

Arrays in Go are read and written to using bracket syntax:
```go
x[0] = 10
fmt.Println(x[2])
```
You cannot read or write past the end of an array or use a negative index(🐍). If you do this with a constant or literal index, it is a compile time error. An out-of-bounds read or write with a variable index will fail at runtime with a *panic*.

The built-in function `len` takes in an array and returns its length:
```go
fmt.Println(len(x))
```

Arrays in Go have an unusual limitation, Go considers the *size* of the array to be part of the *type* of the array. An array declared to be `[3]int` is a different type than `[4]int`. This also means you cannot specify the size of an array with a variable because types must be resolve at compile time.

Additionally, you can't use a type conversion to directly convert array of different sizes to identical types. This means that you can't write a function that works with arrays of any size and you can't assign arrays of different sizes to the same variable.

> **_NOTE:_** We'll cover how arrays work behind the scenes when we cover memory layout in [chapter 6](../06/6_pointers.md).

Because of these restrictions only use arrays if you know the exact length you need ahead of time.

The main reason why arrays exist in Go is to provide the backing store for *slices*.

## Slices
Most of the time when you want a data structure that holds a sequence of values you'll use a slice. The length of a slice is *not* part of its type, removing the biggest limitation of arrays and allowing you to write a single function that processes slices of any size.

Working with slices is very similar to working with arrays but with some subtle differences:
```go
var x = []int{10, 20, 30} // creates a slice of ints using a slice literal
var y = []int{1, 5: 4, 6, 10: 100, 15} // [1 0 0 0 0 4 6 0 0 0 100 15]
x[0] = 24
fmt.Println(x[2])
```

Another difference between slices and arrays can be seen when declaring a slice without using a literal:
```go
var x []int
```
This creates as lice of `int`s, and since no value is assigned, `x` is assigned the zero value for a slice,`nil`. In Go, `nil` is an identifier that represents the lack of a value for some types. Like numeric constants covered in the previous chapter, `nil` has no type, so it can be assigned or compared against values of different types. A `nil` slice contains nothing.

A slice is the first type we've covered that isn't *comparable*. It is a compile time error to use `==` or `!=` to compare slices. The only thing you can compare a slice with using `==` is `nil`.

Since Go 1.21, the `slices` package in the standard library includes two functions to compare slices. The `slices.Equal` function takes in two slices and returns `true` if the slices are the same length, and all of the elements are equal. It requires the elements of the slice to be comparable.

The other function, `slices.EqualFunc`, lets you pass in a function to determine equality and does not require the slice elements to be comparable.
```go
x := []int{1, 2, 3, 4, 5}
y := []int{1, 2, 3, 4, 5}
z := []int{1, 2, 3, 4, 5, 6}
s := []string{"a", "b", "c"}
fmt.Println(slices.Equal(x, y)) // prints true
fmt.Println(slices.Equal(x, z)) // prints false
fmt.Println(slices.Equal(x, s)) // does not compile
```

> **_WARNING:_** The `reflect` package contains a function called `DeepEqual` that can compare almost anything, including slices. Before the inclusion of `slices.Equal` and `slices.EqualFunc`, `reflect.DeepEqual` was often used to compare slices. Don't use it in new code as it is slower and less safe than using the functions in the `slices` package.

### len
Go provides several built-in functions that work with slices. We saw the built-in `len` function earlier when looking at arrays. It works for slices too. Passing a `nil` slice to `len` returns 0.

> **_NOTE:_** Functions like `len` are built into Go because they can do things that can't be done by the functions that you can write. `len`'s parameter can be any array, slice, string, map, or even a channel. Trying to pass a variable of any other type to `len` is a compile-time error. We'll see in [chapter 5](../05/5_functions.md), Go doesn't let developers write a function that accepts any string, array, slice, channel, or map, but rejects other types.

### append
The built-in `append` function is used to grow slices:
```go
var x []int
x = append(x, 10) // assign result to the variable that's passed in
x = append(x, 5, 6, 7, 8) // can append multiple values at once
y := []int{20, 24, 44}
x = append(x, y...) // slices can appended onto another using ... 
```
The `append` function takes at least two parameters, a slice of any type and a value of that type. It returns a slice of the same type, which is assigned to the variable that was passed to `append`.

It is a compile time error if you forget to assign the value returned from `append`. This is because every time you pass a parameter to a function, Go makes a copy of the value that's passed in. Passing a slice to the `append` function actually passes a copy of the slice to the function. The functions adds the values to the copy of the slice and returns the copy. We then assign the returned slice back to the variable in the calling function. We;ll talk about this in more detail in [chapter 5](../05/5_functions.md) when we talk about Go being a *call-by-value* language.

### Capacity
Each element in a slice is assigned to consecutive memory locations, making it quick to read or write values. The length of a slice is the number of consecutive memory locations that have been assigned a value. Every slice also has a *capacity*, which is the number of consecutive memory locations reserved. The capacity can be larger than the length. Each time you append to a slice, one or more values are appended to the end of the slice. If you try to add values when the length equals the capacity, the `append` function uses the Go runtime to allocate a new backing array for the slice with a larger capacity. The values in the original backing array are copies to the new one, the new values are added to the end of the new backing array, and the slice is updated to refer to the new backing array. Finally the updated slice is returned.

> **_NOTE:_** The Go runtime provides services like memory allocation and garbage collection, concurrency support, networking, and implementations of built-in types and functions.
>
> The Go runtime is compiled into every Go binary. This is different than languages that use a virtual machine, which must be installed separately to allow programs written in those languages to function. Including the runtime in the binary makes it easier to distribute Go programs and avoids worries about compatibility issues between the runtime and the program. The drawback is that even the simplest Go program produces a binary that's about 2 MB.

When slices grow via `append`, it takes time for the Go runtime to allocate new memory and copy the existing data from the old memory to the new. The old memory also needs to be garbage collected. For this reason, the Go runtime usually increases a slice by more than one each time it runs out of capacity. As of Go 1.18 the rule is to double the capacity of a slice when the current capacity is less than 256. A bigger slice increases by `(current_capacity + 768) / 4`. This slowly converges to a 25% growth rate.

The built-in `cap` function returns the current capacity of a slice. 

While it's nice that slices grow automatically, it's far more efficient to size them once. If you know how many things you plan to put into a slice, create it with the correct initial capacity with the `make` function.

### make
The built-in `make` function allows you to specify the type, length, and optionally, the capacity of a slice:
```go
// Creates an int slice with a length of 5 and a capacity of 5
// x[0] through x[4] are initialized to 0
x := make([]int, 5) 
```

A common beginner mistake is to try to populate those initial elements using `append`:
```go
x := make([]int, 5)
// 10 is placed after the zeros, append always increases the length of the slice
x = append(x, 10) // [0 0 0 0 0 10]
```

You can also create a slice with zero length but a capacity that's greater than zero:
```go
x := make([]int, 0, 10) // non-nil slice with a length 0 and capacity of 10
// Since the length is 0 you can't directly index into it, but you can append values
x = append(x, 5, 6, 7, 8) // [5 6 7 8] len: 4 cap: 10
```

### Emptying a Slice
Go 1.21 added a `clear` function that takes in a slice and sets all of the slice's elements to their zero value. The length of the slice remains unchanged:
```go
s := []string{"first", "second", "third"}
fmt.Println(s, len(s)) // [first second third] 3
clear(s)
fmt.Println(s, len(s)) // [   ] 3
```

### Declaring Your Slice
When choosing how to declare a slice the primary goal should be to reduce the number of times the slice needs to grow.

If it's possible that the slice won't need to grow at all use a `var` declaration with no assigned value to create a `nil` slice:
```go
var data []int
```

If you have some starting values, or if a slice's values aren't going to change, then a slice literal is a good choice:
```go
data := []int{2, 4, 6, 8}
```

If you have an idea of how large your slice needs to be but don't know the values you can use `make`. The question is whether to specify a nonzero length in the call to `make` or specify a zero length and a nonzero capacity:
  - If you are using a slice as a buffer, then specify a nonzero length.
  - If you are *sure* you know the exact size you want, you can specify the length and index into the slice to set the values. The downside to this approach is that if you have the size wrong, you'll either end up with zero values at the end of your slice or a panic from trying to access elements that don't exist.
  - In other situations, use `make` with a zero length and specified capacity. This allows you to use `append` to add items the slice. If the number of items turns out to be smaller you won't have extra zero values at the end. If the number of items is larger, you code will not panic.

### Slicing Slices
A *slice expression* creates a slice from a slice. It's written inside brackets and consists of a starting offset and ending offset, separated by a colon(:). The starting offset is the first position in the slice that is included in the new slice, and the ending offset is one past the last position to include. If the starting offset or ending offset are not specified 0 and the end of the slice are used respectively.
```go
x := []string{"a", "b", "c", "d"}
y := x[:2]  // [a b]
z := x[1:]  // [b c d]
d := x[1:3] // [b c]
e := x[:]   // [a b c d]
```

When you take a slice from a slice you are *not* making a copy of the data. The slices are sharing memory and changes to one slice will affect all slices that share that memory:
```go
x := []string{"a", "b", "c", "d"}
y := x[:2]
z := x[1:]
x[1] = "y"
y[0] = "x"
z[1] = "z"
fmt.Println(x) // [x y z d]
fmt.Println(y) // [x y]
fmt.Println(z) // [y z d]
```

This can get extra confusing when using `append` to add elements to slices that share memory. For this reason you should either never use `append` with a sub-slice or make sure that `append` doesn't cause an overwrite by using a *full slice expression*. This makes it clear how much memory is shared between the parent slice and sub-slice:
```go
x := make([]string, 0, 5)
x = append(x, "a", "b", "c", "d")
y := x[:2:2]  //
z := x[2:4:4]
```

In the above example both `y` and `z` have a capacity of 2. This because we limited the capacity to of the sub-slices to their lengths, appending additional elements onto `y` and `z` creates new slices that don't interact with the other slices.

### copy
If we need to create a slice that's independent of the original we can use the built-in `copy` function:
```go
x := []int{1, 2, 3, 4}
y := make([]int, 4)
num := copy(y, x)
fmt.Println(y, num) // [1 2 3 4] 4
```

The `copy` function takes two parameters. The first is the destination slice, and the second is the source slice. The function copies as many values as it can from source to destination, limited by whichever slice is smaller, and returns the number of elements copied.

You can also copy from the middle of the source slice:
```go
x := []int{1, 2, 3, 4}
y := make([]int, 2)
copy(y, x[2:])
```

### Converting Arrays to Slices
If you have an array, you can take a slice from it using a slice expression:
```go
xArray := [4]int{5, 6, 7, 8}
xSlice := xArray[:] // convert an entire array into a slice
y := xArray[:2]
z := xArray[2:]
```

Be aware that taking a slice from an array has the same memory-sharing properties as taking a slice from a slice.

### Converting Slices to Arrays
When you convert a slice to an array, the data in the slice is copied to new memory, meaning changes to the slice won't affect the array and vice versa.
```go
xSlice := []int{1, 2, 3, 4}
xArray := [4]int(xSlice)
smallArray := [2]int(xSlice)
xSlice[0] = 10
fmt.Println(xSlice)     // [10 2 3 4]
fmt.Println(xArray)     // [1 2 3 4]
fmt.Println(smallArray) // [1 2]
```

The size of the array must be specified at compile time. It's a compile time error to use `[...]` in a slice to array type conversion. And while the array can be smaller than the slice, it cannot be larger, this will cause a runtime panic.

## Strings, Runes, and Bytes
Go uses a sequence of bytes to represent a string, not runes. These bytes don't have to be in any particular character encoding, but several Go library functions assume that a string is composed of a sequence of UTF-8 encoded code points.

Just as you can extract a single value from an array or slice, you can extract a single value form a string. You can even use slice expression notation with strings:
```go
var s string = "Hello there"
var b byte = s[6]
var s2 string = s[4:7] // "o t"
var s3 string = s[:5]  // "Hello"
var s4 string = s[6:]  // "there"
```

While it can be nice that use index and slice notation with strings we have to be careful when doing so. Strings are immutable so they don't have the modification problem that slices do. The problem is that a string is composed of a sequence of bytes, while a code point in UTF-8 can be anywhere from one to four bytes long. The previous example was composed of code points that were all one byte long so everything worked. That isn't the case in our next example:
```go
var s string = "Hello 🌞"
var s2 string = s[4:7]
var s3 string = s[:5]
var s4 string = s[6:]
```

If you print out these values you'll notice a strange value for `s2`, this is because here we are only copying the first byte of the emoji's code point, which is not a valid code point on its own. If we were to pass `s` to the built-in `len` function we'll see that it returns 10 instead of 7 because it takes four bytes to represent the emoji in this example.

Because of this relationship between runes, strings, and bytes, Go has some interesting type conversions between these types. A single rune or byte can be converted to a string:
```go
var a rune = 'x'
var s string = string(a)
var b byte = 'y'
s = string(b)
```

> **_WARNING:_** A common bug for new Go developers i to try to make an `int` into a string using a type conversion:
>
> ```go
> var x int = 65
> var y = string(x)
> fmt.Println(y) // This prints "A" not "65"
> ```

A string can be converted back and forth to a slice of bytes or a slice of runes:
```go
var s string = "Hello, 🌞"
var bs []byte = []byte(s)
var rs []rune = []rune(s)
fmt.Println(bs) // [72 101 108 108 111 44 32 240 159 140 158]
fmt.Println(rs) // [72 101 108 108 111 44 32 127774]
```

Rather than use the slice and index expressions with strings, you should extract substrings and code points from strings using the functions in the `strings` and `unicode/utf8` packages in the standard library. In the [next chapter](../04/4_blocks_shadows_and_control_structures.md) we'll see how to use the `for-range` loop to iterate over the code points in a string.

## Maps
Map is a built-in type for situations where you want to associate one value to another and is written as `map[keyType]valueType`. There are a few different ways to declare a map
```go
// the zero value for a map is nil. A nil map has a length of 0, attempting to read from
// a nil map always returns the zero value for the map's value type. However attempting 
// to write to a nil map variable causes a panic.
var nilMap map[string]int // nil maps can be used for lazy initialization, saving memory

totalWins := map[string]int{} // initialization with empty map literal.

// Nonempty map literal
teams := map[string][]string {
  "Lakers": []string{"Kobe Bryant", "Jerry West", "Elgin Baylor"},
  "Suns": []string{"Steve Nash", "Amar'e Stoudemire", "Boris Diaw"},
  "Mavericks": []string{"Dirk Nowitzki", "Luka Doncic", "Jason Kidd"}
}
```

If you know how many key-value pairs you want to put in a map but don't know the exact values you can use `make` to create a map with a default size:
```go
ages := make(map[int][]string, 10)
```

Maps created with `make` still have a length of 0 and can grow past the initially specified size.

Here are some properties of maps that are similar to slices:
- Maps automatically grow as you add key-value pairs to them.
- If you know how many key-value pairs you plan to insert into a map, you can use `make` to create a map with a specific initial size.
- Passing a map to the `len` function tells you the number of key-value pairs in a map.
- The zero value for a map is `nil`.
- Maps are not comparable. You can check if they are equal to `nil`, but you cannot check if two maps have identical keys and values using `==` or differ using `!=`.

The key for a map can be any comparable type. This means you *cannot* use a slice or a map as the key for a map.

### Reading and Writing a Map
Let's look at an example to see how to write and read from a map:
```go
totalWins := map[string]int{}
totalWins["Warriors"] = 1
totalWins["Nuggets"] = 2
fmt.Println(totalWins["Warriors"]) // 1
fmt.Println(totalWins["Pistons"])  // 0
totalWins["Pistons"]++
fmt.Println(totalWins["Pistons"]) // 1
totalWins["Nuggests"] = 3
fmt.Println(totalWins["Nuggests"]) // 3
```

When you try to read the value assigned to a map key that was never set, the map returns the zero value for the map's value type. You can use the `++` operator to increment the numeric value for a map key.

### The comma ok Idiom
We've seen that maps return the zero value if you ask for the value associated with a key that's not in the map. While this behavior can be useful, sometimes we need to find out if a key is in a map. Go provides the *comma ok idiom* to tell the difference between a key that's associated with a zero value and a key that's not in the map:
```go
m := map[string]int{
  "hello": 5,
  "world": 0,
}
v, ok := m["hello"]
fmt.Println(v, ok) // 5 true

v, ok = m["world"]
fmt.Println(v, ok) // 0 true

v, ok = m["goodbye"]
fmt.Println(v, ok) // 0 false
```

The result of the map reads are assigned to two variables. The first gets the value associated with the key and the second value is a `bool` which indicates whether the key is present in the map or not. By convention the second variable is named `ok`.

### Deleting from Maps
Key-value pairs are removed from a map via the built-in `delete` function:
```go
m := map[string]int{
  "hello": 5,
  "world": 0,
}
delete(m, "hello")
v, ok := m["hello"]
fmt.Println(v, ok) // 0 false
```

If the key isn't present in the map or if the map is `nil`, nothing happens. The `delete` function doesn't return a value.

### Emptying from Maps
The `clear` function we saw earlier with slices also works on maps. A cleared map has its length set to 0 unlike a cleared slice:
```go
m := map[string]int{
  "hello": 5,
  "world": 10,
}
fmt.Println(m, len(m)) // map[hello:5 world:10] 2
clear(m)
fmt.Println(m, len(m)) // map[] 0
```

### Comparing Maps
Go 1.21 added a package to the standard library called `maps` that contains helper functions for working with maps. Two functions in the package are useful for comparing if two maps are equal, `maps.Equal` and `maps.EqualFunc`. They are analogous to the `slices.Equal` and `slices.EqualFunc` functions:
```go
m := map[string]int{
    "hello": 5,
    "world": 10,
}
n := map[string]int{
    "world": 10,
    "hello": 5,
}
fmt.Println(maps.Equal(m, n)) // true
```

### Using Maps as Sets
A *set* is a data type that ensures there is at most one of a value, but doesn't guarantee that the values are in any particular order. Checking to see if an element is in a set is fast, no matter how many elements are in the set.

Go doesn't include a set, but you can use a map to simulate some of its features. Use the key of the map for the type that you want to put into the set and use a `bool` for the value:
```go
intSet := map[int]bool{}
vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
for _, v := range vals {
  intSet[v] = true
}
fmt.Println(len(vals), len(intSet)) // 11 8
fmt.Println(intSet[5])   // true
fmt.Println(intSet[500]) // false
if intSet[10] {
  fmt.Println("10 is in the set") // This prints because intSet[10] returns true
}
```

If you need sets that provide operations like union, intersection, and subtraction, you can either write either write one yourself or use a third-party library that provides that functionality(We'll cover how to use third-party libraries in [chapter 10](../10/10_modules_packages_and_imports.md).)

## Structs
While maps are great for storing many kinds of data they have some limitations. The don't define an API since there's no way to constrain a map to allow only certain keys. Also, all values in a map must be of the same type. When you have related data that you want to group together, you should define a *struct*:
```go
type person struct {
  name string
  pet  string
  age  int
}
```

A struct type is defined with the keyword `type`, the name of the struct type, the keyword `struct`, and a pair of braces(`{}`). Within the braces, you list the fields in the struct. You can define a struct type inside or outside of a function. A struct type that's defined within a function can be used only within that function(more on functions in [chapter 5](../05/5_functions.md)).

Once a struct type is declared you define variables of that type:
```go
var fred person
bob := person{}
```

Unlike with maps there is no difference between assigning an empty struct literal and not assigning a value at all. Both initialize all fields in the struct to their zero values.

There are two styles for a nonempty struct literal. First is a comma separated list of values inside the braces:
```go
julia := person{
  "Julia",
  "cat",
  40,
}
```

When using this format, a value for every field in the struct must be specified and the value are assigned to the fields in the order they were declared in the struct definition.

The second struct literal style looks like the map literal style:
```go
beth := person{
  age:  30,
  name: "Beth",
}
```

In this style you use the names of the fields in the struct to specify the values. This allows you to specify the fields in any order, and you don't need to provide a value for every field. Any field not specified is set to its zero value.

A field in a struct is accessed with dot notation:
```go
bob.name = "Bob"
fmt.Println(bob.name)
```

### Anonymous Structs
You can declare that a variable implements a struct type without first giving the struct type a name. This is called an *anonymous struct*:
```go
var person struct {
  name string
  pet  string
  age  int
}

person.name = "Bob"
person.age = 50
person.pet = "dog"

pet := struct {
  name string
  kind string
}{
  name: "Aria",
  kind: "Cat",
}
```

In this example, the types of the variables `person` and `pet` are anonymous structs. 

Anonymous structs are handy in two common situations. The first is when you translate external data into a struct or a struct into external data(like JSON or Protocol Buffers). This is called *unmarshaling* and *marshaling* data, respectively. We'll see how to do this in [chapter 13](../13/13_standard_library.md).

Writing tests is another place where anonymous structs often pop up. We'll use a slice of anonymous structs when writing table-driven tests in [chapter 15](../15/15_writing_tests.md).

### Comparing and Converting Structs
Whether a struct is comparable or not depends on its fields. Structs composed entirely of comparable types are comparable. Those with slice, maps, or other fields that aren't comparable are not.

Go doesn't allow comparisons between variables that represent structs of different types, however, you can perform a type conversion from one struct type to another *if the fields of both structs have the same names, order, and types:
```go
type firstPerson struct {
	name string
	age  int
}

type secondPerson struct {
	name string
	age  int
}

type thirdPerson struct {
	age  int
	name string
}
```

You can use a type conversion to convert an instance of `firstPerson` to `secondPerson`, but you can't use `==` to compare them because they are of different types. You can't convert an instance of `firstPerson` to `thirdPerson` because the fields are in a different order.

If two struct variables are being compared and at least one has a type that's an anonymous struct, you can compare them without a type conversion if the fields of both structs have the same names, order, and types. You can also assign between named and anonymous struct types if the fields of both structs have the same names, order, and types:
```go
type firstPerson struct {
  name string
  age  int
}

f := firstPerson{
  name: "Bob",
  age:  50,
}

var g struct {
  name string
  age  int
}

g = f
fmt.Println(f == g)
```

## Exercises
1. Write a program that defines a variable named `greetings` of type slice of strings with the following values: `"Hello"`, `"Hola"`, `"नमस्कार"`, `"こんにちは"`, and `"Привіт"`. Create a sub-slice containing the first two values, a second subslice with the second, third, and fourth values, and a third subslice with the fourth and fifth values. Print out all four slices. [Solution](./exercises/ex1/ex1.go)
2. Write a program that defines a string variable called `message` with the value `"Hi 👩 and 👨"` and prints the 4th rune in it as a character, not a number. [Solution](./exercises/ex2/ex2.go)
3. Write a program that defines a struct called `Employee` with three fields: `firstName`, `lastName`, and `id`. The first two fields are of type `string` and the last field (`id`) is of type `int`. Create three instances of this struct using whatever values you'd like. Initialize the first one using the struct literal style without keys, the second using the struct literal style with keys, and the third with var declaration. Use dot notation to populate the fields in the third struct. Print out all three structs. [Solution](./exercises/ex3/ex3.go)

## Wrapping Up
This chapter covered strings in more detail and how to use the built-in generic container types: slices and maps. We also created our own composite types via structs. The [next chapter](../04/4_blocks_shadows_and_control_structures.md) will cover Go's control structures: `for`, `if/else`, and `switch`. We'll also look at how Go organizes code into blocks and how different block levels change the way our code behaves.
