package terminate

import (
	events "RideBooking/events"
)

func TerminateService(terminationEventQueue <-chan any, eventBus chan<- any) {
	for event := range terminationEventQueue {
		// e has the value as well as type
		switch e := event.(type) {
		case events.PaymentEvent:
			handlePaymentOverEvent(e, eventBus)
		}
	}
}

func handlePaymentOverEvent(event events.PaymentEvent, eventBus chan<- any) {
	switch event.Status {
	case "fail":
		notificationEvent := events.TerminationEvent{
			UserName: event.UserName,
			Status:   "fail",
		}
		eventBus <- notificationEvent
		return
	case "success":
		notificationEvent := events.TerminationEvent{
			UserName: event.UserName,
			Status:   "success",
		}
		eventBus <- notificationEvent
	}
}
