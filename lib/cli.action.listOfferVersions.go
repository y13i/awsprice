package awsprice

import (
	"github.com/bradfitz/slice"
	"github.com/urfave/cli"
)

func listOfferVersions() cli.Command {
	return cli.Command{
		Name:  "listOfferVersions",
		Usage: "list all versions of specified offer",

		Flags: []cli.Flag{
			flags["offerCode"],
		},

		Action: func(c *cli.Context) error {
			oc := OfferCode(c.String("offerCode"))
			ovi, e := getOfferVersionIndex(oc)

			if e != nil {
				return e
			}

			vs := VersionSlice{}

			for k := range ovi.Versions {
				vs = append(vs, k)
			}

			slice.Sort(vs, func(i, j int) bool {
				return vs[i] < vs[j]
			})

			print(vs)

			return nil
		},
	}
}
