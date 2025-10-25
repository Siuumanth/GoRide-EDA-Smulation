package pubsub

// Trip initiation
type TripEvent struct {
	UserID string
	Amount float64
	Driver string
}

// ask payment detials with amount
type PaymentAskEvent struct {
	Amount float64
	Card   CardDetails
}

type CardDetails struct {
	CardNumber string
	CardHolder string
	Expiration string
	CVV        string
}

type PaymentEvent struct {
	TransactionID string
	Status        string
}

type NotificationEvent struct {
	UserName string
	Amount   float64
	Status   string
}
