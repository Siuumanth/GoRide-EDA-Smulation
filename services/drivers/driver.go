package drivers

import (
	events "RideBooking/events"
	"RideBooking/utils"
	"context"
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

func DriverService(driverEventQueue <-chan any, eventBus chan<- any, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case req, ok := <-driverEventQueue:
			if !ok {
				return // channel closed
			}

			switch event := req.(type) {
			case events.TripRequestedEvent:
				// calculate nearest driver
				nearestDriver := assignNearestDriver(event, &mu, eventBus)
				if nearestDriver == nil {
					nearestDriver = &utils.Driver{}
				}

				eta := 1 + rand.Intn(2)
				driverMatchedEvent := events.DriverMatchedEvent{
					DriverName: nearestDriver.Name,
					UserName:   event.UserName,
					Amount:     event.Amount,
					ETA:        float64(eta),
				}

				eventBus <- driverMatchedEvent

			case events.TripCompletedEvent:
				releaseDriver(event.DriverName, &mu)

			default:
				log.Printf("MatchDriver Received event of type %T", event)
			}
		}
	}
}
