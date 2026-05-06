package main

import "fmt"

const NMAX int = 1000
type arrFloat [NMAX]float64

func main() {
	var n int
	var data arrFloat

	// input jumlah data
	fmt.Scan(&n)

	// input data berat kelinci
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	// inisialisasi min dan max
	min := data[0]
	max := data[0]

	// proses pencarian
	for i := 1; i < n; i++ {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}

	// output
	fmt.Printf("%.2f %.2f\n", min, max)
}