package main

import (
    "fmt"
    "sort"
)

func main() {
    strs := []string{"c", "a", "d", "b"}
    sort.Strings(strs)
    fmt.Println("strings:", strs)

    ints := []int{1, 5, 2, 4, 3 }
    sort.Ints(ints)
    fmt.Println("ints:", ints)

    flts := []float64{1.4, 1.9, 1.5, 1.8, 1.2}
    sort.Float64s(flts)
    fmt.Println("floats:", flts)

    intsAreSorted := sort.IntsAreSorted(ints)
    stringsAreSorted := sort.StringsAreSorted(strs)
    fltsAreSorted := sort.Float64sAreSorted(flts)
    fmt.Println("Ints are sorted:", intsAreSorted)
    fmt.Println("Strings are sorted:", stringsAreSorted)
    fmt.Println("Floats are sorted:", fltsAreSorted)
}

