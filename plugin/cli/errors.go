package cli

import (
	"errors"
	"fmt"

	// external
	impi "github.com/sniperkit/snk.golang.impi/pkg"
)

var (
	errSchemeIsNil = errors.New("Invalid scheme is not defined")
)

type consoleErrorReporter struct{}

func (cer *consoleErrorReporter) Report(err impi.VerificationError) {
	fmt.Printf("%s: %s\n", err.FilePath, err.Error())
}
