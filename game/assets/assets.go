package assets

import (
	_ "embed"
)

//go:embed bug_1.qoi
var Bug1 string

//go:embed bug_2.qoi
var Bug2 string

//go:embed running1.qoi
var PlayerSprite1 string

//go:embed running2.qoi
var PlayerSprite2 string
