package main

import "fmt"

func main() {
	var total, usaha, nDoa int
	var doaOrtu bool
	var nilai string

	BacaData(&usaha, &nDoa, &doaOrtu, &nilai)

	TabungUsahaDoa(usaha, nDoa, &total)

	total = TabungDoaOrtu(doaOrtu, total)

	HasilNilaiAlpro(nilai, &total)

	fmt.Println(HasilAkhir(total))
}

func BacaData(usaha, jumlahDoa *int, doaOrtu *bool, nilai *string) {
	var usahaLocal, doaLocal int
	var doaOrtuLocal bool
	var nilaiLocal string

	fmt.Print("Banyaknya Usaha? ")
	fmt.Scanln(&usahaLocal)
	*usaha = usahaLocal

	fmt.Print("Banyaknya Doa? ")
	fmt.Scanln(&doaLocal)
	*jumlahDoa = doaLocal

	fmt.Print("Doa orang tua? ")
	fmt.Scanln(&doaOrtuLocal)
	*doaOrtu = doaOrtuLocal

	fmt.Print("Nilai Algoritma? ")
	fmt.Scanln(&nilaiLocal)
	*nilai = nilaiLocal
}

func TabungUsahaDoa(usaha int, doa int, total *int) {
	*total = usaha + doa
}

func TabungDoaOrtu(doa bool, total int) int {
	if doa {
		return total * 2
	} 

	return total
}

func HasilNilaiAlpro(nilai string, total *int) {
	if nilai == "A" {
		*total -= 150
	} else if nilai == "B" {
		*total -= 130
	} else if nilai == "C" {
		*total -= 100
	}
}

func HasilAkhir(poin int) string {
	if poin >= 130 {
		return "Lulus langsung dapat kerja gaji 2 digit"
	} else if poin >= 50 && poin < 130 {
		return "Lulus langsung dapat kerja"
	} else {
		return "Jangan lelah berdoa dan berusaha, tidak ada yang sia – sia dari usaha dan doa"
	}
}