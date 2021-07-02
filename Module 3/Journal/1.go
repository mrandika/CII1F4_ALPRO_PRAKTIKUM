package main

import "fmt"

func main() {
	var angka_1, angka_2, angka_3, angka_4, angka_5 int
	var huruf string

	fmt.Scanln(&angka_1, &angka_2, &angka_3, &angka_4, &angka_5)
	fmt.Scanln(&huruf)

	fmt.Printf("%c", angka_1)
	fmt.Printf("%c", angka_2)
	fmt.Printf("%c", angka_3)
	fmt.Printf("%c", angka_4)
	fmt.Printf("%c\n", angka_5)

	fmt.Printf("%c", huruf[0]+1)
	fmt.Printf("%c", huruf[1]+1)
	fmt.Printf("%c\n", huruf[2]+1)
}