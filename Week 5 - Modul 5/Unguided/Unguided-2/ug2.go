package main

import "fmt"

func faktor(n, current int) {
    if current > n {
        return
    }
    if n%current == 0 {
        fmt.Printf("%d ", current)
    }
    faktor(n, current+1)
}

func main() {
    var n int
    fmt.Print("Masukkan N: ")
    fmt.Scan(&n)
    faktor(n, 1)
    fmt.Println()
}