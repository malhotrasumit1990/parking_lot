package businesslayer

import (
	"fmt"

	"github.com/ParkingLot/src/datalayer"
)

//ParkingLot Interface
type ParkingLot interface {
	parkVehicle(vehicleNum string) string
	unparkVehicle(vehicleNum string) string
	getVehicleParkedAt(parkingNum string) string
	getRegNum(color string) []regNumbers
	getSlotNum(regNum string) int
	getSlots(color string) []int
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

	carPark := datalayer.CarPark{vehicleNum, color, parkingNum}
	datalayer.Parkings[parkingNum] = carPark
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

func (impl ParkingLotImpl) GetVehicleParkedAt(parkingNum string) string {
	vehicleNum := datalayer.Parkings[parkingNum]
	if vehicleNum == "" {
		return parkingAvailable
	}
	return vehicleNum
}
