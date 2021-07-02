package main

import (
	"fmt"
	"strconv"
)

func main() {
	var dataCount int
	var angka int
	var hasil string

	var isValidInput bool

	fmt.Scan(&dataCount)

	for i := 1; i <= dataCount; i++ {
		fmt.Scan(&angka)

		isValidInput = angka >= 0 && angka <= 9

		for isValidInput == false {
			fmt.Scan(&angka)
			isValidInput = angka >= 0 && angka <= 9
		}

		hasil = hasil + strconv.Itoa(angka)
	}

	fmt.Println(reverse(hasil))
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
	  rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}