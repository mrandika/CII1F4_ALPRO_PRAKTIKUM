package main

import "fmt"

type arrData struct {
	data []int
	count int
}

func main() {
	var n, d, u int
	var arrNum arrData

	fmt.Scanln(&n, &u, &d)

	for i := 0; i < n; i++ {
		isiArray(&arrNum.data, &arrNum.count)
	}

	sorting(&arrNum.data, u, d)

	fmt.Println("\n")
	tampilArray(arrNum.data, arrNum.count)
}

func isiArray(arr *[]int, count *int) {
	var num int

	fmt.Scan(&num)

	*arr = append(*arr, num)
	*count++
}

func tampilArray(arr []int, count int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}

	fmt.Println("\n")
}

func sorting(arr *[]int, u, d int) {
	var n = len(*arr)

	for i := 1; i < n; i++ {
		j := i

		for j > 0 {
			if (*arr)[j-1] < (*arr)[j] && j-1 >= u && j < d {
				(*arr)[j-1] , (*arr)[j]  = (*arr)[j] , (*arr)[j-1] 
			}

			j = j - 1
		}
	}
}