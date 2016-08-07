package main

import (
    "fmt"
    "sort"
)

type ByLength []string

func (s ByLength) Len() int{
    return len(s)
}

func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool{
    return len(s[i]) < len(s[j])
}

func main() {
    langs := []string{"java", "php", "go", "javascript"}
    sort.Sort(ByLength(langs))

    fmt.Println("Sorted by length:", langs)
}