package utils

import (
	"fmt"
	"math/rand/v2"
)

var NUM_DRIVERS = 2000

type Driver struct {
	Name      string
	Lat       float64
	Long      float64
	Available bool
	Rating    float64
}

func GenerateDrivers() *[]Driver {
	drivers := make([]Driver, NUM_DRIVERS)
	for i := 0; i < NUM_DRIVERS; i++ {
		drivers[i] = Driver{
			Name:      fmt.Sprintf("driver-%d", i+1),
			Lat:       rand.Float64()*100 - 50, // random location
			Long:      rand.Float64()*100 - 50,
			Available: true,
			Rating:    3.5 + rand.Float64()*1.5, // rating between 3.5–5.0
		}
	}
	return &drivers
}
