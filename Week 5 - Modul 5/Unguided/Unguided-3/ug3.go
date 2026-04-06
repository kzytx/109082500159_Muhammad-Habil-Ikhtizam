package main

import "fmt"

func ganjil(n, current int) {
    if current > n {
        return
    }
    if current%2 != 0 {
        fmt.Printf("%d ", current)
    }
    ganjil(n, current+1)
}

func main() {
    var n int
    fmt.Print("Masukkan N: ")
    fmt.Scan(&n)
    ganjil(n, 1)
    fmt.Println()
}