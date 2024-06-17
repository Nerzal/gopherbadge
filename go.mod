module github.com/conejoninja/gopherbadge

go 1.21

require (
	github.com/acifani/vita v1.2.0
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	golang.org/x/image v0.15.0
	tinygo.org/x/drivers v0.27.0
	tinygo.org/x/tinydraw v0.4.0
	tinygo.org/x/tinyfont v0.4.0
	tinygo.org/x/tinyterm v0.3.0
)

require github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect

replace tinygo.org/x/drivers => github.com/conejoninja/drivers v0.0.0-20240124083359-dcfbdc0db7ae
