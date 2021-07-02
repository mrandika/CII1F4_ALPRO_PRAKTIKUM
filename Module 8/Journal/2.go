package main

import "fmt"

func main() {
	var calonA, calonB, calonC int
	var value int

	calonA, calonB, calonC = 0, 0, 0

	fmt.Scan(&value)

	for value != -1 {
		hitungSuara(value, &calonA, &calonB, &calonC)

		fmt.Scan(&value)
	}

	fmt.Printf("%d %d %d\n", calonA, calonB, calonC)
}

func hitungSuara(suara int, calonA, calonB, calonC *int) {
	if suara == 1 {
		*calonA++
	} else if suara == 2 {
		*calonB++
	} else if suara == 3 {
		*calonC++
	}
}
