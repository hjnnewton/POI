package main

import "fmt"
import "api"

func main(){
	fmt.Println("hello world")
	address, _ := location.GetLocation("39.983424",
		"116.322987", "4zbsyOHMfdK5BDnhnkthr53Z")
	location.GetSurrounding("39.983424",
		"116.322987", "4zbsyOHMfdK5BDnhnkthr53Z")
	fmt.Println(address)
}