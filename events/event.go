package events

// User initiates trip
type UserEvent struct {
	UserName string
}

// Trip initiation
type TripEvent struct {
	UserName    string
	Amount      float64
	Lat         float64
	Long        float64
	Destination string
}

type DriverMatchedEvent struct {
	DriverName  string
	UserName    string
	Amount      float64
	ETA         float64
	Destination string
}

type RideCompletedEvent struct {
	UserName    string
	DriverName  string
	Amount      float64
	Destination string
}

// ask payment detials with amount
type PaymentAskEvent struct {
	Amount      float64
	CardNumber  string
	Destination string
	UserName    string
}

type PaymentEvent struct {
	TransactionID string
	Status        string
	Destination   string
	UserName      string
	Amount        float64
}

type NotificationDoneEvent struct {
	UserName    string
	Amount      float64
	Destination string
}

type TerminationEvent struct {
	UserName  string
	Status    string
	EventType string
}
