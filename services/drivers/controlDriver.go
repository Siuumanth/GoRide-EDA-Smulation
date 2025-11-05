package drivers

import (
	"RideBooking/events"
	"RideBooking/utils"
	"math"
	"os"
	"sync"
)

var drivers *[]utils.Driver = utils.GenerateDrivers()

var mu sync.Mutex

// Mutex for safe updates of driver info

func assignNearestDriver(event events.TripRequestedEvent, mu *sync.Mutex) utils.Driver {

	for i := 0; i < 1e5; {
		i++
	}
	var driverID int = -1
	var minDist float64 = math.MaxFloat64

	var driver utils.Driver
	var i int
	// lock mutex
	mu.Lock()
	defer mu.Unlock()

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
		os.Exit(1)
	}

	// dereference first, then index
	nearestDriver := &(*drivers)[driverID]
	nearestDriver.Available = false

	return driver
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
