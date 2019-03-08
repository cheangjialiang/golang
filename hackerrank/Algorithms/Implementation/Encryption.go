package main

import (
    "fmt"
    "math"
)

func main() {
    var s string
    fmt.Scanln(&s)

    m, n := detectGridSize(len(s))
    g := make([][]byte, m)
    for i := range g {
        g[i] = make([]byte, n)
    }

    populateGrid(g, s)
    printGridColumns(g)
}

func detectGridSize(l int) (int, int) {
    s := math.Sqrt(float64(l))

    f := int(math.Floor(s))
    if f * f >= l {
        return f, f
    }

    c := int(math.Ceil(s))
    if f * c >= l {
        return f, c
    }

    return c, c
}

func populateGrid(g [][]byte, s string) {
    for i := range g {
        for j := range g[i] {
            if k := i * len(g[i]) + j; k < len(s) {
                g[i][j] = s[k]
            }
        }
    }
}

func printGridColumns(g [][]byte) {
    for j := range g[0] {
        for i := range g {
            if (g[i][j] != 0) {
                fmt.Print(string(g[i][j]))
            }
        }
        fmt.Print(" ")
    }
}
