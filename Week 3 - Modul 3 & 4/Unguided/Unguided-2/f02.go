package main
import "fmt"

// fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// fungsi g(x) = x - 2
func g(x int) int {
	return x - 2
}

// fungsi h(x) = x + 1
func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai z: ")
	fmt.Scan(&c)

	// (f o g o h)(a) = f(g(h(a)))
	hasil1 := f(g(h(a)))

	// (g o h o f)(b) = g(h(f(b)))
	hasil2 := g(h(f(b)))

	// (h o f o g)(c) = h(f(g(c)))
	hasil3 := h(f(g(c)))

	fmt.Println("F(G(H(", a, "))) :", hasil1)
	fmt.Println("G(H(F(", b, "))) :", hasil2)
	fmt.Println("H(F(G(", c, "))) :", hasil3)
}