package awsprice

import "github.com/urfave/cli"

var flags = map[string]cli.Flag{
	"offerCode": cli.StringFlag{
		Name:  "offerCode, o",
		Usage: "A unique code for the product of an AWS service",
		Value: "AmazonEC2",
	},

	"region": cli.StringFlag{
		Name:   "region, r",
		Usage:  "AWS region code",
		EnvVar: "AWS_REGION",
	},

	"offerVersion": cli.StringFlag{
		Name:  "offerVersion",
		Usage: "the version of the offer file (default: current version)",
	},

	"productFamily": cli.StringSliceFlag{
		Name:  "productFamily, p",
		Usage: "filter products by the name of product family",
	},

	"attribute": cli.StringSliceFlag{
		Name:  "attribute, a",
		Usage: "filter products by the attribute in \"KEY[=VALUE]\" format",
	},

	"termType": cli.StringFlag{
		Name:  "termType",
		Usage: "type of the terms. \"OnDemand\" or \"Reserved\"",
		Value: "OnDemand",
	},

	"priceUnit": cli.StringFlag{
		Name:  "priceUnit",
		Usage: "unit of prices",
		Value: "USD",
	},

	"termAttribute": cli.StringSliceFlag{
		Name:  "termAttribute, t",
		Usage: "filter products by the term attribute in \"KEY[=VALUE]\" format",
	},
}

var globalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "cacheTTL",
		Usage: "max age of cache. if existing cache file is older than this value, cache will be recreated. using time.ParseDuration(s string)",
		Value: "24h",
	},

	cli.StringFlag{
		Name:  "format, f",
		Usage: "output format (\"pp\" or \"json\")",
		Value: "pp",
	},

	cli.BoolFlag{
		Name:  "clearCache",
		Usage: "remove all cache files before exiting",
	},
}
