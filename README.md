[![Build Status](https://travis-ci.org/summed/goromans.svg?branch=master)](https://travis-ci.org/summed/goromans)

# Go-Romans
A tiny package for converting between arabic and roman numerals

    go get -u "github.com/summed/goromans"

    ```go
    package main

    import (
        "fmt"

        "github.com/summed/goromans"
    )

    func main() {
        var (
            r      = "MMDCCCLVII" // 2857
            a uint = 1426         // MCCCCXXVI
        )

        fmt.Printf("Arabic numerals: '%s'\n", romans.AtoR(a))

        fmt.Printf("IsRomanNumerals: '%t'\n", romans.IsRomanNumerals(r))

        if i, err := romans.RtoA(r); err == nil {
            fmt.Printf("Roman numerals: '%d'\n", i)
        } else {
            fmt.Printf("Error: '%s'\n", err)
        }
    }
    ```

