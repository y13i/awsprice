package awsprice

import (
	"github.com/bradfitz/slice"
	"github.com/urfave/cli"
)

func listOfferRegions() cli.Command {
	return cli.Command{
		Name:  "listOfferRegions",
		Usage: "list all regions of specified offer",

		Flags: []cli.Flag{
			flags["offerCode"],
		},

		Action: func(c *cli.Context) error {
			oc := OfferCode(c.String("offerCode"))
			ori, e := getOfferRegionIndex(oc)

			if e != nil {
				return e
			}

			rcs := []RegionCode{}

			for k := range ori.Regions {
				rcs = append(rcs, k)
			}

			slice.Sort(rcs, func(i, j int) bool {
				return rcs[i] < rcs[j]
			})

			print(rcs)

			return nil
		},
	}
}
