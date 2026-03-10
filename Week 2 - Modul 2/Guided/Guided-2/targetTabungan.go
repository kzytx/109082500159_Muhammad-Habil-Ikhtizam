package main
import "fmt"

func main() {
	var target,tabungan, total, hari int

	fmt.Print("Masukkan target tabungan: ")
	fmt.Scan(&target)
	
	total = 0
	hari = 0

	for total < target {
		hari++
		fmt.Printf("Masukkan jumlah tabungan hari ke-%d: ", hari)
		fmt.Scan(&tabungan)
		
		total = total + tabungan
	}

	fmt.Printf("Selamat! Target tercapai dalam %d hari\n", hari)
	fmt.Printf("Total tabungan Anda adalah: %d\n", total)
}