package main

import "fmt"

// Fungsi untuk mencari elemen dalam slice
func Contains[T comparable](slice []T, element T) bool {
    for _, e := range slice {
        if e == element {
            return true
        }
    }
    return false
}

func main() {
    intSlice := []int{1, 2, 3, 4, 5}
    stringSlice := []string{"apple", "banana", "cherry"}

    fmt.Println(Contains(intSlice, 3))       // Output: true
    fmt.Println(Contains(intSlice, 6))       // Output: false
    fmt.Println(Contains(stringSlice, "banana")) // Output: true
    fmt.Println(Contains(stringSlice, "grape"))  // Output: false
}
