package main

import "fmt"

var data []int

func main() {
	var n, k int
	var val int

	i := 0

	fmt.Scanln(&n, &k)

	fmt.Scan(&val)
	isiArray(val)
	i++

	for i < n {
		fmt.Scan(&val)
		isiArray(val)

		i++
	}

	posisi := posisi(len(data), k)

	if posisi < 0 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(posisi)	
	}
}

func isiArray(n int) {
	/* IS. Data n sudah siap pada piranti masukan.
	FS. Array data berisi n (<=NMAX) bilangan */

	data = append(data, n)
}

func posisi(n, k int) int {
	/* mengembalikan posisi k dalam array data dengan n elemen. Posisi dimulai dari
	posisi 0. Jika tidak ada kembalikan -1 */

	found := false
	pos := -1

	for i := 0; i < n && !found; {
		if data[i] == k {
			pos = i
			found = true
		} else {
			i++
		}
	}

	return pos
}