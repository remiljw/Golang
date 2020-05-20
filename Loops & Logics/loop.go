package main

import (
	"fmt"
	"math/rand"
)
func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	names := []string{"Jim","Jane","Joe", "June"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	// Loops with maps
	config := map[string]string {
		"debug": "1",
		"logLevel": "warn",
		"version": "1.2.1",
	}
	for key, value := range config {
		fmt.Println(key, "=", value)
	}
	// Using break and continue to control loops
	for {
		r := rand.Intn(8)
		if r % 3 == 0{
			fmt.Println("Skip")
			continue
		}else if r % 2 == 0 {
			fmt.Println("Stop")
			break
		} 
		fmt.Println(r)
	}

}