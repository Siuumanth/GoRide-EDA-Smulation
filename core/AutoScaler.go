package core // AutoScaler

import (
	"context"
	"log"
	"time"
)

/*
Functions of auto scaler:
analyses event bus size every 200 ms using time.tick, if the size is close to the limit, it will increase the number of goroutines
*/

const (
	EVENTBUS_CAPACITY    = 500
	SCALE_UP_THRESHOLD   = 300
	SCALE_DOWN_THRESHOLD = 100
)

var count int = 0

// array to store active q
var activeCancels []context.CancelFunc

/*
Scale up : Add a start worker pools with the fcuntion, and add the context.Cancel() to a list
Scale down : pop back from that list, and cancel the context
*/

func InitAutoScaler(eventBus chan<- any) {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		currLoad := len(eventBus)

		if currLoad > SCALE_UP_THRESHOLD {
			log.Printf("[AutoScaler] Scaling up: current load %d", currLoad)
			scaleUpFunc(eventBus, &activeCancels)
		} else if len(activeCancels) == 1 { // if len 1 skip
			continue
		} else if currLoad < SCALE_DOWN_THRESHOLD {
			log.Printf("[AutoScaler] Scaling down: current load %d", currLoad)
			scaleDownFunc(eventBus, &activeCancels)
		}
	}
}

func scaleUpFunc(eventBus chan<- any, activeCancels *[]context.CancelFunc) {
	// start worker pools
	ctx, cancel := context.WithCancel(context.Background())
	*activeCancels = append(*activeCancels, cancel)

	StartWorkerPools(ctx, eventBus)
	count++
}

func scaleDownFunc(eventBus chan<- any, activeCancels *[]context.CancelFunc) {

	last := len(*activeCancels) - 1
	(*activeCancels)[last]() // cancel it
	*activeCancels = (*activeCancels)[:last]

	count--
}
