package main

import "fmt"

func isPalindromo(text string){
	var textReverse string

	for i:=len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("its a palindromo")
	} else {
		fmt.Println("its not a palindromo")
	}
}

func main() {
	slice := []string{"Hi", "whats", "up?"}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("casa")


}
