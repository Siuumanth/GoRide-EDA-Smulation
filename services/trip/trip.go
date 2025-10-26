package tripService

import (
	pubsub "RideBooking/pubsub"
	"fmt"
	"log"
	"math"
	"time"
)

/*
Trip service logic:
Get trip details from user
forward event to match driver
*/

var locationMap = map[string][]float64{
	"New York":    {40.7128, -74.0060},
	"Los Angeles": {34.0522, -118.2437},
	"Chicago":     {41.8781, -87.6298},
	"Houston":     {29.7633, -95.3632},
	"Phoenix":     {33.4484, -112.0739},
}

func TripService(tripEventQueue <-chan pubsub.UserEvent, eventBus chan<- any) {
	// Get user Details
	for userReq := range tripEventQueue {
		log.Printf("Hello, %v!", userReq.UserName)
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
		time.Sleep(3 * time.Second)

		// Fire and forget
		tripEvent := pubsub.TripEvent{
			UserName:    userReq.UserName,
			Lat:         lat,
			Long:        long,
			Destination: destination,
			Amount:      amount,
		}

		// send event to event bus
		eventBus <- tripEvent
	}
}
