package cli

const UsageTemplate = `
## Overview

Verify proper golang import directives, beyond the capability of gofmt and goimports. Rather than just verifying import order, it classifies imports to three types:
1. "Std": Standard go imports like "fmt", "os", "encoding/json"
2. "Local": Packages which are part of the current project
3. "Third party": Packages which are not standard and not local

It can then verify, according to the chosen scheme, that each import resides in the proper import group. 
Import groups are declared in the 'import()' directive and are separated by an empty line:

'''
import(
    "Group #1 import #1"
    "Group #1 import #2"

    "Group #2 import #1"
    "Group #2 import #2"
    // comments are allowed within a group
    "Group #2 import #3"

    "Group #3 import #1"
    "Group #3 import #2"
)
'''

Note that impi does not support regenerating the files, only warns of infractions. 

## Usage

$ impi [--local <local import prefix>] --scheme <scheme> <packages>

example with args:
$ impi --local github.com/nuclio/nuclio/ --scheme stdLocalThirdParty ./cmd/... ./pkg/...

example with config file:
$ impi --config ~/impi.yaml ./cmd/... ./pkg/...

## Schemes configuration:



`
