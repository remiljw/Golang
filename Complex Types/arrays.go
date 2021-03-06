package main

import (
	"fmt"
)
// Defining an Array
func defineArray() [10]int {
	var arr [10]int
	return arr
}
// Comparing Arrays
func compArrays() (bool, bool, bool){
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0,0,0,0,0}
	arr4 := [5]int{0,0,0,0,9}

	return arr1 == arr2, arr1 == arr3, arr1 == arr4
}
// Initializing Array using keys
func compArray() (bool, bool, [10]int){
	var arr1 [10]int
	arr2 := [...]int{9: 0}
	arr3 := [10]int{1,9: 10, 4: 5}

	return arr1 == arr2, arr1 == arr3, arr3
}
// Reading a Single Item from an Array
	func message ()string{
		arr := [...]string{
			"ready", "Get", "Go", "to",
		}
		return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
	}
	// Writing to an Array
	func words()string{
		arr := [4]string{"ready", "Get", "Go", "to"}

		arr[1] = "It's"
		arr[0] = "time"

		return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
	 }
	//  Looping over an Array
	func loop() string{
		m := ""
		arr := [4]int{1,2,3,4}
		for i := 0; i < len(arr); i++ {
			arr[i] = arr[i] * arr[i]
			m += fmt.Sprintf("%v : %v\n", i, arr[i])
		}
		return m
	}
	// Modifying contents of an array in a loop
	// fills an array with numbers from 1 to 10
	func fillArray(arr [10]int) [10] int {
		for i := 0; i < len(arr); i++{
			arr [i] = i + 1
		}
		return arr
	}
			// multiplies an array bt itself then sets the result back to the array
	func opArray(arr [10]int) [10]int {
		for i := 0; i < len(arr); i++{
			arr[i] = arr[i] * arr[i]
		}
		return arr
	}
func main() {
	// Defines Array
	fmt.Printf("%#v\n", defineArray())
	// Compares Array
	comp1, comp2, comp3 := compArrays()
	fmt.Println("[5]int == [5]int{0}  :", comp1)
	fmt.Println("[5]int == [...]int{0,0,0,0,0}  :", comp2)
	fmt.Println("[5]int == [5]int{0,0,0,0,9}  :", comp3)
	// Initializes Array and Compares
	res1, res2, arr3 := compArray()
	fmt.Println("[10]int == [...]int{9: 0}  :", res1)
	fmt.Println("[10]int == [10]int{1,9: 10, 4: 5}", res2)
	fmt.Println("arr3        :", arr3)
	// Reading items drom an array
	fmt.Print(message())
	// Writing to an Array
	fmt.Print(words())
	//  Looping over an Array
	fmt.Print(loop())
	// Modifying Contents of an Array
	var arr [10]int
	arr = fillArray(arr)
	arr = opArray(arr)
	fmt.Println(arr)
}
