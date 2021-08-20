package main

func palindrome(str string) bool {
	l := len(str)
	for i := 0; i <= l / 2; i++ {
		if str[i] != str[l - i - 1] {
			return false
		}
	}
	return true
}
