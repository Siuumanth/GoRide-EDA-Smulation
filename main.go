package main

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

import (
	events "RideBooking/Events"
	"fmt"
	"log"
)

func main() {
	eventBus := make(chan any)
	pubsubs := InitPubSub()
	// start eventBus
	go events.StartEventBus(eventBus, pubsubs)

}

func PromptUser() {
	var userName string
	fmt.Print("Welcome to GoRide, Enter your username: ")
	_, err := fmt.Scanf("%s\n", &userName)

	if err != nil {
		log.Fatal(err)
	}

	//eventBus <- events.UserEvent{UserName: userName, Event: "Starting"}
}
