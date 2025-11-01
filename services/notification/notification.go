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
		case events.TripRequestedEvent:
			handleTripNotificationEvent(e, eventBus)
		case events.DriverMatchedEvent:
			handleDriverMatchedNotificationEvent(e, eventBus)
		case events.RideCompletedEvent:
			handleRideCompletedNotificationEvent(e, eventBus)
		case events.PaymentEvent:
			handlePaymentDoneNotificationEvent(e, eventBus)
		case events.TripCompletedEvent:
			handleTripCompletedNotificationEvent(e, eventBus)
		default:
			log.Printf("Notification Service Received event of type %T", e)
		}

	}
}

func handleTripNotificationEvent(tripEvent events.TripRequestedEvent, eventBus chan<- any) {
	str := fmt.Sprintf("REQUEST:    %s has requested a trip to %s for amount: $%f.", tripEvent.UserName, tripEvent.Destination, tripEvent.Amount)
	SaveNotification(str)
}
func handleDriverMatchedNotificationEvent(driverMatchedEvent events.DriverMatchedEvent, eventBus chan<- any) {
	str := fmt.Sprintf("MATCH:    %s, has been matched to a driver %s for a trip to %s for amount: $%f.\nETA is %f seconds\n", driverMatchedEvent.UserName, driverMatchedEvent.DriverName, driverMatchedEvent.Destination, driverMatchedEvent.Amount, driverMatchedEvent.ETA)
	time.Sleep(time.Duration(driverMatchedEvent.ETA+1) * time.Second)

	SaveNotification(str)
}
func handleRideCompletedNotificationEvent(rideCompletedEvent events.RideCompletedEvent, eventBus chan<- any) {
	str := fmt.Sprintf("RIDECOMPLETED:    %s's trip with driver %s to %s has been completed for amount: $%f.\n", rideCompletedEvent.UserName, rideCompletedEvent.DriverName, rideCompletedEvent.Destination, rideCompletedEvent.Amount)

	SaveNotification(str)
}

func handlePaymentDoneNotificationEvent(paymentDoneEvent events.PaymentEvent, eventBus chan<- any) {
	str := fmt.Sprintf("PAYMENTCOMPLETED:    %s's payment of %f with transaction ID %s has been %t.\n", paymentDoneEvent.UserName, paymentDoneEvent.Amount, paymentDoneEvent.TransactionID, paymentDoneEvent.Success)

	SaveNotification(str)
}

func handleTripCompletedNotificationEvent(tripCompletedEvent events.TripCompletedEvent, eventBus chan<- any) {
	str := fmt.Sprintf("TRIPCOMPLETED:    %s has completed a trip to %s for amount: $%f.\n", tripCompletedEvent.UserName, tripCompletedEvent.Destination, tripCompletedEvent.Amount)
	SaveNotification(str)
}
