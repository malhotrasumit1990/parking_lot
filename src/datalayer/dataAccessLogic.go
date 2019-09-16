package datalayer

import (
	"strconv"
	"sync"
)

//Parkings is exported to be accessed in other packages
var Parkings map[string]string
var singleton sync.Once

//CreateParkingSpace : creates parking space once
//This map consists of mapping between parking number and vehicle parked there.
func CreateParkingSpace() map[string]string {

	singleton.Do(func() {

		for parkingNumber := 0; parkingNumber < ParkingCapacity; parkingNumber++ {
			key := "P" + strconv.Itoa(parkingNumber)
			Parkings[key] = ""
		}

	})
	return Parkings
}
