package notification

import (
	events "RideBooking/events"
	"fmt"
	"log"
	"time"
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
		case events.TerminationEvent:
			handleTerminationNotificationEvent(e, eventBus)
		default:
			log.Printf("Notification Service Received event of type %T", e)
		}

	}
}

func handleTripNotificationEvent(tripEvent events.TripEvent, eventBus chan<- any) {
	fmt.Printf("NOTIFICATION: Hey %s, you have requested a trip to %s for amount: $%f. We will find you a driver soon.\n", tripEvent.UserName, tripEvent.Destination, tripEvent.Amount)
}

func handleDriverMatchedNotificationEvent(driverMatchedEvent events.DriverMatchedEvent, eventBus chan<- any) {
	fmt.Printf("NOTIFICATION: Hey %s, we have matched you with a driver %s for a trip to %s for amount: $%f.\nETA is %f seconds\n", driverMatchedEvent.UserName, driverMatchedEvent.DriverName, driverMatchedEvent.Destination, driverMatchedEvent.Amount, driverMatchedEvent.ETA)
	time.Sleep(time.Duration(driverMatchedEvent.ETA+1) * time.Second)

	fmt.Println("NOTIFICATION: Trip has started.")
}

func handleRideCompletedNotificationEvent(rideCompletedEvent events.RideCompletedEvent, eventBus chan<- any) {
	fmt.Printf("NOTIFICATION: Hey %s, your trip with driver %s to %s has been completed for amount: $%f.\n", rideCompletedEvent.UserName, rideCompletedEvent.DriverName, rideCompletedEvent.Destination, rideCompletedEvent.Amount)
}

func handlePaymentDoneNotificationEvent(paymentDoneEvent events.PaymentEvent, eventBus chan<- any) {
	fmt.Println("handling payment done noti")
	if paymentDoneEvent.Status == "fail" {
		fmt.Printf("NOTIFICATION: Hey %s, your payment of %f has been failed.\n", paymentDoneEvent.UserName, paymentDoneEvent.Amount)
	} else {

		// payment success notification
		fmt.Printf("NOTIFICATION: Hey %s, your payment of %f with transaction ID %s has been successful.\n", paymentDoneEvent.UserName, paymentDoneEvent.Amount, paymentDoneEvent.TransactionID)
		// send event to event bus
	}

}

func handleTerminationNotificationEvent(terminationEvent events.TerminationEvent, eventBus chan<- any) {
	switch terminationEvent.Status {
	case "success":
		fmt.Printf("NOTIFICATION: Trip Successful")
	case "fail":
		fmt.Printf("NOTIFICATION: Trip Failed")
	case "cancel":
		fmt.Printf("NOTIFICATION: Trip Cancelled")
	}
	eventBus <- terminationEvent
}
