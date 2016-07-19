package i3gobar

import "time"

func DateTime(uc chan<- []I3Block) {
	b := make([]I3Block, 3)
	t := time.Now()
	tc := time.Tick(1 * time.Second)
	for {
		b[0].FullText = t.UTC().Format("2006-01-02")
		b[1].FullText = t.UTC().Format("15:04:05 MST")
		b[2].FullText = t.Local().Format("15:04:05 MST")
		uc <- b
		t = <-tc
	}
}
