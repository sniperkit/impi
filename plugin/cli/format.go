package cli

import (
	"fmt"
	"strings"

	// external
	"github.com/k0kubun/pp"
	"github.com/sniperkit/snk.golang.cobra"

	// internal - core
	impi "github.com/sniperkit/snk.golang.impi/pkg"
	// conf "github.com/sniperkit/snk.golang.impi/pkg/config"
)

var (
	localPrefix = ""
	scheme      = ""
	skipPaths   = ""
	rootPaths   = ""
)

// FormatCmd trigger the current workidr pkg imported and replace the matched pattern
var FormatCmd = &cobra.Command{
	Use:     "fmt",
	Aliases: []string{"fmt", "format", "f"},
	Short:   "Format and group package imports by type",
	// Example: fmt.Sprintf(" %s ^github.com/google/go-github github.com/sniperkit/go-github/pkg ", conf.ProgramName),
	Run: func(cmd *cobra.Command, args []string) {
		output := getOutput()

		if len(args) == 0 {
			output.Fatal("You must specify a package directory/path at least...")
		}

		var scheme *Scheme
		if scheme == nil {
			scheme = defaultSchemeConfig
		}

		err := formatPackage(scheme)
		fatalOnError(err)

		// Dump config
		// fatalOnError(configuration.WriteConfig())
	},
}

func init() {
	FormatCmd.Flags().StringVarP(&rootPaths, "root-paths", "r", "", "root paths")
	FormatCmd.Flags().StringVarP(&localPrefix, "local-prefix", "l", "", "prefix of the local repository")
	FormatCmd.Flags().StringVarP(&scheme, "scheme", "s", "", "verification scheme to enforce. one of stdLocalThirdParty/stdThirdPartyLocal")
	FormatCmd.Flags().StringVarP(&skipPaths, "skip-paths", "p", "", "paths to skip (regex)")
	RootCmd.AddCommand(FormatCmd)
}

func formatPackage(scheme *Scheme) error {
	if scheme == nil {
		return errSchemeIsNil
	}

	pp.Println(scheme)

	// var skipPaths stringArrayFlags
	// flag.Var(&skipPaths, "skip", "paths to skip (regex)")

	// verificationScheme, err := getVerificationSchemeType(scheme)
	// if err != nil {
	// 	return err
	// }

	// TODO: can parallelize across root paths
	// for argIndex := 0; argIndex < flag.NArg(); argIndex++ {

	skipPathsList := strings.Split(skipPaths, ",")
	rootPathsArr := strings.Split(rootPaths, ",")
	for _, rootPath := range rootPathsArr {
		// rootPath := flag.Arg(argIndex)
		impiInstance, err := impi.NewImpi(options.numCPUs)
		if err != nil {
			return fmt.Errorf("Failed to create impi: %s", err.Error())
		}

		err = impiInstance.Verify(
			rootPath,
			&impi.VerifyOptions{
				SkipTests:   false,
				LocalPrefix: localPrefix,
				// Scheme:      verificationScheme,
				SkipPaths: skipPathsList,
			},
			&consoleErrorReporter{},
		)

		if err != nil {
			return err
		}
	}

	return nil
}
