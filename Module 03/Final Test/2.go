package main

import "fmt"

func main() {
	var i, n int
	var tendangan int
	var status bool

	i = 0
	n = 0
	status = true

	for status {

		fmt.Printf("Tendangan ke-%d: ", i+1)
		fmt.Scanln(&tendangan)

		if (tendangan == 0 || tendangan < 0) {
			status = false
		} else if (tendangan % 4 == 0) {
			n++
		}

		i++
	}

	if n > 0 && i/2 != n {
		fmt.Println("Penyerang menang")
	} else if n == 0 {
		fmt.Println("Kiper menang")
	} else if n > 0 && i/2 == n {
		fmt.Println("Draw")
	}

	fmt.Printf("Skor gol: %d dari %d tendangan\n", n, i-1)
}