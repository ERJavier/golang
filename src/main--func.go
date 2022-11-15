package main

import "fmt"

func nomalFunction(message string) {
	fmt.Println(message)

}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 7
}

func doubleReturn(a int) (c, d int) {
	return a, a*3
}
func main() {
	nomalFunction("Hello world")
	tripleArgument(1, 3, "hello")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(3)
	fmt.Println("Value1 and Value2:", value1, value2)
}