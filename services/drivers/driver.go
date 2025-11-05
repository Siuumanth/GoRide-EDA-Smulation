package drivers

import (
	events "RideBooking/events"
	"log"
	"math/rand"
)

/*
This service will provide the nearest driver to the user
Get nearest driver will take input from the event bus and process it
Trip event has langitude and logitude

Calculation of nearest driver:
- Calc sum of userLat - driverlat and long for all and find minimum
*/

func DriverService(driverEventQueue <-chan any, eventBus chan<- any) {

	for tripReq := range driverEventQueue {
		// calculate nearest driverW
		switch event := tripReq.(type) {
		case events.TripRequestedEvent:
			nearestDriver := assignNearestDriver(event, &mu)
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
