module github.com/conejoninja/gopherbadge

go 1.20

require (
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	golang.org/x/image v0.7.0
	tinygo.org/x/drivers v0.25.0
	tinygo.org/x/tinydraw v0.3.0
	tinygo.org/x/tinyfont v0.3.0
	tinygo.org/x/tinyterm v0.2.1-0.20221003142551-ae75982d313f
)

require (
	github.com/go-gl/mathgl v1.0.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
)

replace tinygo.org/x/drivers => /home/conejo/go/src/tinygo.org/x/drivers
