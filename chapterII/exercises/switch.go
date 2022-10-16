package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Wednesday
	switch dayBorn {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Not Monday or Tuesday")
	}
}
