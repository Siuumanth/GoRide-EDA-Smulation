package matchDriverService

import (
	events "RideBooking/events"
	utils "RideBooking/utils"
	"log"
	"math"
	"math/rand"
	"os"
)

/*
This service will provide the nearest driver to the user
Get nearest driver will take input from the event bus and process it
Trip event has langitude and logitude

Calculation of nearest driver:
- Calc sum of userLat - driverlat and long for all and find minimum
*/

var drivers *[]utils.Driver = utils.GenerateDrivers()

func DriverService(driverEventQueue <-chan any, eventBus chan<- any) {

	for tripReq := range driverEventQueue {
		// calculate nearest driverW
		switch event := tripReq.(type) {
		case events.TripRequestedEvent:
			for i := 0; i < 1e5; {
				i++
			}
			var driverID int = -1
			var minDist float64 = math.MaxFloat64

			for i, driver := range *drivers {
				if driver.Available == false {
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

			eta := 1 + rand.Intn(2)

			driverMatchedEvent := events.DriverMatchedEvent{
				DriverName: nearestDriver.Name,
				UserName:   event.UserName,
				Amount:     event.Amount,
				ETA:        float64(eta),
			}

			eventBus <- driverMatchedEvent

		case events.TripCompletedEvent:
			for i := range *drivers {
				if (*drivers)[i].Name == event.DriverName {
					(*drivers)[i].Available = true
					//	log.Printf("✅ Driver %v is now available again", event.DriverName)
					break
				}
			}

		default:
			log.Printf("MatchDriver Received event of type %T", event)

		}
	}

}
