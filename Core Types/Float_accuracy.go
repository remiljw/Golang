// To see the accuracy of float type
package main

import "fmt"

func main() {
	var x int = 100
	var y float32 = 100
	var z float64 = 100
	fmt.Println((x/3)*3)
	fmt.Println((y/3)*3)
	fmt.Println((z/3)*3)
// Number Wraparound
	var a int8 = 125
	var b uint8 = 253
	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8", a, "uint8", b)
	}
}