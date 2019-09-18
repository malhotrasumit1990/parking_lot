package businesslayer

import (
	"errors"
	"strings"

	"github.com/parking_lot/src/dataLayer"
)

//returns parking that is nearest to the entry
var closestEmptyParking = func() (int, error) {

	for k, v := range dataLayer.Parkings {
		if v.Slot == 0 {
			return k, nil
		}
	}
	return -1, errors.New("Parking full")
}

//just a nominal validation
var validateVehicleNumber = func(number string) bool {

	if strings.Contains(number, "-") && len(number) >= 10 && len(number) <= 13 {
		return true
	}
	return false
}

var isEmpty = func() bool {
	if SpaceOccupied == 0 {
		return true
	}
	return false
}

var isFull = func() bool {
	if SpaceOccupied == dataLayer.ParkingCapacity {
		return true
	}
	return false
}
