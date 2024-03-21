package main

import monitor "dont-kill-pls/pkg"

func main() {
  // TODO: we should handle graceful shutdown here and close the channel, but no time for it now
  ch := make(chan bool)
	m := monitor.NewMonitor(monitor.MonitorInput{
		MaxAllowedMemory: 0.05,
		Frequency:        2,
    Close: ch,
	})
	m.Run()
  <- ch
}
