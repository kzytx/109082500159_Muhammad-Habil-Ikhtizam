package main

import (
	"fmt"
)

const NMAX = 100

type tabelHasil [NMAX]string

func main() {
	var hasil tabelHasil
	var n int
	var klubA, klubB string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pertandingan := 1
	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
		} else {
			hasil[n] = "Draw"
		}
		n++
		pertandingan++
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}