package main

import "fmt"

func main() {
	var BDG, JKT, L int
	var value string

	BDG, JKT, L = 0, 0, 0

	fmt.Scan(&value)

	for value != "." {
		hitungNopol(value, &BDG, &JKT, &L)

		fmt.Scan(&value)
	}

	fmt.Printf("%d %d %d\n", BDG, JKT, L)
}

func hitungNopol(noPol string, bdg, jkt, l *int) {
	if noPol == "D" || noPol == "Z" {
		*bdg++
	} else if noPol == "B" || noPol == "F" {
		*jkt++
	} else {
		*l++
	}
}