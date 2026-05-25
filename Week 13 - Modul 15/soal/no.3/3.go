package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func InsertionSort(t *tabPartai, n int) {
	for i := 1; i < n; i++ {
		key := t[i]
		j := i - 1
		for j >= 0 && t[j].suara < key.suara {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = key
	}
}

func main() {
	var p tabPartai
	var n int = 0
	var x int

	fmt.Println("Masukkan proses input suara :")
	fmt.Scan(&x)
	for x != -1 {
		idx := posisi(p, n, x)
		if idx == -1 {
			p[n].nama = x
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}
		fmt.Scan(&x)
	}

	InsertionSort(&p, n)

	fmt.Println("Hasil Perhitungan suara :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d)", p[i].nama, p[i].suara)
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}