package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%v\n ", fizzbuzz(i))
	}
}

func fizzbuzz(input int) string {

	if input%3 == 0 && input%5 == 0 {
		return "FizzBuzz"
	} else if input%3 == 0 {
		return "Fizz"

	} else if input%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(input)
	}
}
