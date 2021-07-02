package main

import "fmt"

const nMAX = 1000000

type partai struct {
	nama int
	suara int
}

type perolehan struct {
	data []partai
	count int
}

func main() {
	var vote int
	var dataPartai partai
	var dataPerolehan perolehan

	for vote != -1 {
		fmt.Scan(&vote)

		if vote == -1 {
			break
		}

		posisi := posisi(dataPerolehan, dataPerolehan.count, vote)

		if posisi == -1 {
			dataPartai.nama = vote
			dataPartai.suara = 1

			dataPerolehan.data = append(dataPerolehan.data, dataPartai)
			dataPerolehan.count++
		} else {
			dataPerolehan.data[posisi].suara++
		}
	}

	sortByVote(&dataPerolehan.data)
	
	for _, v := range dataPerolehan.data {
		fmt.Printf("%d(%d) ", v.nama, v.suara)
	}
	fmt.Print("\n")	
}

func sortByVote(data *[]partai) {
	var len = len(*data)

	for i := 1; i < len; i++ {
		j := i

		for j > 0 {
			if (*data)[j-1].suara < (*data)[j].suara {
				(*data)[j-1], (*data)[j] = (*data)[j], (*data)[j-1]
			}
			
			j--
		}
	}
}

func posisi(dataPerolehan perolehan, count int, nama int) int {
	i := 0
	found := -1

	for i < count && found == -1 {
		if dataPerolehan.data[i].nama == nama {
			found = i
		} else {
			i++
		}
	}

	return found
}