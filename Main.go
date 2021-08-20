package main

import (
	"bufio"
	"fmt"
	"os"
)

//func WordCount(s string) map[string]int {
//	str := strings.Fields(s)
//	counter := make(map[string]int)
//	for _, el := range str {
//		counter[el] += 1
//	}
//	return counter
//}
//
var num int
var str string
var reader = bufio.NewReader(os.Stdin)


func main() {
	// Uncomment one or another program to check

	// Fibonacci
	f := fibonacci()
	fmt.Print("Enter number of fibonacci series: ")
	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		fmt.Print(f(), " ")
	}

	// Fizzbuzz
	//fmt.Print("Enter number for fizzbuzz: ")
	//fmt.Scanln(&num)
	//fmt.Println(fizzbuzz(num))

	// Palindrome
	//fmt.Print("Enter a word to check if it is palindrome: ")
	//fmt.Scanln(&str)
	//fmt.Println(palindrome(str))

	// Odd or even
	//fmt.Print("Enter number to check if it is odd or even: ")
	//fmt.Scanln(&num)
	//fmt.Println(oddEven(num))

	// Duplicates
	//fmt.Print("Enter a string to check if it contains duplicates (example: hello world hello): ")
	//str, _ := reader.ReadString('\n')
	//arr := strings.Fields(str)
	//ans := duplicate(arr)
	//if len(ans) == 0 {
	//	fmt.Println("No duplicates")
	//} else {
	//	for key := range ans {
	//		fmt.Printf("%s ", key)
	//	}
	//}
}