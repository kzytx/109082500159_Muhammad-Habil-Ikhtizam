package main

import "fmt"

const nProv = 10

type NamaProv [nProv + 1]string
type PopProv [nProv + 1]int
type TumbuhProv [nProv + 1]float64

func inputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi, Angka Pertumbuhan Provinsi ===")
	for i := 1; i <= nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idxMaks := 1
	for i := 2; i <= nProv; i++ {
		if tumbuh[i] > tumbuh[idxMaks] {
			idxMaks = i
		}
	}
	return idxMaks
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 1; i <= nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := int((tumbuh[i] + 1) * float64(pop[i]))
			fmt.Println(prov[i], prediksi)
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 1; i <= nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	inputData(&prov, &pop, &tumbuh)

	idx := ProvinsiTercepat(tumbuh)
	fmt.Println("\nProvinsi dengan angka pertumbuhan tercepat :", prov[idx])

	var cari string
	fmt.Print("\nData provinsi yang dicari : ")
	fmt.Scan(&cari)

	idxCari := IndeksProvinsi(prov, cari)
	fmt.Println(idxCari)

	Prediksi(prov, pop, tumbuh)
}