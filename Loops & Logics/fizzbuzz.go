package main

import (
	"fmt"
)
func main() {
	for i := 1; i <= 100; i++{
		if i % 3 == 0{
			fmt.Println("Fizz")
			continue
		} else if i % 5 == 0{
			fmt.Println("Buzz")
			continue
		} else if i % 15 == 0{
			fmt.Println("FizzBuzz")
			continue
		} 
		fmt.Println(i)
	}
}