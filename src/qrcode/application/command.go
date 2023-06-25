package application

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"

	"qrcode_gen/src/qrcode/domain"
)

func Generate(qrService domain.QRCodeServiceInterface) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate a qr code",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("################################################")
			commandArguments, err := ProcessInput(args)
			if err != nil {
				fmt.Printf("Encountered an error while processing arguments: %s \n", err.Error())
				fmt.Println("################################################")

				return
			}

			err = qrService.Generate(*commandArguments)
			if err != nil {
				fmt.Printf("Encountered an error while generating the qr code: %s \n", err.Error())
				fmt.Println("################################################")

				return
			}

			fmt.Println("################################################")
		},
	}

	return cmd
}

func ProcessInput(args []string) (*domain.CommandArguments, error) {
	resultArgs := &domain.CommandArguments{}

	// 0 - URI
	_, err := url.ParseRequestURI(args[0])
	if err != nil {
		return nil, fmt.Errorf("validation error: invalid url: %s", err.Error())
	}
	resultArgs.Url = args[0]

	// 1 - errorRecovery (0 - 7%, 1 - 15%, 2 - 20%, 3 - 30%)
	errorRecovery, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return nil, errors.New("validation error: error recovery is not a number")
	}
	if errorRecovery > 3 {
		return nil, errors.New("validation error: error recovery out of bounds, maixmum is 3")
	}
	resultArgs.ErrorRecovery = int(errorRecovery)

	// 2 - width
	width, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return nil, errors.New("validation error: qr code width not a number")
	}
	resultArgs.Width = int(width)

	targetExtension := filepath.Ext(args[3])
	if targetExtension != ".png" {
		return nil, fmt.Errorf("validation error: target image is not a png: %s", targetExtension)
	}
	resultArgs.TargetFilename = args[3]

	if len(args) >= 5 {
		centerExtension := filepath.Ext(args[4])
		if centerExtension != ".png" {
			return nil, fmt.Errorf("validation error: center image is not a png: %s", centerExtension)
		}

		resultArgs.CenterImage = args[4]
	}

	return resultArgs, nil
}
