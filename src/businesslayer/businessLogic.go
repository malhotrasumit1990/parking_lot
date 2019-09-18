package businesslayer

import (
	"fmt"
	"strconv"

	"github.com/ParkingLot/src/dataLayer"
)

//ParkingLot Interface
type ParkingLot interface {
	ParkVehicle(regNum string) string
	UnparkVehicle(slot int) string
	GetStatus() []dataLayer.CarPark
	GetRegNumsWithColor(color string) []string
	GetSlotsForColor(color string) []int
	GetSlotforRegistrationNumber(regNum string) string
}

//ParkingLotImpl impementation
type ParkingLotImpl struct {
}

//SpaceOccupied in parking lot
var SpaceOccupied int

func (impl ParkingLotImpl) ParkVehicle(regNum string, color string) string {

	if isFull() {
		return parkingFull
	}

	valid := validateVehicleNumber(regNum)
	if !valid {
		return "enter valid number"
	}

	index, err := closestEmptyParking()
	if err != nil {
		return err.Error()
	}
	//carPark := dataLayer.Parkings[index]

	carPark := dataLayer.CarPark{index + 1, regNum, color}
	dataLayer.Parkings[index] = carPark
	SpaceOccupied = SpaceOccupied + 1
	return fmt.Sprintf("Allocated slot number: %d \n", index+1)
}

func (impl ParkingLotImpl) UnparkVehicle(parkingNum int) string {

	if isEmpty() {
		return parkingEmpty
	}
	dataLayer.Parkings[parkingNum-1] = dataLayer.CarPark{}
	SpaceOccupied = SpaceOccupied - 1
	return fmt.Sprintf("Slot number %d is free ", parkingNum)
}

func (impl ParkingLotImpl) GetStatus() []dataLayer.CarPark {

	carsParked := []dataLayer.CarPark{}

	for _, v := range dataLayer.Parkings {
		if v.Slot != 0 {
			carsParked = append(carsParked, v)
		}
	}
	return carsParked
}

func (impl ParkingLotImpl) GetRegNumsWithColor(color string) []string {

	regNums := make([]string, 0)

	for _, v := range dataLayer.Parkings {
		if v.Slot != 0 {
			car := v
			if car.Color == color {
				regNums = append(regNums, car.RegNum)
			}
		}
	}
	return regNums
}

func (impl ParkingLotImpl) GetSlotforRegistrationNumber(regNum string) string {

	for _, v := range dataLayer.Parkings {
		if v.Slot != 0 {
			car := v
			if car.RegNum == regNum {
				return strconv.Itoa(car.Slot)
			}
		}
	}
	return "Not found"

}

func (impl ParkingLotImpl) GetSlotsForColor(color string) []int {

	slots := make([]int, 0)

	for _, v := range dataLayer.Parkings {
		if v.Slot != 0 {
			car := v
			if car.Color == color {
				slots = append(slots, car.Slot)
			}
		}
	}
	return slots
}
