package services

import (
	"RideBooking/events"
	"context"
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
func PaymentService(PaymentEventQueue <-chan any, eventBus chan<- any, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-PaymentEventQueue:
			if !ok {
				return
			}
			switch e := event.(type) {
			case events.PaymentAskEvent:
				time.Sleep(100 * time.Millisecond)

				eventBus <- events.PaymentEvent{
					Amount:        e.Amount,
					Destination:   e.Destination,
					UserName:      e.UserName,
					TransactionID: fmt.Sprintf("%s-%s", e.UserName, e.Destination),
					Success:       true,
					DriverName:    e.DriverName,
				}
			default:
				log.Printf("Received event of type %T", event)
			}
		}
	}
}
