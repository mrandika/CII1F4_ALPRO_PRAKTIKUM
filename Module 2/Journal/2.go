package main

import "fmt"

func main() {
	var r, luas_lingkaran float64

    fmt.Scanln(&r)
    luas_lingkaran = r * 22/7

    fmt.Println("Luas lingkaran dengan jari-jari =", r, "adalah", luas_lingkaran)
}