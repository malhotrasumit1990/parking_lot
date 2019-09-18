package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/parking_lot/src/businesslayer"
	"github.com/parking_lot/src/dataLayer"
)

//setup Parking lot

var parkingLot businesslayer.ParkingLotImpl

func init() {
	//initialize parking map
	dataLayer.Parkings = []dataLayer.CarPark{}
	parkingLot = businesslayer.ParkingLotImpl{}
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
			dataLayer.CreateParkingSpace(capacityInt)
			fmt.Printf("Created a parking lot with %d slots\n", capacityInt)
			break
		case "park":
			fmt.Print(parkingLot.ParkVehicle(param1, param2))
			break
		case "leave":
			slot, err := strconv.Atoi(param1)
			if err != nil {
				fmt.Println(err.Error() + "Enter a proper integer for slot")
				break
			}
			fmt.Println(parkingLot.UnparkVehicle(slot))
			break
		case "status":
			fmt.Println(parkingLot.GetStatus())
			break
		case "registration_numbers_for_cars_with_colour":
			fmt.Printf("%s \n", parkingLot.GetRegNumsWithColor(param1))
			break
		case "slot_numbers_for_cars_with_colour":
			fmt.Printf("%d \n", parkingLot.GetSlotsForColor(param1))
			break
		case "slot_number_for_registration_number":
			fmt.Printf("%s \n", parkingLot.GetSlotforRegistrationNumber(param1))
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
