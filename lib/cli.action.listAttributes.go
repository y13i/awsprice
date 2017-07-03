package awsprice

import (
	"github.com/bradfitz/slice"
	"github.com/urfave/cli"
)

func listAttributes() cli.Command {
	return cli.Command{
		Name:  "listAttributes",
		Usage: "list attributes of the offer products",

		Flags: []cli.Flag{
			flags["offerCode"],
			flags["region"],
			flags["offerVersion"],
			flags["productFamily"],
			flags["attribute"],
		},

		Action: func(c *cli.Context) error {
			ps, e := getProductSlice(
				OfferCode(c.String("offerCode")),
				RegionCode(c.String("region")),
				Version(c.String("offerVersion")),
				stringsToProductFamilies(c.StringSlice("productFamily")),
				stringsToAttributes(c.StringSlice("attribute")),
			)

			if e != nil {
				return e
			}

			am := map[AttributeKey][]AttributeValue{}
			vm := map[AttributeKey]map[AttributeValue]bool{}

			for _, p := range ps {
				for ak, av := range p.Attributes {
					if !vm[ak][av] {
						am[ak] = append(am[ak], av)

						_, ok := vm[ak]

						if !ok {
							vm[ak] = map[AttributeValue]bool{}
						}

						vm[ak][av] = true
					}
				}
			}

			for _, av := range am {
				slice.Sort(av, func(i, j int) bool {
					return av[i] < av[j]
				})
			}

			print(am)

			return nil
		},
	}
}
