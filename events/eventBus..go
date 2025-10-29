package events

import (
	"fmt"
	"log"
)

/*
this will have:

*/

// Event bus is a goroutine that will run and look for events in the eventChan channel
// Map will have a string: list of channels of all subcribers
func StartEventBus(eventChan <-chan any, pubsub map[string][]chan any) {

	for event := range eventChan {
		switch event.(type) {
		case UserEvent:
			for _, subscriber := range pubsub["UserEvent"] {
				subscriber <- event
			}
		case TripEvent:
			for _, subscriber := range pubsub["TripEvent"] {
				subscriber <- event
			}
		case DriverMatchedEvent:
			for _, subscriber := range pubsub["DriverMatchedEvent"] {
				subscriber <- event
			}
		case PaymentAskEvent:
			for _, subscriber := range pubsub["PaymentAskEvent"] {
				subscriber <- event
			}

		case PaymentEvent:
			fmt.Println("PaymentEvent recieved in event bus")
			for _, subscriber := range pubsub["PaymentEvent"] {
				subscriber <- event
			}
		case NotificationDoneEvent:
			for _, subscriber := range pubsub["NotificationEvent"] {
				subscriber <- event
			}
		case RideCompletedEvent:
			for _, subscriber := range pubsub["RideCompletedEvent"] {
				subscriber <- event
			}

		case TerminationEvent:
			for _, subscriber := range pubsub["TerminationEvent"] {
				subscriber <- event
			}

		default:
			log.Printf("EventBus Received event of type %T", event)
		}
	}

}
