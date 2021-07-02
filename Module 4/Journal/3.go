package main

import "fmt"

func main() {
	var key int
	var a, b int
	var isDone bool

	var privateInfo []int

	fmt.Scan(&key)

	for isDone == false {
		fmt.Scanln(&a, &b)

		if a == -1 && b == -1 {
			isDone = true
		}

		if a % key == 0 || b % key == 0 {
			if a % key == 0 {
				privateInfo = append(privateInfo, b)
			} else if b % key == 0 {
				privateInfo = append(privateInfo, a)
			}
		}
	}

	for _, value := range privateInfo {
		fmt.Printf("%d\n", value)
	}
}