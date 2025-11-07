package main

import (
	events "RideBooking/events"
	"fmt"
	"math/rand"
	"time"
)

// Generate Users function

func SimulateRandomUsers(eventBus chan<- any, numUsers int) {
	destinations := []string{"BLR", "HK", "NY", "DC", "AMS", "LA", "CH", "LS", "AM", "SF", "MUM"}

	generated := make(map[string]bool) // ensure unique usernames
	count := 0
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			fmt.Printf("Generated %d users\n", count)
		}
	}()

	for i := 0; i < numUsers; i++ {
		count++
		var userName string
		for {
			userName = fmt.Sprintf("user-%d", rand.Intn(1000000))
			if !generated[userName] {
				generated[userName] = true
				break
			}
		}

		lat := rand.Float64() * 100
		long := rand.Float64() * 100
		destination := destinations[rand.Intn(len(destinations))]

		eventBus <- events.UserEvent{
			UserName:    userName,
			Lat:         lat,
			Long:        long,
			Destination: destination,
		}

		// sleep random duration between users
		time.Sleep(time.Duration(rand.Intn(2500)+50) * time.Microsecond)
	}

	fmt.Println("Users generated.")
}
