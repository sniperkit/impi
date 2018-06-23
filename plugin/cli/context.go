package cli

import (
	"fmt"
	"os"

	// internal - core
	conf "github.com/sniperkit/snk.golang.impi/pkg/config"
	ver "github.com/sniperkit/snk.golang.impi/pkg/version"
)

var (
	configDirectoryPath string = "."
	cwd, _                     = os.Getwd()
)

func printContextInfo() {
	fmt.Println("ProgramName=", conf.ProgramName)
	fmt.Println("ProgramVersion=", ver.Version)
	fmt.Println("CurrentWorkDirectory=", cwd)
	fmt.Println("ConfigDirectoryPath=", configDirectoryPath)
}
