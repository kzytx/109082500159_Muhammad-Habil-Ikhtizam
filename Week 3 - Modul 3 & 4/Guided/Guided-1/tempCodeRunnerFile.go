package main
import "fmt"

func hitungLuasPersegiPanjang(panjang int, lebar int) int {
	luas := panjang * lebar
	return luas
}

func main() {
	p:= 10
	l:= 5
	luas := hitungLuasPersegiPanjang(p, l)
	fmt.Printf("Luas persegi panjang: %d\n", luas)
}