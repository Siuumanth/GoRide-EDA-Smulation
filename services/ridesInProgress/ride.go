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
	sleepTime := 5 + rand.Intn(5)                      // int
	time.Sleep(time.Duration(sleepTime) * time.Second) // convert to Duration
}

func RideService(ridesEventQueue <-chan any, eventBus chan<- any) {
	for e := range ridesEventQueue {
		switch event := e.(type) {
		case events.DriverMatchedEvent:

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
