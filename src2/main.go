package main

import (
	"curso_golang_platzi/src2/mypackage"
	"fmt"
)

func main(){
	var myCar mypackage.CarPublic
	myCar.Brand = "Tesla"
	fmt.Println(myCar)
}