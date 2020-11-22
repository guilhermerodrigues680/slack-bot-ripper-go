package stopwatch

import (
	"fmt"
	"time"
)

var isPaused = false
var ticker = time.NewTicker(1000 * time.Millisecond)
var startTime time.Time
var currentTime time.Time
var endTime time.Time

func Cronometro() {
	ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			currentTime = t.UTC()
			fmt.Printf("Ticker at: %s\n", t.UTC())
			fmt.Println(currentTime.Sub(startTime))
			fmt.Println(endTime)

			if currentTime.After(endTime) {
				ticker.Stop()
			}
		}
	}
}

func Start() {
	ticker.Reset(1000 * time.Millisecond)
	startTime = time.Now().UTC()
	endTime = startTime.Add(10 * time.Second)
	isPaused = false
	fmt.Println(isPaused)
}

func Pause() {
	ticker.Stop()
	isPaused = true
	fmt.Println(isPaused)
}

func Reset() {
	ticker.Stop()
	fmt.Println(isPaused)
}

func GetStartTime() time.Time {
	return startTime
}

func GetCurrentTime() time.Time {
	return currentTime
}

func GetEndTime() time.Time {
	return endTime
}
