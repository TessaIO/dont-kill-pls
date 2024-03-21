package main

import (
	monitor "dont-kill-pls/pkg"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
  c := make(chan bool)
  m := monitor.NewMonitor(monitor.MonitorInput{
    Frequency: 2,
    MaxAllowedMemory: 10,
    Close: c,
  })
  m.Run()

  signalChan := make(chan os.Signal, 1)
  signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
  go func() {
    sig := <-signalChan
    switch sig {
    case os.Interrupt:
      fmt.Println("OS INTERRUPT")
      m.Stop()
    case syscall.SIGTERM:
      fmt.Println("OS SIGTERM")
      m.Stop()
    }
  }()

  size := 1024 * 1024 * 8
  for j := 0; j < 100; j += 1 {
      a := make([]int, size)
      for i := 0; i < size; i += 1 {
          a[i] = i
      }
      a = nil
  }

  <- c
}
