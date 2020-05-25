package main

import "fmt"
// Creating Slices from a Slice
func message() string{
	s := []int{1,2,3,4,5,6,7,8,9}
	// extract the first value
	m := fmt.Sprintln("First:", s[0], s[0:1], s[:1])
	// extract the last element
	m += fmt.Sprintln("Last :", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
	// get the first five values and add to the message
	m += fmt.Sprintln("First 5:", s[:5])
	// get the last four values
	m += fmt.Sprintln("Last 4:", s[5:])
	// get 5 values from  the middle
	m += fmt.Sprintln("Middle 5:", s[2:7])
	return m
}
func main(){
	fmt.Print(message())
}