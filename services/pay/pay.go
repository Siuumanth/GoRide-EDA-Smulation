package services

import (
	"RideBooking/events"
	"fmt"
	"log"
	"time"
)

/*
Ask payment details(card number only)
*/

// PaymentAskService is a function that asks the user for their card number.
// It takes in a channel of payment ask events and a channel to send the events to.
// It reads the card number from the user and sends it to the event bus.
func PaymentService(PaymentEventQueue <-chan any, eventBus chan<- any) {
	for event := range PaymentEventQueue {
		switch e := event.(type) {
		case events.PaymentAskEvent:
			fmt.Printf("Confirm payment of %f to %s? (y/n): ", e.Amount, e.Destination)
			var confirmation string
			time.Sleep(500 * time.Millisecond)
			fmt.Println("y")

			if confirmation == "n" {
				eventBus <- events.PaymentEvent{
					Amount:        e.Amount,
					Destination:   e.Destination,
					UserName:      e.UserName,
					TransactionID: "",
					Status:        "fail",
				}
				continue
			}

			eventBus <- events.PaymentEvent{
				Amount:        e.Amount,
				Destination:   e.Destination,
				UserName:      e.UserName,
				TransactionID: fmt.Sprintf("%s-%s", e.UserName, e.Destination),
				Status:        "success",
			}
		default:
			log.Printf("Received event of type %T", event)
		}
	}
}
