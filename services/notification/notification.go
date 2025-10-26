package notification

import (
	events "RideBooking/events"
	"fmt"
)

func NotificationService(NotificationEventQueue <-chan any, eventBus chan<- any) {
	for event := range NotificationEventQueue {
		// e has the value as well as type
		switch e := event.(type) {
		case events.TripEvent:
			handleTripNotificationEvent(e, eventBus)
		case events.DriverMatchedEvent:
			handleDriverMatchedNotificationEvent(e, eventBus)
		case events.RideCompletedEvent:
			handleRideCompletedNotificationEvent(e, eventBus)
		case events.PaymentEvent:
			handlePaymentDoneNotificationEvent(e, eventBus)
		}
	}
}

func handleTripNotificationEvent(tripEvent events.TripEvent, eventBus chan<- any) {
	fmt.Printf("Hey %s user, you have requested a trip to %s for amount %f. We will find you a driver soon.\n", tripEvent.UserName, tripEvent.Destination, tripEvent.Amount)
	// send event to event bus
	eventBus <- tripEvent
}

func handleDriverMatchedNotificationEvent(driverMatchedEvent events.DriverMatchedEvent, eventBus chan<- any) {
	fmt.Printf("Hey %s user, we have matched you with a driver %s for a trip to %s for amount %f.\n", driverMatchedEvent.UserName, driverMatchedEvent.DriverName, driverMatchedEvent.Destination, driverMatchedEvent.Amount)
	// send event to event bus
	eventBus <- driverMatchedEvent
}

func handleRideCompletedNotificationEvent(rideCompletedEvent events.RideCompletedEvent, eventBus chan<- any) {
	fmt.Printf("Hey %s user, your trip with driver %s to %s has been completed for amount %f.\n", rideCompletedEvent.UserName, rideCompletedEvent.DriverName, rideCompletedEvent.Destination, rideCompletedEvent.Amount)
	// send event to event bus
	eventBus <- rideCompletedEvent
}

func handlePaymentDoneNotificationEvent(paymentDoneEvent events.PaymentEvent, eventBus chan<- any) {
	if paymentDoneEvent.Status == "fail" {
		fmt.Printf("Hey %s user, your payment of %f has been failed.\n", paymentDoneEvent.UserName, paymentDoneEvent.Amount)
		// send event to event bus
		eventBus <- events.NotificationDoneEvent{
			UserName:    paymentDoneEvent.UserName,
			Amount:      paymentDoneEvent.Amount,
			Destination: paymentDoneEvent.Destination,
		}
		return
	}

	// payment success notification
	fmt.Printf("Hey %s user, your payment of %f with transaction ID %s has been successful.\n", paymentDoneEvent.UserName, paymentDoneEvent.Amount, paymentDoneEvent.TransactionID)
	// send event to event bus
	eventBus <- events.NotificationDoneEvent{
		UserName:    paymentDoneEvent.UserName,
		Amount:      paymentDoneEvent.Amount,
		Destination: paymentDoneEvent.Destination,
	}
}
