package abstract

import (
	"fmt"
)

const (
	LuxuryCarType   = 1
	FamiliarCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamiliarCarType:
		return new(FamiliarCar), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}
