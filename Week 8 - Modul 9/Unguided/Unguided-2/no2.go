package main

import (
	"fmt"
	"math"
)

const NMAX = 100

type tabel [NMAX]int

func isiArray(t *tabel, n *int) {
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&t[i])
	}
}

// a. Tampilkan semua elemen
func cetakSemua(t tabel, n int) {
	fmt.Print("Semua elemen: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", t[i])
	}
	fmt.Println()
}

// b. Tampilkan elemen indeks ganjil
func cetakGanjil(t tabel, n int) {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", t[i])
	}
	fmt.Println()
}

// c. Tampilkan elemen indeks genap
func cetakGenap(t tabel, n int) {
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", t[i])
	}
	fmt.Println()
}

// d. Tampilkan elemen indeks kelipatan x
func cetakKelipatanX(t tabel, n, x int) {
	fmt.Printf("Indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", t[i])
		}
	}
	fmt.Println()
}

// e. Hapus elemen pada indeks tertentu
func hapusElemen(t *tabel, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		t[i] = t[i+1]
	}
	*n--
}

// f. Rata-rata
func rataRata(t tabel, n int) float64 {
	sum := 0
	for i := 0; i < n; i++ {
		sum += t[i]
	}
	return float64(sum) / float64(n)
}

// g. Standar deviasi
func standarDeviasi(t tabel, n int) float64 {
	mean := rataRata(t, n)
	var sumKuadrat float64
	for i := 0; i < n; i++ {
		diff := float64(t[i]) - mean
		sumKuadrat += diff * diff
	}
	return math.Sqrt(sumKuadrat / float64(n))
}

// h. Frekuensi suatu bilangan
func frekuensi(t tabel, n, cari int) int {
	freq := 0
	for i := 0; i < n; i++ {
		if t[i] == cari {
			freq++
		}
	}
	return freq
}

func main() {
	var t tabel
	var n int

	isiArray(&t, &n)
	fmt.Println()

	// a
	cetakSemua(t, n)

	// b
	cetakGanjil(t, n)

	// c
	cetakGenap(t, n)

	// d
	var x int
	fmt.Print("Masukkan nilai x untuk kelipatan: ")
	fmt.Scan(&x)
	cetakKelipatanX(t, n, x)

	// e
	var idx int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)
	hapusElemen(&t, &n, idx)
	fmt.Print("Array setelah penghapusan: ")
	cetakSemua(t, n)

	// f
	fmt.Printf("Rata-rata: %.2f\n", rataRata(t, n))

	// g
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi(t, n))

	// h
	var cari int
	fmt.Print("Masukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)
	fmt.Printf("Frekuensi %d: %d\n", cari, frekuensi(t, n, cari))
}