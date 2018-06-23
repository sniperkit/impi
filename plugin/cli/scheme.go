package cli

import (
	"fmt"

	impi "github.com/sniperkit/snk.golang.impi/pkg"
)

type Scheme struct {
	AutoSave   bool                     `json:"auto_save" yaml:"auto_save" toml:"auto_save"`
	StrictMode bool                     `json:"strict_mode" yaml:"strict_mode" toml:"strict_mode"`
	OrderBy    string                   `json:"order_by" yaml:"order_by" toml:"order_by"`
	Groups     map[string]*SchemeConfig `json:"groups" yaml:"groups" toml:"groups"`
}

type SchemeConfig struct {
	Aliases string         `json:"aliases" yaml:"aliases" toml:"aliases"`
	Desc    string         `json:"schemeDesc" yaml:"schemeDesc" toml:"schemeDesc"`
	Pattern string         `json:"pattern" yaml:"pattern" toml:"pattern"`
	Comment *SchemeComment `json:"comment" yaml:"comment" toml:"comment"`
}

type SchemeComment struct {
	Header    string `json:"header" yaml:"header" toml:"header"`
	Footer    string `json:"footer" yaml:"footer" toml:"footer"`
	Separator string `json:"separator" yaml:"separator" toml:"separator"`
}

func getVerificationSchemeTypeStr(scheme string) (impi.ImportGroupVerificationScheme, error) {
	switch scheme {
	case "stdLocalThirdParty":
		return impi.ImportGroupVerificationSchemeStdLocalThirdParty, nil
	case "stdThirdPartyLocal":
		return impi.ImportGroupVerificationSchemeStdThirdPartyLocal, nil
	default:
		return 0, fmt.Errorf("Unsupported verification scheme: %s", scheme)
	}
}

func dumpConfig(prefixPath string) error {
	// fatalOnError(configuration.WriteConfig())
	return nil
}
