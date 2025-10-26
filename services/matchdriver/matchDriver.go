package matchDriverService

import (
	events "RideBooking/events"
	utils "RideBooking/utils"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

/*
This service will provide the nearest driver to the user
Get nearest driver will take input from the event bus and process it
Trip event has langitude and logitude

Calculation of nearest driver:
- Calc sum of userLat - driverlat and long for all and find minimum
*/

var drivers *[]utils.Driver = utils.GenerateDriverData()

func MatchDriver(driverEventQueue <-chan any, eventBus chan<- any) {

	for tripReq := range driverEventQueue {
		// calculate nearest driverW
		switch event := tripReq.(type) {
		case events.TripEvent:
			time.Sleep(3 * time.Second)
			var driverID int = -1
			var minDist float64 = 1000
			for i, driver := range *drivers {
				if driver.Status == "busy" {
					continue
				}
				dist := math.Abs(event.Lat-driver.Lat) + math.Abs(event.Long-driver.Long)

				if dist < minDist {
					driverID = i
					minDist = dist
				}
			}
			if driverID == -1 {
				log.Println("No driver available")
				os.Exit(1)
			}

			// dereference first, then index
			nearestDriver := (*drivers)[driverID]
			(*drivers)[driverID].Status = "busy"

			eta := 2 + rand.Intn(3)

			driverMatchedEvent := events.DriverMatchedEvent{
				DriverName: nearestDriver.Name,
				UserName:   event.UserName,
				Amount:     event.Amount,
				ETA:        float64(eta),
			}

			eventBus <- driverMatchedEvent

		default:
			log.Printf("MatchDriver Received event of type %T", event)

		}
	}

}
