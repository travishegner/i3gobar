package left

import (
	"github.com/travishegner/i3gobar"
)

func main() {
	f := []func(chan<- []i3gobar.I3Block){
		i3gobar.Date,
		i3gobar.TimeUTC,
		i3gobar.TimeLocal,
	}

	i3gobar.Run(f, false, 25)
}
