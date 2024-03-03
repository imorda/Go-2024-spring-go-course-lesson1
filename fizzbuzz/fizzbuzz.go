package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(i int) (result string) {
	if i%3 == 0 {
		result += "Fizz"
	}
	if i%5 == 0 {
		result += "Buzz"
	}
	if result == "" {
		result = strconv.Itoa(i)
	}
	return
}
