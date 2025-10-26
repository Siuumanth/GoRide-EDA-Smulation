package ridesinprogress

import (
	pubsub "RideBooking/pubsub"
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

func RideService(ridesEventQueue <-chan pubsub.DriverMatchedEvent, eventBus chan<- any) {
	for event := range ridesEventQueue {

		simulateRide()
		// sending RideCompletedEvent
		rideEvent := pubsub.RideCompletedEvent{
			DriverName:  event.DriverName,
			UserName:    event.UserName,
			Amount:      event.Amount,
			Destination: event.Destination,
		}

		// Ride completed send ride to Event
		eventBus <- rideEvent
	}
}
