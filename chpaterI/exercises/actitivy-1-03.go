package main

import "fmt"

func main() {
	count := 5
	message := ""
	if count > 5 {
		message = "count is greater than 5"
	} else {
		message = "Not greater than 5"
	}
	fmt.Println(message)
}
