package core // AutoScaler

import (
	"context"
	"log"
	"sync"
	"time"
)

/*
Functions of auto scaler:
analyses event bus size every 200 ms using time.tick, if the size is close to the limit, it will increase the number of goroutines
*/

const (
	EVENTBUS_CAPACITY    = 1000
	SCALE_UP_THRESHOLD   = 650
	SCALE_DOWN_THRESHOLD = 100
	IDLE_SHUTDOWN_TICKS  = 200
)

/*
Scale up : Add a start worker pools with the fcuntion, and add the context.Cancel() to a list
Scale down : pop back from that list, and cancel the context
*/

var scalerMutex sync.Mutex // <-- NEW: Mutex for updated activeCancels
// array to store active q
var activeCancels []context.CancelFunc

func InitAutoScaler(eventBus chan<- any, firstctx context.Context, firstCancel context.CancelFunc) {
	// start worker pools when length is 0
	activeCancels = append(activeCancels, firstCancel)
	StartWorkerPools(firstctx, eventBus)

	// then tick every time to check for increasing
	var idleTicks int = 0
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		currLoad := len(eventBus)
		wplength := len(activeCancels)
		// iF zero for a long time, cancel the final goroutine
		if idleTicks > IDLE_SHUTDOWN_TICKS {
			scaleDownFunc(&activeCancels)
			return // exit
		}

		if currLoad > SCALE_UP_THRESHOLD {
			log.Printf("[AutoScaler] Scaling up: current load %d, Current Worker Pools Count %d", currLoad, wplength)
			scaleUpFunc(eventBus, &activeCancels)
			idleTicks = 0
		} else if currLoad < SCALE_DOWN_THRESHOLD && wplength > 1 {
			log.Printf("[AutoScaler] Scaling down: current load %d, Current Worker Pools Count %d", currLoad, wplength)
			idleTicks = 0
			scaleDownFunc(&activeCancels)
		} else if len(activeCancels) == 1 {

			if currLoad == 0 {
				idleTicks++ // increment if load is 0
			} else {
				idleTicks = 0 // reset if load present
			}
			if idleTicks > IDLE_SHUTDOWN_TICKS {
				log.Printf("[AutoScaler] Idle shutdown: load is 0 for %d ticks. Exiting.", idleTicks)
				scaleDownFunc(&activeCancels)
				return
			}
		}
	}
}

func scaleUpFunc(eventBus chan<- any, activeCancels *[]context.CancelFunc) {
	scalerMutex.Lock()
	defer scalerMutex.Unlock()

	// start worker pools
	ctx, cancel := context.WithCancel(context.Background())
	*activeCancels = append(*activeCancels, cancel)
	StartWorkerPools(ctx, eventBus)
}

func scaleDownFunc(activeCancels *[]context.CancelFunc) {
	scalerMutex.Lock()
	defer scalerMutex.Unlock()

	last := len(*activeCancels) - 1

	// Safety check: Don't panic if called when empty
	if last < 0 {
		return
	}

	(*activeCancels)[last]() // cancel it
	*activeCancels = (*activeCancels)[:last]

}
