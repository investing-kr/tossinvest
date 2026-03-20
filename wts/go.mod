module github.com/investing-kr/tossinvest/wts

go 1.25.5

replace github.com/investing-kr/tossinvest => ../

require (
	github.com/investing-kr/tossinvest v0.0.0-00010101000000-000000000000
	github.com/xtdlib/chttp v0.0.0-20260320164302-cfb478b082f6
	github.com/xtdlib/rat v0.0.0-20260110121614-1dc662dee071
)

require (
	github.com/benbjohnson/clock v1.3.5 // indirect
	github.com/coder/websocket v1.8.14 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	go.uber.org/ratelimit v0.3.1 // indirect
)
