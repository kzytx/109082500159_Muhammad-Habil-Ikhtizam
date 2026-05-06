package main

import (
	"fmt"
)

// tipe array
type arrBalita [100]float64

// subprogram untuk cari min & max
func hitungMinMax(arrBerat arrBalita, n int, bMin *float64, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// fungsi untuk hitung rata-rata
func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max float64

	// input jumlah data
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// input data berat
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	// proses
	hitungMinMax(data, n, &min, &max)
	rata := rerata(data, n)

	// output
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}