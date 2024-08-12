package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	p := MakePerson("Fred", "Williamson", 25)
	fmt.Println(p)
	p2 := MakePersonPointer("Wilma", "Fredson", 32)
	fmt.Println(p2)
}

/*
* Compile the code with -gcflags="m" to show when values escape to the heap.
* It also shows when functions are inlined, this is a compiler optimization.
*
❯ go build -gcflags="-m" ./main.go
# command-line-arguments
./main.go:11:6: can inline MakePerson
./main.go:19:6: can inline MakePersonPointer
./main.go:28:17: inlining call to MakePerson
./main.go:29:13: inlining call to fmt.Println
./main.go:30:25: inlining call to MakePersonPointer
./main.go:31:13: inlining call to fmt.Println
./main.go:11:17: leaking param: firstName to result ~r0 level=0
./main.go:11:28: leaking param: lastName to result ~r0 level=0
./main.go:19:24: leaking param: firstName
./main.go:19:35: leaking param: lastName
./main.go:20:9: &Person{...} escapes to heap
./main.go:29:13: ... argument does not escape
./main.go:29:14: p escapes to heap
./main.go:30:25: &Person{...} escapes to heap
./main.go:31:13: ... argument does not escape
*
* The &Person{} returned from MakePersonPointer escapes to the heap. Any time
* a pointer is returned from a function, a pointer is returned on the stack, but
* the value it points to must be stored on the heap.
*
* Surprisingly, it also says that p escapes to the heap. The reason for this is
* that it is passed to fmt.Println. This is because the paramters to fmt.Println
* are ...any. The current Go compiler moves to the heap any value that is passed
* in to a function via a paramter that is of an interface type.
*/
