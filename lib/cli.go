package awsprice

import "github.com/urfave/cli"

var outputFormat string
var backend backendClient

// GetCLIApp returns cli.App instance
func GetCLIApp() *cli.App {
	app := cli.NewApp()

	app.Name = "awsprice"
	app.Usage = "CLI for AWS Price List API."
	app.Version = "0.0.3"
	app.EnableBashCompletion = true
	app.Flags = globalFlags

	app.Before = func(c *cli.Context) error {
		outputFormat = c.GlobalString("format")

		backend = backendClient{
			CacheTTL: c.GlobalString("cacheTTL"),
		}

		return nil
	}

	app.After = func(c *cli.Context) error {
		if c.GlobalBool("clearCache") {
			return clearCache()
		}

		return nil
	}

	app.Commands = []cli.Command{
		listOffers(),
		listOfferRegions(),
		listOfferVersions(),
		listProductFamilies(),
		listAttributes(),
		listProducts(),
		listTermTypes(),
		listProductTerms(),
	}

	return app
}
