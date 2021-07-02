package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	var nCount, pCount int
	var status bool

	status = true
	nCount = 0
	pCount = 0

	for status {
		fmt.Scanln(&a, &b)

		fmt.Println("----")

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(a < b)
		fmt.Println(b < a)
		fmt.Println(a < b || b < a)

		fmt.Println("----")

		if math.Abs(a) == math.Abs(b) {
			nCount++
		} else if a > b || b > a {
			pCount++
		}

		if a < 0 && math.Abs(a) < math.Abs(b) {
			status = false
		}

		if b < 0 && math.Abs(b) < math.Abs(a) {
			status = false
		}
	}

	fmt.Printf("Netral: %d\n", nCount)
	fmt.Printf("Positive: %d\n", pCount)
}