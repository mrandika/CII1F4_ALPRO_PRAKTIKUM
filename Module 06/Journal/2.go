package main

import "fmt"

func main() {
	var x float64
	var y float64

	fmt.Scanln(&x)

	fmt.Printf("f(%.1f) = %.2f\n", x, f(x))
	fmt.Printf("g(%.1f) = %.2f\n", x, g(x))
	fmt.Printf("h(%.1f) = %.2f\n", x, h(x))

	komposisi(x, &y)

	fmt.Printf("(fogoh)(%.1f) = %.2f\n", x, y)
}

func f(x float64) float64 {
	return x * x
}

func g(x float64) float64 {
	return x - 2
}

func h(x float64) float64 {
	return x + 1
}

func komposisi(x float64, y *float64) {
	fog := f(g(h(x)))

	*y = fog
}