package cli

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
