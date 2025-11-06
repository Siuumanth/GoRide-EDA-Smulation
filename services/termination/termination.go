package termination

import "RideBooking/events"

func TerminationService(terminationEventQueue <-chan any, eventBus chan<- any) {
	for event := range terminationEventQueue {
		// e has the value as well as type
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
			eventBus <- events.TerminationEvent{
				UserName: e.UserName,
				Message:  "User Terminated Process",
				Status:   "fail",
			}

		}
	}
}
