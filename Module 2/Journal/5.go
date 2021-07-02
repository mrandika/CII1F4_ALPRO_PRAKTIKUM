package main

import "fmt"

func main() {
	var kata string

	fmt.Scanln(&kata)

	for kata != "selesai" {
		fmt.Println(kata)

		fmt.Scanln(&kata)
	}
}