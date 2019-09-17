package businesslayer

import (
	"fmt"

	"github.com/ParkingLot/src/dataLayer"
	datalayer "github.com/ParkingLot/src/dataLayer"
)

//ParkingLot Interface
type ParkingLot interface {
	ParkVehicle(vehicleNum string) string
	UnparkVehicle(vehicleNum string) string
	//GetVehicleParkedAt(parkingNum string) string
	//getSlotNum(regNum string) int
	//getSlots(color string) []int
}

//ParkingLotImpl impementation
type ParkingLotImpl struct {
}

//SpaceOccupied in parking lot
var SpaceOccupied int

func (impl ParkingLotImpl) ParkVehicle(vehicleNum string, color string) string {

	if isFull() {
		return parkingFull
	}

	valid := validateVehicleNumber(vehicleNum)
	if !valid {
		return "enter valid number"
	}

	parkingNum, err := closestEmptyParking()
	if err != nil {
		return err.Error()
	}

	carPark := dataLayer.CarPark{vehicleNum, color, parkingNum}
	dataLayer.Parkings[parkingNum] = carPark
	SpaceOccupied = SpaceOccupied + 1
	return fmt.Sprintf("Allocated slot number: %d \n", parkingNum)
}

func (impl ParkingLotImpl) UnparkVehicle(parkingNum int) string {

	if isEmpty() {
		return parkingEmpty
	}
	datalayer.Parkings[parkingNum] = ""
	SpaceOccupied = SpaceOccupied - 1
	return fmt.Sprintf("Slot number %d is free ", parkingNum)
}
