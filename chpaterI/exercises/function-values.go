package main

import "fmt"

func add5Value(count int) {
	count += 5
	fmt.Printf("add5value :%#v\n", count)
}
func add5Point(count *int) {
	*count += 5
	fmt.Printf("add5point :%#v\n", *count)
}

func main() {

	var count int
	add5Value(count)
	add5Point(&count)
	fmt.Printf("add5Point post:%#v", count)
}
