package main
import "fmt"

func main() {
	var berat, tinggi, bmi float64

	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scan(&berat)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&tinggi)
	
	bmi = berat / (tinggi * tinggi)

	fmt.Printf("Nilai BMI Anda adalah: %.2f\n", bmi)
}