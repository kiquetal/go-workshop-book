package main

import (
	"errors"
	"fmt"
)

func validate(input int) error {

	if input < 0 {
		return errors.New("input must be greater than 0")
	} else if input > 100 {
		return errors.New("input must be less than 100")
	} else if input%7 == 0 {
		return errors.New("input must not be divisible by 7")
	} else {
		return nil
	}

}
func main() {

	input := 22
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("input is valid")
	}

}
