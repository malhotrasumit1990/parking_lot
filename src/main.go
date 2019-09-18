package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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
	//for {
	var command string
	var param1 string
	var param2 string

	//accepts file path as a flag "filePath"
	filePath := flag.String("filePath", "", "Read inputs from file")
	flag.Parse()

	//reads from input file
	if *filePath != "" {
		file, err := os.Open(*filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Sscan(line, &command, &param1, &param2)
			executeLogicforParking(command, param1, param2)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		//reads from command line
		//fmt.Scanf("%s %s %s", &command, &param1, &param2)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Sscan(line, &command, &param1, &param2)
			executeLogicforParking(command, param1, param2)
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}

}

//}
func executeLogicforParking(command string, param1 string, param2 string) {

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
		fmt.Print(parkingLot.UnparkVehicle(slot))
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
