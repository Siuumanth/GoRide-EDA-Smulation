package tripService

import (
	events "RideBooking/events"
	"fmt"
	"log"
	"math"
)

/*
Trip service logic:
Get trip details from user
forward event to match driver
*/

var locationMap = map[string][]float64{
	"NY":  {40.7128, -74.0060},
	"LA":  {34.0522, -118.2437},
	"CH":  {41.8781, -87.6298},
	"LS":  {29.7633, -95.3632},
	"AM":  {33.4484, -112.0739},
	"SF":  {37.7749, -122.4194},
	"DC":  {38.9072, -77.0369},
	"BLR": {12.9716, -79.8773},
	"MUM": {-17.8239, 31.0465},
}

func TripService(tripEventQueue <-chan any, eventBus chan<- any) {
	// Get user Details
	for userReq := range tripEventQueue {
		switch event := userReq.(type) {
		case events.UserEvent:
			log.Printf("Hello, %v!", event.UserName)
			var lat, long float64
			fmt.Print("Enter your latitude: ")
			_, err := fmt.Scan(&lat)
			if err != nil {
				log.Fatalf("Failed to read latitude: %v", err)
			}

			fmt.Print("Enter your longitude: ")
			_, err = fmt.Scan(&long)
			if err != nil {
				log.Fatalf("Failed to read longitude: %v", err)
			}

			fmt.Print("Enter your destination: ")
			var destination string
			_, err = fmt.Scan(&destination)
			if err != nil {
				log.Fatalf("Failed to read destination: %v", err)
			}

			latDiff := math.Abs(lat - locationMap[destination][0])
			longDiff := math.Abs(long - locationMap[destination][1])
			amount := float64(latDiff + longDiff) // 1 latlong = $1

			fmt.Println("Finding Nearest Driver......")

			// Fire and forget
			tripEvent := events.TripEvent{
				UserName:    event.UserName,
				Lat:         lat,
				Long:        long,
				Destination: destination,
				Amount:      amount,
			}

			// send event to event bus
			eventBus <- tripEvent
		}
	}

}
