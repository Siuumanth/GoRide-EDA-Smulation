package tripCompleted

import (
	events "RideBooking/events"
)

func TripCompletedService(terminationEventQueue <-chan any, eventBus chan<- any) {
	for event := range terminationEventQueue {
		// e has the value as well as type
		switch e := event.(type) {
		case events.PaymentEvent:
			handlePaymentOverEvent(e, eventBus)
		}
	}
}

func handlePaymentOverEvent(event events.PaymentEvent, eventBus chan<- any) {
	tripCompletedEvent := events.TripCompletedEvent{
		UserName:   event.UserName,
		DriverName: event.DriverName,
	}
	eventBus <- tripCompletedEvent
}
