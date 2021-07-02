package main

import "fmt"

func main() {
	var day1, day2, year1, year2 int
	var day, month1, month2 string
	var m2 int

	fmt.Scanln(&day1, &month1, &year1, &day)

	Pengambilan(day1, Bulan2Angka(month1), year1, day, &day2, &m2, &year2)

	month2 = Angka2Bulan(m2)

	fmt.Println(day2, month2, year2)
}

func Kabisat(tahun int) bool {
	return (tahun % 400 == 0 || tahun % 4 == 0) && tahun % 100 != 0
}

func Bulan2Angka(bulan string) int {
	if bulan == "januari" {
		return 1
	} else if bulan == "februari" {
		return 2
	} else if bulan == "maret" {
		return 3
	} else if bulan == "april" {
		return 4
	} else if bulan == "mei" {
		return 5
	} else if bulan == "juni" {
		return 6
	} else if bulan == "juli" {
		return 7
	} else if bulan == "agustus" {
		return 8
	} else if bulan == "september" {
		return 9
	} else if bulan == "oktober" {
		return 10
	} else if bulan == "november" {
		return 11
	} else {
		return 12
	}
}

func Angka2Bulan(angka int) string {
	if angka == 1 {
		return "januari"
	} else if angka == 2 {
		return "februari"
	} else if angka == 3 {
		return "maret"
	} else if angka == 4 {
		return "april"
	} else if angka == 5 {
		return "mei"
	} else if angka == 6 {
		return "juni"
	} else if angka == 7 {
		return "juli"
	} else if angka == 8 {
		return "agustus"
	} else if angka == 9 {
		return "september"
	} else if angka == 10 {
		return "oktober"
	} else if angka == 11 {
		return "november"
	} else {
		return "desember"
	}
}

func JumlahHari(bln, thn int) int {
	if bln == 1 || bln == 3 || bln == 5 || bln == 7 || bln == 8 || bln == 10 || bln == 12 {
		return 31
	} else if bln == 4 || bln == 6 || bln == 9 || bln == 11 {
		return 30
	} else {
		if Kabisat(thn) {
			return 29
		} else {
			return 28
		}
	}
}

func Pengambilan(tgl1, bln1, thn1 int, hari string, tgl2, bln2, thn2 *int) {
	var waktu_proses int

	if hari == "kamis" || hari == "jumat" {
		waktu_proses = 4
	} else {
		waktu_proses = 2
	}

	*tgl2 = tgl1 + waktu_proses
	totalHari := JumlahHari(bln1, thn1)

	if *tgl2 <= totalHari {
		*bln2 = bln1
		*thn2 = thn1
	} else {
		*tgl2 = *tgl2 - totalHari

		if bln1 == 12 {
			*bln2 = 1
			*thn2 = thn1 + 1
		} else {
			*bln2 = bln1 + 1
			*thn2 = thn1
		}
	}
}