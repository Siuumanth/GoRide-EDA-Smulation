package notification

import (
	pubsub "RideBooking/pubsub"
	"fmt"
)

func NotificationService(NotificationEventQueue <-chan any, eventBus chan<- any) {
	for event := range NotificationEventQueue {
		// e has the value as well as type
		switch e := event.(type) {
		case pubsub.TripEvent:
			handleTripNotificationEvent(e, eventBus)
		case pubsub.DriverMatchedEvent:
			handleDriverMatchedNotificationEvent(e, eventBus)
		case pubsub.RideCompletedEvent:
			handleRideCompletedNotificationEvent(e, eventBus)
		case pubsub.PaymentEvent:
			handlePaymentDoneNotificationEvent(e, eventBus)
		}
	}
}

func handleTripNotificationEvent(tripEvent pubsub.TripEvent, eventBus chan<- any) {
	fmt.Printf("Hey %s user, you have requested a trip to %s for amount %f. We will find you a driver soon.\n", tripEvent.UserName, tripEvent.Destination, tripEvent.Amount)
	// send event to event bus
	eventBus <- tripEvent
}

func handleDriverMatchedNotificationEvent(driverMatchedEvent pubsub.DriverMatchedEvent, eventBus chan<- any) {
	fmt.Printf("Hey %s user, we have matched you with a driver %s for a trip to %s for amount %f.\n", driverMatchedEvent.UserName, driverMatchedEvent.DriverName, driverMatchedEvent.Destination, driverMatchedEvent.Amount)
	// send event to event bus
	eventBus <- driverMatchedEvent
}

func handleRideCompletedNotificationEvent(rideCompletedEvent pubsub.RideCompletedEvent, eventBus chan<- any) {
	fmt.Printf("Hey %s user, your trip with driver %s to %s has been completed for amount %f.\n", rideCompletedEvent.UserName, rideCompletedEvent.DriverName, rideCompletedEvent.Destination, rideCompletedEvent.Amount)
	// send event to event bus
	eventBus <- rideCompletedEvent
}

func handlePaymentDoneNotificationEvent(paymentDoneEvent pubsub.PaymentEvent, eventBus chan<- any) {
	fmt.Printf("Hey %s user, your payment of %f with transaction ID %s has been successful.\n", paymentDoneEvent.UserName, paymentDoneEvent.Amount, paymentDoneEvent.TransactionID)
	// send event to event bus
	eventBus <- pubsub.NotificationDoneEvent{
		UserName:    paymentDoneEvent.UserName,
		Amount:      paymentDoneEvent.Amount,
		Destination: paymentDoneEvent.Destination,
	}
}
