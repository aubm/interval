package interval

import (
	"time"
)

// Start invokes an action every interval, it returns a stop function
// that must be invoked to stop the interval
func Start(action func(), interval time.Duration) func() {
	t := time.NewTicker(interval)
	q := make(chan struct{})
	go func() {
		for {
			select {
			case <-t.C:
				action()
			case <-q:
				t.Stop()
				return
			}
		}
	}()
	return func() {
		close(q)
	}
}
