package main

import "fmt"

type Tabel struct {
	char []string
}

const nMax = 127

func main() {
	var teks string

	// _=palindrom

	// isiArray(&teks, 0)
	// cetakArray(&teks, 0)
	// balikanArray(&teks, &palindrom, 0)

	fmt.Scanln(&teks)

	fmt.Println(palindrom(teks))
}

func isiArray(t *Tabel, n int) {
	var value string

	for value != "." {
		fmt.Scan(&value)

		if value != "." {
			*&t.char = append(t.char, value)
		}
	}
}

func cetakArray(t *Tabel, n int) {
	for _, c := range t.char {
		fmt.Println(c)
	}
}

func balikanArray(t *Tabel, c *Tabel, n int) {
	// :)
}

func palindrom(t string) bool {
	for i := 0; i < len(t)/2; i++ {
		if t[i] != t[len(t)-i-1] {
			return false
		}
	}

	return true
}