package core

import (
	driverService "RideBooking/services/drivers"
	notiService "RideBooking/services/notification"
	payService "RideBooking/services/pay"
	paymentAskService "RideBooking/services/paymentask"
	rides "RideBooking/services/ridesinprogress"
	tripService "RideBooking/services/trip"
	tripCompletedService "RideBooking/services/tripCompleted"
)

func StartWorkerPools(eventBus chan<- any) {
	for i := 0; i < 5; i++ {
		go driverService.DriverService(DriverEventChan, eventBus)
		go tripService.TripService(TripRequestedEventChan, eventBus)
		go paymentAskService.PaymentAskService(PaymentAskEventChan, eventBus)
		go payService.PaymentService(PaymentEventChan, eventBus)
		go notiService.NotificationService(NotificationEventChan, eventBus)
		go rides.RideService(RideEventChan, eventBus)
		go tripCompletedService.TripCompletedService(TripCompletedEventChan, eventBus)
	}
}
