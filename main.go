package main

import (
	core "RideBooking/core"
	events "RideBooking/events"
	"context"
	"fmt"
	"runtime"
	"time"
)

/*
7 Services:
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
var NUM_USERS = 20000
var EVENTBUS_CAPACITY = 1000

func main() {
	runtime.GOMAXPROCS(10)
	fmt.Println("Starting the system...")
	eventBus := make(chan any, EVENTBUS_CAPACITY)
	pubsubs := core.InitPubSub()

	// initiate context for the first Worker Pool
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go events.StartEventBus(eventBus, pubsubs)    // start eventBus
	go core.InitAutoScaler(eventBus, ctx, cancel) // start worker pools, auto scalers

	fmt.Println("Simulating Users")
	go SimulateRandomUsers(eventBus, NUM_USERS) // start user generation

	// waits for cancel Fucntion to be called
	<-ctx.Done()
	fmt.Println("Shutting down the system...")

	time.Sleep(3 * time.Second)
}
