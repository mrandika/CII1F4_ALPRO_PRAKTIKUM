package main

import "fmt"

func main() {
	var nama string
	var soal, skor int

	var tempNama string
	var tempSoal, tempSkor int

	fmt.Scan(&nama)

	for nama != "Selesai" {
		hitungSkor(&soal, &skor)

		if soal > tempSoal {
			tempNama = nama
			tempSoal = soal
			tempSkor = skor
		} else if soal == tempSoal && skor < tempSkor {
			tempNama = nama
			tempSoal = soal
			tempSkor = skor
		}

		fmt.Scan(&nama)
	}

	fmt.Println(tempNama, tempSoal, tempSkor)
}

func hitungSkor(soal, skor *int) {
	var tempSoal int

	*skor = 0
	*soal = 0

	for i := 0; i <= 7; i++ {
		fmt.Scan(&tempSoal)
		
		if tempSoal < 301 {
			*skor += tempSoal
			*soal++
		}
	}
}