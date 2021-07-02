package main

import "fmt"

func main() {
	var a, b, c, d int
	var min, max int

	fmt.Scanln(&a, &b, &c, &d)

	// Progress 1
	minmax(a, b, &min, &max)

	// Progress 2
	tempMin := min
	tempMax := max

	minmax(c, d, &min, &max)

	if tempMin < min {
		min = tempMin
	} 

	if tempMax > max {
		max = tempMax
	}

	// Output
	fmt.Printf("%d %d\n", max, min)
}

func minmax(f1, f2 int, min, max *int) {
	if f1 > f2 {
		*min = f2
		*max = f1
	} else {
		*min = f1
		*max = f2
	}
}