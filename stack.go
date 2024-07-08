package main

import "fmt"

// Struktur Stack
type Stack[T any] struct {
    elements []T
}

// Metode untuk menambah elemen ke stack
func (s *Stack[T]) Push(element T) {
    s.elements = append(s.elements, element)
}

// Metode untuk mengeluarkan elemen dari stack
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }
    element := s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1]
    return element, true
}

func main() {
    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)
    
    fmt.Println(intStack.Pop()) // Output: 3 true
    fmt.Println(intStack.Pop()) // Output: 2 true
    fmt.Println(intStack.Pop()) // Output: 1 true
    fmt.Println(intStack.Pop()) // Output: 0 false

    stringStack := Stack[string]{}
    stringStack.Push("a")
    stringStack.Push("b")
    stringStack.Push("c")
    
    fmt.Println(stringStack.Pop()) // Output: c true
    fmt.Println(stringStack.Pop()) // Output: b true
    fmt.Println(stringStack.Pop()) // Output: a true
    fmt.Println(stringStack.Pop()) // Output:  false
}
