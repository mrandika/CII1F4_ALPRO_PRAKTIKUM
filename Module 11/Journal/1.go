package main

import "fmt"

const countMax = 1000

type waktu struct {
	menit int
	detik int
}

type lagu struct {
	judul    string
	penyanyi string
	durasi   waktu
}

type playlist struct {
	data  []lagu
	count int
}

func main() {
	var arrPlaylist playlist

	buatPlaylist(&arrPlaylist, arrPlaylist.count)
	cetakPlaylist(arrPlaylist, arrPlaylist.count)
}

func cariLagu(judul, penyanyi string, arrPlaylist playlist, count int) bool {
	for _, lagu := range arrPlaylist.data {
		if (lagu.judul== judul) && (lagu.penyanyi == penyanyi) {
			return true
		} 
	}

	return false
}

func buatPlaylist(dataPlaylist *playlist, count int) {
	var tmpLagu lagu

	fmt.Scanln(&tmpLagu.judul, &tmpLagu.penyanyi, &tmpLagu.durasi.menit, &tmpLagu.durasi.detik)

	for (tmpLagu.judul != "#") && (tmpLagu.penyanyi != "#") {
		if cariLagu(tmpLagu.judul, tmpLagu.penyanyi, *dataPlaylist, dataPlaylist.count) == false {
			dataPlaylist.data = append(dataPlaylist.data, tmpLagu)
			dataPlaylist.count++
		}

		fmt.Scanln(&tmpLagu.judul, &tmpLagu.penyanyi, &tmpLagu.durasi.menit, &tmpLagu.durasi.detik)
	}
}

func terlama(arrPlaylist playlist, count int) int {
	var index, totalDuration int

	index = 0
	totalDuration = 0

	for i := 0; i < arrPlaylist.count; i++ {
		lagu := arrPlaylist.data[i]
		tmpDuration := (lagu.durasi.menit * 60) + lagu.durasi.detik

		if totalDuration < tmpDuration {
			index = i
			totalDuration = tmpDuration
		}
	}

	return index
}

func cetakPlaylist(arrPlaylist playlist, count int) {
	var longestIndex int

	longestIndex = terlama(arrPlaylist, arrPlaylist.count)

	for i, lagu := range arrPlaylist.data {
		if i == longestIndex {
			fmt.Printf("*%s %d menit %d detik\n", lagu.judul, lagu.durasi.menit, lagu.durasi.detik)
		} else {
			fmt.Println(lagu.judul)
		}
	}
}
