package main

import (
	driverService "RideBooking/services/matchdriver"
	notiService "RideBooking/services/notification"
	payService "RideBooking/services/pay"
	paymentAskService "RideBooking/services/paymentask"
	rides "RideBooking/services/ridesinprogress"
	tripService "RideBooking/services/trip"
)

func StartWorkerPools(eventBus chan<- any) {
	for i := 0; i < 5; i++ {
		go driverService.MatchDriver(DriverMatchedEventChan, eventBus)
		go tripService.TripService(TripEventChan, eventBus)
		go paymentAskService.PaymentAskService(PaymentAskEventChan, eventBus)
		go payService.PaymentService(PaymentEventChan, eventBus)
		go notiService.NotificationService(NotificationEventChan, eventBus)
		go rides.RideService(RideEventChan, eventBus)
	}
}
