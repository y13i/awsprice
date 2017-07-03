package awsprice

import (
	"github.com/bradfitz/slice"
	"github.com/urfave/cli"
)

func listProductFamilies() cli.Command {
	return cli.Command{
		Name:  "listProductFamilies",
		Usage: "list product families of the offer products",

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

			pfs := []ProductFamily{}
			vm := map[ProductFamily]bool{}

			for _, p := range ps {
				if !vm[p.ProductFamily] {
					pfs = append(pfs, p.ProductFamily)
					vm[p.ProductFamily] = true
				}
			}

			slice.Sort(pfs, func(i, j int) bool {
				return pfs[i] < pfs[j]
			})

			print(pfs)

			return nil
		},
	}
}
