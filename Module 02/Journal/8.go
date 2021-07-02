package main

import "fmt"

func main() {
	var r, t int

	fmt.Scanln(&r, &t)
	fmt.Println(hitungVolume(r, t))

}

func hitungVolume(r int, t int) float64 {
	var pi float64 = 3.14
    return float64(r) * float64(r) * pi * float64(t)
}