package businesslayer

import (
	"errors"
	"fmt"
	"log"
	"regexp"

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

	match, _ := regexp.MatchString("[a-z][0-9]", number)
	if match {
		if !(len(number) > 5 && len(number) < 8) {
			log.Println("Length of vehicle registeration must be greater than 5 and less than 8")
			return false
		}
	} else {
		log.Println("Vehicle number should be alphanumeric only")
		return false
	}
	return true
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
