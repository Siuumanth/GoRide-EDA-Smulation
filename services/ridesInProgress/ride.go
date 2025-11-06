package ridesinprogress

import (
	events "RideBooking/events"
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

func RideService(ridesEventQueue <-chan any, eventBus chan<- any) {
	for e := range ridesEventQueue {
		switch event := e.(type) {
		case events.DriverMatchedEvent:
			// end if no river forund
			if event.DriverName == "" {
				continue
			}

			time.Sleep(time.Duration(event.ETA) * time.Second)

			simulateRide()
			// sending RideCompletedEvent
			rideEvent := events.RideCompletedEvent{
				DriverName:  event.DriverName,
				UserName:    event.UserName,
				Amount:      event.Amount,
				Destination: event.Destination,
			}

			// Ride completed send ride to Event
			eventBus <- rideEvent
		}
	}
}
