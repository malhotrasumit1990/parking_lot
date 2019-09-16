package main

import (
	"fmt"

	"github.com/ParkingLot/src/datalayer"
)

//setup Parking lot
func init() {

	Parkings := datalayer.CreateParkingSpace()

	if len(Parkings) == 0 {
		panic("Parking space not created Successfully")
	}

}

func main() {

	fmt.Println("Hello World !!")

	fmt.Print("Enter First String: ") //Print function is used to display output in same line
	var first string
	fmt.Scanln(&first) // Take input from user
	fmt.Print("Enter Second String: ")
	var second string
	fmt.Scanln(&second)
	fmt.Print(first + second)

}
