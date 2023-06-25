package application

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"

	"qrcode_gen/src/qrcode/domain"
)

type (
	CommandTestSuite struct {
		suite.Suite
	}

	argFixture struct {
		Title       string
		Args        []string
		Expectation *domain.CommandArguments
		wantErr     bool
	}
)

func TestCommandTestSuite(t *testing.T) {
	suite.Run(t, &CommandTestSuite{})
}

func (s *CommandTestSuite) TestProcessInput() {
	argFixtures := []argFixture{
		{
			Title: "success",
			Args:  []string{"http://example.com", "2", "500", "target_qr_code_image.png", "superimposed_image.png"},
			Expectation: &domain.CommandArguments{
				Url:            "http://example.com",
				ErrorRecovery:  2,
				Width:          500,
				TargetFilename: "target_qr_code_image.png",
				CenterImage:    "superimposed_image.png",
			},
			wantErr: false,
		},
		{
			Title:       "invalid url",
			Args:        []string{"blah", "2", "500", "target_qr_code_image.png", "superimposed_image.png"},
			Expectation: nil,
			wantErr:     true,
		},
		{
			Title:       "error correction out of bounds",
			Args:        []string{"http://example.com", "4", "500", "target_qr_code_image.png", "superimposed_image.png"},
			Expectation: nil,
			wantErr:     true,
		},
		{
			Title:       "error correction not a number",
			Args:        []string{"http://example.com", "AAA", "500", "target_qr_code_image.png", "superimposed_image.png"},
			Expectation: nil,
			wantErr:     true,
		},
		{
			Title:       "qr width not a number",
			Args:        []string{"http://example.com", "3", "GFD", "target_qr_code_image.png", "superimposed_image.png"},
			Expectation: nil,
			wantErr:     true,
		},
		{
			Title:       "target qr code extension not png",
			Args:        []string{"http://example.com", "1", "200", "target_qr_code_image.bmp", "superimposed_image.png"},
			Expectation: nil,
			wantErr:     true,
		},
		{
			Title:       "image to superimpose is not a png",
			Args:        []string{"http://example.com", "1", "200", "target_qr_code_image.png", "superimposed_image.jpeg"},
			Expectation: nil,
			wantErr:     true,
		},
		{
			Title: "call without additional superimposed image",
			Args:  []string{"http://example.com", "1", "200", "target_qr_code_image.png"},
			Expectation: &domain.CommandArguments{
				Url:            "http://example.com",
				ErrorRecovery:  1,
				Width:          200,
				TargetFilename: "target_qr_code_image.png",
				CenterImage:    "",
			},
			wantErr: false,
		},
	}

	for _, fixture := range argFixtures {
		actual, err := ProcessInput(fixture.Args)
		if fixture.wantErr {
			s.NotNil(err)
		}
		if fixture.Expectation != nil {
			s.True(reflect.DeepEqual(*actual, *fixture.Expectation))
		}
	}
}
