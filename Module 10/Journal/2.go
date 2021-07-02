package main

import "fmt"

func main() {
	var i int
	var daftar, wisuda []string

	var iteration int
	var nim string

	// Daftar wisuda
	fmt.Scan(&i)
	iteration = 0

	fmt.Scan(&nim)
	iteration++
	daftar = append(daftar, nim)

	for iteration < i {
		iteration++
		fmt.Scan(&nim)
		daftar = append(daftar, nim)
	}

	fmt.Println(daftar)

	// Wisuda
	fmt.Scan(&i)
	iteration = 0

	fmt.Scan(&nim)
	iteration++
	wisuda = append(wisuda, nim)

	for iteration < i {
		iteration++

		fmt.Scan(&nim)
		wisuda = append(wisuda, nim)
	}

	fmt.Println(wisuda)

	updateKelulusan(&daftar, wisuda, 0, 0)
}

func posisi(tab []string, n int, x string) int {
	pos := 0

	for _, v := range tab {
		pos++
		if x == v {
			fmt.Println("found at", pos)
			tab[pos] = "0"
			break
		}
	}

	return pos
}

func hapus(tab *[]string, n int, x string) {
	posisi(*tab, 0, x)
}

func updateKelulusan(mhs *[]string, wisuda []string, n int, m int) {
	var newMhs []string

	for _, v := range wisuda {
		fmt.Printf("FOUND %s\n", v)

		hapus(mhs, 0, v)
	}

	counter := 0

	for _, v := range *mhs {
		if v != "0" {
			counter++
			newMhs = append(newMhs, v)
		}
	}

	*mhs = newMhs

	fmt.Printf("%d ", counter)

	for _, v := range *mhs {
		fmt.Printf("%s ", v)
	}
}