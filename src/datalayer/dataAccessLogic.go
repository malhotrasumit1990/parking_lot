package dataLayer

type CarPark struct {
	Slot   int
	RegNum string
	Color  string
}

//Parkings is exported to be accessed in other packages. mapping between parking slot and vehicle parked
//var Parkings map[int]interface{}
var Parkings []CarPark

//ParkingCapacity is total space available
var ParkingCapacity int

//CreateParkingSpace : creates parking space
func CreateParkingSpace(capacity int) {
	for parkingNumber := 1; parkingNumber <= capacity; parkingNumber++ {
		Parkings = append(Parkings, CarPark{})
	}
	ParkingCapacity = ParkingCapacity + capacity
}
