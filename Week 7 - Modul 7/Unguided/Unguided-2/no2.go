package main

import "fmt"

// Tipe data alias: angka merepresentasikan int
type angka int

// Tipe data alias: kata merepresentasikan string
type kata string

// Struct Buku dengan 5 field menggunakan tipe alias
type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	// Deklarasi variabel bertipe Buku
	var buku Buku

	// Input data buku dari user
	fmt.Println("=== INPUT BIODATA BUKU ===")

	fmt.Print("Masukkan judul buku : ")
	fmt.Scan(&buku.judul)

	fmt.Print("Masukkan nama penulis : ")
	fmt.Scan(&buku.penulis)

	fmt.Print("Masukkan penerbit : ")
	fmt.Scan(&buku.penerbit)

	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scan(&buku.tahunTerbit)

	fmt.Scan(&buku.jumlahHalaman)

	// Output biodata buku
	fmt.Println()
	fmt.Println("=== BIODATA BUKU ===")
	fmt.Println("Judul Buku :", buku.judul)
	fmt.Println("Penulis :", buku.penulis)
	fmt.Println("Penerbit :", buku.penerbit)
	fmt.Println("Tahun Terbit :", buku.tahunTerbit)
	fmt.Println("Jumlah Halaman :", buku.jumlahHalaman)
}