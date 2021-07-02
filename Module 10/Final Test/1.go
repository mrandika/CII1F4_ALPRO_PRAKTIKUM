package main

import (
	"fmt"
	"sort"
)

func main() {
	var suara, totalSuara, suaraSah int
	var sliceSuara []int
	var finalis []int

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
		if v > 1 {
			finalis = append(finalis, k)
		}
	}

	sort.Ints(finalis)

	if len(finalis) < 2 {
		finalis = append(finalis, sliceSuara[0])
	}

	fmt.Printf("Ketua RT: %d\n", finalis[0])
	fmt.Printf("Wakil Ketua: %d\n", finalis[1])

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