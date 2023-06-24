package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"

	"qrcode_gen/src/qrcode"
)

func main() {
	flamingo.App(
		[]dingo.Module{
			new(qrcode.Module),
		},
	)
}
