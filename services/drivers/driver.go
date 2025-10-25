package drivers

import (
	pubsub "RideBooking/pubsub"
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

var drivers []utils.Driver = utils.GenerateDriverData()

func MatchDriver(driverEventQueue <-chan pubsub.TripEvent, eventBus chan<- any) {
	for userReq := range driverEventQueue {
		var driverID int
		var minDist float64 = 100
		for i, driver := range drivers {
			dist := math.Abs(userReq.Lat-driver.Lat) + math.Abs(userReq.Long-driver.Long)

			if dist < minDist {
				driverID = i
				minDist = dist
			}
		}

		nearestDriver := drivers[driverID]
		eta := minDist

		driverMatchedEvent := pubsub.DriverMatchedEvent{
			DriverName: nearestDriver.Name,
			UserName:   userReq.UserName,
			Amount:     userReq.Amount,
			ETA:        eta,
		}

		eventBus <- driverMatchedEvent

	}
}
