package main

import (
	"fmt"
	"math"
)

var jL1, jL2 float64

func main() {
	var tpxL1, tpyL1 float64
	var tpxL2, tpyL2 float64

	var x, y float64

	fmt.Scanln(&tpxL1, &tpyL1, &jL1)
	fmt.Scanln(&tpxL2, &tpyL2, &jL2)
	fmt.Scanln(&x, &y)

	isOnL1 := lingkaran(x, y, tpxL1, tpyL1, jL1)
	isOnL2 := lingkaran(x, y, tpxL2, tpyL2, jL2)

	cariTeks(isOnL1, isOnL2)
}

func jarak(x1, y1, x2, y2 float64) float64 {
	ac := x1 - x2*x1 - x2
	bd := y1 - y2*y1 - y2

	return math.Sqrt(ac + bd)
}

func lingkaran(x, y, tpx, tpy, r float64) bool {
	jarakKePusat := jarak(tpx, tpy, x, y)
	return jarakKePusat <= r
}

func cariTeks(L1, L2 bool) {
	if L1 && L2 {
		fmt.Println("di dalam lingkaran 1 dan 2")
	} else if L1 {
		fmt.Println("di dalam lingkaran 1")
	} else if L2 {
		fmt.Println("di dalam lingkaran 2")
	} else {
		fmt.Println("di luar lingkaran 1 dan 2")
	}
}
