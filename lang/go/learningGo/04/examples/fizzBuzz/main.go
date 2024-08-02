package main

import "fmt"

func main() {
	fmt.Println("confusingFizzBuzz:")
	confusingFizzBuzz()
	fmt.Println("simpleFizzBuzz:")
	simpleFizzBuzz()
}

func confusingFizzBuzz() {
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
}

func simpleFizzBuzz() {
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
}
