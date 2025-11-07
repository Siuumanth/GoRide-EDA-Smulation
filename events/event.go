package events

// User initiates trip
type UserEvent struct {
	UserName    string
	Lat         float64
	Long        float64
	Destination string
}

// Trip initiation
type TripRequestedEvent struct {
	UserName    string
	Amount      float64
	Lat         float64
	Long        float64
	Destination string
	Status      string
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
	DriverName  string
}

type PaymentEvent struct {
	TransactionID string
	Success       bool
	Destination   string
	UserName      string
	Amount        float64
	DriverName    string
}

type NotificationDoneEvent struct {
	UserName    string
	Amount      float64
	Destination string
}

// type TerminationEvent struct {
// 	UserName   string
// 	Status     string
// 	DriverName string
// }

type TripCompletedEvent struct {
	UserName    string
	DriverName  string
	Amount      float64
	Destination string
}

type TerminationEvent struct {
	UserName string
	Status   string
	Message  string
}
