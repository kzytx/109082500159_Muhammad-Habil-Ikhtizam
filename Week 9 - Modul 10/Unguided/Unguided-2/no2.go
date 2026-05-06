package main

import "fmt"

const NMAX = 1000

func main() {
	var x, y int
	var ikan [NMAX]float64

	// input
	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var totalWadah []float64
	var sum float64
	count := 0

	// proses pembagian ke wadah
	for i := 0; i < x; i++ {
		sum += ikan[i]
		count++

		if count == y || i == x-1 {
			totalWadah = append(totalWadah, sum)
			sum = 0
			count = 0
		}
	}

	// output baris 1 (total tiap wadah)
	for i := 0; i < len(totalWadah); i++ {
		fmt.Printf("%.2f ", totalWadah[i])
	}
	fmt.Println()

	// hitung rata-rata wadah
	var total float64
	for i := 0; i < len(totalWadah); i++ {
		total += totalWadah[i]
	}
	rata := total / float64(len(totalWadah))

	// output baris 2
	fmt.Printf("%.2f\n", rata)
}