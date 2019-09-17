package dataLayer

type CarPark struct {
	RegNum string
	Color  string
	Slot   int
}

//Parkings is exported to be accessed in other packages. mapping between parking slot and vehicle parked
var Parkings map[int]interface{}

//ParkingCapacity is total space available
var ParkingCapacity int

//CreateParkingSpace : creates parking space
func CreateParkingSpace(capacity int) {
	for parkingNumber := 1; parkingNumber <= capacity; parkingNumber++ {
		Parkings[parkingNumber+ParkingCapacity] = ""
	}
	ParkingCapacity = ParkingCapacity + capacity
}
