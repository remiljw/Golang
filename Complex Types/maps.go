// Creating, Reading and Writing a Map
package main
import (
	"fmt"
	"os"
)
var users = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
  }
func getUsers() map[string] string{
	users := map[string]string{
		"305" : "Sue",
		"204" : "Bob",
		"631" : "Jake",
	}
	users["073"] = "Tracy"
	return users
}
func getUser(id string) (string,bool){
	users := getUsers()
	user, exists := users[id]
	return user, exists
}
func delUser(id string){
	delete(users, id)
}

func main(){
	if len(os.Args) < 2{
		fmt.Println("User ID not passed")
		os.Exit(1)
	}
	userID := os.Args[1]
	name, exists := getUser(userID)
	if !exists {
		fmt.Printf("Passed user ID (%v) not found. \nUsers:\n", userID)
		for key, value := range getUsers() {
			fmt.Println(" ID:", key, "Name:", value)
		}
		os.Exit(1)
	}
	fmt.Println("Name:", name)
	// fmt.Println("Users:", getUsers())
	// ///// to delete an element from a map
	// userID := os.Args[1]
	// delUser(userID)
	// fmt.Println("Users:", users)
}