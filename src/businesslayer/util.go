package businesslayer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ParkingLot/src/dataLayer"
)

//returns first empty parking
var closestEmptyParking = func() (string, error) {

	for k, v := range dataLayer.Parkings {
		if v == "" {
			fmt.Println("Closest empty parking : " + k)
			return k, nil
		}
	}
	return "", errors.New("Parking full")
}

var validateVehicleNumber = func(number string) bool {

	if strings.Contains("-") && len(number) == 13 {
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
