package main

import "fmt"

// LoginToken - Public declaration of a variable begins with an upper case character
const LoginToken string = "dfkj544534erjk"

func main() {
	var username string = "steve2640"
	fmt.Printf("Variable name is of type: %T \n ", username)

	var isUserLoggedIn bool = true
	fmt.Printf("Variable name is of type: %T \n ", isUserLoggedIn)

	var smallVal int = 256
	fmt.Printf("Variable name is of type: %T \n ", smallVal)

	var smallFloat float32 = 256.00
	fmt.Printf("Variable name is of type: %T \n ", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable name is of type: %T \n ", anotherVariable)

	var website = "string"
	fmt.Println(website)

	//Declaring & initializing variables without var and type. But its allowed only within function block.
	//Outside of function block var keyword & type must be defined.
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable name is of type: %T \n ", LoginToken)
}
