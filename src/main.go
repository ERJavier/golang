package main

import "fmt"

func main() {
	// Constant
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi, " p2:", pi2)
	fmt.Println()

	// Variables 
	base := 12
	var height int = 14
	var area int

	fmt.Print(base, height, area)
	fmt.Println()

	// Zero Value

	var a int // 0
	var b float64 //0
	var c string  // empty 
	var d bool // true | false

	fmt.Println(a, b, c, d)


	// area cube

	const baseCube = 10
	areaCube := baseCube * baseCube
	fmt.Println("cube is: ", areaCube)
}
