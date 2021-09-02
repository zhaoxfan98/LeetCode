package easy

import "strconv"

func fizzBuzz(n int) (strs []string) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			strs = append(strs, "FizzBuzz")
		} else if i%3 == 0 {
			strs = append(strs, "Fizz")
		} else if i%5 == 0 {
			strs = append(strs, "Buzz")
		} else {
			strs = append(strs, strconv.Itoa(i))
		}
	}

	return
}
