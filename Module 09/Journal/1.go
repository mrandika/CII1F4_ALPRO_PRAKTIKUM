package main

import "fmt"

func main() {
	var dataTim1, dataTim2, dataTim3 []int

	// Input data menyamping dipisahkan spasi, diakhiri dengan -1.
	inputData(&dataTim1, 0)
	inputData(&dataTim2, 0)
	inputData(&dataTim3, 0)

	fmt.Printf("Rata tim 1: %.2f\nRata tim 2: %.2f\nRata tim 3: %.2f\n", 
	rataan(dataTim1), rataan(dataTim2), rataan(dataTim3))
}

func inputData(t *[]int, n int) {
	for n != -1 {
		fmt.Scan(&n)
		
		if n != -1 {
			*t = append(*t, n)
		}
	}
}

func rataan(t []int) float64 {
	var total float64

	for i := 0; i < len(t); i++ {
		total += float64(t[i])
	}

	return total / float64(len(t))
}