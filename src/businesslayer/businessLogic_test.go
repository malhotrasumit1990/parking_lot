package businesslayer

import (
	"testing"

	"github.com/ParkingLot/src/dataLayer"
	"github.com/stretchr/testify/assert"
)

var dummyParking []dataLayer.CarPark
var parkingImpl ParkingLotImpl

func initializeParking(t *testing.T) {

	carPark1 := dataLayer.CarPark{1, "KA-01-HH-1234", "White"}
	carPark2 := dataLayer.CarPark{2, "KA-01-HH-8888", "Red"}
	carPark3 := dataLayer.CarPark{3, "KA-01-MM-1234", "Green"}
	carPark4 := dataLayer.CarPark{}
	carPark5 := dataLayer.CarPark{5, "KA-01-LL-9999", "White"}
	carPark6 := dataLayer.CarPark{}

	dataLayer.ParkingCapacity = 6

	dummyParking = append(dummyParking, carPark1, carPark2, carPark3, carPark4, carPark5, carPark6)
	dataLayer.Parkings = dummyParking
}

func TestParkVehicle(t *testing.T) {
	initializeParking(t)
	SpaceOccupied = 3
	regNum := "DL-05-MM-8653"
	color := "Black"
	parkingImpl = ParkingLotImpl{}
	parkingImpl.ParkVehicle(regNum, color)

	assert.Equal(t, "Black", dataLayer.Parkings[3].Color)
	assert.Equal(t, "DL-05-MM-8653", dataLayer.Parkings[3].RegNum)
	assert.Equal(t, 4, SpaceOccupied)

}
func TestIsFull(t *testing.T) {
	initializeParking(t)
	SpaceOccupied = 3
	val := isFull()
	assert.Equal(t, false, val)

}
func TestUnparkVehicle(t *testing.T) {
	initializeParking(t)
}
func TestGetStatus(t *testing.T) {
	initializeParking(t)
}
func TestGetRegNumsWithColor(t *testing.T) {
	initializeParking(t)
}

func TestGetSlotsForColor(t *testing.T) {
	initializeParking(t)
}

func TestGetSlotforRegistrationNumber(t *testing.T) {
	initializeParking(t)
}
