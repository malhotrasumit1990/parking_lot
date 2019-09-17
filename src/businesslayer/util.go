package businesslayer

import (
	"errors"
	"strings"

	"github.com/ParkingLot/src/dataLayer"
)

//returns first empty parking
var closestEmptyParking = func() (int, error) {

	for k, v := range dataLayer.Parkings {
		if v == "" {
			return k, nil
		}
	}
	return 0, errors.New("Parking full")
}

var validateVehicleNumber = func(number string) bool {

	if strings.Contains(number, "-") && len(number) == 13 {
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
