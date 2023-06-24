package application

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"

	"qrcode_gen/src/qrcode/domain"
)

func Generate(qrService domain.QRCodeServiceInterface) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate a qr code",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("################################################")
			commandArguments, err := processInput(args)
			if err != nil {
				fmt.Println("Encountered an error while processing arguments ", err.Error())

				return
			}

			err = qrService.Generate(*commandArguments)
			if err != nil {
				fmt.Println("Encountered an error while generating the qr code", err.Error())

				return
			}

			fmt.Println("################################################")
		},
	}

	return cmd
}

func processInput(args []string) (*domain.CommandArguments, error) {
	if len(args) < 5 {
		return nil, errors.New("argument missing")
	}

	errorRecovery, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, err
	}

	width, err := strconv.Atoi(args[2])
	if err != nil {
		return nil, err
	}

	resultArgs := &domain.CommandArguments{
		Url:            args[0],
		ErrorRecovery:  errorRecovery,
		Width:          width,
		TargetFilename: args[3],
		CenterImage:    args[5],
	}

	if args[4] == "true" {
		resultArgs.Border = true
	}

	return resultArgs, nil
}
