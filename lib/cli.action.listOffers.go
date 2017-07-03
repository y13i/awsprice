package awsprice

import (
	"github.com/bradfitz/slice"
	"github.com/urfave/cli"
)

func listOffers() cli.Command {
	return cli.Command{
		Name:  "listOffers",
		Usage: "list all offer codes",

		Action: func(c *cli.Context) error {
			oi, e := getOfferIndex()

			if e != nil {
				return e
			}

			var ocs OfferCodes

			for k := range oi.Offers {
				ocs = append(ocs, k)
			}

			slice.Sort(ocs, func(i, j int) bool {
				return ocs[i] < ocs[j]
			})

			print(ocs)

			return nil
		},
	}
}
