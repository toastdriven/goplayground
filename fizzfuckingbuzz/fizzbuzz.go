package fizzbuzz

import (
	"fmt"
)

func CheckFizzBuzz(value int) string {
	var result string
	switch {
	case value % 15 == 0:
		result = "FizzBuzz"
	case value % 3 == 0:
		result = "Fizz"
	case value % 5 == 0:
		result = "Buzz"
	default:
		result = fmt.Sprintf("%v", value)
	}
	return result
}

func FizzBuzz(max int) {
	for i := 1; i <= max; i++ {
		fmt.Println(CheckFizzBuzz(i))
	}
}
