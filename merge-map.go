package main

import "fmt"

func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
    result := make(map[K]V)
    for _, m := range maps {
        for k, v := range m {
            result[k] = v
        }
    }
    return result
}

func main() {
    currentMap := map[string]int{"a": 1, "b": 2}
    newMap := map[string]int{"b": 3, "c": 4}

    mergedMap := MergeMaps(currentMap, newMap)

    fmt.Println(mergedMap) // Output: map[a:1 b:3 c:4]
}
