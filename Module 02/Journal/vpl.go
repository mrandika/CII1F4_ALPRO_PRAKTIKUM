package main

import "fmt"

func main() {
	var start, end int

	fmt.Scanln(&start, &end)

	for i := start; i <= end; i++ {
		fmt.Printf("Simbol ASCII dari %d adalah %c\n", i, i)
	}
}