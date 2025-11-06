package main

import (
	core "RideBooking/core"
	events "RideBooking/events"
	"fmt"
	"math/rand"
	"time"
)

/*

   5 Services:
   1. Driver Service
   2. Trip Service
   3. PaymentAskService
   4. PaymentService
   5. NotificationService

   Event bus: goroutine with input channel, map of subscribers and publishers
   Event is in the form of data and publisher

*/

func main() {
	eventBus := make(chan any, 1000) // initial size = 100
	pubsubs := core.InitPubSub()
	// start eventBus
	go events.StartEventBus(eventBus, pubsubs)

	// start worker pools
	core.InitAutoScaler(eventBus)
	// start user prompt
	PromptUser(eventBus)

	time.Sleep(200 * time.Second)
}

func PromptUser(eventBus chan<- any) {
	time.Sleep(1 * time.Second)
	userName := fmt.Sprintf("user-%d", rand.Intn(100000))

	lat := rand.Float64() * 100
	long := rand.Float64() * 100
	destinations := []string{"BLR", "HK", "NY", "DC", "AMS", "LA", "CH", "LS", "AM", "SF", "MUM"}
	destination := destinations[rand.Intn(len(destinations))]

	eventBus <- events.UserEvent{
		UserName:    userName,
		Lat:         lat,
		Long:        long,
		Destination: destination,
	}
}
