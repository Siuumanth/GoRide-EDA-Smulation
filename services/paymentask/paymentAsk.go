package paymentAsk

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
func PaymentAskService(PaymentAskEventQueue <-chan any, eventBus chan<- any) {
	reader := bufio.NewReader(os.Stdin)
	for event := range PaymentAskEventQueue {
		switch e := event.(type) {
		case events.RideCompletedEvent:
			var cardNumber string
			fmt.Print("Please enter your card number: ")
			cardNumber, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Error reading CN:", err)
				continue // Skip processing if input fails
			}

			eventBus <- events.PaymentAskEvent{
				CardNumber:  cardNumber,
				Amount:      e.Amount,
				Destination: e.Destination,
				UserName:    e.UserName,
			}
		default:
			log.Printf("PaymentAsk Received event of type %T", event)
		}
	}
}
