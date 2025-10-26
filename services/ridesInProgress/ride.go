package ridesinprogress

import (
	events "RideBooking/events"
	"fmt"
	"math/rand"
	"time"
)

/*
function to simulate rides in progress
will run as goroutine
*/

func simulateRide() {
	sleepTime := 2 + rand.Intn(3)                      // int
	time.Sleep(time.Duration(sleepTime) * time.Second) // convert to Duration
}

func RideService(ridesEventQueue <-chan any, eventBus chan<- any) {
	for e := range ridesEventQueue {
		switch event := e.(type) {
		case events.DriverMatchedEvent:
			time.Sleep(time.Duration(event.ETA) * time.Second)
			fmt.Println("Driver has arrived at your location!")

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
