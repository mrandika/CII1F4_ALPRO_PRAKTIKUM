package main

import "fmt"
import "strings"

func main() {
	var nomor string
	var tebak_1, tebak_2, tebakan string

	fmt.Scanln(&nomor)
	fmt.Scanln(&tebak_1)
	fmt.Scanln(&tebak_2)
	fmt.Scanln(&tebakan)

	isCorrect := nomor == tebakan

	isFirstContain := strings.Contains(nomor, tebak_1)
	isSecondContain := strings.Contains(nomor, tebak_2)

	fmt.Println(isCorrect)
	fmt.Println(isFirstContain)
	fmt.Println(isSecondContain)
}