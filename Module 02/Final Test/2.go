package main

import "fmt"

func main() {
	var n int
	var a, b int

	var privateInfo []int

	fmt.Scan(&n)

	for i := 1; i < n+1; i++ {
		fmt.Scanln(&a, &b)

		if (a+b) % 2 == 0 {
			privateInfo = append(privateInfo, a)
		} else {
			privateInfo = append(privateInfo, b)
		}
	}

	for _, value := range privateInfo {
		fmt.Printf("%d\n", value)
	}
}