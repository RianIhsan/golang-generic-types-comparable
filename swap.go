package main

import "fmt"

// Fungsi untuk menukar dua nilai
func Swap[T any](a, b T) (T, T) {
    return b, a
}

func main() {
    a, b := 1, 2
    fmt.Println("Before swap:", a, b) // Output: Before swap: 1 2
    a, b = Swap(a, b)
    fmt.Println("After swap:", a, b)  // Output: After swap: 2 1

    x, y := "hello", "world"
    fmt.Println("Before swap:", x, y) // Output: Before swap: hello world
    x, y = Swap(x, y)
    fmt.Println("After swap:", x, y)  // Output: After swap: world hello
}
