package main

import "fmt"

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
