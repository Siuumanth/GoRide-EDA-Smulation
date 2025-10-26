package events

import (
	"log"
)

/*
this will have:

*/

// Event bus is a goroutine that will run and look for events in the eventChan channel
// Map will have a string: list of channels of all subcribers
func eventBus(eventChan <-chan any, pubsub map[string][]chan any) {

	for event := range eventChan {
		switch event.(type) {
		case TripEvent:
			for _, subscriber := range pubsub["TripEvent"] {
				subscriber <- event
			}
		case PaymentAskEvent:
			for _, subscriber := range pubsub["PaymentAskEvent"] {
				subscriber <- event
			}

		case PaymentEvent:
			for _, subscriber := range pubsub["PayEvent"] {
				subscriber <- event
			}
		case NotificationDoneEvent:
			for _, subscriber := range pubsub["NotificationEvent"] {
				subscriber <- event
			}

		default:
			log.Printf("Received event of type %T", event)
		}
	}

}
