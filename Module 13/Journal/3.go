package main

import "fmt"

type arrBelanja struct {
	harga []int
	count int
}

func main() {
	var dataBelanja arrBelanja
	
	var ha int
	var hp, diskon float64

	entri(&dataBelanja, &dataBelanja.count)
	ha = total(dataBelanja, dataBelanja.count)

	if ha > 100000 {
		urut(&dataBelanja, dataBelanja.count)

		for i := 0; i < dataBelanja.count; i++ {
			harga := dataBelanja.harga[i]
			diskon = float64(i+1) / 100
			
			hp += float64(harga) - (float64(harga) * diskon)
		}

		fmt.Println(ha, hp)
	} else {
		fmt.Println(ha, ha)
	}

}

func entri(dataBelanja *arrBelanja, count *int) {
	var h, m int

	fmt.Scanln(&h, &m)

	for h != 0 && m != 0 {
		dataBelanja.harga = append(dataBelanja.harga, h*m)
		dataBelanja.count++

		fmt.Scanln(&h, &m)
	}
}

func urut(dataBelanja *arrBelanja, count int) {
	var n = len(dataBelanja.harga)

	for i := 1; i < n; i++ {
		j := i

		for j > 0 {
			if dataBelanja.harga[j-1] > dataBelanja.harga[j] {
				dataBelanja.harga[j-1], dataBelanja.harga[j] = dataBelanja.harga[j], dataBelanja.harga[j-1]
			}

			j = j - 1
		}
	}
}

func total(dataBelanja arrBelanja, count int) int {
	var total int

	for _, v := range dataBelanja.harga {
		total += v
	}

	return total
}