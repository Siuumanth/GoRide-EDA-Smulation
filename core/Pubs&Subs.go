package core

var n int = 10
var (
	UserEventChan          = make(chan any, n)
	TripRequestedEventChan = make(chan any, n)
	DriverEventChan        = make(chan any, n)
	RideEventChan          = make(chan any, n)
	PaymentAskEventChan    = make(chan any, n)
	PaymentEventChan       = make(chan any, n)
	NotificationEventChan  = make(chan any, n)
	TripCompletedEventChan = make(chan any, n)
	TerminationEventChan   = make(chan any, n)
)

// init and assign a channel for each event type
func InitPubSub() map[string][]chan any { // map of publisher event → list of subscriber channels
	var PubSub = map[string][]chan any{
		"UserEvent":          {TripRequestedEventChan},
		"TripRequestedEvent": {DriverEventChan, NotificationEventChan, TerminationEventChan},
		"DriverMatchedEvent": {RideEventChan, NotificationEventChan, TerminationEventChan},
		"RideCompletedEvent": {PaymentAskEventChan, NotificationEventChan},
		"PaymentAskEvent":    {PaymentEventChan},
		"PaymentEvent":       {NotificationEventChan, TripCompletedEventChan},
		"TripCompletedEvent": {DriverEventChan, NotificationEventChan},
		"TerminationEvent":   {NotificationEventChan},
	}

	return PubSub
}
