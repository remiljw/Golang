
package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// To see the accuracy of float type
	{
		var x int = 100
		var y float32 = 100
		var z float64 = 100
		fmt.Println((x/3)*3)
		fmt.Println((y/3)*3)
		fmt.Println((z/3)*3)
	}
// Number Wraparound
	{
		var a int8 = 125
		var b uint8 = 253
		for i := 0; i < 5; i++ {
			a++
			b++
			fmt.Println(i, ")", "int8", a, "uint8", b)
		}
	}
// Big Numbers
	{
		intA := math.MaxInt64
		intA = intA + 1
		bigA := big.NewInt(math.MaxInt64)
		bigA.Add(bigA, big.NewInt(1))
		fmt.Println("Max Int64: ", math.MaxInt64)
		fmt.Println("Int :", intA)
		fmt.Println("Big Int: ", bigA.String())
	}
//  Looping over a string
	{
		logLevel := "デバッグ"
		for index, runeVal := range logLevel {
			fmt.Println(index, string(runeVal))
		}
		username := "Sir_King_Über"
		// Length of a string
		fmt.Println("Bytes:", len(username))
		fmt.Println("Runes:", len([]rune(username)))
		// Limit to 10 Characters
		fmt.Println(string(username[:10]))
		fmt.Println(string([]rune(username)[:10]))
	}
	// Rune
	username := "Sir_King_Über"
	runes := []rune(username)
	for i := 0; i < len(runes); i++{
		fmt.Print(string(runes[i]), " ")
	}
}