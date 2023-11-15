package main

import (
	"fmt"
	"myapp/packageone"
)

var myVar = "Var"

func main() {

	fmt.Println(myVar)

	myFunc()

	newString := packageone.PackageVar
	fmt.Println("from packagevar:", newString)
}

func myFunc() {

}
