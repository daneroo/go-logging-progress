package main

import (
    "fmt"
    "strings"
    "time"

    log "github.com/Sirupsen/logrus"
    "gopkg.in/cheggaaa/pb.v1"
)

// const llog = log.New(os.Stderr, "", log.LstdFlags)
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
    // do something here to set environment depending on an environment variable
    // or command-line flag
    if true {
        log.SetFormatter(new(MyProgressPreventingFormatter))
        // log.SetFormatter(&log.JSONFormatter{})
    } else {
        // The TextFormatter is default, you don't actually have to do this.
        log.SetFormatter(&log.TextFormatter{})
    }
}

func main() {
    // offset:= 0
    count := 100
    bar := pb.New(100)

    // bar.Prefix("First ")
    // show percents (by default already true)
    bar.ShowPercent = true

    // show bar (by default already true)
    bar.ShowBar = true

    bar.ShowCounters = true

    bar.ShowTimeLeft = true

    go func() {
        time.Sleep(3 * time.Second)
        log.Printf("Hello, here is a delayed message")
    }()

    // and start
    bar.Start()
    for i := 0; i < count; i++ {
        if i%17 == 5 {
            log.Printf("Hello, here is a message")
        }
        if i > 10 {
            bar.Total = 100
        }
        if i > 50 {
            // bar.Total = 200
        }
        bar.Set(i)
        // bar.Increment()
        time.Sleep(100 * time.Millisecond)
    }
    bar.FinishPrint("The End!")
}
