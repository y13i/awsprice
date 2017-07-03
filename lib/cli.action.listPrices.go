package awsprice

import "github.com/urfave/cli"

// ProductTerm conbines Product and OfferTerms
type ProductTerm struct {
	Product    Product
	OfferTerms OfferTerms
}

func listProductTerms() cli.Command {
	return cli.Command{
		Name:  "listProductTerms",
		Usage: "list offer terms of products match given filters",

		Flags: []cli.Flag{
			flags["offerCode"],
			flags["region"],
			flags["offerVersion"],
			flags["productFamily"],
			flags["attribute"],
			flags["termType"],
			flags["priceUnit"],
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

			ov, e := getOfferVersion(
				OfferCode(c.String("offerCode")),
				RegionCode(c.String("region")),
				Version(c.String("offerVersion")),
			)

			if e != nil {
				return e
			}

			ptm := ov.Terms[TermType(c.String("termType"))]

			pts := []ProductTerm{}

			for _, p := range ps {
				productPrice := ProductTerm{
					Product:    p,
					OfferTerms: ptm[p.SKU],
				}

				pts = append(pts, productPrice)
			}

			print(pts)

			return nil
		},
	}
}
