package drivers

import (
	"RideBooking/events"
	"RideBooking/utils"
	"math"
	"sync"
	"time"
)

var drivers *[]utils.Driver = utils.GenerateDrivers()

var mu sync.Mutex

// Mutex for safe updates of driver info

func assignNearestDriver(event events.TripRequestedEvent, mu *sync.Mutex, eventBus chan<- any) *utils.Driver {

	for i := 0; i < 1e5; {
		i++
	}
	var driverID int = -1
	var minDist float64 = math.MaxFloat64

	var driver utils.Driver
	var i int
	tries := 10
	// lock mutex
	mu.Lock()
	defer mu.Unlock()

	for driverID == -1 && tries != 0 {
		for i, driver = range *drivers {
			if !driver.Available {
				continue
			}
			dist := math.Abs(event.Lat-driver.Lat) + math.Abs(event.Long-driver.Long)

			if dist < minDist {
				driverID = i
				minDist = dist
			}
		}
		if driverID == -1 {
			time.Sleep(500 * time.Millisecond)
			tries--
		}

	}

	// If number of tries exhausted, the driver struct will be NULL, and termination service will handle it
	if driverID == -1 {
		return &utils.Driver{}
	}

	// dereference first, then index
	nearestDriver := &(*drivers)[driverID]
	nearestDriver.Available = false

	return nearestDriver // already has pointer to driver
}

func releaseDriver(driverName string, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	for i := range *drivers {
		if (*drivers)[i].Name == driverName {
			(*drivers)[i].Available = true
			break
		}
	}
}
