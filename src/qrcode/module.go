package qrcode

import (
	"flamingo.me/dingo"
	"github.com/spf13/cobra"

	"qrcode_gen/src/qrcode/application"
	"qrcode_gen/src/qrcode/domain"
)

type (
	Module struct{}
)

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(cobra.Command)).ToProvider(application.Generate)
	injector.Bind(new(domain.QRCodeServiceInterface)).To(new(application.QrCodeServiceImpl))
}
