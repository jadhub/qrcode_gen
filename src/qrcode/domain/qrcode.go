package domain

type (
	CommandArguments struct {
		Url            string
		ErrorRecovery  int
		Width          int
		TargetFilename string
		Border         bool
		CenterImage    string
	}

	QRCodeServiceInterface interface {
		Generate(CommandArguments) error
	}
)
