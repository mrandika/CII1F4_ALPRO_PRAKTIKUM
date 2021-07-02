package main

import "fmt"

type permainan struct {
	namaDepan string
	namaBelakang string

	gol int
	assist int
}

func main() {
	var n int
	var hasilPermainan permainan
	var dataPermainan []permainan

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&hasilPermainan.namaDepan, &hasilPermainan.namaBelakang, &hasilPermainan.gol, &hasilPermainan.assist)
		dataPermainan = append(dataPermainan, hasilPermainan)
	}

	sortByGoalAndAssist(&dataPermainan)

	for _, v := range dataPermainan {
		fmt.Println(v.namaDepan, v.namaBelakang, v.gol, v.assist)
	}
}

func sortByGoalAndAssist(data *[]permainan) {
	var len = len(*data)

	for i := 1; i < len; i++ {
		j := i

		for j > 0 {
			if (*data)[j-1].gol < (*data)[j].gol {
				(*data)[j-1], (*data)[j] = (*data)[j], (*data)[j-1]
			}

			if (*data)[j-1].gol == (*data)[j].gol && (*data)[j-1].assist < (*data)[j].assist {
				(*data)[j-1], (*data)[j] = (*data)[j], (*data)[j-1]
			}
			
			j--
		}
	}
} 