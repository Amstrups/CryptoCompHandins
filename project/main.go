package main

import (
	Scheme "crycomproj/big"
	"fmt"
	"math"
)

func main() {
	var q float64 = math.Pow(2, 27)
	n := 64
	m := n
	ex := 3

	scheme := Scheme.New(q, n, m, ex)
	results := scheme.Run(10)

	for result := range results {
		fmt.Printf("%-5v - %2.4f - Message: %s\n", result.Success, result.Time.Seconds(), result.Message)
	}
}
