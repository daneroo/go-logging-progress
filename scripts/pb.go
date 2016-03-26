package main

import (
	"math/rand"
	"time"

	"github.com/cheggaaa/pb"
)

func main() {
	count := 100000
	bar := pb.New(count)

	bar.Prefix("Prefix ")

	// refresh info every second (default 200ms)
	bar.SetRefreshRate(200 * time.Millisecond)

	// show percents (by default already true)
	bar.ShowPercent = true

	// show bar (by default already true)
	bar.ShowBar = true

	// no counters
	// bar.ShowCounters = false

	// show "time left"
	bar.ShowTimeLeft = true

	// show average speed
	bar.ShowSpeed = true

	// sets the width of the progress bar
	// bar.SetWidth(80)

	// sets the width of the progress bar, but if terminal size smaller will be ignored
	// bar.SetMaxWidth(80)

	// convert output to readable format (like KB, MB)
	// bar.SetUnits(pb.U_BYTES)

	// and start
	bar.Start()
	// and start
	bar.Start()
	for i := 0; i < count; i += 1 + rand.Intn(10) {
		// bar.Increment()
		bar.Set(i)
		time.Sleep(1 * time.Millisecond)
	}
	bar.Set(count)
	bar.FinishPrint("The End!")
}
