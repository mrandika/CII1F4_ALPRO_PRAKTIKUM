package main

import "fmt"

func main() {
	var m, n int
	var result int

	fmt.Scanln(&m, &n)

	result = m*n / fpb(m, n)

	fmt.Println(result)
}

// we need to look for the fpb
func fpb(a, b int) int {
	var temp int
	temp = 0

	for b != 0 {
		temp = a % b
		a = b
		b = temp
	}

	return a
}