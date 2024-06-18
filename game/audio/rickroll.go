package audio

import (
	"tinygo.org/x/drivers/tone"
)

var (
	rickroll = Song{
		notes: []tone.Note{
			tone.G4, tone.C4, tone.DS4, tone.F4, tone.G4, tone.C4, tone.DS4, tone.F4,
			tone.G4, tone.C4, tone.DS4, tone.F4, tone.G4, tone.C4, tone.DS4, tone.F4,
			tone.G4, tone.C4, tone.E4, tone.F4, tone.G4, tone.C4, tone.E4, tone.F4,
			tone.G4, tone.C4, tone.E4, tone.F4, tone.G4, tone.C4, tone.E4, tone.F4,
			tone.G4, tone.C4,

			tone.DS4, tone.F4, tone.G4, tone.C4, tone.DS4, tone.F4,
			tone.D4,
			tone.F4, tone.AS3,
			tone.DS4, tone.D4, tone.F4, tone.AS3,
			tone.DS4, tone.D4, tone.C4,

			tone.G4, tone.C4,

			tone.DS4, tone.F4, tone.G4, tone.C4, tone.DS4, tone.F4,
			tone.D4,
			tone.F4, tone.AS3,
			tone.DS4, tone.D4, tone.F4, tone.AS3,
			tone.DS4, tone.D4, tone.C4,
			tone.G4, tone.C4,
			tone.DS4, tone.F4, tone.G4, tone.C4, tone.DS4, tone.F4,

			tone.D4,
			tone.F4, tone.AS3,
			tone.D4, tone.DS4, tone.D4, tone.AS3,
			tone.C4,
			tone.C5,
			tone.AS4,
			tone.C4,
			tone.G4,
			tone.DS4,
			tone.DS4, tone.F4,
			tone.G4,

			tone.C5,
			tone.AS4,
			tone.C4,
			tone.G4,
			tone.DS4,
			tone.DS4, tone.D4,
			tone.C5, tone.G4, tone.GS4, tone.AS4, tone.C5, tone.G4, tone.GS4, tone.AS4,
			tone.C5, tone.G4, tone.GS4, tone.AS4, tone.C5, tone.G4, tone.GS4, tone.AS4,

			tone.Note(0), tone.GS5, tone.AS5, tone.C6, tone.G5, tone.GS5, tone.AS5,
			tone.C6, tone.G5, tone.GS5, tone.AS5, tone.C6, tone.G5, tone.GS5, tone.AS5,
		},
		durations: []int{
			8, 8, 16, 16, 8, 8, 16, 16,
			8, 8, 16, 16, 8, 8, 16, 16,
			8, 8, 16, 16, 8, 8, 16, 16,
			8, 8, 16, 16, 8, 8, 16, 16,
			4, 4,

			16, 16, 4, 4, 16, 16,
			1,
			4, 4,
			16, 16, 4, 4,
			16, 16, 1,

			4, 4,

			16, 16, 4, 4, 16, 16,
			1,
			4, 4,
			16, 16, 4, 4,
			16, 16, 1,
			4, 4,
			16, 16, 4, 4, 16, 16,

			2,
			4, 4,
			8, 8, 8, 8,
			1,
			2,
			2,
			2,
			2,
			2,
			4, 4,
			1,

			2,
			2,
			2,
			2,
			2,
			4, 4,
			8, 8, 16, 16, 8, 8, 16, 16,
			8, 8, 16, 16, 8, 8, 16, 16,

			4, 16, 16, 8, 8, 16, 16,
			8, 16, 16, 16, 8, 8, 16, 16,
		},
	}
)
