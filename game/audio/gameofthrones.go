package audio

import (
	"tinygo.org/x/drivers/tone"
)

var (
	GameOfThrones = Song{
		notes: []tone.Note{
			tone.G6, tone.C6, tone.DS6, tone.F6, tone.G6, tone.C6, tone.DS6, tone.F6,
			tone.G6, tone.C6, tone.DS6, tone.F6, tone.G6, tone.C6, tone.DS6, tone.F6,
			tone.G6, tone.C6, tone.E6, tone.F6, tone.G6, tone.C6, tone.E6, tone.F6,
			tone.G6, tone.C6, tone.E6, tone.F6, tone.G6, tone.C6, tone.E6, tone.F6,
			tone.G6, tone.C6,

			tone.DS6, tone.F6, tone.G6, tone.C6, tone.DS6, tone.F6,
			tone.D6,
			tone.F6, tone.AS5,
			tone.DS6, tone.D6, tone.F6, tone.AS5,
			tone.DS6, tone.D6, tone.C6,

			tone.G6, tone.C6,

			tone.DS6, tone.F6, tone.G6, tone.C6, tone.DS6, tone.F6,
			tone.D6,
			tone.F6, tone.AS5,
			tone.DS6, tone.D6, tone.F6, tone.AS5,
			tone.DS6, tone.D6, tone.C6,
			tone.G6, tone.C6,
			tone.DS6, tone.F6, tone.G6, tone.C6, tone.DS6, tone.F6,

			tone.D6,
			tone.F6, tone.AS5,
			tone.D6, tone.DS6, tone.D6, tone.AS5,
			tone.C6,
			tone.C7,
			tone.AS6,
			tone.C6,
			tone.G6,
			tone.DS6,
			tone.DS6, tone.F6,
			tone.G6,

			tone.C7,
			tone.AS6,
			tone.C6,
			tone.G6,
			tone.DS6,
			tone.DS6, tone.D6,
			tone.C7, tone.G6, tone.GS6, tone.AS6, tone.C7, tone.G6, tone.GS6, tone.AS6,
			tone.C7, tone.G6, tone.GS6, tone.AS6, tone.C7, tone.G6, tone.GS6, tone.AS6,

			tone.Note(0), tone.GS7, tone.AS7, tone.C8, tone.G7, tone.GS7, tone.AS7,
			tone.C8, tone.G7, tone.GS7, tone.AS7, tone.C8, tone.G7, tone.GS7, tone.AS7,
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
