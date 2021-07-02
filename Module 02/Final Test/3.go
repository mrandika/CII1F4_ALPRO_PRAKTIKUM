package main

import "fmt"
import "math/rand"

func main() {
	var angkaRahasia int64
	var tebakan, tebakanBot, random int
	var status string

	fmt.Print("Angka rahasia: ")
	fmt.Scanln(&angkaRahasia)
	rand.Seed(angkaRahasia)

	random = rand.Intn(6) + 1
	tebakanBot = rand.Intn(6) + 1

	fmt.Print("Anda: ")
	fmt.Scanln(&tebakan)

	fmt.Println("Dadang: ", tebakanBot)

	if tebakan == random {
		status = "Pemenang adalah anda"
	} else if tebakanBot == random {
		status = "Pemenang adalah Dadang"
	} else if tebakan == random && tebakanBot == random {
		status = "Seri"
	} else {
		status = "Tidak ada pemenang"
	}

	fmt.Printf("Nilai dadu %d, %s\n", random, status)
}