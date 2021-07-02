package main

import "fmt"

func main() {
	var n, r int
	var p, c int

	fmt.Scanln(&n, &r)

	p = permutation(n, r)
	c = combination(n, r)

	fmt.Printf("%d %d\n", p, c)
}

func findFactorial(n int) int {
	factorial := 1

	if n > 0 {
		for i := 1; i <= n; i++ {
			factorial *= i
		}
	}

	return factorial
}

func permutation(n int, r int) int {
	return findFactorial(n)/findFactorial(n-r)
}

func combination(n int, r int) int {
	return findFactorial(n)/(findFactorial(r)*findFactorial(n-r))
}