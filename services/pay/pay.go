package services

import (
	"RideBooking/events"
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
Ask payment details(card number only)
*/

// PaymentAskService is a function that asks the user for their card number.
// It takes in a channel of payment ask events and a channel to send the events to.
// It reads the card number from the user and sends it to the event bus.
func PaymentService(PaymentEventQueue <-chan any, eventBus chan<- any) {
	reader := bufio.NewReader(os.Stdin)
	for event := range PaymentEventQueue {
		switch e := event.(type) {
		case events.PaymentAskEvent:
			fmt.Printf("Confirm payment of %f to %s? (y/n): ", e.Amount, e.Destination)
			var confirmation string
			confirmation, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Error reading confirmation:", err)
				continue // Skip processing if input fails
			}

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
