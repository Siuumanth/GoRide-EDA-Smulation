package tripService

import (
	events "RideBooking/events"
	"context"
	"math"
	"time"
)

/*
Trip service logic:
Get trip details from user
forward event to match driver
*/

var locationMap = map[string][]float64{
	"BLR": {12.34, 56.78},
	"HK":  {23.45, 67.89},
	"NY":  {34.56, 78.90},
	"DC":  {45.67, 89.01},
	"AMS": {56.78, 90.12},
	"LA":  {67.89, 101.23},
	"CH":  {78.90, 112.34},
	"LS":  {89.01, 123.45},
	"AM":  {90.12, 134.56},
	"SF":  {101.23, 145.67},
	"MUM": {112.34, 156.78},
}

func TripService(tripEventQueue <-chan any, eventBus chan<- any, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case req, ok := <-tripEventQueue:
			if !ok {
				return
			}
			switch event := req.(type) {
			case events.UserEvent:

				lat := event.Lat
				long := event.Long

				// valid destinations from map
				destination := event.Destination

				// distance & cost calc
				destCoords := locationMap[destination]
				latDiff := math.Abs(lat - destCoords[0])
				longDiff := math.Abs(long - destCoords[1])
				amount := latDiff + longDiff

				// simulate amount calculation
				time.Sleep(10 * time.Millisecond)

				tripEvent := events.TripRequestedEvent{
					UserName:    event.UserName,
					Lat:         lat,
					Long:        long,
					Destination: destination,
					Amount:      amount,
				}

				eventBus <- tripEvent
			}
		}

	}

}
