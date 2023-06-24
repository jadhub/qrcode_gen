package application

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	qrCode "github.com/skip2/go-qrcode"

	"flamingo.me/flamingo/v3/framework/flamingo"

	"qrcode_gen/src/qrcode/domain"
)

type (
	QrCodeServiceImpl struct {
		logger flamingo.Logger
	}
)

var (
	_ domain.QRCodeServiceInterface = new(QrCodeServiceImpl)
)

func (s *QrCodeServiceImpl) Generate(args domain.CommandArguments) error {
	recoveryLeveL := qrCode.RecoveryLevel(args.ErrorRecovery)

	newQRCode, err := qrCode.New(args.Url, recoveryLeveL)
	if err != nil {
		return err
	}

	newQRCode.DisableBorder = true

	png, err := newQRCode.PNG(args.Width)
	if err != nil {
		return err
	}

	if args.CenterImage != "" {
		png, err = s.superImpose(png, args.Width, args.CenterImage)
		if err != nil {
			return err
		}
	}

	err = s.saveFile(png, args.TargetFilename)
	if err != nil {
		return err
	}

	fmt.Printf("QR Code generated, saved at %s \n", args.TargetFilename)

	return nil
}

func (s *QrCodeServiceImpl) saveFile(imgByte []byte, targetFileName string) error {
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		return err
	}

	out, _ := os.Create("./" + targetFileName)
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	var opts jpeg.Options
	opts.Quality = 80

	err = jpeg.Encode(out, img, &opts)
	if err != nil {
		return err
	}

	return nil
}

func (s *QrCodeServiceImpl) superImpose(imgByte []byte, width int, centerImage string) ([]byte, error) {
	imposeFile, err := os.Open(centerImage)
	if err != nil {
		return nil, err
	}

	imposeImg, _, err := image.Decode(imposeFile)
	if err != nil {
		return nil, err
	}

	qrReader := bytes.NewReader(imgByte)

	qrImg, _, err := image.Decode(qrReader)
	if err != nil {
		return nil, err
	}

	superImposeSize := imposeImg.Bounds().Size()
	qrSize := qrImg.Bounds().Size()

	startX := (qrSize.X - superImposeSize.X) / 2
	startY := (qrSize.Y - superImposeSize.Y) / 2

	startingPointSuperImpose := image.Point{startX, startY}
	imposeRectangle := image.Rectangle{startingPointSuperImpose, startingPointSuperImpose.Add(imposeImg.Bounds().Size())}

	// init target image
	resultRectangle := image.Rectangle{image.Point{0, 0}, qrSize}
	resultImage := image.NewRGBA(resultRectangle)
	draw.Draw(resultImage, qrImg.Bounds(), qrImg, image.Point{0, 0}, draw.Src)
	draw.Draw(resultImage, imposeRectangle, imposeImg, image.Point{0, 0}, draw.Over)

	buf := new(bytes.Buffer)
	err = png.Encode(buf, resultImage)
	if err != nil {
		return nil, fmt.Errorf("encode error: %s", err.Error())
	}

	return buf.Bytes(), err
}
