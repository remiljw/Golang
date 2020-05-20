package main

// Using comaprison and logical operators o see what level of membership 
// a user has based on the number of visits they've had.
 import "fmt"
 func main() {
	 visits := 15
	 fmt.Println("First visit :", visits == 1)
	 fmt.Println("Return visit :", visits != 1)
	 fmt.Println("Silver member :", visits >= 10 && visits < 21)
	 fmt.Println("Gold member :", visits > 20 && visits <= 30)
	 fmt.Println("Platinum member :", visits > 30)


 }