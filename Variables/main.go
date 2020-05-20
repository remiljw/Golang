package main

import (
	"fmt"
	"time"
)

var defaultOffset = 10
var foo string = "bar"

func main() {
	var baz string = "qux"
	fmt.Println(foo, baz)
}

// Declaring multiple variables at once
var (
	Debug       bool
	LogLevel    string    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}

//Shorthand variable Declaration

func main() {
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}

// Declaring Multiple Variables with a Short Variable Declaration
func main() {
	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}

// Declaring Multiple Variables from a Function
func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}

// Changing the Value of a Variable
func main() {
	offset := defaultOffset
	fmt.Println(offset)
	offset = offset + defaultOffset
	fmt.Println(offset)
}

// Changing Multiple Values at Once
func main() {
	query, limit, offset := "bat", 10, 0
	query, limit, offset = "ball", offset, 20
	fmt.Println(query, limit, offset)
}
