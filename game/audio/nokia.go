package audio

import "tinygo.org/x/drivers/tone"

var (
	Nokia = Song{
		notes: []tone.Note{
			tone.E7, tone.D7, tone.FS6, tone.GS6,
			tone.CS7, tone.B6, tone.D6, tone.E6,
			tone.B6, tone.A6, tone.CS6, tone.E6,
			tone.A6,
		},
		durations: []int{
			8, 8, 4, 4,
			8, 8, 4, 4,
			8, 8, 4, 4,
			2,
		},
	}
)
