package main

import (
	"github.com/travishegner/i3gobar"
)

func main() {
	f := []func(chan<- []i3gobar.I3Block){
		i3gobar.Uptime,
		i3gobar.LoadAvg,
		i3gobar.CPUGraph,
		i3gobar.CPUTemp,
		i3gobar.MemFree,
		i3gobar.SwapUsed,
	}

	i3gobar.Run(f, false, 25)
}
