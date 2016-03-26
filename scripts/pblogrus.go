package main

import (
	"fmt"
	"strings"
	"time"

	// log "github.com/Sirupsen/logrus"
	log "github.com/Sirupsen/logrus"
	"github.com/cheggaaa/pb"
	"github.com/daneroo/go-logging-progress"
)

// TODO
// type Hook interface {
//     Levels() []Level
//     Fire(*Entry) error
// }

// HACK
// -create a new bar just to get screen width
// -clear line before every ouput
type MyProgressPreventingFormatter struct{}

var bbbar = pb.New(0)

func (f *MyProgressPreventingFormatter) Format(entry *log.Entry) ([]byte, error) {
	std := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05.9",
	}

	clear := "\r" + strings.Repeat(" ", bbbar.GetWidth()) + "\r"
	fmt.Print(clear)

	return std.Format(entry)
}

func init() {
	hook := new(coco.PBHook)
	log.AddHook(hook)

	log.SetFormatter(new(MyProgressPreventingFormatter))
}

func main() {
	// offset:= 0
	count := 1000
	bar := pb.New(count)

	bar.Prefix("Prefix ")
	bar.ShowPercent = true
	bar.ShowBar = true
	bar.ShowCounters = true
	bar.ShowTimeLeft = true
	bar.ShowSpeed = true

	go func() {
		time.Sleep(2 * time.Second)
		for i := 1; i <= 5; i++ {
			time.Sleep(1 * time.Second)
			log.Printf("Hello, here is a delayed message: %d/5", i)
		}
	}()

	// and start
	bar.Start()
	for i := 0; i < count; i++ {
		// double the speed in the begining
		if i > count/4 && i < 3*count/4 {
			i++
		}
		bar.Set(i)
		// bar.Increment()
		time.Sleep(50 * time.Millisecond)
	}
	// TODO(daneroo): on finish not show percent if 100%
	bar.ShowPercent = false
	bar.Set(count)
	bar.FinishPrint("The End!")
}
