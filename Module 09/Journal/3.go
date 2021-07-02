package main

import "fmt"

type Buku struct {
	judul, penulis string
	tahunTerbit int
}

const nMax = 5
type TabBuku = [nMax]Buku

func main() {
	var kardus TabBuku
	var atas int

	var judul string

	kardus = TabBuku{}
	atas = 0

	fmt.Println("Masukan data buku (Judul Penulis Tahun), dipisah dengan spasi.")

	for i := 0; i < nMax; i++ {
		tambahBuku(&kardus, &atas)
	}

	fmt.Print("Cari judul buku: ")
	fmt.Scan(&judul)

	cariBuku(&kardus, &atas, judul)

	fmt.Println("Masukan data buku (Judul Penulis Tahun), dipisah dengan spasi. (Mengisi sisa array)")
	atas--

	for i := atas; i < nMax; i++ {
		tambahBuku(&kardus, &atas)	
	}

	// Cari judul yang tidak ada

	fmt.Print("Cari judul buku: ")
	fmt.Scan(&judul)

	cariBuku(&kardus, &atas, judul)
}

func tambahBuku(kardus *TabBuku, atas *int) {
	var j, p string
	var t int

	fmt.Scanln(&j, &p, &t)

	buku := Buku{
		judul: j,
		penulis: p,
		tahunTerbit: t,
	}

	kardus[*atas] = buku

	if j != "" && p != "" && t != 0 {
		*atas++
	}
}

func ambilBuku(kardus *TabBuku, atas *int, ambil Buku) {
	fmt.Println(kardus[*atas-1].judul)
}

func cariBuku(kardus *TabBuku, atas *int, X string) {
	var ketemu bool
	var ada bool

	// Periksa apakah ada bukunya
	for _, buku := range kardus {
		if buku.judul == X {
			ada = true
		}
	}

	if ada {
		for !ketemu {
			for i := 4; i > 0; i-- {
				if kardus[i].judul == X {
					ketemu = true
				} else {
					ambilBuku(kardus, atas, kardus[i])
					*atas--
				}
			}
		}

		fmt.Println("KETEMU")
	} else {
		fmt.Println("Buku Tidak Ada!")
	}

	for _, buku := range kardus {
		if buku.judul != "" && buku.penulis != "" && buku.tahunTerbit != 0 {
			*atas++
		}
	}
}