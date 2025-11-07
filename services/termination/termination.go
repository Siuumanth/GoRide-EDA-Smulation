package termination

import (
	"RideBooking/events"
	"context"
	"log"
)

func TerminationService(terminationEventQueue <-chan any, eventBus chan<- any, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case event, ok := <-terminationEventQueue:
			if !ok {
				return
			}

			switch e := event.(type) {
			case events.DriverMatchedEvent:
				if e.DriverName == "" {
					eventBus <- events.TerminationEvent{
						UserName: e.UserName,
						Message:  "Driver Not Found",
						Status:   "fail",
					}
				}

			case events.UserEvent:
				if e.Destination == "" {
					eventBus <- events.TerminationEvent{
						UserName: e.UserName,
						Message:  "User Terminated Process",
						Status:   "fail",
					}

				}

			case events.TripRequestedEvent:
				if e.Status == "fail" {
					eventBus <- events.TerminationEvent{
						UserName: e.UserName,
						Message:  "User Terminated Process",
						Status:   "fail",
					}
				}

			default:
				log.Printf("TerminationService received unknown event type: %T", e)
			}
		}
	}
}
