package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var angkaRahasia int64
	var tebakanParitas, tebakanAngka string
	var statusParitas, statusAngka bool

	var random int

	fmt.Scanln(&angkaRahasia)
	fmt.Scanln(&tebakanParitas, &tebakanAngka)

	rand.Seed(angkaRahasia)
	random = rand.Intn(6) + 1

	fmt.Println(random)

	fmt.Println(tebakanParitas)
	fmt.Println(tebakanAngka)

	statusParitas = tebakanParitas == "genap" && random % 2 == 0
	statusAngka = tebakanAngka == "besar" && random >= 4

	if statusParitas && statusAngka {
		fmt.Println("Nilai dadu", random, ", anda menang")
	} else {
		fmt.Println("Nilai dadu", random, ", anda kalah")
	}
}
