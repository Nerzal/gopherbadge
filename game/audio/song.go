package audio

import (
	"time"
	"tinygo.org/x/drivers/tone"
)

type Song struct {
	durations []int
	notes     []tone.Note
}

func (s *Song) Play(speaker *tone.Speaker) {
	for idx := range s.notes {
		speaker.SetNote(s.notes[idx])
		time.Sleep(time.Duration(1000/s.durations[idx]) * time.Millisecond)
	}
	speaker.Stop()
}
