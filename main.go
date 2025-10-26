package main

import (
	events "RideBooking/events"
	"fmt"
	"log"
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
	eventBus := make(chan any)
	pubsubs := InitPubSub()
	// start eventBus
	go events.StartEventBus(eventBus, pubsubs)

	// start worker pools
	StartWorkerPools(eventBus)
	// start user prompt
	PromptUser(eventBus)

	time.Sleep(200 * time.Second)
}

func PromptUser(eventBus chan<- any) {
	var userName string
	fmt.Print("Welcome to GoRide!!\nEnter your username to start your journey: ")
	_, err := fmt.Scanf("%s\n", &userName)

	if err != nil {
		log.Fatal(err)
	}

	eventBus <- events.UserEvent{UserName: userName}
}
