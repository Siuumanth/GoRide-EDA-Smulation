package main

var (
	UserEventChan          = make(chan any)
	TripEventChan          = make(chan any)
	DriverMatchedEventChan = make(chan any)
	RideCompletedEventChan = make(chan any)
	PaymentAskEventChan    = make(chan any)
	PaymentEventChan       = make(chan any)
	NotificationEventChan  = make(chan any)
	TerminationEventChan   = make(chan any)
)

// init and assign a channel for each event type
func InitPubSub() map[string][]chan any { // map of publisher event → list of subscriber channels
	var PubSub = map[string][]chan any{
		"UserEvent":          {TripEventChan},
		"TripEvent":          {DriverMatchedEventChan, NotificationEventChan},
		"DriverMatchedEvent": {RideCompletedEventChan, NotificationEventChan},
		"RideCompletedEvent": {PaymentAskEventChan, NotificationEventChan},
		"PaymentAskEvent":    {PaymentEventChan},
		"PaymentEvent":       {NotificationEventChan, TerminationEventChan},
	}

	return PubSub
}
