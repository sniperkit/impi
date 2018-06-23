package cli

import (
	"fmt"
	"os"

	// external
	homedir "github.com/mitchellh/go-homedir"

	// internal - core
	conf "github.com/sniperkit/snk.golang.impi/pkg/config"
	ver "github.com/sniperkit/snk.golang.impi/pkg/version"
)

var (
	configDirectoryPath  string = "."
	currentProjectGitURI string = "."
)

type VCS interface {
	Info()
}

type Git struct {
	Context *VCSContext
}

func (g *Git) Info() interface{} {
	return g.Context
}

type Gitlab struct {
	Context *VCSContext
}

func (g *Gitlab) Info() interface{} {
	return g.Context
}

type Bitbucket struct {
	Context *VCSContext
}

func (b *Bitbucket) Info() interface{} {
	return b.Context
}

type VCSType string

const (
	GIT VCSType = "git"
	SVN VCSType = "svn"
	HG  VCSType = "hg"
)

func (t VCSType) String() string {
	return fmt.Sprintf("%v", t)
}

type VCSProvider string

const (
	GITHUB    VCSProvider = "github"
	GITLAB    VCSProvider = "gitlab"
	BITBUCKET VCSProvider = "bitbucket"
	PRIVATE   VCSProvider = "private"
)

func (p VCSProvider) String() string {
	return fmt.Sprintf("%v", p)
}

type VCSContext struct {
	Type      VCSType     `json:"paths" yaml:"paths" toml:"paths"`
	Provider  VCSProvider `json:"paths" yaml:"paths" toml:"paths"`
	Owner     string      `json:"paths" yaml:"paths" toml:"paths"`
	Name      string      `json:"paths" yaml:"paths" toml:"paths"`
	URI       string      `json:"paths" yaml:"paths" toml:"paths"`
	RemoteURL string      `json:"paths" yaml:"paths" toml:"paths"`

	Paths struct {
		Prefix   string `json:"prefix" yaml:"prefix" toml:"prefix"`
		Absolute string `json:"absolute" yaml:"absolute" toml:"absolute"`
		Relative string `json:"relative" yaml:"relative" toml:"relative"`
	} `json:"paths" yaml:"paths" toml:"paths"`

	Build struct {
		Version  string `json:"version" yaml:"version" toml:"version"`
		Tag      string `json:"tag" yaml:"tag" toml:"tag"`
		Total    string `json:"total" yaml:"total" toml:"total"`
		TimePrev string `json:"prevCommitTime" yaml:"prevCommitTime" toml:"prevCommitTime"`
	} `json:"commit" yaml:"commit" toml:"commit"`

	Commit struct {
		ID       string
		Hash     string
		TimePrev string `json:"prevCommitTime" yaml:"prevCommitTime" toml:"prevCommitTime"`
	} `json:"commit" yaml:"commit" toml:"commit"`
}

var (
	currentWorkDirectory, _ = os.Getwd()
)

func printContextInfo() {
	fmt.Println("ProgramName=", conf.ProgramName)
	fmt.Println("ProgramVersion=", ver.Version)
	fmt.Println("CurrentWorkDirectory=", currentWorkDirectory)
	fmt.Println("ConfigDirectoryPath=", configDirectoryPath)
	fmt.Println("CurrentProjectGitURI=", currentProjectGitURI)

}
