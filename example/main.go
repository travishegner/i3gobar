package main

import (
	"github.com/clinta/i3gobar"
)

func main() {
	f := []func(chan<- []i3gobar.I3Block){
		i3gobar.LoadAvg,
		i3gobar.CPU,
		i3gobar.MemFree,
		i3gobar.SwapUsed,
		i3gobar.DateTime,
	}

	i3gobar.Run(f, false, 25)
}
