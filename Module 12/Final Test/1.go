package main

import (
	"fmt"
	"math"
)

const nMAX = 100

type provinsi struct {
	nama string
	populasi int
	tumbuh float64
}

type dataProvinsi struct {
	data []provinsi
	count int
}

func main() {
	var data provinsi
	var arrProvinsi dataProvinsi
	var namaProvinsi string

	fmt.Println("Data pertumbuhan provinsi:")
	fmt.Scanln(&data.nama, &data.populasi, &data.tumbuh)
	
	for data.nama != "NONE" && data.populasi != 0 && data.tumbuh != 0.0 {
		if data.nama == "NONE" && data.populasi == 0 && data.tumbuh == 0.0 {
			break
		}

		arrProvinsi.data = append(arrProvinsi.data, data)
		arrProvinsi.count++

		fmt.Scanln(&data.nama, &data.populasi, &data.tumbuh)
	}

	fmt.Print("Nama provinsi? ")
	fmt.Scanln(&namaProvinsi)

	searchProvinsi := cariProvinsi(arrProvinsi, namaProvinsi)

	fmt.Printf("%s %d %.4f\n", searchProvinsi.nama, searchProvinsi.populasi, searchProvinsi.tumbuh)

	fmt.Print("Prediksi populasi tahun depan provinsi? ")
	fmt.Scanln(&namaProvinsi)

	fmt.Printf("Populasi Provinsi %s tahun depan: %d\n", namaProvinsi, prediksi(arrProvinsi, namaProvinsi))

	urutData(&arrProvinsi)
	fmt.Println("Urutan data pertumbuhan provinsi berdasarkan populasi terurut menurun:")

	for _, v := range arrProvinsi.data {
		fmt.Printf("%s %d %.4f\n", v.nama, v.populasi, v.tumbuh)
	}
}

func cariProvinsi(dataProvinsi dataProvinsi, nama string) provinsi {
	i := 0
	found := false

	for i < dataProvinsi.count && found == false {
		if dataProvinsi.data[i].nama == nama {
			found = true
		} else {
			i++
		}
	}

	return dataProvinsi.data[i]
}

func prediksi(data dataProvinsi, nama string) int {
	provinsi := cariProvinsi(data, nama)

	return provinsi.populasi + int(math.Ceil(float64(provinsi.populasi) * provinsi.tumbuh))
}

func urutData(dataProvinsi *dataProvinsi) {
	var len = len(dataProvinsi.data)

	for i := 1; i < len; i++ {
		j := i

		for j > 0 {
			if (dataProvinsi.data)[j-1].populasi < (dataProvinsi.data)[j].populasi {
				(dataProvinsi.data)[j-1], (dataProvinsi.data)[j] = (dataProvinsi.data)[j], (dataProvinsi.data)[j-1]
			}
			
			j--
		}
	}
}