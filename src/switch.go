package main

import "fmt"

func main(){
	module := 5 % 2
	switch module {
	case 0:
		fmt.Println("Its even")
	default:
		fmt.Println("Its odd")
	}

	// without condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Its bigger than 100")
	case value < 0:
		fmt.Println("Its lower than 0")
	default:
		fmt.Println("Dont apply")
	}
}