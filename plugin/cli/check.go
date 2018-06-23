package cli

import (
	"context"

	// external
	"github.com/sniperkit/snk.golang.cobra"
)

var checker = map[string]func(ctx context.Context, args []string){
	"check-errors":         listErrors,
	"list-skip-paths":      listSkipPaths,
	"list-conf-schemes":    listSchemeFromConf,
	"list-conf-skip-paths": listSkipPathsFromConf,
	"validate-schemes":     validateScheme,
}

var CheckCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"checker", "c", "chk"},
	Short:   "List/validate schemes and skpi-paths from config file or cli arguments.",
	// Example: fmt.Sprintf(" %s ^github.com/google/go-github github.com/sniperkit/go-github/pkg ", conf.ProgramName),
	Run: func(cmd *cobra.Command, args []string) {
		output := getOutput()

		if len(args) == 0 {
			output.Fatal("You must specify a package directory/path at least...")
		}

		err := formatPackage(&Scheme{})
		fatalOnError(err)

		// Dump config
		// fatalOnError(configuration.WriteConfig())
	},
}

func init() {
	// CheckCmd.Flags().StringVarP(&skipPaths, "skip-paths", "p", "", "skip", "paths to skip (regex)")
	RootCmd.AddCommand(CheckCmd)
}

func listSchemeFromConf(ctx context.Context, _ []string) {}

func listSkipPathsFromConf(ctx context.Context, _ []string) {}

func validateScheme(_ context.Context, args []string) {}

func listSkipPaths(_ context.Context, args []string) {}

func listErrors(_ context.Context, args []string) {}
