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

func MatchDriver(driverEventQueue <-chan any, eventBus chan<- any) {

	for tripReq := range driverEventQueue {
		// calculate nearest driver
		switch event := tripReq.(type) {
		case events.TripEvent:
			var driverID int
			var minDist float64 = 100
			for i, driver := range *drivers {
				if driver.Status == "busy" {
					continue
				}
				dist := math.Abs(event.Lat-driver.Lat) + math.Abs(event.Long-driver.Long)

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
				UserName:   event.UserName,
				Amount:     event.Amount,
				ETA:        eta,
			}

			eventBus <- driverMatchedEvent

		}
	}

}
