package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(T *arrayMahasiswa, n *int) {
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}
}

func nilaiPertama(T arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}
	return -1
}

func nilaiTerbesar(T arrayMahasiswa, n int, nim string) int {
	maks := -1
	for i := 0; i < n; i++ {
		if T[i].NIM == nim {
			if T[i].nilai > maks {
				maks = T[i].nilai
			}
		}
	}
	return maks
}

func main() {
	var data arrayMahasiswa
	var n int

	inputData(&data, &n)

	var nimCari string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	pertama := nilaiPertama(data, n, nimCari)
	terbesar := nilaiTerbesar(data, n, nimCari)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, pertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, terbesar)
}