package paymentAsk

import (
	"RideBooking/events"
	"fmt"
	"log"
)

/*
Ask payment details(card number only)
*/

// PaymentAskService is a function that asks the user for their card number.
// It takes in a channel of payment ask events and a channel to send the events to.
// It reads the card number from the user and sends it to the event bus.
func PaymentAskService(PaymentAskEventQueue <-chan any, eventBus chan<- any) {
	for event := range PaymentAskEventQueue {
		switch e := event.(type) {
		case events.DriverMatchedEvent:
			cardNumber := ""
			fmt.Print("Please enter your card number: ")
			_, err := fmt.Scanf("%s\n", &cardNumber)

			if err != nil {
				log.Fatal(err)
			}

			eventBus <- events.PaymentAskEvent{
				CardNumber:  cardNumber,
				Amount:      e.Amount,
				Destination: e.Destination,
				UserName:    e.UserName,
			}
		default:
			log.Printf("Received event of type %T", event)
		}
	}
}
