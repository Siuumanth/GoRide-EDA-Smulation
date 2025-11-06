package events

import (
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
			//	fmt.Printf("EventBus Received event of type %T\n", event)
			for _, subscriber := range pubsub["UserEvent"] {
				subscriber <- event
			}
		case TripRequestedEvent:
			//fmt.Printf("EventBus Received event of type %T\n", event)
			for _, subscriber := range pubsub["TripRequestedEvent"] {
				subscriber <- event
			}
		case DriverMatchedEvent:
			//	fmt.Printf("EventBus Received event of type %T\n", event)
			for _, subscriber := range pubsub["DriverMatchedEvent"] {
				subscriber <- event
			}
		case PaymentAskEvent:
			//	fmt.Printf("EventBus Received event of type %T\n", event)
			for _, subscriber := range pubsub["PaymentAskEvent"] {
				subscriber <- event
			}

		case PaymentEvent:
			//	fmt.Printf("EventBus Received event of type %T\n", event)
			for _, subscriber := range pubsub["PaymentEvent"] {
				subscriber <- event
			}
		case NotificationDoneEvent:
			//	fmt.Printf("EventBus Received event of type %T\n", event)
			for _, subscriber := range pubsub["NotificationEvent"] {
				subscriber <- event
			}
		case RideCompletedEvent:
			//	fmt.Printf("EventBus Received event of type %T\n", event)
			for _, subscriber := range pubsub["RideCompletedEvent"] {
				subscriber <- event
			}

		case TripCompletedEvent:
			//	fmt.Printf("EventBus Received event of type %T\n", event)
			for _, subscriber := range pubsub["TripCompletedEvent"] {
				subscriber <- event
			}
		case TerminationEvent:
			//	fmt.Printf("EventBus Received event of type %T\n", event)
			for _, subscriber := range pubsub["TerminationEvent"] {
				subscriber <- event
			}

		default:
			log.Printf("EventBus Received event of type %T", event)
		}
	}

}
