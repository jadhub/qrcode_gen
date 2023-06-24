package domain

type (
	CommandArguments struct {
		Url            string
		ErrorRecovery  int
		Width          int
		TargetFilename string
		CenterImage    string
	}

	QRCodeServiceInterface interface {
		Generate(CommandArguments) error
	}
)
