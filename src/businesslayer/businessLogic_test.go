package businesslayer

import (
	"testing"

	"github.com/parking_lot/src/dataLayer"
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
	SpaceOccupied = 4

	dummyParking = append(dummyParking, carPark1, carPark2, carPark3, carPark4, carPark5, carPark6)
	dataLayer.Parkings = dummyParking
}

func demolishParking(t *testing.T) {
	dataLayer.Parkings = []dataLayer.CarPark{}
	dataLayer.ParkingCapacity = 0
	SpaceOccupied = 0
}

func TestParkingLotApp(t *testing.T) {
	initializeParking(t)
	parkingImpl = ParkingLotImpl{}

	t.Run("Parked", func(t *testing.T) {
		regNum := "DL-05-MM-8653"
		color := "Black"
		parkingImpl.ParkVehicle(regNum, color)

		assert.Equal(t, "Black", dataLayer.Parkings[3].Color)
		assert.Equal(t, "DL-05-MM-8653", dataLayer.Parkings[3].RegNum)
		assert.Equal(t, 5, SpaceOccupied)
	})
	t.Run("UnparkVehicle", func(t *testing.T) {
		slot := 3
		msg := parkingImpl.UnparkVehicle(slot)
		assert.Equal(t, "Slot number 3 is free \n", msg)
		slotVal := dataLayer.Parkings[2].Slot
		assert.Equal(t, 0, slotVal)

	})
	t.Run("IsParkingFull", func(t *testing.T) {
		val := isFull()
		assert.Equal(t, false, val)
	})
	t.Run("GetStatus", func(t *testing.T) {
		msg := parkingImpl.GetStatus()
		assert.Equal(t, 4, len(msg))
	})
	t.Run("GetRegNumsWithColor", func(t *testing.T) {
		color := "White"
		msg := parkingImpl.GetRegNumsWithColor(color)
		assert.Equal(t, 2, len(msg))
		assert.Equal(t, "KA-01-HH-1234", msg[0])
		assert.Equal(t, "KA-01-LL-9999", msg[1])

	})
	t.Run("GetSlotsForColor", func(t *testing.T) {
		color := "White"
		msg := parkingImpl.GetSlotsForColor(color)
		assert.Equal(t, 2, len(msg))
		assert.Equal(t, 1, msg[0])
		assert.Equal(t, 5, msg[1])

	})
	t.Run("GetSlotsForColor", func(t *testing.T) {
		regNum := "KA-01-HH-1234"
		msg := parkingImpl.GetSlotforRegistrationNumber(regNum)
		assert.Equal(t, "1", msg)
	})
	t.Run("Parking Full", func(t *testing.T) {
		SpaceOccupied = 6
		regNum := "DL-05-MM-8653"
		color := "Black"
		msg := parkingImpl.ParkVehicle(regNum, color)
		assert.Equal(t, parkingFull, msg)
	})
	demolishParking(t)

}
