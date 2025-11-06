package core

import (
	driverService "RideBooking/services/drivers"
	notiService "RideBooking/services/notification"
	payService "RideBooking/services/pay"
	paymentAskService "RideBooking/services/paymentask"
	rides "RideBooking/services/ridesinprogress"
	terminationService "RideBooking/services/termination"
	tripService "RideBooking/services/trip"
	tripCompletedService "RideBooking/services/tripCompleted"
	"context"
)

// Constants for Number of goroutines for each service
const (
	NUM_TRIP_SERVICES           = 3
	NUM_DRIVER_SERVICES         = 3
	NUM_PAYMENT_SERVICES        = 3
	NUM_PAYMENTASK_SERVICES     = 3
	NUM_NOTIFICATION_SERVICES   = 8
	NUM_RIDES_SERVICES          = 15
	NUM_TRIP_COMPLETED_SERVICES = 5
	NUM_TERMINATION_SERVICES    = 1
)

// call for all services
func StartWorkerPools(ctx context.Context, eventBus chan<- any) {
	StartServiceWorkers(TripRequestedEventChan, eventBus, NUM_TRIP_SERVICES, tripService.TripService, ctx)
	StartServiceWorkers(DriverEventChan, eventBus, NUM_DRIVER_SERVICES, driverService.DriverService, ctx)
	StartServiceWorkers(PaymentAskEventChan, eventBus, NUM_PAYMENTASK_SERVICES, paymentAskService.PaymentAskService, ctx)
	StartServiceWorkers(PaymentEventChan, eventBus, NUM_PAYMENT_SERVICES, payService.PaymentService, ctx)
	StartServiceWorkers(NotificationEventChan, eventBus, NUM_NOTIFICATION_SERVICES, notiService.NotificationService, ctx)
	StartServiceWorkers(RideEventChan, eventBus, NUM_RIDES_SERVICES, rides.RideService, ctx)
	StartServiceWorkers(TripCompletedEventChan, eventBus, NUM_TRIP_COMPLETED_SERVICES, tripCompletedService.TripCompletedService, ctx)
	StartServiceWorkers(TerminationEventChan, eventBus, 1, terminationService.TerminationService, ctx) // usually 1
}

// Generic function to start multiple goroutines for a service
func StartServiceWorkers(eventChan <-chan any, eventBus chan<- any, numWorkers int, service func(<-chan any, chan<- any, context.Context), ctx context.Context) {
	for i := 0; i < numWorkers; i++ {
		go service(eventChan, eventBus, ctx)
	}
}
