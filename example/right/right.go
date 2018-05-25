package main

import (
	"github.com/travishegner/i3gobar"
)

func main() {
	f := []func(chan<- []i3gobar.I3Block){
		i3gobar.CurrentTrack,
	}

	i3gobar.Run(f, false, 25)
}
