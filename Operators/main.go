package main

import (
	"fmt"
	"time"
)

func main() {
	// Concatenation
	givenName := "John"
	familyName := "Smith"
	fullName := givenName + " " + familyName
	fmt.Println("Hello,", fullName)
	// Shorthand Operators
	count := 5
	count += 5
	fmt.Println(count)
	count++
	fmt.Println(count)
	count--
	fmt.Println(count)
	count -= 5
	fmt.Println(count)
	name := "John"
	name += " Smith"
	fmt.Println("Hello,", name)
	// Zero values
	var counts int
	fmt.Printf("Count : %#v \n", counts)
	var discount float64
	fmt.Printf("Discount : %#v \n", discount)
	var debug bool
	fmt.Printf("Debug : %#v \n", debug)
	var message string
	fmt.Printf("Message : %#v \n", message)
	var emails []string
	fmt.Printf("Emails : %#v \n", emails)
	var startTime time.Time
	fmt.Printf("Start : %#v \n", startTime)
	// Getting Pointers
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}
	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time: %#v\n", t)
	// Getting a Value from a Pointer

	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("t: %#v\n", *t)
		fmt.Printf("time : %#v\n", t.String())
	}
	
}
// Function Design with Pointers
func add5Value(count int) {
	count += 5
	fmt.Println("add5Value :", count)
}
func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point :", *count)
}
func addmore() {
	var count int
	add5Value(count)
	fmt.Println("add5Value post:", count)
	add5Point(&count)
	fmt.Println("add5Point post:", count)
}