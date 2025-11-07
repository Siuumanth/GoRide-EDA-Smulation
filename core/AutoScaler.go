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

var count int = 0
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
		//	fmt.Printf("Current load is %d \n", currLoad)

		// iF zero for a long time, cancel the final goroutine
		if idleTicks > IDLE_SHUTDOWN_TICKS {
			scaleDownFunc(&activeCancels)
			return // exit
		}

		if currLoad > SCALE_UP_THRESHOLD {
			log.Printf("[AutoScaler] Scaling up: current load %d, Current Worker Pools Count %d", currLoad, count)
			scaleUpFunc(eventBus, &activeCancels)
			idleTicks = 0
		} else if len(activeCancels) == 1 { // if len 1 skip
			idleTicks++
			continue
		} else if currLoad < SCALE_DOWN_THRESHOLD {
			log.Printf("[AutoScaler] Scaling down: current load %d, Current Worker Pools Count %d", currLoad, count)
			idleTicks = 0
			scaleDownFunc(&activeCancels)
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
	count++
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

	count--
}
