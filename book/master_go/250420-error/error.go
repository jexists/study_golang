package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	err := check(0, 10)
	if err == nil {
		fmt.Println("check() ended")
	} else {
		fmt.Println(err)
	}
	// check() ended

	err = check(0, 0)
	if err.Error() == "this is a custom error" {
		fmt.Println("custom error")
	}
	// custom error

	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	// a 0 and b 0. UserID: 50

	i, err := strconv.Atoi("-123")
	if err == nil {
		fmt.Println("int value is", i)
	}
	// int value is -123

	_, err = strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	}
	// strconv.Atoi: parsing "Y123": invalid syntax

}

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is a custom error")
	}
	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}
