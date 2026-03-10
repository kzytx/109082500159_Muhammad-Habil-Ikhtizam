package main
import "fmt"

func main(){
    var tiket, harga int

    fmt.Print("Masukkan jenis tiket (1-5): ")
    fmt.Scanln(&tiket)

    switch tiket {
    case 1:
        harga = 10000
    case 2:
        harga = 20000
    case 3:
        harga = 30000
    case 4:
        harga = 40000
    case 5:
        harga = 50000
    default:
        fmt.Println("Jenis tiket tidak valid.")
        return
    }
    fmt.Printf("Harga tiket untuk jenis %d adalah Rp%d\n", tiket, harga)
}