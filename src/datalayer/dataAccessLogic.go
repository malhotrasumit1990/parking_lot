package datalayer

type carPark struct {
	RegNum string
	Color  string
	Slot   int
}

//Parkings is exported to be accessed in other packages. mapping between parking slot and vehicle parked
var Parkings map[int]string

//ParkingCapacity is total space available
var ParkingCapacity int

//CreateParkingSpace : creates parking space
func CreateParkingSpace(capacity int) {
	for parkingNumber := 1; parkingNumber <= capacity; parkingNumber++ {
		Parkings[parkingNumber+ParkingCapacity] = ""
	}
	ParkingCapacity = ParkingCapacity + capacity
}
