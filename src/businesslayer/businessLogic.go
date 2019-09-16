package businesslayer

import "github.com/ParkingLot/src/datalayer"

//ParkingLot Interface
type ParkingLot interface {
	parkVehicle(vehicleNum string) string
	unparkVehicle(vehicleNum string) string
	getVehicleParkedAt(parkingNum string) string
}

//ParkingLotImpl impementation
type ParkingLotImpl struct {
}

//SpaceOccupied in parking lot
var SpaceOccupied int

func (impl ParkingLotImpl) parkVehicle(vehicleNum string) string {

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
	datalayer.Parkings[parkingNum] = vehicleNum
	SpaceOccupied = SpaceOccupied + 1
	return "Vehicle Parked"
}

func (impl ParkingLotImpl) unparkVehicle(vehicleNum string) string {

	if isEmpty() {
		return parkingEmpty
	}
	valid := validateVehicleNumber(vehicleNum)

	if !valid {
		return "enter valid number"
	}

	for k, v := range datalayer.Parkings {
		if v == vehicleNum {
			datalayer.Parkings[k] = ""
			SpaceOccupied = SpaceOccupied - 1
			break
		}
	}
	return "Vehicle UnParked"
}

func (impl ParkingLotImpl) getVehicleParkedAt(parkingNum string) string {
	vehicleNum := datalayer.Parkings[parkingNum]
	if vehicleNum == "" {
		return parkingAvailable
	}
	return vehicleNum
}
