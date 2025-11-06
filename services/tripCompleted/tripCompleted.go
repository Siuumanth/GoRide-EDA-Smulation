package tripCompleted

import (
	events "RideBooking/events"
	"context"
	"log"
)

func TripCompletedService(tripCompletedQueue <-chan any, eventBus chan<- any, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case event, ok := <-tripCompletedQueue:
			if !ok {
				return
			}

			switch e := event.(type) {
			case events.PaymentEvent:
				handlePaymentOverEvent(e, eventBus)
			default:
				log.Printf("TripCompletedService received unknown event type: %T", e)
			}
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
