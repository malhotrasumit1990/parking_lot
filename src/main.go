package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ParkingLot/src/datalayer"
)

//setup Parking lot

func init() {
	//initialize parking map
	datalayer.Parkings = make(map[int]string)
}

func main() {
	for {
		var command string
		var param1 string
		var param2 string

		fmt.Scanf("%s %s %s", &command, &param1, &param2)

		switch command {
		case "create_parking_lot":
			capacityInt, err := strconv.Atoi(param1)
			if err != nil {
				fmt.Println(err.Error() + "Enter a proper integer for parking numbers")
				break
			}
			datalayer.CreateParkingSpace(capacityInt)
			fmt.Printf("Created a parking lot with %d slots\n", capacityInt)
			break
		case "park":
			fmt.Println(param1, param2)
			break
		case "leave":
			fmt.Println("three")
			break
		case "status":
			fmt.Println("four")
			break
		case "registration_numbers_for_cars_with_colour":
			fmt.Println("five")
			break
		case "slot_numbers_for_cars_with_colour":
			fmt.Println("six")
			break
		case "slot_number_for_registration_number":
			fmt.Println("seven")
			break
		case "exit":
			os.Exit(1)
		default:
			fmt.Print("")
			break
		}
	}

}

/*fmt.Print("Enter First String: ") //Print function is used to display output in same line
var first string
fmt.Scanln(&first) // Take input from user
fmt.Print("Enter Second String: ")
var second string
fmt.Scanln(&second)
fmt.Print(first + second)*/
