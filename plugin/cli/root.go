package cli

import (
	"fmt"
	"os"
	"runtime"

	// external
	"github.com/sniperkit/snk.golang.cobra"

	// internal - core
	conf "github.com/sniperkit/snk.golang.impi/pkg/config"
	out "github.com/sniperkit/snk.golang.impi/pkg/output"
)

// configuration is...
var configuration *conf.Config

var (
	programName string
	numCPUs     int
)

// options defines...
var options struct {
	mapping   map[string]string
	numCPUs   int
	match     string
	scheme    string
	output    string
	dirConf   string
	writeConf bool
	dryMode   bool
	version   bool
	debug     bool
}

// RootCmd is the root command for limo
var RootCmd = &cobra.Command{
	Use:   conf.ProgramName,
	Short: "A CLI for formating import packages.",
	Long:  fmt.Sprintf(`%s allows you to formating import packages and group them by section in your code.`, conf.ProgramName),
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func setRuntimeLimit(max int) {
	switch {
	case numCPUs > max:
		runtime.GOMAXPROCS(max)
	case max == -1:
		return
	default:
		runtime.GOMAXPROCS(numCPUs)
	}

}

// 	newPackage := flag.Arg(0)

func init() {

	numCPUs := runtime.NumCPU()

	flags := RootCmd.PersistentFlags()

	flags.StringVarP(&options.scheme, "scheme", "s", "", "verification scheme to enforce. one of stdLocalThirdParty/stdThirdPartyLocal")
	flags.StringVarP(&options.output, "output", "o", "color", "output type")
	flags.StringVarP(&options.dirConf, "conf-dir", "c", ".goimpi.yaml", "write config to prefix dir, default XGDB Base directory.")

	flags.IntVarP(&options.numCPUs, "max-procs", "m", numCPUs, "max parallel processes. (default: number of cores of the local machine).")

	flags.BoolVarP(&options.writeConf, "write-conf", "w", true, "write config file to xgd dir")

	flags.BoolVarP(&options.dryMode, "dry-mode", "n", false, "don't make any changes; perform checks only")
	flags.BoolVarP(&options.version, "version", "v", false, "display version")
	flags.BoolVarP(&options.debug, "debug", "d", false, "debug mode")
}

func getConfiguration() (*conf.Config, error) {
	if configuration == nil {
		var err error
		if configuration, err = conf.ReadConfig(); err != nil {
			return nil, err
		}
	}
	return configuration, nil
}

func getOutput() out.Output {
	o := out.ForName(options.output)
	oc, err := getConfiguration()
	if err == nil {
		o.Configure(oc.GetOutput(options.output))
	}
	return o
}

func fatalOnError(err error) {
	if err != nil {
		if options.debug {
			log.WithFields(logFields{
				"prog-name": conf.ProgramName,
			}).Fatalln("error", err.Error())
		} else {
			getOutput().Fatal(err.Error())
		}
	}
}
