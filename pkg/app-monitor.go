package monitor

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
)

// Monitor defines the monitor that is executed to check the memory usage of the internal application
type Monitor struct {
  // MaxAllowedMemory defines the maximum allowed memory for an application to use
  // If the application memory usage is bigger than MaxAllowedMemory a SIGTERM signal would be sent
  // to the application. This should be defined in MiB
  maxAllowedMemory float64
  // Frequency defines the frequency by which we check the memory usage
  frequency int
  // Close is the channel that is used to wait for the monitor 
  c chan(bool)
}

type MonitorInput struct {
  MaxAllowedMemory float64
  Frequency int
  Close chan(bool)
}

// NewMonitor instantiates a new Monitor with the predefined configuration
func NewMonitor(input MonitorInput) *Monitor {
  return &Monitor{
    maxAllowedMemory: input.MaxAllowedMemory,
    frequency: input.Frequency,
    c: input.Close,
  }
}

// Run starts the monitor which run in the background as a goroutine
func (m *Monitor) Run() {
  fmt.Println("Running monitor...")
  ticker := time.NewTicker(time.Duration(m.frequency)*time.Second)

  go func() {
    for {
      select {
        case <- ticker.C:
          stats := &runtime.MemStats{}
          runtime.ReadMemStats(stats)
          currentMemoryUsage := float64(stats.HeapAlloc) / 1024 / 1024
          if currentMemoryUsage > m.maxAllowedMemory {
            fmt.Printf("Should stop application because it exceeds the memory limit, Current Usage is %v, Max Memory is %v\n", currentMemoryUsage, m.maxAllowedMemory)
            err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
            if err != nil {
              fmt.Println(err)
            }
        }
      default:
    }
    }
  }()
}

func (m *Monitor) Stop() {
  fmt.Println("Stopping monitor...")
  close(m.c)
}

