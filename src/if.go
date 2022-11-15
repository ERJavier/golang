package main

import (
	"fmt"
	"log"
	"strconv"
)

func main(){
	value1 := 1
	value2 := 2

	if value1 == 1 {
		fmt.Println("its 1")
	}else {
		fmt.Println("its not 1")
	}

	// 	with and 
	if value1 == 1 && value2 == 2 {
		fmt.Println("its true")
	}

	// with or
	if value1 == 0 || value2 ==2{
		fmt.Println("its true")
	}

	// string to number 
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)
}