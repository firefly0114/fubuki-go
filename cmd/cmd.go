package cmd

import "flag"

var (
	Debug          bool
	ApiListingPort int
)

func init() {
	flag.BoolVar(&Debug, "d", false, "enable debug mode")
	flag.IntVar(&ApiListingPort, "api-port", 9900, "manage api port")
	flag.Parse()
}
