package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, radius float64
	var cx, cy float64

	fmt.Scanln(&x, &y, &radius)
	fmt.Scanln(&cx, &cy)

	if posisi(cx, cy, radius, x, y) {
		fmt.Println("Anda berada di dalam taman")
	} else {
		fmt.Println("Anda berada di luar taman")
	}
}

func jarak(a, b, c, d float64) float64 {
	ac := a-c * a-c
	bd := b-d * b-d

	return math.Sqrt(ac + bd)
}

func posisi(cx, cy, r, x, y float64) bool {

	jarakKePusat := jarak(cx, cy, x, y)
	jarakKePagar := r - jarakKePusat

	if jarakKePusat < jarakKePagar {
		return true
	}

	return false
}