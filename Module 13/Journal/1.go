package main

import "fmt"

type mahasiswa struct {
	inisial string
	tinggiBadan int
}

type arrMahasiswa struct {
	data []mahasiswa
	count int
}

func main() {
	var nData int
	var dataMhs arrMahasiswa

	fmt.Scanln(&nData)

	for i := 0; i < nData; i++ {
		bacaData(&dataMhs, &dataMhs.count)
	}

	urutData(&dataMhs, dataMhs.count)

	fmt.Println("\n")
	cetakData(dataMhs, dataMhs.count)
}

func bacaData(arrMhs *arrMahasiswa, count *int) {
	var mhs mahasiswa

	fmt.Scanln(&mhs.inisial, &mhs.tinggiBadan)

	arrMhs.data = append(arrMhs.data, mhs)
	arrMhs.count++
}

func cetakData(arrMhs arrMahasiswa, count int) {
	for _, v := range arrMhs.data {
		fmt.Println(v.inisial, v.tinggiBadan)
	}
}

func urutData(arrMhs *arrMahasiswa, count int) {
	var n = len(arrMhs.data)

	for i := 0; i < n; i++ {
		var lowestIndex = i
		
		for j := i; j < n; j++ {
			if arrMhs.data[j].tinggiBadan < arrMhs.data[lowestIndex].tinggiBadan {
				lowestIndex = j
			}
		}

		arrMhs.data[i], arrMhs.data[lowestIndex] = arrMhs.data[lowestIndex], arrMhs.data[i]
	}
}