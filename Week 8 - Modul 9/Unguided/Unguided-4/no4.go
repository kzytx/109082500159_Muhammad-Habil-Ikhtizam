package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	// Input selama karakter bukan TITIK dan n <= NMAX
	var ch rune
	*n = 0
	for *n < NMAX {
		fmt.Scanf("%c", &ch)
		if ch == '.' {
			break
		}
		// Skip newline/spasi
		if ch == '\n' || ch == ' ' {
			continue
		}
		t[*n] = ch
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	// Salin array ke variabel sementara
	var temp tabel
	for i := 0; i < n; i++ {
		temp[i] = t[i]
	}
	// Balik array temp
	balikanArray(&temp, n)
	// Bandingkan dengan array asli
	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Hves : ")
	isiArray(&tab, &m)

	// Tampilkan array asli
	fmt.Print("Array asli     : ")
	cetakArray(tab, m)

	// Balik array
	balikanArray(&tab, m)
	fmt.Print("Rvvvusv tves   : ")
	cetakArray(tab, m)

	// Balik kembali ke posisi semula untuk cek palindrom
	balikanArray(&tab, m)

	// Cek palindrom
	fmt.Print("flalitauou o   : ")
	fmt.Println(palindrom(tab, m))
}