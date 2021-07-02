package main

import "fmt"

func main() {
	var s string
    var a, b int
    var jumlah int

    fmt.Scanln(&s, &a, &b)
    jumlah = a + b

    fmt.Println("Kata =", s)
    fmt.Println("Jumlah =", jumlah)
}