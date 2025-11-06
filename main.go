package main

import (
	core "RideBooking/core"
	events "RideBooking/events"
	"context"
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
   6. Trip Completed Service
   7. Termination Serivce

   Event bus: goroutine with input channel, map of subscribers and publishers
   Event is in the form of data and publisher

*/

func main() {
	eventBus := make(chan any, 500) // initial size = 100
	pubsubs := core.InitPubSub()

	// initiate context for the first goroutine
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go events.StartEventBus(eventBus, pubsubs)    // start eventBus
	go core.InitAutoScaler(eventBus, ctx, cancel) // start worker pools, auto scalers
	go PromptUser(eventBus)                       // start user prompt

	// Waits for cancel Fucntion to be called
	<-ctx.Done()
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
