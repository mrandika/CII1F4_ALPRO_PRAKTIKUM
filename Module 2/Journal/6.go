package main

import "fmt"

func main() {
	var nilai, jumlah, n int
	var rata_rata float64

	jumlah = 0
	n = 0

	fmt.Scanln(&nilai)

	for nilai != -1 {
		n = n + 1
		jumlah = jumlah + nilai

		fmt.Scanln(&nilai)
	}

	if n == 0 {
		rata_rata = 0.0
	} else {
		rata_rata = float64(jumlah) / float64(n)
	}

	fmt.Println(rata_rata)
}