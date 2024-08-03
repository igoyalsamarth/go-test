package main

import "fmt"

const LoginToken string = "ThisIsALoginToken"

func main() {
	var username string = "Samarth"
	fmt.Println(username)
	fmt.Printf("variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type %T \n", smallVal)

	var smallFloat float32 = 255.455678945678
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type %T \n", smallFloat)

	var randomVal int
	fmt.Println(randomVal)
	fmt.Printf("variable is of type %T \n", randomVal)

	numberOfUsers := 3000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type %T \n", LoginToken)

}
