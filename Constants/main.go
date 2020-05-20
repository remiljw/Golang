// Miscellaneous Codes
package main 
import "fmt" 
func main() { 
  count := 5 
  if count > 5 { 
	message := "Greater than 5" 
	fmt.Println(message) 
  } else { 
	message := "Not greater than 5" 
	fmt.Println(message) 
  } 
//   ////////////
  firstName := "Bob"
  familyname := "Smith"
  Age := 34
  peanuntAllergy := false 
  fmt.Println(firstName)
  fmt.Println(familyname)
  fmt.Println(Age)
  fmt.Println(peanuntAllergy)
//   /////////////////////
  {count := 0
	if count < 5 {
	  count := 10
	  count++
	  fmt.Println(count == 11)
	}}
// /////////////////
	a, b := 5, 10
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}


func swap(a *int, b *int) {
	*a, *b = *b, *a
}
