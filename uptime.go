package i3gobar

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/host"
)

func Uptime(uc chan<- []I3Block) {
	b := make([]I3Block, 1)
	for {
		ut, err := host.Uptime()
		if err != nil {
			b[0].FullText = err.Error()
		}
		d := time.Duration(int64(ut)) * time.Second
		bt := time.Now().Add(-1 * d)
		b[0].FullText = fmt.Sprintf("Booted: %v", humanize.Time(bt))
		time.Sleep(time.Second)
		uc <- b
	}
}
