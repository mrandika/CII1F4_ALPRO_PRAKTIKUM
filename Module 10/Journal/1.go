package main

import (
	"fmt"
	"sort"
)

func main() {
	var suara, totalSuara, suaraSah int
	var sliceSuara []int

	fmt.Scan(&suara)
	if suara > 0 && suara <= 20 {
		suaraSah++
		sliceSuara = append(sliceSuara, suara)
	}


	for suara != 0 {
		totalSuara++
		fmt.Scan(&suara)

		if suara > 0 && suara <= 20 {
			suaraSah++
			sliceSuara = append(sliceSuara, suara)
		}
	}

	sort.Ints(sliceSuara)
	
	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	hasil := hitung_suara(sliceSuara)

	for k, v := range hasil {
		fmt.Printf("%d: %d\n", k, v)
	}

}

func hitung_suara(sortedSlices []int) map[int]int {
	counter := make(map[int]int)

 	for _, i := range sortedSlices {
 		_, exist := counter[i]

 		if exist {
			counter[i] += 1
 		} else {
			counter[i] = 1
 		}
 	}
	
 	return counter
}