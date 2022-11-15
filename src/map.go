package main

import (
	"fmt"
)

func main(){
	m := make(map[string]int)

	m["Javi"] = 29
	m["pepito"] = 20

	fmt.Println(m)

	// Run map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Find Value
	value := m["javi"]
	fmt.Println(value)

	value2, ok := m["pepe"]
	fmt.Println(value2, ok)
}