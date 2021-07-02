package main

import "fmt"

func main() {
	var n, total int

	fmt.Scan(&n)

	total = 0
	
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			fmt.Print(i, " ")
			total++
		}
	}

	if total > 2 {
		fmt.Println("\nBukan Oleole")
	} else {
		fmt.Println("\nOleole")
	}
}