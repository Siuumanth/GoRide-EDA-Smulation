package matchDriverService

import (
	events "RideBooking/events"
	utils "RideBooking/utils"
	"math"
)

/*
This service will provide the nearest driver to the user
Get nearest driver will take input from the event bus and process it
Trip event has langitude and logitude

Calculation of nearest driver:
- Calc sum of userLat - driverlat and long for all and find minimum
*/

var drivers *[]utils.Driver = utils.GenerateDriverData()

func MatchDriver(driverEventQueue <-chan events.TripEvent, eventBus chan<- any) {

	for tripReq := range driverEventQueue {
		// calculate nearest driver
		var driverID int
		var minDist float64 = 100
		for i, driver := range *drivers {
			if driver.Status == "busy" {
				continue
			}
			dist := math.Abs(tripReq.Lat-driver.Lat) + math.Abs(tripReq.Long-driver.Long)

			if dist < minDist {
				driverID = i
				minDist = dist
			}
		}

		// dereference first, then index
		nearestDriver := (*drivers)[driverID]
		(*drivers)[driverID].Status = "busy"

		eta := minDist

		driverMatchedEvent := events.DriverMatchedEvent{
			DriverName: nearestDriver.Name,
			UserName:   tripReq.UserName,
			Amount:     tripReq.Amount,
			ETA:        eta,
		}

		eventBus <- driverMatchedEvent

	}
}
