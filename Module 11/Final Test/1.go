package main

import "fmt"

func main() {
	var totalData, cariData, dataInput, count, i int
	var data []int
	var found bool

	count = 0

	fmt.Scanln(&totalData, &cariData)

	for count < totalData {
		fmt.Scan(&dataInput)

		data = append(data, dataInput)
		count++
	}

	// Even or odd?
	isEven := paritycheck(cariData)

	// Look for min max
	min := searchlowest(data)
	max := searchhighest(data)

	if isEven {
		i = max

		for i < count && found == false {
			if data[i] == cariData {
				found = true
			} else {
				i++
			}
		}

	} else {
		i = 0

		for i <= min && found == false {
			if data[i] == cariData {
				found = true
			} else {
				i++
			}
		}

		i = max

		for i < count && found == false {
			if data[i] == cariData {
				found = true
			} else {
				i++
			}
		}
	}

	if found {
		fmt.Println("ditemukan")
	} else {
		fmt.Println("tidak ditemukan")
	}
}

func searchlowest(data []int) int {
	value := data[0]
	index := 0

	for i, v := range data {
		if v < value {
			value = v
			index = i
		}
	}

	return index
}

func searchhighest(data []int) int {
	value := 0
	index := 0

	for i, v := range data {
		if value < v {
			value = v
			index = i
		}
	}

	return index
}

func paritycheck(a int) bool {
	return a%2 == 0
}
