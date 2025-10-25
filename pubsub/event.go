package pubsub

// Trip initiation
type TripEvent struct {
	UserID   string
	UserName string
	Amount   float64
	Lat      float64
	Long     float64
}

type DriverMatchedEvent struct {
	DriverName string
	UserName   string
	Amount     float64
	ETA        float64
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
