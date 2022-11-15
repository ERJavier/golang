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

	x := 10 
	y := 50

	// Addition 

	result := x + y
	fmt.Println("Addition is ", result)

	// Subtraction
	
	result = y - x
	fmt.Println("Subtraction is ", result)

	// Multiplication

	result = x * y
	fmt.Println("Multiplication is ", result)

	// Division

	result = x / y
	fmt.Println("Division ", result)

	// Module 

	result = x % y
	fmt.Println("Module  ", result)

	// Incremental 
	x++
	fmt.Println("Incremental ", x)

	// Decremental
	x--
	fmt.Print("Decremental ", x)
 }
