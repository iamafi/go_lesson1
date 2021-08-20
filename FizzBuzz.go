package main

import "strconv"

func fizzbuzz(n int) string {
	var str string
	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			str += "fizzbuzz "
		} else if i % 3 == 0 {
			str += "fizz "
		} else if i % 5 == 0 {
			str += "buzz "
		} else {
			str += strconv.Itoa(i) + " "
		}
	}
	return str
}
