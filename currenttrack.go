package i3gobar

import (
	"os/exec"
	"time"
)

func CurrentTrack(uc chan<- []I3Block) {
	b := make([]I3Block, 1)
	b[0].FullText = ""
	for {
		artist, err := exec.Command("/usr/bin/playerctl", "metadata", "artist").Output()
		if err != nil {
			b[0].FullText = err.Error()
		}

		title, err := exec.Command("/usr/bin/playerctl", "metadata", "title").Output()
		if err != nil {
			b[0].FullText = err.Error()
		}

		if err == nil {
			b[0].FullText = string(artist) + " - " + string(title)
		}

		uc <- b
		time.Sleep(500 * time.Millisecond)
	}
}
