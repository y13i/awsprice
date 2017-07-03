package awsprice

import "github.com/urfave/cli"

func listProducts() cli.Command {
	return cli.Command{
		Name:  "listProducts",
		Usage: "list products match given filters",

		Flags: []cli.Flag{
			flags["offerCode"],
			flags["region"],
			flags["offerVersion"],
			flags["productFamily"],
			flags["attribute"],
		},

		Action: func(c *cli.Context) error {
			as, e := stringsToAttributes(c.StringSlice("attribute"))

			if e != nil {
				return e
			}

			ps, e := getProductSlice(
				OfferCode(c.String("offerCode")),
				RegionCode(c.String("region")),
				Version(c.String("offerVersion")),
				stringsToProductFamilies(c.StringSlice("productFamily")),
				as,
			)

			if e != nil {
				return e
			}

			print(ps)

			return nil
		},
	}
}
