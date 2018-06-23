package cli

import (
	"fmt"
	"strings"

	// external
	"github.com/sniperkit/snk.golang.cobra"
	// conf "github.com/sniperkit/snk.golang.impi/pkg/config"
)

const defaultSchemeSperator = "\n"

var (
	configFormat string = ""

	currentPackageURI string = ""

	validConfigExportFormats = []string{"yaml"}

	defaultSchemeConfig = &Scheme{
		AutoSave:   true,
		StrictMode: false,                              // will check if an a group alias match
		OrderBy:    "standard,local,plugin,thirdparty", // if empty order by alphabetical order
		Groups: map[string]*SchemeConfig{
			"standard": &SchemeConfig{
				Aliases: "std",
				Desc:    "All golang default packages (`fmt`, `os`, `ioutil`, ...)",
				Pattern: "^", // no tld in the package namespace

				Comment: &SchemeComment{
					Header:    "// golang - standard package",
					Footer:    "",
					Separator: defaultSchemeSperator,
				},
			},
			"local": &SchemeConfig{
				Aliases: "core|internal-core",
				Desc:    "All packages under pkg",
				Pattern: "",
				Comment: &SchemeComment{
					Header:    "// local - core package(s)",
					Footer:    "",
					Separator: defaultSchemeSperator,
				},
			},
			"plugin": &SchemeConfig{
				Aliases: "core-plugin|local-plugin|internal-plugin",
				Desc:    "",
				Pattern: "",
				Comment: &SchemeComment{
					Header:    "// local - additional plugin package(s)",
					Footer:    "",
					Separator: defaultSchemeSperator,
				},
			},
			"thirdparty": &SchemeConfig{
				Aliases: "3rparty|external|vendor",
				Desc:    "All packages providing from other packages/librairies",
				Pattern: "",
				Comment: &SchemeComment{
					Header:    "// external - thirdparty package(s)",
					Footer:    "",
					Separator: defaultSchemeSperator,
				},
			},
		},
	}
)

var ConfigCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"conf", "c", "config", "prefs"},
	Short:   "Config or dump schemes and skpi-paths into a local/global config file.",
	// Example: fmt.Sprintf(" %s ^github.com/google/go-github github.com/sniperkit/go-github/pkg ", conf.ProgramName),
	Run: func(cmd *cobra.Command, args []string) {
		output := getOutput()

		if len(args) == 0 {
			output.Fatal("You must specify a package directory/path at least...")
		}

		// Dump config
		fatalOnError(configuration.WriteConfig())
	},
}

func init() {
	validFormatInfo := fmt.Sprintf("export config file to formats. (valid formats: %s)", strings.Join(validConfigExportFormats, ","))

	ConfigCmd.Flags().StringVarP(&configFormat, "skip-paths", "p", "", validFormatInfo)
	RootCmd.AddCommand(ConfigCmd)
}

type skipPathsList []string

func (s *skipPathsList) String() string {
	return strings.Join(*s, ",")
}

func (s *skipPathsList) Set(value string) error {
	*s = append(*s, value)
	return nil
}
