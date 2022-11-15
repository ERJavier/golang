package main

import "fmt"

func main() {

	helloMessage := "Hello"
	worldMessage := "world"

	//println

	fmt.Println(helloMessage, worldMessage)

	// printf

	name := "Javier"
	lastName := "Cepeda"
	fmt.Printf("Mi nombre es %s y mi apellido es %s  \n", name, lastName)

	// Sprintf

	message := fmt.Sprintf("Mi nombre es %s y mi apellido es %s ", name, lastName)
	fmt.Println(message)


	// typeof 

	fmt.Printf("helloMessage: %T\n", helloMessage )

}
