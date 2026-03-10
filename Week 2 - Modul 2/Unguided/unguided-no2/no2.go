package main

import "fmt"

func main() {
	var warna [5][4]string
	berhasil := true

	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %d : ", i+1)
		fmt.Scan(&warna[i][0], &warna[i][1], &warna[i][2], &warna[i][3])

		if warna[i][0] != "merah" || warna[i][1] != "kuning" || warna[i][2] != "hijau" || warna[i][3] != "ungu" {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL :", berhasil)
}