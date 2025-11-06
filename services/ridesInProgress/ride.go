package ridesinprogress

import (
	events "RideBooking/events"
	"context"
	"log"
	"math/rand"
	"time"
)

/*
function to simulate rides in progress
will run as goroutine
*/

func simulateRide() {
	sleepTime := 1 + rand.Intn(2)                      // int
	time.Sleep(time.Duration(sleepTime) * time.Second) // convert to Duration
}

func RideService(ridesEventQueue <-chan any, eventBus chan<- any, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case e, ok := <-ridesEventQueue:
			if !ok {
				return // channel closed
			}

			switch event := e.(type) {
			case events.DriverMatchedEvent:
				// skip if no driver found
				if event.DriverName == "" {
					continue
				}

				time.Sleep(time.Duration(event.ETA) * time.Second)
				simulateRide()

				rideEvent := events.RideCompletedEvent{
					DriverName:  event.DriverName,
					UserName:    event.UserName,
					Amount:      event.Amount,
					Destination: event.Destination,
				}

				eventBus <- rideEvent
			default:
				log.Printf("RideService received unknown event type: %T", e)
			}
		}
	}
}
