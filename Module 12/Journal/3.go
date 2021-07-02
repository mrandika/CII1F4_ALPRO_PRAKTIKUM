package main

import "fmt"

const nMAX = 1000000

type arrDate struct {
	data []string
}

func main() {
	var t arrDate
	var n int

	addData(&t, &n)
	urut(&t, &n)
	printData(t, n)
}

func addData(t *arrDate, n *int) {
	tempData := ""

	for tempData != "####.##.##" {
		fmt.Scanln(&tempData)

		if tempData == "####.##.##" {
			break
		}

		t.data = append(t.data, tempData)
	}

	fmt.Println("\n")
}

func urut(t *arrDate, n *int) {
	var len = len(t.data)

	for i := 0; i < len; i++ {
		var minIndex = i

		for j := i; j < len; j++ {
			if t.data[j] > t.data[minIndex] {
				minIndex = j
			}
		}

		t.data[i], t.data[minIndex] = t.data[minIndex], t.data[i]
	}
}

func maxPos(t arrDate, start, n int) int {
	value := "0000-00-00"
	index := 0

	for i, v := range t.data {
		if value < v {
			value = v
			index = i
		}
	}

	return index
}

func swap(x, y *string) {
	x, y = y, x
}

func printData(t arrDate, n int) {
	for _, v := range t.data {
		fmt.Println(v)
	}
}